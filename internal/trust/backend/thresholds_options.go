package backend

import "capture/internal/app"

type ThresholdsOptions struct {
	LowTrust     int
	NeutralTrust int
	MaxTrust     int
	Penalty      int
	Reward       int
}

func NewThresholdsOptions() *ThresholdsOptions {
	app.Log.Trace("trust (thresholds): new options")
	return &ThresholdsOptions{
		LowTrust:     app.Cfg.TrustThresholdsLow,
		NeutralTrust: app.Cfg.TrustThresholdsNeutral,
		MaxTrust:     app.Cfg.TrustThresholdsMax,
		Penalty:      app.Cfg.TrustThresholdsPenalty,
		Reward:       app.Cfg.TrustThresholdsReward,
	}
}

func (o ThresholdsOptions) Copy() ThresholdsOptions {
	app.Log.Trace("trust (thresholds): copy options")
	return ThresholdsOptions{
		LowTrust:     o.LowTrust,
		NeutralTrust: o.NeutralTrust,
		MaxTrust:     o.MaxTrust,
		Penalty:      o.Penalty,
		Reward:       o.Reward,
	}
}
