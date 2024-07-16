package cmd

import (
	"davidshefcik/tesla-cli/cmd/handlers"

	"github.com/spf13/cobra"
)

var clientIdArg string
var cfgFile string

var SetupCmd = &cobra.Command{
	Use: "setup",
	Short: "Setup Tesla OAuth application configuration for CLI.",
	ValidArgs: []string{"client-id", "c"},
	Run: func(cmd *cobra.Command, args []string) {
		handlers.SetupHandler(cmd, args, &clientIdArg)
	},
}

func init() {
	SetupCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	SetupCmd.Flags().StringVarP(&clientIdArg, "client-id", "c", "", "Client ID for Tesla OAuth application.")
}