package trust

import (
	"capture/internal/api"
	"capture/internal/app"
	"capture/internal/trust/backend"
)

type Trust interface {
	Update()
	Reward(node string)
	Penalize(node, reason string)
}

func New(trustBackend, apiBackend string) Trust {
	switch backendFromString(trustBackend) {
	case ThresholdsBackend:
		app.Log.Info("trust: backend: thresholds")
		opts := backend.NewThresholdsOptions()
		return backend.NewThresholdsTrust(opts, api.New(apiBackend))
	default:
		app.Log.Info("trust: backend: thresholds")
		opts := backend.NewThresholdsOptions()
		return backend.NewThresholdsTrust(opts, api.New(apiBackend))
	}
}
