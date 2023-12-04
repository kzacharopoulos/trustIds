package api

import (
	"capture/internal/app"
	"fmt"
)

type Backend int

const (
	LogBackend Backend = iota
	AmqpBackend
	NoneBackend
)

func AvailableBackends() {
	app.Log.Trace("api: available backends")

	fmt.Println("API Backends:")
	fmt.Println("- amqp")
	fmt.Println("- log (default)")
}

func backendFromString(apiBackend string) Backend {
	app.Log.Trace("api: string to backend")

	switch apiBackend {
	case "amqp":
		return AmqpBackend
	case "log":
		return LogBackend
	default:
		app.Log.Warning("api: unknown backend: " + apiBackend)
		return NoneBackend
	}
}
