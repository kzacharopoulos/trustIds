package cli

import (
	"capture/internal/app"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:       app.Name,
	Short:     "Network traffic sniffing tool",
	ValidArgs: []string{"info", "conf", "open"},
	Args:      cobra.OnlyValidArgs,
	Version:   app.Version,
}

func parseBoolFlag(cmd *cobra.Command, name string, viperKey string) {
	val, err := cmd.Flags().GetBool(name)
	if err != nil {
		app.Log.Warningf("cli: parse '%s' flag error: %s", name, err.Error())
		return
	}

	viper.Set(viperKey, val)
}

func parseStringFlag(cmd *cobra.Command, name string, viperKey string) {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		app.Log.Warningf("cli: parse '%s' flag error: %s", name, err.Error())
		return
	}

	viper.Set(viperKey, val)
}

func Execute() {
	app.Log.Trace("cli: execute")

	confCmd.AddCommand(confPrintCmd)
	confCmd.AddCommand(confSaveCmd)

	openCmd.PersistentFlags().String("filter", "", "Sets capture filter")
	openFileCmd.PersistentFlags().Bool("replay", false, "Replays pcap traffic honoring the packets timestamps")
	openFileCmd.PersistentFlags().Bool("stats", false, "Prints all stats after pcap file parsing")
	openCmd.AddCommand(openFileCmd)
	openCmd.AddCommand(openInterfaceCmd)

	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(confCmd)
	rootCmd.AddCommand(openCmd)

	if err := rootCmd.Execute(); err != nil {
		app.Log.Err(err.Error())
		os.Exit(1)
	}
}
