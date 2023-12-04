package cli

import (
	"capture/internal/api"
	"capture/internal/app"
	"capture/internal/kit"
	"capture/internal/trust"
	"os"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "Shows release info and exits",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		app.Log.Trace("cli: info")

		app.ShowBuildInfo()
		api.AvailableBackends()
		trust.AvailableBackends()
		kit.AvailableKits()

		os.Exit(0)
	},
}
