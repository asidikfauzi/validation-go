package utils

import (
	"encoding/json"
	"net/http"
	"test-prepare/model/web/response"
)

func Response(w http.ResponseWriter, statusCode int, message string, data interface{}) {

	responses := response.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}

	payload, _ := json.Marshal(responses)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(payload)
}
