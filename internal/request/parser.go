package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func GetStructFromJsonBody(res http.ResponseWriter, req *http.Request, structPointer interface{}) bool {
	if reflect.ValueOf(structPointer).Type().Kind() != reflect.Ptr {
		log.Fatal("GetStructFromJsonBody must receive a struct pointer")
	}

	bodyBytes, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("server error"))
		return false
	}

	err = json.Unmarshal(bodyBytes, structPointer)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("invalid json data format"))
		return false
	}

	return true
}
