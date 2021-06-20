package model

import (
	"time"
	CONSTANT "video-parser/constant"
	DB "video-parser/database"
)

// Video .
type Video struct {
	VideoID      string    `json:"video_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ThumbnailURL string    `json:"thumbnail_url"`
	PublishedAt  time.Time `json:"published_at"`
}

// InsertVideos - add videos into database
func InsertVideos(videos []Video) error {
	ids := []string{}
	videosConvert := []interface{}{}
	for _, video := range videos {
		videosConvert = append(videosConvert, video)
		ids = append(ids, video.VideoID)
	}
	return DB.InsertIntoElastisearch(CONSTANT.VideosIndex, videosConvert, ids)
}
