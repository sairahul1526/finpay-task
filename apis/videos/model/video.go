package model

import (
	"encoding/json"
	"fmt"
	"time"
	CONSTANT "video-api/constant"
	DB "video-api/database"
)

// Video .
type Video struct {
	VideoID      string    `json:"video_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ThumbnailURL string    `json:"thumbnail_url"`
	PublishedAt  time.Time `json:"published_at"`
}

// SearchVideos - serch for title, description in videos via query
func SearchVideos(query string) ([]Video, string, bool) {

	searchResult, err := DB.SearchInElastisearch(query, CONSTANT.VideosIndex)
	if err != nil {
		fmt.Println("SearchVideos", query, err)
		return []Video{}, CONSTANT.StatusCodeServerError, false // default
	}
	if searchResult.Responses == nil || len(searchResult.Responses) == 0 {
		fmt.Println("SearchVideos", query, "response is nil")
		return []Video{}, CONSTANT.StatusCodeServerError, false // default
	}
	fmt.Println(searchResult.TookInMillis, searchResult.Responses[0].Status)

	sres := searchResult.Responses[0]
	if sres.Hits == nil {
		fmt.Println("SearchVideos", query, "hits are nil")
		return []Video{}, CONSTANT.StatusCodeServerError, false // default
	}

	fmt.Println(len(sres.Hits.Hits))
	videos := []Video{}
	video := Video{}
	for _, hit := range sres.Hits.Hits {
		err = json.Unmarshal(hit.Source, &video)
		if err != nil {
			fmt.Println("SearchVideos", query, hit.Source, "unable to unmarshal")
			return []Video{}, CONSTANT.StatusCodeServerError, false // default
		}
		videos = append(videos, video)
	}

	return videos, CONSTANT.StatusCodeOk, true
}
