package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadConfig - load .env file from given path for local, else will be getting from env var
func LoadConfig() {
	// load .env file from given path for local, else will be getting from env var
	if len(os.Getenv("lambda")) == 0 {
		err := godotenv.Load(".test-env")
		if err != nil {
			panic("Error loading .env file")
		}
	}

	ElastisearchConfig = os.Getenv("ELASTI_SEARCH_CONFIG")
	RedisAddress = os.Getenv("REDIS_ADDRESS")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	RedisDB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	SearchQueryYoutube = os.Getenv("SEARCH_QUERY_YOUTUBE")
}
