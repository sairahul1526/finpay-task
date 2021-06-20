package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	CONFIG "video-parser/config"
	CONSTANT "video-parser/constant"
	MODEL "video-parser/model"
)

// GetYoutubeVideoResults - get search results from youtube api
func GetYoutubeVideoResults(apiKey, publishedAfter string) (MODEL.SearchResponse, error) {
	// get request for youtube search api
	res, err := http.Get(CONSTANT.YoutubeSearchURL + "?part=snippet&maxResults=100&order=date&q=" + CONFIG.SearchQueryYoutube + "&key=" + apiKey + "&type=video&publishedAfter=" + publishedAfter)
	if err != nil {
		fmt.Println("GetYoutubeVideoResults", err)
		return MODEL.SearchResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("GetYoutubeVideoResults", err)
		return MODEL.SearchResponse{}, err
	}

	// parse body
	searchResponse := MODEL.SearchResponse{}
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		fmt.Println("GetYoutubeVideoResults", err)
		return MODEL.SearchResponse{}, err
	}

	return searchResponse, nil
}
