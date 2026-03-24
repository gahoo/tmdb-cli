package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage TMDB CLI configuration",
}

var setAuthCmd = &cobra.Command{
	Use:   "set-auth [TOKEN]",
	Short: "Set TMDB API Auth Token (v4 Read Access Token or v3 API Key)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]
		viper.Set("token", token)
		
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			return
		}

		configPath := homeDir + "/.tmdb.json"
		err = viper.WriteConfigAs(configPath)
		if err != nil {
			fmt.Println("Error writing config:", err)
			return
		}
		fmt.Println("Auth token saved successfully to", configPath)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setAuthCmd)
}
