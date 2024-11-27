package api

import "fmt"

// DisplayMovies formats and prints movie information to the console
func DisplayMovies(movies []Movie) {
	fmt.Println("\nMovies List:")
	fmt.Println("============")
	for i, movie := range movies {
		fmt.Printf("\n%d. %s\n", i+1, movie.Title)
		fmt.Printf("   Release Date: %s\n", movie.ReleaseDate)
		fmt.Printf("   Rating: %.1f/10\n", movie.VoteAverage)
		fmt.Printf("   Overview: %s\n", movie.Overview)
		fmt.Println("   ---")
	}
}
