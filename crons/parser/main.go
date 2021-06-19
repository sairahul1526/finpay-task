package main

import (
	"os"
	"strings"
	CONFIG "video-parser/config"
	CRON "video-parser/cron"
	DB "video-parser/database"
)

func main() {

	CONFIG.LoadConfig()
	DB.ConnectElastisearch()
	DB.ConnectRedis()

	if strings.EqualFold(os.Getenv("lambda"), "1") {
		CRON.Start(true)
	} else {
		CRON.Start(false)
	}
}
