package utils

import (
	"encoding/json"
	"net/http"
)

func JsonDecode(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	SetPanicError(err)
}

func JsonEncode(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	SetPanicError(err)
}
