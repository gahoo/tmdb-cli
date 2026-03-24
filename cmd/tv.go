package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/gahoolee/tmdb-cli/api"
	"github.com/gahoolee/tmdb-cli/formatter"
)

var seasonNum int

var tvCmd = &cobra.Command{
	Use:   "tv [ID]",
	Short: "Get details for a TV series by its TMDB ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid tv ID. It must be a number.")
			return
		}

		token := viper.GetString("token")
		if token == "" {
			fmt.Println("Error: API Token is not set. Use 'tmdb config set-auth <TOKEN>'")
			return
		}

		client := api.NewClient(token)

		if seasonNum > 0 {
			result, err := client.GetTVSeasonDetails(id, seasonNum)
			if err != nil {
				fmt.Println("Error fetching TV season:", err)
				return
			}
			err = formatter.OutputResult(result, outputFormat, "season")
			if err != nil {
				fmt.Println("Formatting error:", err)
			}
			if outputFormat == "nfo" && downloadPoster {
				destPath := fmt.Sprintf("season%02d-poster.jpg", result.SeasonNumber)
				err = api.DownloadImage(result.PosterPath, destPath)
				if err != nil {
					fmt.Println("Warning: Failed to download poster:", err)
				} else {
					fmt.Printf("Poster downloaded to %s\n", destPath)
				}
			}
		} else {
			result, err := client.GetTVDetails(id)
			if err != nil {
				fmt.Println("Error fetching TV series:", err)
				return
			}
			err = formatter.OutputResult(result, outputFormat, "tvshow")
			if err != nil {
				fmt.Println("Formatting error:", err)
			}
			if outputFormat == "nfo" && downloadPoster {
				err = api.DownloadImage(result.PosterPath, "poster.jpg")
				if err != nil {
					fmt.Println("Warning: Failed to download poster:", err)
				} else {
					fmt.Println("Poster downloaded to poster.jpg")
				}
			}
		}
	},
}

func init() {
	tvCmd.Flags().IntVarP(&seasonNum, "season", "s", 0, "Fetch specific season details by season number")
	rootCmd.AddCommand(tvCmd)
}
