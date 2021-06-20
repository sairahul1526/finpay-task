package main

import (
	API "video-api/api"
	CONFIG "video-api/config"
	CONSTANT "video-api/constant"
	DB "video-api/database"

	_ "video-api/docs"
)

// @title Finpay Task - Video API
// @version 1.0
// @description This is an api to show videos
func main() {

	CONFIG.LoadConfig()
	DB.ConnectElastisearch()
	DB.ConnectRedis()

	DB.AddIndex(CONSTANT.VideosIndex, CONSTANT.VideosIndexMapping) // add videos index if not available

	API.StartServer()
}
