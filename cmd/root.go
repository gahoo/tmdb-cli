package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	outputFormat   string
	downloadPoster bool
	outputFile     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tmdb",
	Short: "A CLI tool for querying The Movie Database (TMDB) API",
	Long: `TMDB CLI allows you to search for movies, TV shows, and people on The Movie Database.
It supports exporting data to JSON, Markdown, and NFO formats.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
    
	// Global flags
	rootCmd.PersistentFlags().StringVar(&outputFormat, "format", "json", "Output format: json, markdown, nfo, table")
	rootCmd.PersistentFlags().BoolVar(&downloadPoster, "poster", false, "Download the poster image locally when output format is NFO")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "Export JSON output to the specified file")
}

func initConfig() {
	homeDir, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(homeDir)
	}
	viper.SetConfigName(".tmdb")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	
	_ = viper.ReadInConfig() // ignore error if config does not exist
}
