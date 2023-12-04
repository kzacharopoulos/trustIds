package cli

import "github.com/spf13/cobra"

var confCmd = &cobra.Command{
	Use:       "conf",
	Short:     "Shows effective configuration",
	ValidArgs: []string{"print", "save"},
	Args:      cobra.OnlyValidArgs,
	Aliases:   []string{"c"},
}
