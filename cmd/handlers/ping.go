package handlers

import (
	"fmt"

	"github.com/spf13/cobra"
)

func PingHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Pong!")
}