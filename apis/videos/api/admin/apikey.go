package admin

import (
	"fmt"
	"net/http"
	CONSTANT "video-api/constant"

	DB "video-api/database"
	UTIL "video-api/util"
)

// APIKeyGet godoc
// @Tags Get APIKeys
// @Summary Get all youtube apikeys from list
// @Router /admin/apikey [get]
// @Produce json
// @Success 200
func APIKeyGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response = make(map[string]interface{})

	// get all keys from list
	keys, err := DB.GetFromRedisSortedSets(CONSTANT.APIKeysRedisKey, "-inf", "+inf", 0, 10000)
	if err != nil {
		fmt.Println("APIKeyGet", err)
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, "", CONSTANT.ShowDialog, response)
		return
	}

	response["keys"] = keys
	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", CONSTANT.ShowDialog, response)
}

// APIKeyAdd godoc
// @Tags Add APIKey
// @Summary Add youtube apikey to the list
// @Router /admin/apikey [post]
// @Param apikey query string true "API key to be added"
// @Produce json
// @Success 200
func APIKeyAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response = make(map[string]interface{})

	// add to list of keys with current time as score
	DB.AddToRedisSortedSets(CONSTANT.APIKeysRedisKey, float64(UTIL.GetCurrentTime().Unix()), r.FormValue("apikey"))

	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", CONSTANT.ShowDialog, response)
}
