package endpoints

import (
	"net/http"
	"github.com/gorilla/mux"
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
		urlToProcess := url_processor.AddUrls(urls)
		utils.ResponseJson(urlToProcess, w)
	}
}

func GetResultById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId :=  vars["request_id"]
	utils.ResponseJson(url_processor.UrlsDone[requestId], w)
}

func GetAllResults(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJson(url_processor.UrlsDone, w)
}