package request

import (
	"fmt"
	"net/http"
)

func HasValidContentType(res http.ResponseWriter, req *http.Request) bool {
	requestContentType := req.Header.Get("content-type")
	if requestContentType != "application/json" {
		res.WriteHeader(http.StatusUnsupportedMediaType)
		res.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", requestContentType)))
		return false
	}

	return true
}
