package broker

import (
	"capture/internal/app"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"sync"

	"github.com/streadway/amqp"
	"golang.org/x/crypto/ssh"
)

const (
	ErrAmqpOpenConnection  string = "amqp: failed to open connection"
	ErrAmqpCloseConnection string = "amqp: failed to close connection"
	ErrAmqpOpenChannel     string = "amqp: failed to open channel"
	ErrAmqpCloseChannel    string = "amqp: failed to close channel"
	ErrAmqpConfirmMode     string = "amqp: failed to set channel into confirm mode"
	ErrAmqpPublish         string = "amqp: failed to publish message"
)

type AmqpBroker struct {
	sync.RWMutex
	cfg  amqp.Config
	scfg *tls.Config
	conn *amqp.Connection
	chnl *amqp.Channel
	opts AmqpOptions
}

func New(opts *AmqpOptions) (*AmqpBroker, error) {
	app.Log.Trace("amqp: new")
	var err error

	br := new(AmqpBroker)
	br.opts = opts.Copy()

	if br.opts.UseTls {
		err = br.tlsConnection()
	} else {
		err = br.plainConnection()
	}

	if err != nil {
		return nil, fmt.Errorf(ErrAmqpOpenConnection + ": " + err.Error())
	}
	app.Log.Info("amqp: connected to '" + br.opts.SafeUrl() + "' as " + br.opts.Username)

	br.chnl, err = br.conn.Channel()
	if err != nil {
		br.conn.Close()
		return nil, fmt.Errorf(ErrAmqpOpenChannel + ": " + err.Error())
	}

	// Should this be fatal?
	err = br.chnl.Confirm(false)
	if err != nil {
		return br, fmt.Errorf(ErrAmqpConfirmMode + ": " + err.Error())
	}

	return br, nil
}

func (br *AmqpBroker) Close() error {
	app.Log.Trace("amqp: close")

	br.chnl.Close()

	err := br.conn.Close()
	if err != nil {
		return fmt.Errorf(ErrAmqpCloseConnection + ": " + err.Error())
	}
	app.Log.Info("amqp: disconnected from '" + br.opts.SafeUrl() + "'")

	return nil
}

func (br *AmqpBroker) Publish(topic string, msg []byte) error {
	app.Log.Tracef("amqp: publishing to '%s':\n%s", topic, string(msg))

	err := br.chnl.Publish(
		br.opts.Exchange, // exchange
		topic,            // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: br.opts.ContentType,
			Body:        msg,
		},
	)
	if err != nil {
		return fmt.Errorf(ErrAmqpPublish + ": " + err.Error())
	}
	app.Log.Debugf("amqp: publish success")

	return nil
}

// http://www.rabbitmq.com/ssl.html
func (br *AmqpBroker) tlsConnection() error {
	app.Log.Trace("amqp: tls connection")
	var err error

	br.scfg = new(tls.Config)
	br.scfg.RootCAs = x509.NewCertPool()

	// Load CA certificate as trusted
	ca, err := os.ReadFile(br.opts.CACertFile)
	app.DieOnErr(err)

	br.scfg.RootCAs.AppendCertsFromPEM(ca)

	// Load own certificate
	// If a passphrase is given, assume key is encrypted
	if false {
		keyBytes, err := os.ReadFile(br.opts.ClientKeyFile)
		app.DieOnErr(err)

		key, err := ssh.ParseRawPrivateKeyWithPassphrase(keyBytes, []byte(""))
		app.DieOnErr(err)

		// accept only rsa
		switch k := key.(type) {
		case *rsa.PrivateKey:
		default:
			return fmt.Errorf("unexpected key type: %T", k)
		}

		decrypted, err := x509.MarshalPKCS8PrivateKey(key)
		app.DieOnErr(err)

		certBytes, err := os.ReadFile(br.opts.ClientCertFile)
		app.DieOnErr(err)

		cert, err := tls.X509KeyPair(certBytes, decrypted)
		app.DieOnErr(err)

		br.scfg.Certificates = append(br.scfg.Certificates, cert)
	} else {
		cert, err := tls.LoadX509KeyPair(br.opts.ClientCertFile, br.opts.ClientKeyFile)
		app.DieOnErr(err)

		br.scfg.Certificates = append(br.scfg.Certificates, cert)
	}

	// Try connecting
	br.conn, err = amqp.DialTLS(br.opts.Url(), br.scfg)
	if err != nil {
		return err
	}

	return nil
}

func (br *AmqpBroker) plainConnection() error {
	app.Log.Trace("amqp: plain connection")
	var err error

	br.conn, err = amqp.DialConfig(br.opts.Url(), br.cfg)
	if err != nil {
		return err
	}

	return nil
}
