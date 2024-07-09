package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "tesla",
	Short: "Tesla CLI is a CLI application to control and view information about your Tesla.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	Version: "0.0.0", // TODO: Update using release please
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(PingCmd)
}