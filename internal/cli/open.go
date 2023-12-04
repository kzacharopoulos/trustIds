package cli

import "github.com/spf13/cobra"

var openCmd = &cobra.Command{
	Use:       "open",
	Short:     "Opens a packet source to sniff from",
	ValidArgs: []string{"file", "interface"},
	Args:      cobra.OnlyValidArgs,
	Aliases:   []string{"o"},
}
