package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL   = "https://api.themoviedb.org/3/movie"
	userAgent = "TMDB-CLI-Tool/1.0"
)

// TMDBClient handles API communication with TMDB
type TMDBClient struct {
	apiKey string
	client *http.Client
}

// NewTMDBClient creates a new TMDB API client
func NewTMDBClient(apiKey string) *TMDBClient {
	return &TMDBClient{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

// GetMovies fetches movies from the specified endpoint
func (c *TMDBClient) GetMovies(endpoint string) (*MovieResponse, error) {
	url := fmt.Sprintf("%s%s", baseURL, endpoint)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var movieResp MovieResponse
	if err := json.NewDecoder(resp.Body).Decode(&movieResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &movieResp, nil
}
