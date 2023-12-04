package api

import (
	"capture/internal/api/backend"
	"capture/internal/app"
	"capture/internal/broker"
)

type API interface {
	BlockNode(ip, reason string)
	DropClient(ip string)
	SendTrustValues(payload []byte)
}

func New(apiBackend string) API {
	app.Log.Trace("api: new")

	switch backendFromString(apiBackend) {
	case AmqpBackend:
		app.Log.Info("api: backend=amqp")
		return backend.NewAmqp(backend.NewAmqpOptions(), broker.NewAmqpOptions())
	case LogBackend:
		app.Log.Info("api: backend=log")
		return backend.NewLog()
	default:
		app.Log.Info("api: backend=log")
		return backend.NewLog()
	}
}
