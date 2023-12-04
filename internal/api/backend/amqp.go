package backend

import (
	"capture/internal/api/message"
	"capture/internal/app"
	"capture/internal/broker"
)

const (
	ErrAmqpConnection string = "api: amqp: connection failure"
)

type Amqp struct {
	broker *broker.AmqpBroker
	opts   AmqpOptions
}

func NewAmqp(apiOptions *AmqpOptions, brokerOptions *broker.AmqpOptions) *Amqp {
	app.Log.Trace("api: amqp: new")
	var err error
	a := new(Amqp)

	a.broker, err = broker.New(brokerOptions)
	if err != nil {
		app.Log.Err(ErrAmqpConnection + ": " + err.Error())
		a = nil
		return nil
	}

	a.opts = apiOptions.Copy()

	return a
}

func (a *Amqp) BlockNode(ip, reason string) {
	app.Log.Trace("api: amqp: block node")

	m := message.NewWarn(ip, reason)
	a.pubWarn(m.ToJson())
}

func (a *Amqp) DropClient(ip string) {
	app.Log.Trace("api: amqp: drop client")

	m := message.NewMtdWarn(ip)
	a.pubMtdWarn(m.ToJson())
}

func (a *Amqp) SendTrustValues(payload []byte) {
	app.Log.Trace("api: amqp: send trust values")

	a.pubTrust(payload)
}

func (a *Amqp) pubWarn(payload []byte) {
	app.Log.Trace("api: amqp: pub warn")

	if len(payload) > 0 {
		err := a.broker.Publish(a.opts.WarnTopic, payload)
		if err != nil {
			app.Log.Warningf("api: amqp: error publishing: %s", err)
			app.Log.Warningf("api: amqp: payload:\n%s", string(payload))
			return
		}
	} else {
		app.Log.Debug("api: amqp: pub warn: skipping empty payload")
	}
}

func (a *Amqp) pubMtdWarn(payload []byte) {
	app.Log.Trace("api: amqp: pub mtdwarn")

	if len(payload) > 0 {
		err := a.broker.Publish(a.opts.MtdTopic, payload)
		if err != nil {
			app.Log.Warningf("api: amqp: error publishing: %s", err)
			app.Log.Warningf("api: amqp: payload:\n%s", string(payload))
			return
		}
	} else {
		app.Log.Debug("api: amqp: pub mtdwarn: skipping empty payload")
	}
}

func (a *Amqp) pubTrust(payload []byte) {
	app.Log.Trace("api: amqp: pub trust")

	if len(payload) > 0 {
		err := a.broker.Publish(a.opts.TrustTopic, payload)
		if err != nil {
			app.Log.Warningf("api: amqp: error publishing: %s", err)
			app.Log.Warningf("api: amqp: payload:\n%s", string(payload))
			return
		}
	} else {
		app.Log.Debug("api: amqp: pub trust: skipping empty payload")
	}
}
