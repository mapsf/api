package controllers

import (
	"net/http"
	"encoding/json"
)

func JSON(w http.ResponseWriter, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte("error"))
	}
	w.Write(bytes)
}

type successResponse struct {
	Data string `json:"data"`
}

type errorResponse struct {
	Message   string `json:"message"`
	ErrorCode uint   `json:"error_code"`
}
