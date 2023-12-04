package cli

import (
	"capture/internal/app"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var confPrintCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints effective configuration",
	Run: func(cmd *cobra.Command, args []string) {
		app.Log.Trace("cli: conf print")

		fmt.Println(viper.AllSettings())
		os.Exit(0)
	},
}
