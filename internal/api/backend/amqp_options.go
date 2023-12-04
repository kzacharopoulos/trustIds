package backend

import "capture/internal/app"

type AmqpOptions struct {
	WarnTopic  string
	TrustTopic string
	MtdTopic   string
}

func NewAmqpOptions() *AmqpOptions {
	app.Log.Trace("api: amqp: options: new")
	return &AmqpOptions{
		WarnTopic:  app.Cfg.AmqpWarnTopic,
		TrustTopic: app.Cfg.AmqpTrustTopic,
		MtdTopic:   app.Cfg.AmqpMtdTopic,
	}
}

func (o AmqpOptions) Copy() AmqpOptions {
	app.Log.Trace("api: amqp: options: options")
	return AmqpOptions{
		WarnTopic:  o.WarnTopic,
		TrustTopic: o.TrustTopic,
		MtdTopic:   o.MtdTopic,
	}
}
