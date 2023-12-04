package backend

import (
	"capture/internal/app"
	"capture/internal/trust"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/google/gopacket"
)

type TrustIds struct {
	ch              chan TrustIdsFeature
	winSize         int
	packetThreshold int
	src             *gopacket.PacketSource
	opts            Options
	trust           trust.Trust
	trustLock       sync.Mutex
}

// chan []TrustIdsFeature
// or
// chan TrustIdsWindow
// struct TrustIdsWindow {...}

// TODO: create a structure that fits what we collect
type TrustIdsFeature struct {
	Timestamp   time.Time
	SrcIp       gopacket.Endpoint
	PayloadSize int
}

func (f TrustIdsFeature) IsValid() bool {
	var ep gopacket.Endpoint

	return ep != f.SrcIp
}

func NewTrustIdsFeature(p gopacket.Packet) TrustIdsFeature {
	var ts time.Time
	var src gopacket.Endpoint
	var size int

	ts = p.Metadata().Timestamp

	// ethernet, arp, etc
	if p.NetworkLayer() != nil {
		src = p.NetworkLayer().NetworkFlow().Src()
		size = len(p.NetworkLayer().LayerPayload())
	}

	return TrustIdsFeature{
		Timestamp:   ts,
		SrcIp:       src,
		PayloadSize: size,
	}
}

func NewTrustIds(src *gopacket.PacketSource, opts *Options, trust trust.Trust) *TrustIds {
	app.Log.Trace("kit (trustIds): new")

	iot := new(TrustIds)

	iot.ch = make(chan TrustIdsFeature, 10)
	iot.winSize = 1000
	iot.packetThreshold = 5
	iot.src = src
	iot.opts = opts.Copy()
	iot.trust = trust

	app.Stats["kit.trustIds"] = iot

	return iot
}

func (iot *TrustIds) Run(wg *sync.WaitGroup) {
	app.Log.Trace("kit (trustIds): run")
	wg.Add(1)

	go iot.accumulate()

	if iot.opts.PcapReplay {
		go iot.replay(wg)
	} else {
		go iot.read(wg)
	}
}

func (iot *TrustIds) replay(wg *sync.WaitGroup) {
	app.Log.Trace("kit (trustIds): replay")

	app.Log.Info("kit: reader started")
	var prev, next time.Time
	for curr := range iot.src.Packets() {
		next = curr.Metadata().Timestamp
		app.Log.Tracef("next timestamp: %s\n", next.String())
		if !prev.IsZero() {
			diff := next.Sub(prev)
			app.Log.Tracef("next packet at: %s\n", diff)
			time.Sleep(diff)
		}

		feat := NewTrustIdsFeature(curr)

		if feat.IsValid() {
			app.Log.Tracef("kit (trustIds): packet from %s", feat.SrcIp.String())
			iot.ch <- feat
		} else {
			app.Log.Debug("kit (trustIds): skipped packet: lower than network layer")
		}

		prev = next
	}

	wg.Done()
	app.Log.Info("kit: reader finished")
	if iot.opts.AppStats {
		app.PrintStats()
	}
}

func (iot *TrustIds) read(wg *sync.WaitGroup) {
	app.Log.Trace("kit (trustIds): read")

	app.Log.Info("kit: reader started")
	for curr := range iot.src.Packets() {
		feat := NewTrustIdsFeature(curr)

		if feat.IsValid() {
			app.Log.Tracef("kit (trustIds): packet from %s", feat.SrcIp.String())
			iot.ch <- feat
		} else {
			app.Log.Debug("kit (trustIds): skipped packet: lower than network layer")
		}
	}

	wg.Done()
	app.Log.Info("kit: reader finished")
	if iot.opts.AppStats {
		app.PrintStats()
	}
}

func (iot *TrustIds) accumulate() {
	app.Log.Trace("kit (trustIds): accumulate")
	window := [][]TrustIdsFeature{}
	size := 0
	curr := 0

	window = append(window, make([]TrustIdsFeature, iot.winSize))

	// TODO: add timer also
	for {
		feat := <-iot.ch

		window[curr] = append(window[curr], feat)

		size++
		if size == iot.winSize {
			// TODO: check this works as expected
			go iot.process(window[curr])
			curr = curr + 1
			window = append(window, make([]TrustIdsFeature, iot.winSize))
			size = 0
		}
	}
}

func (iot *TrustIds) process(data []TrustIdsFeature) {
	app.Log.Trace("kit (trustIds): process")
	winNodes := make(map[gopacket.Endpoint][]TrustIdsFeature)

	// split features per node
	for _, d := range data {
		// why?
		if d.PayloadSize == 0 {
			continue
		}
		winNodes[d.SrcIp] = append(winNodes[d.SrcIp], d)
	}

	var wg sync.WaitGroup

	// for each node
	// compute packetrate and throughput
	// check thresholds and either penalize or reward
	for node, packets := range winNodes {
		wg.Add(1)
		go iot.networkMetrics(&wg, node, packets)
	}

	wg.Wait()

	// publish trust values
	iot.trust.Update()
}

func (iot *TrustIds) networkMetrics(wg *sync.WaitGroup, node gopacket.Endpoint, packets []TrustIdsFeature) {
	app.Log.Trace("kit (trustIds): networkMetrics")
	defer wg.Done()

	last := len(packets) - 1
	tic := packets[0].Timestamp
	toc := packets[last].Timestamp
	duration := float64(toc.Sub(tic).Nanoseconds())

	if duration < float64(time.Microsecond) {
		app.Log.Infof("kit (trustIds): skip %s, %f nsec", node.String(), duration)
		return
	}

	if len(packets) < iot.packetThreshold {
		app.Log.Infof("kit (trustIds): skip %s, %d packets", node.String(), len(packets))
		return
	}

	throughput := float64(0)
	packetrate := float64(len(packets)*1000000000) / duration

	for _, p := range packets {
		throughput += float64(p.PayloadSize)
	}

	throughput = throughput * 1000000 / duration

	app.Log.Infof("kit (trustIds): %s: packets=%d, duration=%f, packetrate=%f", node.String(), len(packets), duration, packetrate)
	app.Log.Tracef("kit (trustIds): %s: packets=%d, duration=%f, throughput=%f", node.String(), len(packets), duration, throughput)

	iot.trustLock.Lock()
	if packetrate > float64(iot.opts.TrustIdsOpts.PacketRate) {
		iot.trust.Penalize(node.String(), fmt.Sprintf("packet rate %f", packetrate))
	} else {
		iot.trust.Reward(node.String())
	}

	if throughput > float64(iot.opts.TrustIdsOpts.Throughput) {
		iot.trust.Penalize(node.String(), fmt.Sprintf("throughput %f", throughput))
	} else {
		iot.trust.Reward(node.String())
	}
	iot.trustLock.Unlock()
}

func (iot *TrustIds) Report(w io.Writer) {
	app.Log.Trace("kit (trustIds): report stats")
}

// TODO: work with cross correlation
