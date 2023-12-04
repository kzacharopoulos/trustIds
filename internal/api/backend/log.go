package backend

import (
	"capture/internal/api/message"
	"capture/internal/app"
)

type Log struct{}

func NewLog() *Log {
	return new(Log)
}

func (a Log) BlockNode(ip, reason string) {
	app.Log.Trace("api: log: block node")

	m := message.NewWarn(ip, reason)
	payload := m.ToJson()
	app.Log.Infof("api: log: payload: \n%s", string(payload))
}

func (a Log) DropClient(ip string) {
	app.Log.Trace("api: log: drop client")

	m := message.NewMtdWarn(ip)
	payload := m.ToJson()
	app.Log.Infof("api: log: payload: \n%s", string(payload))
}

func (a Log) SendTrustValues(payload []byte) {
	app.Log.Trace("api: log: send trust values")

	app.Log.Infof("api: log: payload: \n%s", string(payload))
}
