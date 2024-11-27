package cmd

import (
	"tmdb-cli-tool/internal/api"
	"tmdb-cli-tool/internal/config"

	"github.com/spf13/cobra"
)

// topCmd represents the top command
var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Get top rated movies",
	Long:  `Fetch and display a list of top-rated movies on TMDB.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		client := api.NewTMDBClient(cfg.APIKey)
		movies, err := client.GetMovies("/top_rated")
		if err != nil {
			return err
		}

		api.DisplayMovies(movies.Results)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(topCmd)
}
