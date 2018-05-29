package router

import (
	"github.com/gorilla/mux"
	"endpoints"

)

func GetRouter() *mux.Router {

	var router = mux.NewRouter()

	router.HandleFunc("/upload/urls", 	endpoints.UploadUrls).Methods("POST")
	router.HandleFunc("/result/{request_id}", endpoints.GetResultById).Methods("GET")
	router.HandleFunc("/result", endpoints.GetAllResults).Methods("GET")

	return router
}

