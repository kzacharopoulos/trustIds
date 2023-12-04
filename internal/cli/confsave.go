package cli

import (
	"capture/internal/app"
	"capture/internal/conf"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var confSaveCmd = &cobra.Command{
	Use:   "save",
	Short: "Writes effective configuration on disk",
	Run: func(cmd *cobra.Command, args []string) {
		app.Log.Trace("cli: conf save")

		filename := conf.Save(app.Name)
		fmt.Println("Configuration saved at: " + filename)
		os.Exit(0)
	},
}
