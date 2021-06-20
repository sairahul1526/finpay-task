package main

import (
	"os"
	"strings"
	"time"
	CONFIG "video-parser/config"
	CONSTANT "video-parser/constant"
	CRON "video-parser/cron"
	DB "video-parser/database"
	UTIL "video-parser/util"
)

func main() {

	CONFIG.LoadConfig()
	DB.ConnectElastisearch()
	DB.ConnectRedis()

	DB.SetRedisValue(CONSTANT.LastPublishedAtKey, UTIL.GetCurrentTime().Format(time.RFC3339)) // init published date

	if strings.EqualFold(os.Getenv("lambda"), "1") {
		CRON.Start(true)
	} else {
		CRON.Start(false)
	}
}
