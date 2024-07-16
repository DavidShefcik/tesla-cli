package handlers

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetClientIdHandler(cmd *cobra.Command, args []string) {
	clientId := viper.GetString("client_id")

	if clientId == "" {
		fmt.Println("No client ID set.")
	} else {
		fmt.Println("Client ID:", clientId)
	}
}