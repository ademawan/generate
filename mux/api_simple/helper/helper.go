package helper

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func ResponseWithErr(w http.ResponseWriter, code int, message string, payload interface{}) {

	responseFailed := ResponseError{
		Status:  code,
		Message: message,
		Data:    payload,
	}
	js, _ := json.Marshal(responseFailed)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(js)
}
