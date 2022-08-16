package utils

import (
	"encoding/json"
	"net/http"
)

type defaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func setJSONResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}

func SetResponse(w http.ResponseWriter, statusCode int, message string, response interface{}) {
	data := defaultResponse{
		Message: message,
		Data:    response,
	}
	setJSONResponse(w, statusCode)
	json.NewEncoder(w).Encode(data)
}
