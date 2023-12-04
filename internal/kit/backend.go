package kit

import (
	"capture/internal/app"
	"fmt"
)

type Backend int

const (
	TrustIdsBackend Backend = iota
	LogBackend
	NoneBackend
)

func AvailableKits() {
	app.Log.Trace("kit: available backend")

	fmt.Println("Available Kits:")
	fmt.Println("- trustIds")
	fmt.Println("- log (default)")
}

func backendFromString(kitBackend string) Backend {
	app.Log.Trace("kit: string to backend")

	switch kitBackend {
	case "trustIds":
		return TrustIdsBackend
	case "log":
		return LogBackend
	default:
		app.Log.Warning("kit: invalid backend: " + kitBackend)
		return NoneBackend
	}
}
