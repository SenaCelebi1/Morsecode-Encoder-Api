package router

import (
	"morsecode/service"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/morse/{text}", service.MorseCodeFunc).Methods("POST", "OPTIONS")

	return router
}
