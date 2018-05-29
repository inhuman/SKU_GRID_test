package server

import (
	"github.com/gorilla/mux"
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/handlers"
)

func Start(router *mux.Router, port *int) {

	log.Println("Starting sku grid server on port :" + strconv.Itoa(*port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), handlers.CORS(getCORS())(router)))
}

func getCORS() (handlers.CORSOption) {

	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return methodsOk
}