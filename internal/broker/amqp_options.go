package broker

import (
	"capture/internal/app"
	"fmt"
)

type AmqpOptions struct {
	Host           string
	Port           int
	ClientId       string
	Username       string
	Password       string
	Exchange       string
	ContentType    string
	UseTls         bool
	CACertFile     string
	ClientCertFile string
	ClientKeyFile  string
}

func NewAmqpOptions() *AmqpOptions {
	app.Log.Trace("amqp: new options")
	return &AmqpOptions{
		Host:           app.Cfg.AmqpHost,
		Port:           app.Cfg.AmqpPort,
		ClientId:       app.Cfg.AppClientId,
		Username:       app.Cfg.AmqpUsername,
		Password:       app.Cfg.AmqpPassword,
		Exchange:       app.Cfg.AmqpExchange,
		ContentType:    app.Cfg.AmqpContentType,
		UseTls:         app.Cfg.AmqpUseTLS,
		CACertFile:     app.Cfg.AmqpCACertFile,
		ClientCertFile: app.Cfg.AmqpClientCertFile,
		ClientKeyFile:  app.Cfg.AmqpClientKeyFile,
	}
}

func (o AmqpOptions) Copy() AmqpOptions {
	app.Log.Trace("amqp: copy options")
	return AmqpOptions{
		Host:           o.Host,
		Port:           o.Port,
		ClientId:       o.ClientId,
		Username:       o.Username,
		Password:       o.Password,
		Exchange:       o.Exchange,
		ContentType:    o.ContentType,
		UseTls:         o.UseTls,
		CACertFile:     o.CACertFile,
		ClientCertFile: o.ClientCertFile,
		ClientKeyFile:  o.ClientKeyFile,
	}
}

func (o AmqpOptions) Url() string {
	app.Log.Trace("options: broker amqp: connection url")
	url := "amqp"

	if o.UseTls {
		url += "s://"
	} else {
		url += "://"
	}

	if o.Username != "" && o.Password != "" {
		url += o.Username + ":" + o.Password + "@"
	}

	url += o.Host
	if o.Port != 0 {
		url += ":" + fmt.Sprint(o.Port)
	}

	return url
}

func (o AmqpOptions) SafeUrl() string {
	app.Log.Trace("options: broker amqp: connection url for safe logging")
	url := "amqp"

	if o.UseTls {
		url += "s://"
	} else {
		url += "://"
	}

	url += o.Host
	if o.Port != 0 {
		url += ":" + fmt.Sprint(o.Port)
	}

	return url
}
