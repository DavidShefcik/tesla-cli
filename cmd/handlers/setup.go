package handlers

import (
	"errors"
	"fmt"
	"os"

	"davidshefcik/tesla-cli/utils"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


func validateStringInput(input string) error {
	utils.RemoveStringWhitespace(&input)

	if len(input) == 0 {
		return errors.New("input cannot be empty")
	}

	return nil
}

func getClientId(clientId *string) (string, error) {
	result := *clientId

	if *clientId == "" {
		clientIDPrompt := &promptui.Prompt{
			Label: "Client ID",
			Validate: validateStringInput,
		}
		
		fmt.Println("Please input your Tesla OAuth application credentials.")

		runResult, err := clientIDPrompt.Run()

		if err != nil {
			return "", err
		}

		utils.RemoveStringWhitespace(&runResult)

		result = runResult
	} else {
		if err := validateStringInput(*clientId); err != nil {
			return "", err
		}
	}

	return result, nil
}

func SaveClientId(clientId string) error {
	viper.Set("client_id", clientId)
	err := viper.WriteConfig()

	return err
}

func SetupHandler(cmd *cobra.Command, args []string, clientIdArg *string) {
	clientId, err := getClientId(clientIdArg)

	if err != nil {
		fmt.Println("failed to get client id:", err)
		os.Exit(1)
	}

	err = SaveClientId(clientId)

	if err != nil {
		fmt.Println("failed to save client id:", err)
		os.Exit(1)
	}

	fmt.Println(clientId)
}

func init() {

}