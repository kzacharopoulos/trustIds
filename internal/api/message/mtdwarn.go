package message

import (
	"capture/internal/app"
	"encoding/json"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type MtdWarn struct {
	Type   string
	Action string
	Body   string
	Time   *timestamppb.Timestamp
}

func NewMtdWarn(ip string) *MtdWarn {
	app.Log.Trace("message: mtdwarn: new " + ip)
	return &MtdWarn{
		Type:   "error",
		Action: "drop",
		Body:   ip,
		Time:   timestamppb.Now(),
	}
}

func (w MtdWarn) ToJson() []byte {
	app.Log.Trace("message: mtdwarn: to json")

	payload, err := json.MarshalIndent(&w, "", "    ")

	if err != nil {
		app.Log.Warningf("message: mtdwarn: error marshaling: %s", err)
		app.Log.Debugf("message: mtdwarn: payload:\n%+v", w)
		return []byte{}
	}

	return payload
}
