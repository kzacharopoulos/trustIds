package backend

import (
	"capture/internal/app"
	"fmt"
	"sync"
	"time"

	"github.com/google/gopacket"
)

type Log struct {
	src  *gopacket.PacketSource
	opts Options
}

func NewLog(src *gopacket.PacketSource, opts *Options) *Log {
	app.Log.Trace("kit (log): new")

	kitLog := new(Log)
	kitLog.src = src
	kitLog.opts = opts.Copy()

	return kitLog
}

func (kitLog Log) Run(wg *sync.WaitGroup) {
	app.Log.Trace("kit (log): run")

	wg.Add(1)

	if kitLog.opts.PcapReplay {
		go kitLog.replay(wg)
	} else {
		go kitLog.read(wg)
	}
}

func (kitLog Log) replay(wg *sync.WaitGroup) {
	defer wg.Done()

	app.Log.Trace("kit (log): replay")

	app.Log.Info("kit: reader started")
	var prev, next time.Time
	i := 0
	for curr := range kitLog.src.Packets() {
		next = curr.Metadata().Timestamp
		if !prev.IsZero() {
			time.Sleep(next.Sub(prev))
		}

		if kitLog.opts.KitCountPackets {
			fmt.Printf("%d %+v\n", i, curr.String())
			i++
		} else {
			fmt.Printf("%+v\n", curr.String())
		}

		prev = next
	}
	app.Log.Info("kit: reader finished")

	if kitLog.opts.AppStats {
		app.PrintStats()
	}
}

func (kitLog Log) read(wg *sync.WaitGroup) {
	defer wg.Done()

	app.Log.Trace("kit (log): read")

	app.Log.Info("kit: reader started")
	i := 0
	for curr := range kitLog.src.Packets() {
		if kitLog.opts.KitCountPackets {
			fmt.Printf("%d %+v\n", i, curr.String())
			i++
		} else {
			fmt.Printf("%+v\n", curr.String())
		}
	}
	app.Log.Info("kit: reader finished")

	if kitLog.opts.AppStats {
		app.PrintStats()
	}
}
