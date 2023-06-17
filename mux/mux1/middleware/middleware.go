package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ValidationChannel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		HEADER_CHANNELID := r.Header.Get("CHANNELID")

		if HEADER_CHANNELID == "" {
			ResponseWithErr(w, http.StatusUnauthorized, "error chanel", nil)
			return
		}
		fmt.Println("Middleware channel di lewati")

		next.ServeHTTP(w, r)
	})
}

func DeviceIdValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		DEVICEID := r.Header.Get("device-id")

		if DEVICEID == "" {
			ResponseWithErr(w, http.StatusUnauthorized, "error device", nil)
			return
		}
		fmt.Println("Middleware device di lewati")

		next.ServeHTTP(w, r)
	})
}

type ResponseError struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
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
