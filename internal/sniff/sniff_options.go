package sniff

import (
	"capture/internal/app"

	"github.com/google/gopacket/pcap"
)

type Options struct {
	BufSize int
	SnapLen int
	// TODO: check that this converts correctly
	Direction pcap.Direction
	Filter    string
}

func NewOptions() *Options {
	app.Log.Trace("sniff: new options")
	return &Options{
		BufSize:   app.Cfg.PcapBufSize,
		SnapLen:   app.Cfg.PcapSnapLen,
		Direction: pcap.Direction(app.Cfg.PcapDirection),
		Filter:    app.Cfg.PcapFilter,
	}
}

func (o Options) Copy() Options {
	app.Log.Trace("sniff: copy options")
	return Options{
		BufSize:   o.BufSize,
		SnapLen:   o.SnapLen,
		Direction: o.Direction,
		Filter:    o.Filter,
	}
}
