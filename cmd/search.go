package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/gahoolee/tmdb-cli/api"
	"github.com/gahoolee/tmdb-cli/formatter"
)

var searchType string

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for movies, tv shows, or multi",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		token := viper.GetString("token")
		if token == "" {
			fmt.Println("Error: API Token is not set. Use 'tmdb config set-auth <TOKEN>'")
			return
		}

		client := api.NewClient(token)
		results, err := client.Search(query, searchType)
		if err != nil {
			fmt.Println("Search error:", err)
			return
		}

		err = formatter.OutputResult(results, outputFormat, searchType)
		if err != nil {
			fmt.Println("Formatting error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&searchType, "type", "t", "movie", "Search type: movie, tv, multi")
}
