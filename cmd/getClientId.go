package cmd

import (
	"davidshefcik/tesla-cli/cmd/handlers"

	"github.com/spf13/cobra"
)

var GetClientIdCmd = &cobra.Command{
	Use: "get-client-id",
	Short: "Prints out the client ID.",
	ValidArgs: []string{"get-client-id", "gc"},
	Run: handlers.GetClientIdHandler,
}