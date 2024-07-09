package cmd

import (
	"davidshefcik/tesla-cli/cmd/handlers"

	"github.com/spf13/cobra"
)

var PingCmd = &cobra.Command{
	Use: "ping",
	Short: "Pong?",
	Run: handlers.PingHandler,
}