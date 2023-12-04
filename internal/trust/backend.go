package trust

import (
	"capture/internal/app"
	"fmt"
)

type Backend int

const (
	ThresholdsBackend Backend = iota
	NoneBackend
)

func AvailableBackends() {
	app.Log.Trace("trust: available backend")

	fmt.Println("Available Trust Backends:")
	fmt.Println("- thresholds (default)")
}

func backendFromString(trustBackend string) Backend {
	app.Log.Trace("trust: string to backend")

	switch trustBackend {
	case "thresholds":
		return ThresholdsBackend
	default:
		app.Log.Warning("trust: invalid backend: " + trustBackend)
		return NoneBackend
	}
}
