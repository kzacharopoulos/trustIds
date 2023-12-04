package cli

import (
	"capture/internal/app"
	"capture/internal/conf"
	"capture/internal/kit"
	"capture/internal/kit/backend"
	"capture/internal/sniff"
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

var openFileCmd = &cobra.Command{
	Use:   "file",
	Short: "Opens a pcap file to sniff from",
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		app.Log.Trace("cli: open file (pre)")
		parseBoolFlag(cmd, "replay", "pcap_replay")
		parseStringFlag(cmd, "filter", "pcap_filter")
		parseBoolFlag(cmd, "stats", "app_show_stats")
		app.Cfg = conf.Update()
	},
	Run: func(cmd *cobra.Command, args []string) {
		app.Log.Trace("cli: open file")
		var wg sync.WaitGroup

		src := sniff.OpenFilename(args[0], app.Cfg.PcapFilter)

		opts := backend.NewOptions()
		k := kit.New(opts, src)
		if k == nil {
			app.DieOnErr(fmt.Errorf("kit unavailable"))
		}

		k.Run(&wg)
		wg.Wait()
	},
}
