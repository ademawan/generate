package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
	// log_dir      = ""
	// start        = time.Now()
	// errorMessage = "system error"
	openFile = os.OpenFile
)

func main() {
	timeNow := time.Now()

	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(fmt.Sprintf("%v", time.Since(timeNow)))

	// test := &Log2{Event: "sub validate", ResponseTime: time.Since(timeNow), Response: "holla"}
	// data, _ := json.Marshal("test")
	// CreateLog(&Log2{
	// 	Event:        "validate",
	// 	ResponseTime: time.Since(timeNow),
	// 	Response:     string(data),
	// })

	// CreateLogV2(&Log2{
	// 	Event:        "validate",
	// 	ResponseTime: time.Since(timeNow),
	// 	Response:     string(data),
	// }, "info", "success")

	// go func(timeNow time.Time) {
	// 	HttpService("info", "GET", "", "purchase", http.StatusOK, "Success", time.Since(timeNow), "Url", "bofy")

	// }(timeNow)
	// time.Sleep(time.Duration(5) * time.Second)
}

func HttpService(level string, method string, request string, service string, status int, message string, responseTime time.Duration, url string, response string) {
	log.Out = os.Stdout

	file, err := openFile("logs/service-halo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr (http log)")
	}

	payload := logrus.Fields{
		"method":           method,
		"request":          request,
		"service digicore": service,
		"status code":      status,
		"response time":    responseTime,
		"url":              url,
		"response":         response,
		"error":            err,
	}

	if level == "info" {
		log.WithFields(payload).Info(message)
	} else if level == "warn" {
		log.WithFields(payload).Warn(message)
	} else if level == "error" {
		log.WithFields(payload).Error(message)
	}

	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: false},
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(30) * time.Second,
	}

	baseUrl := os.Getenv("URL_INVOICE_SERVICE")
	baseUrl += "/" + "276" + "/last-invoice"
	req, _ := http.NewRequest("GET", baseUrl, nil)

	req.Header.Set("content-type", "application/json")
	req.Header.Set("test", "aaaa/json")
	ree, _ := json.Marshal(req.Header)
	fmt.Println(string(ree))
	fmt.Println(client)
	// res, _ := client.Do(req)

	// body, err := ioutil.ReadAll(res.Body)

}

type Log2 struct {
	Event        string
	ResponseTime time.Duration
	Response     interface{}
}

func CreateLog(data *Log2) error {

	log.Out = os.Stdout

	file, err := openFile("logs/purchase.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// You could set this to any `io.Writer` such as a file
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.WithFields(logrus.Fields{
		"event":         data.Event,
		"response_time": data.ResponseTime,
		"response":      data.Response,
	}).Error(data.Response)

	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	// log.Out = os.Stdout

	return nil
}

func CreateLogV2(data interface{}, level, message string) {
	log.Out = os.Stdout

	file, err := openFile("logs/logging.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr (http log)")
	}

	payload := logrus.Fields{}
	j, _ := json.Marshal(data)
	err = json.Unmarshal(j, &payload)
	if err != nil {
		fmt.Println(err.Error(), string(j))
		return
	}

	if level == "info" {
		log.WithFields(payload).Info(message)
	} else if level == "warn" {
		log.WithFields(payload).Warn(message)
	} else if level == "error" {
		log.WithFields(payload).Error(message)
	}
}
