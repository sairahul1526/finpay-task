package video

import (
	"net/http"
	"strconv"
	CONSTANT "video-api/constant"
	MODEL "video-api/model"

	UTIL "video-api/util"
)

// VideoSearch godoc
// @Tags Video Search
// @Summary Get latest videos by query
// @Router /video/search [get]
// @Param query query string true "Query term to search"
// @Param page query string false "Page number"
// @Produce json
// @Success 200
func VideoSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response = make(map[string]interface{})

	// get videos by query
	videos, status, ok := MODEL.SearchVideos(`{"from": ` + strconv.Itoa(UTIL.GetPageNumber(r.FormValue("page"))*CONSTANT.SearchVideosPerPage) + `, "size": ` + strconv.Itoa(CONSTANT.SearchVideosPerPage) + `, "query": {"multi_match": {"query": "` + r.FormValue("query") + `", "fields": ["title", "description"]}},"sort": [{"published_at": {"order": "desc"}}]}`)
	if !ok {
		UTIL.SetReponse(w, status, "", CONSTANT.ShowDialog, response)
		return
	}

	response["videos"] = videos
	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", CONSTANT.ShowDialog, response)
}
