# TMDB CLI Tool

The url for this repository is https://github.com/ysegalb/golang-tmdb-cli-tool
It solves the roadmap example from https://roadmap.sh/projects/tmdb-cli

Submission link: https://roadmap.sh/projects/tmdb-cli/solutions?u=6746e4845039431075f1db70

A command-line interface tool written in Go that fetches and displays movie information from The Movie Database (TMDB) API.

## Features

- Fetch and display Now Playing movies
- Fetch and display Popular movies
- Fetch and display Top Rated movies
- Fetch and display Upcoming movies

## Project Structure

```
tmdb-cli-tool/
├── cmd/                    # Command-line interface components
│   ├── root.go            # Root command and base CLI setup
│   ├── playing.go         # Now playing movies command
│   ├── popular.go         # Popular movies command
│   ├── top.go            # Top rated movies command
│   └── upcoming.go        # Upcoming movies command
├── internal/              # Internal packages
│   ├── api/              # TMDB API interaction
│   │   ├── client.go     # API client implementation
│   │   ├── models.go     # Data structures for API responses
│   │   └── display.go    # Output formatting functions
│   └── config/           # Configuration management
│       └── config.go     # Environment and config handling
├── main.go               # Application entry point
├── .env                  # Environment configuration
└── README.md            # Project documentation
```

### Package Descriptions

- `cmd/`: Contains all the CLI commands and their implementations
  - `root.go`: Defines the base command and global flags

- `internal/api/`: Handles all TMDB API interactions
  - `client.go`: Implements the HTTP client for API communication
  - `models.go`: Defines data structures for API responses
  - `display.go`: Formats and displays movie information

- `internal/config/`: Manages application configuration
  - `config.go`: Handles environment variables and .env file loading

## Prerequisites

- Go 1.16 or higher
- TMDB API Key (get it from [TMDB website](https://www.themoviedb.org/settings/api))

## Installation

1. Clone this repository
2. Create a `.env` file in the root directory with your TMDB API key:
   ```
   TMDB_API_KEY=your_api_key_here
   ```
   Note: Make sure to use your Bearer token from TMDB, not the API key
3. Build the application:
   ```bash
   go build -o tmdb-app
   ```

## Usage

The application provides several commands to fetch different types of movies:

```bash
# Show help and available commands
./tmdb-app --help

# Show Now Playing movies
./tmdb-app playing

# Show Popular movies
./tmdb-app popular

# Show Top Rated movies
./tmdb-app top

# Show Upcoming movies
./tmdb-app upcoming

# Get help for a specific command
./tmdb-app [command] --help
```

## Error Handling

The application includes error handling for:
- Invalid command-line arguments
- API request failures
- Network issues
- Invalid API responses
- Missing or invalid API key in .env file
