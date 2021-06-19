package cron

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	CONFIG "video-parser/config"
	CONSTANT "video-parser/constant"
	DB "video-parser/database"
	MODEL "video-parser/model"
)

func parseVideos() {
	// get first element from api keys
	apiKey, err := DB.GetPopFirstInRedisLists(CONSTANT.APIKeysRedisKey)
	if err != nil {
		fmt.Println("parseVideos", err)
		return
	}

	if len(apiKey) > 0 {
		// get last video published datetime to get videos from it
		lastPublishedAt, err := DB.GetRedisValue("last_published_at")
		if err == nil {
			// get videos from api
			searchResponse := getYoutubeVideoResults(apiKey, lastPublishedAt)
			if len(searchResponse.Items) > 0 {
				lastPublishedAt = searchResponse.Items[0].Snippet.PublishedAt
			}
			videos := []MODEL.Video{}
			var publishedAt time.Time
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
			// insert videos into elasti search
			err = MODEL.InsertVideos(videos)
			if err == nil {
				// update last video published datetime
				DB.SetRedisValue("last_published_at", lastPublishedAt)

				// add api key to end of list
				DB.AddToRedisLists(CONSTANT.APIKeysRedisKey, apiKey)
			} else {
				fmt.Println("parseVideos", err)
			}
		} else {
			fmt.Println("parseVideos", err)
		}
	} else {
		fmt.Println("parseVideos", "there are no api keys")
	}
}

func getYoutubeVideoResults(apiKey, publishedAfter string) MODEL.SearchResponse {
	fmt.Println(CONSTANT.YoutubeSearchURL + "?part=snippet&maxResults=100&order=date&q=" + CONFIG.SearchQueryYoutube + "&key=" + apiKey + "&type=video&publishedAfter=" + publishedAfter)
	// get request from youtube
	res, err := http.Get(CONSTANT.YoutubeSearchURL + "?part=snippet&maxResults=100&order=date&q=" + CONFIG.SearchQueryYoutube + "&key=" + apiKey + "&type=video&publishedAfter=" + publishedAfter)
	if err != nil {
		fmt.Println(err)
		return MODEL.SearchResponse{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return MODEL.SearchResponse{}
	}

	// parse response
	searchResponse := MODEL.SearchResponse{}
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		fmt.Println(err)
		return MODEL.SearchResponse{}
	}

	return searchResponse
}
