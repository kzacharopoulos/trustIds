package api

import (
	"capture/internal/app"
	"capture/internal/conf"
	"capture/internal/logger"
	"testing"
)

func init() {
	app.Log = logger.NewNone()
	app.Cfg = conf.New(app.Name)
}

func TestAvailableBackens(t *testing.T) {
	AvailableBackends()
}

func TestBackendFromString(t *testing.T) {
	test := struct {
		input []string
		want  []Backend
	}{
		input: []string{"amqp", "AMQP", "Amqp", "amgp", "log", "none", "what?"},
		want:  []Backend{AmqpBackend, NoneBackend, NoneBackend, NoneBackend, LogBackend, NoneBackend, NoneBackend},
	}

	for i := range test.input {
		if backendFromString(test.input[i]) != test.want[i] {
			t.Errorf("got %q want %q", test.input[i], test.want[i])
		}
	}
}
