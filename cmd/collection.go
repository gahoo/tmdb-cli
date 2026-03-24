package cmd

import (
	"fmt"
	"strconv"

	"github.com/gahoolee/tmdb-cli/api"
	"github.com/gahoolee/tmdb-cli/formatter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var collectionCmd = &cobra.Command{
	Use:   "collection [ID]",
	Short: "Get details for a collection by its TMDB ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid collection ID. It must be a number.")
			return
		}

		token := viper.GetString("token")
		if token == "" {
			fmt.Println("Error: API Token is not set. Use 'tmdb config set-auth <TOKEN>'")
			return
		}

		client := api.NewClient(token)
		result, err := client.GetCollection(id)
		if err != nil {
			fmt.Println("Error fetching collection:", err)
			return
		}

		err = formatter.OutputResultToFileOrStdout(outputFile, result, outputFormat, "collection")
		if err != nil {
			fmt.Println("Formatting error:", err)
		}

		if outputFormat == "nfo" && downloadPoster {
			err = client.DownloadImage(result.PosterPath, "poster.jpg")
			if err != nil {
				fmt.Println("Warning: Failed to download poster:", err)
			} else {
				fmt.Println("Poster downloaded to poster.jpg")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(collectionCmd)
}
