package cmd

import (
	"tmdb-cli-tool/internal/api"
	"tmdb-cli-tool/internal/config"

	"github.com/spf13/cobra"
)

// upcomingCmd represents the upcoming command
var upcomingCmd = &cobra.Command{
	Use:   "upcoming",
	Short: "Get upcoming movies",
	Long:  `Fetch and display a list of upcoming movies in theaters.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		client := api.NewTMDBClient(cfg.APIKey)
		movies, err := client.GetMovies("/upcoming")
		if err != nil {
			return err
		}

		api.DisplayMovies(movies.Results)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upcomingCmd)
}
