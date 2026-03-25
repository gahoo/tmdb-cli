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
var episodeNum int

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

		lang := outputLanguage
		if lang == "" {
			lang = viper.GetString("language")
		}

		client := api.NewClient(token, lang)

		seasonSet := cmd.Flags().Changed("season")
		episodeSet := cmd.Flags().Changed("episode")

		if seasonSet && episodeSet {
			// Get episode details
			result, err := client.GetTVEpisode(id, seasonNum, episodeNum)
			if err != nil {
				fmt.Println("Error fetching TV episode:", err)
				return
			}

			err = formatter.OutputResultToFileOrStdout(outputFile, result, outputFormat, "episode", outputFields)
			if err != nil {
				fmt.Println("Formatting error:", err)
			}

			if outputFormat == "nfo" && downloadPoster {
				err = client.DownloadImage(result.StillPath, "thumb.jpg")
				if err != nil {
					fmt.Println("Warning: Failed to download thumbnail:", err)
				} else {
					fmt.Println("Thumbnail downloaded to thumb.jpg")
				}
			}
		} else if seasonSet {
			// Get season details
			result, err := client.GetTVSeasonDetails(id, seasonNum)
			if err != nil {
				fmt.Println("Error fetching TV season:", err)
				return
			}

			err = formatter.OutputResultToFileOrStdout(outputFile, result, outputFormat, "season", outputFields)
			if err != nil {
				fmt.Println("Formatting error:", err)
			}

			if outputFormat == "nfo" && downloadPoster {
				destPath := fmt.Sprintf("season%02d-poster.jpg", result.SeasonNumber)
				err = client.DownloadImage(result.PosterPath, destPath)
				if err != nil {
					fmt.Println("Warning: Failed to download poster:", err)
				} else {
					fmt.Printf("Poster downloaded to %s\n", destPath)
				}
			}
		} else {
			// Get TV series details
			result, err := client.GetTVDetails(id)
			if err != nil {
				fmt.Println("Error fetching TV series:", err)
				return
			}

			err = formatter.OutputResultToFileOrStdout(outputFile, result, outputFormat, "tvshow", outputFields)
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
		}
	},
}

func init() {
	tvCmd.Flags().IntVarP(&seasonNum, "season", "s", 0, "Fetch specific season details by season number")
	tvCmd.Flags().IntVarP(&episodeNum, "episode", "e", 0, "Fetch specific episode details by episode number (requires --season)")
	rootCmd.AddCommand(tvCmd)
}
