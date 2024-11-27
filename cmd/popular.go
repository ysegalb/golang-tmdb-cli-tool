package cmd

import (
	"tmdb-cli-tool/internal/api"
	"tmdb-cli-tool/internal/config"

	"github.com/spf13/cobra"
)

// popularCmd represents the popular command
var popularCmd = &cobra.Command{
	Use:   "popular",
	Short: "Get popular movies",
	Long:  `Fetch and display a list of currently popular movies.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		client := api.NewTMDBClient(cfg.APIKey)
		movies, err := client.GetMovies("/popular")
		if err != nil {
			return err
		}

		api.DisplayMovies(movies.Results)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(popularCmd)
}
