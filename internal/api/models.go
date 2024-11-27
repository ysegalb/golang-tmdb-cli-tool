package api

// Movie represents a movie from TMDB API
type Movie struct {
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	VoteAverage float64 `json:"vote_average"`
	Overview    string  `json:"overview"`
}

// MovieResponse represents the TMDB API response
type MovieResponse struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}
