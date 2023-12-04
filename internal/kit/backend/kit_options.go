package backend

import "capture/internal/app"

type TrustIdsOptions struct {
	PacketRate int
	Throughput int
}

type Options struct {
	KitBackend      string
	TrustBackend    string
	ApiBackend      string
	AppStats        bool
	KitCountPackets bool
	PcapReplay      bool
	TrustIdsOpts    TrustIdsOptions
}

func NewOptions() *Options {
	app.Log.Trace("kit: new options")
	return &Options{
		KitBackend:      app.Cfg.AppBackendKit,
		TrustBackend:    app.Cfg.AppBackendTrust,
		ApiBackend:      app.Cfg.AppBackendApi,
		AppStats:        app.Cfg.AppShowStats,
		KitCountPackets: app.Cfg.KitCountPackets,
		PcapReplay:      app.Cfg.PcapReplay,
		TrustIdsOpts: TrustIdsOptions{
			PacketRate: app.Cfg.KitTrustIdsPacketRate,
			Throughput: app.Cfg.KitTrustIdsThroughput,
		},
	}
}

func (o Options) Copy() Options {
	app.Log.Trace("kit: copy options")
	return Options{
		KitBackend:      o.KitBackend,
		TrustBackend:    o.TrustBackend,
		ApiBackend:      o.ApiBackend,
		AppStats:        o.AppStats,
		KitCountPackets: o.KitCountPackets,
		PcapReplay:      o.PcapReplay,
		TrustIdsOpts: TrustIdsOptions{
			PacketRate: o.TrustIdsOpts.PacketRate,
			Throughput: o.TrustIdsOpts.Throughput,
		},
	}
}
