package api

import (
	"capture/internal/app"
	"capture/internal/conf"
	"capture/internal/logger"
	"reflect"
	"testing"
)

func init() {
	app.Log = logger.NewNone()
	app.Cfg = conf.NewDefaults()
}

func TestNew(t *testing.T) {
	test := struct {
		input []string
		want  []string
	}{
		input: []string{"log", "amqp", "none"},
		want:  []string{"*backend.Log", "*backend.Amqp", "*backend.Log"},
	}

	for i := range test.input {
		// TODO: mock amqp
		iType := reflect.TypeOf(New(test.input[i]))

		if iType.String() != test.want[i] {
			t.Errorf("got %q want %q", iType.String(), test.want[i])
		}
	}
}
