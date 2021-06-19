package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// LoadConfig - load .env file from given path for local, else will be getting from env var
func LoadConfig() {
	// load .env file from given path for local, else will be getting from env var
	if !strings.EqualFold(os.Getenv("prod"), "true") {
		err := godotenv.Load(".test-env")
		if err != nil {
			panic("Error loading .env file")
		}
	}

	ElastisearchConfig = os.Getenv("ELASTI_SEARCH_CONFIG")
	Log, _ = strconv.ParseBool(os.Getenv("LOG"))
	Migrate, _ = strconv.ParseBool(os.Getenv("MIGRATE"))
	Port = os.Getenv("PORT")
}
