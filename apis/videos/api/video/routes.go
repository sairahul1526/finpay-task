package video

import "github.com/gorilla/mux"

// LoadVideoRoutes - load all video routes with video prefix
func LoadVideoRoutes(router *mux.Router) {
	videoRoutes := router.PathPrefix("/video").Subrouter()

	// search
	videoRoutes.HandleFunc("/search", VideoSearch).Queries(
		"query", "{query}",
	).Methods("GET")

}
