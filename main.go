package main

import (
	"davidshefcik/tesla-cli/cmd"
	"davidshefcik/tesla-cli/cmd/config"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	// Init config
	defaultConfig := &config.Config{
		ClientId: "",
	}
	
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.teslacli")
	viper.AddConfigPath("./")
	viper.SetConfigType("json")

	viper.SetDefault("client_id", defaultConfig.ClientId)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfigAs("config.json"); err != nil {
				fmt.Println("failed to write viper config", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("failed to read viper config", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("config file loaded successfully")
	}

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("failed to read viper config", err)
		os.Exit(1)
	}
	
	// Execute root command
	cmd.Execute()
}