package cmd

import (
	"tmdb-cli-tool/internal/api"
	"tmdb-cli-tool/internal/config"

	"github.com/spf13/cobra"
)

// playingCmd represents the playing command
var playingCmd = &cobra.Command{
	Use:   "playing",
	Short: "Get now playing movies",
	Long:  `Fetch and display a list of movies that are currently playing in theaters.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		client := api.NewTMDBClient(cfg.APIKey)
		movies, err := client.GetMovies("/now_playing")
		if err != nil {
			return err
		}

		api.DisplayMovies(movies.Results)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(playingCmd)
}
