package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	Find4digit()
}
func createHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 30,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   30,
	}

	return client
}

func Find4digit() {

	// find := "A567"
	// timeNow := time.Now()
	client := createHTTPClient()
	h := NewTec(client)
	r := mux.NewRouter()
	r.HandleFunc("/api/users/login", h.Tests).Methods(http.MethodPost)
	http.Handle("/", r)
	fmt.Println("Listen on PORT :", 8098)

	err2 := http.ListenAndServe(fmt.Sprintf(":%v", 8098), r)
	if err2 != nil {

		fmt.Println("errorrrrr")
	}
	// panjangCharacter := 4

	// Find10digit()
}

type Tec struct {
	client *http.Client
}

func NewTec(client *http.Client) *Tec {
	return &Tec{client: client}
}
func (h Tec) Tests(w http.ResponseWriter, r *http.Request) {
	abjad := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	hasil := ""
	hasilREs := false
	for i := 0; i < len(abjad); i++ {
		for j := 0; j < len(abjad); j++ {
			for k := 0; k < len(abjad); k++ {
				for l := 0; l < len(abjad); l++ {
					fmt.Println("dd")

					hasil = string(abjad[i]) + string(abjad[j]) + string(abjad[k]) + string(abjad[l])
					// fmt.Println(hasil)
					fmt.Println(hasil)
					requestBody, _ := json.Marshal(map[string]interface{}{
						"email":    "ademawan@gmail.com",
						"password": hasil,
					})
					req, err := http.NewRequest("POST", "http://127.0.0.1:8099/api/users/login", bytes.NewBuffer(requestBody))
					if err != nil {
						fmt.Println("error 1")
						return
					}
					// defer req.Body.Close()
					// req.Header.Add("AccessToken", accesToken)
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("Accept", "application/json")
					res, err := h.client.Do(req)
					if err != nil {
						fmt.Println("error 2", err.Error())

						return
					}
					defer res.Body.Close()

					body, err := io.ReadAll(res.Body)

					if err != nil {
						fmt.Println("error 3")

						return
					}
					var ress interface{}
					err = json.Unmarshal(body, &ress)

					if err != nil {

						fmt.Println("error 4")

						return
					}
					if res.StatusCode == 200 {
						fmt.Println("GET PASSWORD", hasil)
						hasilREs = true
					}

				}
			}
		}
	}

	if hasilREs {
		RespondWithJSON(w, 200, "Success "+hasil)

	} else {
		RespondWithJSON(w, 400, "fales")

	}

}
func Find10digit() {
	abjad := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	hasil := ""
	find := "ZXAXHJKHSJ"
	// panjangCharacter := 4
	timeNow := time.Now()
	for i := 0; i < len(abjad); i++ {
		for j := 0; j < len(abjad); j++ {
			for k := 0; k < len(abjad); k++ {
				for l := 0; l < len(abjad); l++ {
					for m := 0; m < len(abjad); m++ {
						for n := 0; n < len(abjad); n++ {
							for o := 0; o < len(abjad); o++ {
								for p := 0; p < len(abjad); p++ {
									for q := 0; q < len(abjad); q++ {
										for r := 0; r < len(abjad); r++ {
											hasil = string(abjad[i]) + string(abjad[j]) + string(abjad[k]) + string(abjad[l]) + string(abjad[m]) + string(abjad[n]) + string(abjad[o]) + string(abjad[p]) + string(abjad[q]) + string(abjad[r])
											// fmt.Println(hasil)
											fmt.Println(hasil)
											if hasil == find {
												fmt.Println("KETEMU", hasil)
												fmt.Println(time.Since(timeNow))
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
