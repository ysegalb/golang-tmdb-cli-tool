package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmdb-app",
	Short: "A CLI tool for fetching movie information from TMDB",
	Long: `TMDB CLI Tool allows you to fetch and display movie information from The Movie Database (TMDB).
You can get information about now playing, popular, top-rated, and upcoming movies.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
