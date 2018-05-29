package endpoints

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	"url_processor"
	"utils"
)


func UploadUrls(w http.ResponseWriter, r *http.Request) {
	var urls url_processor.Urls
	err := json.NewDecoder(r.Body).Decode(&urls)

	if err != nil {
		w.Write([]byte(err.Error()))
		utils.CheckError(err)
	} else {
		url_processor.AddUrls(urls)

	}
}


func GetResultById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId :=  vars["request_id"]
	fmt.Print(requestId)
}

func GetAllResults(w http.ResponseWriter, r *http.Request) {

}