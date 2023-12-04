package app

import (
	"capture/internal/conf"
	"capture/internal/logger"
	"os"
)

const (
	Name string = "capture"
)

var (
	Log   logger.Logger         = nil
	Cfg   *conf.Configuration   = nil
	Stats map[string]Statistics = map[string]Statistics{}
)

func DieOnErr(err error) {
	if err == nil {
		return
	}

	Log.Crit(err.Error())
	os.Exit(1)
}

func DieOnNil(val any, msg string) {
	if val != nil {
		return
	}

	Log.Crit(msg)
	os.Exit(1)
}
