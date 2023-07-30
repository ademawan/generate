package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/NeowayLabs/wabbit"
	"github.com/NeowayLabs/wabbit/amqp"

	"github.com/joho/godotenv"
	amqp_default "github.com/rabbitmq/amqp091-go"

	"strconv"
	"time"
	// amqp "github.com/rabbitmq/amqp091-go"
)

type Publishamqp struct {
	BearerToken string
	TokenID     string
	IsMerchant  bool
}

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

func HTTPLog(h *HTTPLogData) error {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	log := logrus.New()
	log.Out = os.Stdout
	logDir := ""

	file, errorLog := os.OpenFile(logDir+"logs/auth.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if errorLog == nil {
		log.Out = file
	} else {
		panic(errorLog)
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

	// str := fmt.Sprintf("CITY_IS_ALFABETINC:%v|CITY_IS_RESTRICTED:%v|CITY_IS_NUMERIC:%v", !helper.IsAlfabetic(p.City), !helper.RestrictedCharacter(p.City), !helper.IsNumeric(p.City))
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// RabbitMQDefault()
	RabbitMQWABBIT()
}

func RabbitMQWABBIT() {
	// QueueBiasa()
	conn, err := NewConnectionWabbit()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10000; i++ {

		start := time.Now()
		ch, err := conn.Channel()
		index := strconv.Itoa(i)
		defer conn.Close()

		if err != nil {
			data := &HTTPLogData{
				Level:        "error",
				Method:       "POST",
				Service:      "mb.Channel",
				Status:       http.StatusBadRequest,
				Message:      "channel failed to connect",
				ResponseTime: time.Since(start),
				Err:          err,
			}
			HTTPLog(data)
			panic(err)
		}
		publishamqp := Publishamqp{
			BearerToken: "BearerToken_" + index,
			TokenID:     "TokenID_" + index,
			IsMerchant:  true,
		}

		bd, _ := json.Marshal(publishamqp)
		err = ch.Publish(
			"",   // exchange
			"cc", // routing key
			bd,
			wabbit.Option{
				"contentType": "application/json",
			},
		)
		if err != nil {
			data := &HTTPLogData{
				Level:        "error",
				Method:       "POST",
				Service:      "Publish",
				Status:       http.StatusBadRequest,
				Message:      "RabbitMQ Failed",
				ResponseTime: time.Since(start),
				ReqURL:       "urlAccessToken",
				Err:          err,
			}
			HTTPLog(data)
			panic(err)
		}
	}
	timeOut := "1000"
	timeOutInt, _ := strconv.Atoi(timeOut)
	fmt.Println(timeOutInt)
	time.Sleep(time.Duration(timeOutInt) * time.Second)
}

func RabbitMQDefault() {
	conn, err := NewConnection()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		// QueueBiasa()
		start := time.Now()
		ch, err := conn.Channel()

		index := strconv.Itoa(i)
		// conn.Close()

		if err != nil {
			data := &HTTPLogData{
				Level:        "error",
				Method:       "POST",
				Service:      "mb.Channel",
				Status:       http.StatusBadRequest,
				Message:      "channel failed to connect",
				ResponseTime: time.Since(start),
				Err:          err,
			}
			HTTPLog(data)
			panic(err)
		}
		ch.Close()
		// conn.Close()

		publishamqp := Publishamqp{
			BearerToken: "BearerToken_" + index,
			TokenID:     "TokenID_" + index,
			IsMerchant:  true,
		}
		js, _ := json.Marshal(publishamqp)
		if err != nil {
			log.Println(err)
		}
		content := amqp_default.Publishing{
			ContentType: "text/plain",
			Body:        []byte(js),
		}

		err = ch.Publish(
			"",
			"coba",
			false,
			false,
			content,
		)
		if err != nil {
			fmt.Println(err.Error())

			data := &HTTPLogData{
				Level:        "error",
				Method:       "POST",
				Service:      "Publish",
				Status:       http.StatusBadRequest,
				Message:      "RabbitMQ Failed",
				ResponseTime: time.Since(start),
				ReqURL:       "urlAccessToken",
				Err:          err,
			}
			HTTPLog(data)
			panic(err)
		}
	}

	timeOut := "1000"
	timeOutInt, _ := strconv.Atoi(timeOut)
	fmt.Println(timeOutInt)
	time.Sleep(time.Duration(timeOutInt) * time.Second)
}
func NewConnectionWabbit() (*amqp.Conn, error) {

	// uri := os.Getenv("RABBITMQ_URL")
	uri := "amqp://guest:guest@localhost:5672/"

	// not using ssl
	if uri == "" {
		log.Fatal("RabbitMQ URI is invalid")
	}

	connection, err := amqp.Dial(uri)
	log.Println("RabbitMQ accepted connection")

	if err != nil {
		return nil, err
	}

	return connection, err
}
func NewConnection() (*amqp_default.Connection, error) {

	// uri := os.Getenv("RABBITMQ_URL")
	uri := "amqp://guest:guest@localhost:5672/"

	// not using ssl
	if uri == "" {
		log.Fatal("RabbitMQ URI is invalid")
	}

	connection, err := amqp_default.Dial(uri)
	log.Println("RabbitMQ accepted connection")

	if err != nil {
		return nil, err
	}

	return connection, err
}
