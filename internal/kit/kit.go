package kit

import (
	"capture/internal/app"
	"capture/internal/kit/backend"
	"capture/internal/trust"
	"sync"

	"github.com/google/gopacket"
)

type Kit interface {
	Run(wg *sync.WaitGroup)
}

func New(opts *backend.Options, src *gopacket.PacketSource) Kit {
	app.Log.Trace("kit: new")

	switch backendFromString(opts.KitBackend) {
	case TrustIdsBackend:
		app.Log.Info("kit: backend: trustIds")
		tr := trust.New(opts.TrustBackend, opts.ApiBackend)
		return backend.NewTrustIds(src, opts, tr)
	case LogBackend:
		app.Log.Info("kit: backend: log")
		return backend.NewLog(src, opts)
	default:
		app.Log.Info("kit: backend: log")
		return backend.NewLog(src, opts)
	}
}
