package app

import (
	"fmt"
	"io"
	"os"
)

type Statistics interface {
	Report(w io.Writer)
}

func PrintStats() {
	if len(Stats) == 0 {
		Log.Info("app: no stats providers")
		return
	}

	fmt.Println("Stats")

	for key, val := range Stats {
		fmt.Println("\n" + key)
		val.Report(os.Stdout)
	}
}

func WriteStats() {
	if len(Stats) == 0 {
		Log.Info("app: no stats providers")
		return
	}

	for key, val := range Stats {
		f, err := os.Create(key + ".csv")
		if err != nil {
			Log.Warning("stats: " + key + ": failed to open file")
			continue
		}

		val.Report(f)
		f.Close()
	}
}
