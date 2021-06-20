package admin

import "github.com/gorilla/mux"

// LoadAdminRoutes - load all admin routes with admin prefix
func LoadVideoRoutes(router *mux.Router) {
	videoRoutes := router.PathPrefix("/admin").Subrouter()

	// apikey
	videoRoutes.HandleFunc("/apikey", APIKeyGet).Methods("GET")
	videoRoutes.HandleFunc("/apikey", APIKeyAdd).Methods("POST")

}
