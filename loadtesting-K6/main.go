package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("haloo")

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/example", ExampleHandler).Methods("GET")
	http.ListenAndServe(":8098", router)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {

	ip := ReadUserIP(r)

	data := make(map[string]interface{})
	data["success"] = true
	data["message"] = "success hit api" + ip
	BuildResponse(w, 200, data)
}

func BuildResponse(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Header.Get("X-Forwarded-For"))
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
