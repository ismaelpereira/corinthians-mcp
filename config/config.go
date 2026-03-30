package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MeuTimaoFeedURL       string
	CentralDoTimaoFeedURL string
	MatchesAPIURL         string
	MatchesAPIKey         string
	HistoryFilePath       string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found, using system env")
	}

	config := &Config{
		MeuTimaoFeedURL:       os.Getenv("MEU_TIMAO_FEED_URL"),
		CentralDoTimaoFeedURL: os.Getenv("CENTRAL_DO_TIMAO_FEED_URL"),
		MatchesAPIURL:         os.Getenv("MATCHES_API_URL"),
		MatchesAPIKey:         os.Getenv("MATCHES_API_KEY"),
		HistoryFilePath:       os.Getenv("HISTORY_FILE_PATH"),
	}

	// validate(config)
	return config
}

func validate(c *Config) {
	if c.MeuTimaoFeedURL == "" {
		log.Fatal("MEU_TIMAO_FEED_URL is required")
	}
	if c.CentralDoTimaoFeedURL == "" {
		log.Fatal("CENTRAL_DO_TIMAO_FEED_URL is required")
	}
	if c.MatchesAPIURL == "" {
		log.Fatal("MATCHES_API_URL is required")
	}
	if c.MatchesAPIKey == "" {
		log.Fatal("MATCHES_API_KEY is required")
	}
	if c.HistoryFilePath == "" {
		log.Fatal("HISTORY_FILE_PATH is required")
	}
}
