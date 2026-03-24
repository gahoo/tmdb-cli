package cmd

import (
	"encoding/json"
	"fmt"
	"os"
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

			if outputFile != "" {
				data, _ := json.MarshalIndent(result, "", "  ")
				_ = os.WriteFile(outputFile, data, 0644)
				fmt.Printf("Results exported to %s\n", outputFile)
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
			result, err := client.GetTVDetails(id)
			if err != nil {
				fmt.Println("Error fetching TV series:", err)
				return
			}
			err = formatter.OutputResult(result, outputFormat, "tvshow")
			if err != nil {
				fmt.Println("Formatting error:", err)
			}

			if outputFile != "" {
				data, _ := json.MarshalIndent(result, "", "  ")
				_ = os.WriteFile(outputFile, data, 0644)
				fmt.Printf("Results exported to %s\n", outputFile)
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

var episodeCmd = &cobra.Command{
	Use:   "episode [tvID] [seasonNum] [episodeNum]",
	Short: "Get details for a TV episode",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		tvID, _ := strconv.Atoi(args[0])
		seasonNum, _ := strconv.Atoi(args[1])
		epNum, _ := strconv.Atoi(args[2])

		token := viper.GetString("token")
		if token == "" {
			fmt.Println("Error: API Token is not set. Use 'tmdb config set-auth <TOKEN>'")
			return
		}

		client := api.NewClient(token)
		result, err := client.GetTVEpisode(tvID, seasonNum, epNum)
		if err != nil {
			fmt.Println("Error fetching TV episode:", err)
			return
		}

		err = formatter.OutputResult(result, outputFormat, "episode")
		if err != nil {
			fmt.Println("Formatting error:", err)
		}

		if outputFile != "" {
			data, _ := json.MarshalIndent(result, "", "  ")
			_ = os.WriteFile(outputFile, data, 0644)
			fmt.Printf("Results exported to %s\n", outputFile)
		}

		if outputFormat == "nfo" && downloadPoster {
			err = client.DownloadImage(result.StillPath, "thumb.jpg")
			if err != nil {
				fmt.Println("Warning: Failed to download thumbnail:", err)
			} else {
				fmt.Println("Thumbnail downloaded to thumb.jpg")
			}
		}
	},
}

func init() {
	tvCmd.Flags().IntVarP(&seasonNum, "season", "s", 0, "Fetch specific season details by season number")
	tvCmd.AddCommand(episodeCmd)
	rootCmd.AddCommand(tvCmd)
}
