package backend

import (
	"capture/internal/api"
	"capture/internal/api/message"
	"capture/internal/app"
	"fmt"
	"io"
	"sync"
	"time"
)

const NoTrust int = 0

type ThresholdsTrust struct {
	opts  ThresholdsOptions
	api   api.API
	nodes *message.Trust
	// TODO: make a proper data structure
	stats map[string][]trust
	slock sync.Mutex
}

type trust struct {
	Value     int
	Timestamp time.Time
}

func NewThresholdsTrust(opts *ThresholdsOptions, api api.API) *ThresholdsTrust {
	app.Log.Trace("trust (thresholds): new")
	tr := new(ThresholdsTrust)

	tr.opts = opts.Copy()
	tr.api = api
	tr.nodes = message.NewTrust()
	tr.stats = map[string][]trust{}

	app.Stats["trust.thresholds"] = tr

	return tr
}

func (tr *ThresholdsTrust) Reward(node string) {
	app.Log.Info("trust (thresholds): reward " + node)
	var val int

	tr.slock.Lock()
	defer tr.slock.Unlock()

	oldval, _ := tr.nodes.Get(node)

	if n, ok := tr.nodes.Get(node); !ok {
		val = tr.opts.NeutralTrust + tr.opts.Reward
	} else {
		maxedOut := tr.opts.MaxTrust - tr.opts.Reward
		if n <= maxedOut {
			val = n + tr.opts.Reward
		}
	}

	app.Log.Infof("trust (thresholds): %s %d -> %d", node, oldval, val)
	tr.nodes.Set(node, val)

	if _, ok := tr.stats[node]; !ok {
		tr.stats[node] = make([]trust, 0)
	}

	tr.stats[node] = append(tr.stats[node], trust{val, time.Now()})
}

func (tr *ThresholdsTrust) Penalize(node, reason string) {
	app.Log.Info("trust (thresholds): penalize " + node)
	var val int

	tr.slock.Lock()
	defer tr.slock.Unlock()

	oldval, _ := tr.nodes.Get(node)

	if n, ok := tr.nodes.Get(node); !ok {
		val = tr.opts.NeutralTrust - tr.opts.Penalty
	} else {
		maxedOut := NoTrust + tr.opts.Penalty
		if n >= maxedOut {
			val = n - tr.opts.Penalty
		}
	}

	app.Log.Infof("trust (thresholds): %s %d -> %d", node, oldval, val)
	tr.nodes.Set(node, val)

	if val < tr.opts.LowTrust {
		if _, ok := tr.stats[node]; !ok {
			tr.stats[node] = make([]trust, 0)
		}

		tr.stats[node] = append(tr.stats[node], trust{val, time.Now()})
		tr.api.BlockNode(node, reason)
	}
}

func (tr *ThresholdsTrust) Update() {
	app.Log.Trace("trust (thresholds): update")

	tr.api.SendTrustValues(tr.nodes.ToJson())
}

func (tr *ThresholdsTrust) Report(w io.Writer) {
	app.Log.Trace("trust (thresholds): report stats")
	tr.slock.Lock()
	defer tr.slock.Unlock()

	for node, trustList := range tr.stats {
		for _, tr := range trustList {
			_, err := w.Write([]byte(fmt.Sprintf("%s, %d, %v\n", node, tr.Value, tr.Timestamp)))
			if err != nil {
				app.Log.Warning("trust (thresholds): stats: report failed")
			}
		}
	}
}
