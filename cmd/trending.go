package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/gahoolee/tmdb-cli/api"
	"github.com/gahoolee/tmdb-cli/formatter"
)

var trendingType string
var trendingTime string

var trendingCmd = &cobra.Command{
	Use:   "trending",
	Short: "Get trending items for movies, tv, and people",
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")
		if token == "" {
			fmt.Println("Error: API Token is not set. Use 'tmdb config set-auth <TOKEN>'")
			return
		}

		client := api.NewClient(token)
		result, err := client.GetTrending(trendingType, trendingTime)
		if err != nil {
			fmt.Println("Error fetching trending:", err)
			return
		}

		err = formatter.OutputResult(result, outputFormat, "trending")
		if err != nil {
			fmt.Println("Formatting error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(trendingCmd)
	trendingCmd.Flags().StringVarP(&trendingType, "type", "t", "all", "Media type: all, movie, tv, person")
	trendingCmd.Flags().StringVarP(&trendingTime, "time", "w", "day", "Time window: day, week")
}
