package utils

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func CheckError(err error) {

	if err != nil {
		fmt.Println(err)
	}
}

func ResponseJson(i interface{}, w http.ResponseWriter) {

	jsn, err := json.Marshal(i)
	CheckError(err)
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(jsn))
}