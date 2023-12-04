package message

import (
	"capture/internal/app"
	"encoding/json"
)

type Warn struct {
	Ip     string
	Action string
	Key    []byte
	Reason string
}

func NewWarn(ip, reason string) *Warn {
	app.Log.Debug("message: warn: new " + ip)
	return &Warn{
		Ip:     ip,
		Action: "block",
		Key:    app.Cfg.KeyBytes(),
		Reason: reason,
	}
}

func (w Warn) ToJson() []byte {
	app.Log.Trace("message: warn: to json")

	payload, err := json.MarshalIndent(&w, "", "    ")

	if err != nil {
		app.Log.Warningf("message: warn: json error marshaling '%v': %s", w, err)
		return []byte{}
	}

	return payload
}
