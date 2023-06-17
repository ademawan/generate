package helper

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/sirupsen/logrus"
)

var (
	log                   = logrus.New()
	OsOpenFile            = os.OpenFile
	getEnv                = os.Getenv
	sqlOpen               = sql.Open
	pathLogAuth           = "logs/service-auth.log"
	failedLogMessage      = "Failed to log to file, using default stderr (http log)"
	HttpNewRequest        = http.NewRequest
	IoutilReadAll         = ioutil.ReadAll
	JsonUnmarshal         = json.Unmarshal
	JsonMarshal           = json.Marshal
	UrlParseQuery         = url.ParseQuery
	UrlParse              = url.Parse
	JsonNewDecoder        = json.NewDecoder
	IoutilReadFile        = ioutil.ReadFile
	TlsLoadX509KeyPair    = tls.LoadX509KeyPair
	GodotenvLoad          = godotenv.Load
	HttpListenAndServe    = http.ListenAndServe
	HttpListenAndServeTLS = http.ListenAndServeTLS
)

type HTTPLogData struct {
	Email        string
	Level        string
	Method       string
	Service      string
	OrderID      string
	Status       int
	Message      string
	ResponseTime time.Duration
	ReqURL       string
	Response     string
	Err          error
}

// HTTPLog function
func HTTPLog(h *HTTPLogData) error {

	log.Out = os.Stdout
	logDir := ""

	if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "preproduction" {
		logDir = "/app/"
	}

	file, errorLog := OsOpenFile(logDir+pathLogAuth, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if errorLog == nil {
		log.Out = file
	} else {
		log.Info(failedLogMessage)
		return errorLog
	}
	var errorMessage string
	if h.Err != nil {
		errorMessage = h.Err.Error()
	} else {
		errorMessage = ""
	}

	payload := logrus.Fields{
		"method":        h.Method,
		"service":       h.Service,
		"status code":   h.Status,
		"response time": h.ResponseTime,
		"url":           h.ReqURL,
		"response":      h.Response,
		"error":         errorMessage,
		"email":         h.Email,
	}

	if h.Level == "info" {
		log.WithFields(payload).Info(h.Message)
	} else if h.Level == "warn" {
		log.WithFields(payload).Warn(h.Message)
	} else if h.Level == "error" {
		log.WithFields(payload).Error(h.Message)
	}

	return nil
}
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
