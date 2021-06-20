package cron

import (
	"fmt"
	"time"
	CONSTANT "video-parser/constant"
	DB "video-parser/database"
	MODEL "video-parser/model"
	UTIL "video-parser/util"
)

func parseVideos() {
	// get first element from api keys
	apiKey, err := DB.GetFromRedisSortedSets(CONSTANT.APIKeysRedisKey, "-inf", "+inf", 0, 1)
	if err != nil {
		fmt.Println("parseVideos", err)
		return
	}

	if len(apiKey) == 0 {
		fmt.Println("parseVideos", "there are no api keys")
		return
	}

	// get last video published datetime to get videos from that time
	lastPublishedAt, err := DB.GetRedisValue(CONSTANT.LastPublishedAtKey)
	if err != nil {
		fmt.Println("parseVideos", err)
		return
	}

	// get videos from api
	searchResponse, err := UTIL.GetYoutubeVideoResults(apiKey[0], lastPublishedAt)
	if err != nil {
		fmt.Println("parseVideos", err)

		// since we got error, that means current apikey credits got used, so we update its score to next day
		DB.AddToRedisSortedSets(CONSTANT.APIKeysRedisKey, float64(UTIL.GetCurrentTime().AddDate(0, 0, 1).Unix()), apiKey[0])

		// we now get least score apikey and parse youtube api again
		parseVideos()
		return
	}
	if len(searchResponse.Items) == 0 {
		return
	}

	// store last published date of latest video for future api calls
	lastPublishedAt = searchResponse.Items[0].Snippet.PublishedAt

	videos := []MODEL.Video{}
	var publishedAt time.Time

	// convert into video models
	for _, item := range searchResponse.Items {
		publishedAt, _ = time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		videos = append(videos, MODEL.Video{
			VideoID:      item.ID.VideoID,
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			ThumbnailURL: item.Snippet.Thumbnails.Default.URL,
			PublishedAt:  publishedAt,
		})
	}

	// add videos to database
	err = MODEL.InsertVideos(videos)
	if err != nil {

		fmt.Println("parseVideos", err)
		return
	}
	// update last video published datetime
	DB.SetRedisValue(CONSTANT.LastPublishedAtKey, lastPublishedAt)

}
