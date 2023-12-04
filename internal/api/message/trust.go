package message

import (
	"capture/internal/app"
	"encoding/json"
	"sync"
)

type Trust struct {
	Values *map[string]int
	vlock  sync.Mutex
}

func NewTrust() *Trust {
	app.Log.Trace("message: trust: new")

	t := new(Trust)
	t.Values = new(map[string]int)
	*t.Values = make(map[string]int, 1)

	return t
}

func (t *Trust) Get(node string) (int, bool) {
	t.vlock.Lock()
	defer t.vlock.Unlock()

	val, ok := (*t.Values)[node]
	return val, ok
}

func (t *Trust) Set(node string, val int) {
	t.vlock.Lock()
	defer t.vlock.Unlock()

	(*t.Values)[node] = val
}

func (t *Trust) ToJson() []byte {
	t.vlock.Lock()
	defer t.vlock.Unlock()

	app.Log.Trace("message: trust: to json")

	payload, err := json.MarshalIndent(&t.Values, "", "    ")

	if err != nil {
		app.Log.Warningf("message: trust: error marshaling: %s", err)
		app.Log.Debugf("message: trust: payload:\n%+v", *t.Values)
		return []byte{}
	}

	return payload
}
