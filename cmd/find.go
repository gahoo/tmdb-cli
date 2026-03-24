package cmd

import (
	"fmt"

	"github.com/gahoolee/tmdb-cli/api"
	"github.com/gahoolee/tmdb-cli/formatter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var externalSource string

var findCmd = &cobra.Command{
	Use:   "find [external_id]",
	Short: "Find TMDB items by external ID (e.g. imdb_id)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		externalID := args[0]
		if externalSource == "" {
			fmt.Println("Error: --source flag is required (e.g. imdb_id, tvdb_id)")
			return
		}

		token := viper.GetString("token")
		if token == "" {
			fmt.Println("Error: API Token is not set. Use 'tmdb config set-auth <TOKEN>'")
			return
		}

		client := api.NewClient(token)
		result, err := client.FindByExternalID(externalID, externalSource)
		if err != nil {
			fmt.Println("Error finding item:", err)
			return
		}

		err = formatter.OutputResultToFileOrStdout(outputFile, result, outputFormat, "find")
		if err != nil {
			fmt.Println("Formatting error:", err)
		}
	},
}

func init() {
	findCmd.Flags().StringVar(&externalSource, "source", "", "External source (e.g. imdb_id, freebase_mid, freebase_id, tvdb_id, tvrage_id, facebook_id, twitter_id, instagram_id)")
	rootCmd.AddCommand(findCmd)
}
