package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"sync"
	"test-mux/middleware"
	"text/tabwriter"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	t := NewHandler()

	r := mux.NewRouter()

	paymentSubRouter := r.PathPrefix("/v2/mux/").Subrouter()
	paymentSubRouterStatus := r.PathPrefix("/v2/mux/status").Subrouter()

	paymentSubRouter.HandleFunc("/{payment_type}", t.Handler).Methods(http.MethodGet)
	paymentSubRouterStatus.HandleFunc("/{transaction_id}", t.HandlerStatus).Methods(http.MethodGet)
	paymentSubRouter.Use(middleware.ValidationChannel, middleware.DeviceIdValidation)
	paymentSubRouterStatus.Use(middleware.ValidationChannel)
	server := &http.Server{Addr: ":" + "8092", Handler: r, TLSConfig: &tls.Config{InsecureSkipVerify: true}}
	log.Printf("starting service on port %s", "8092")
	server.ListenAndServe()
}

type TestDelivery struct{}

func NewHandler() *TestDelivery {
	return &TestDelivery{}
}

func (t *TestDelivery) Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paymentType := params["payment_type"]

	var res interface{}
	switch paymentType {
	case "ppob":
		res = UseCasePPOB()
	case "credit":
		res = UseCaseCredit()
	}
	example1()
	fmt.Println("hallo")
	ResponseWithJSON(w, "success", http.StatusOK, res)
	return
}
func (t *TestDelivery) HandlerStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paymentType := params["transaction_id"]

	// var res interface{}
	// switch paymentType {
	// case "ppob":
	// 	res = UseCasePPOB()
	// case "credit":
	// 	res = UseCaseCredit()
	// }
	ResponseWithJSON(w, "success", http.StatusOK, paymentType)
	return
}

type Data struct {
	Payment string
}
type ValidateCiamOtpRequest struct {
	AuthId string `json:"auth_id"`
	Msisdn string `json:"msisdn"`
	Token  string `json:"token"`
}

func UseCasePPOB() *Data {

	httpClient := createHTTPClient()
	url := "http://localhost:8098/api/v1/example?ccc=ok"
	method := "GET"

	reqBody := ValidateCiamOtpRequest{
		AuthId: "authId",
		Token:  "token",
	}

	bodyJson, _ := json.Marshal(reqBody)

	http, err := http.NewRequest(method, url, bytes.NewReader(bodyJson))
	if err != nil {
		fmt.Println(err.Error())
	}
	http.Header.Add("token", "token")
	http.Header.Add("code", "code")
	fmt.Println(http)
	res, err := httpClient.Do(http)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	curl := GenerateCURL(http)
	fmt.Println(curl)

	bodySubmit, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("SUCCESS", res.StatusCode, string(bodySubmit), http.URL.Query())

	return &Data{Payment: "PPOB"}
}

func GenerateCURL(http2 *http.Request) string {
	curl := `curl `
	if http2.Header != nil {
		for key, val := range http2.Header {
			curl += fmt.Sprintf(` -H '%s:%s' `, key, val[0])
		}
	}

	curl += ` --request ` + http2.Method

	curl += fmt.Sprintf(` %v`, http2.URL)

	return curl
}
func UseCaseCredit() *Data {
	return &Data{Payment: "Credit"}
}

func ResponseWithJSON(w http.ResponseWriter, message string, code int, payload interface{}) {

	responseSuccess := ResponseSuccess{
		Status:  code,
		Message: message,
		Data:    payload,
	}

	js, _ := json.Marshal(responseSuccess)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(js)
}

type ResponseSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	MaxIdleConns       int  = 100
	MaxIdleConnections int  = 100
	RequestTimeout     int  = 30
	SSL                bool = true
)

func createHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: SSL},
		MaxIdleConns:        MaxIdleConns,
		MaxIdleConnsPerHost: MaxIdleConnections,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(RequestTimeout) * time.Second,
	}

	return client
}

func example1() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) { //1
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1) //2
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup

		wg.Add(count + 1)
		beginTestTime := time.Now()

		go producer(&wg, mutex)

		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Reader\tRWmutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}

}
