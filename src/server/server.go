package server

import (
	"github.com/gorilla/mux"
	"strconv"
	"net/http"
	"fmt"
)

func Start(router *mux.Router, port *int) {
	fmt.Println("Starting sku grid server on port :" + strconv.Itoa(*port))
	http.ListenAndServe(":"+strconv.Itoa(*port), router)
}

