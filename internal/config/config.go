package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Config holds application configuration
type Config struct {
	APIKey string
}

// LoadConfig loads configuration from environment and .env file
func LoadConfig() (*Config, error) {
	if err := loadEnvFile(); err != nil {
		fmt.Printf("Warning: Could not load .env file: %v\n", err)
	}

	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("TMDB_API_KEY environment variable is not set")
	}

	return &Config{
		APIKey: apiKey,
	}, nil
}

// loadEnvFile loads environment variables from .env file
func loadEnvFile() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				os.Setenv(key, value)
			}
		}
	}
	return scanner.Err()
}
