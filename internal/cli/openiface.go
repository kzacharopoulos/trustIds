package cli

import (
	"capture/internal/app"
	"capture/internal/kit"
	"capture/internal/kit/backend"
	"capture/internal/sniff"
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

var openInterfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "Opens a network interface to sniff from",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app.Log.Trace("cli: open interface")
		var wg sync.WaitGroup

		srcOpts := sniff.NewOptions()
		src := sniff.OpenInterface(args[0], srcOpts)

		opts := backend.NewOptions()
		k := kit.New(opts, src)
		if k == nil {
			app.DieOnErr(fmt.Errorf("kit unavailable"))
		}

		k.Run(&wg)
		wg.Wait()
	},
}
