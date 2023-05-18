package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// QueueBiasa()
	ImplementExchangePublisher()

}

func QueueBiasa() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"reqSendUserInvitation", // name
		false,                   // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		nil,                     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

// //IMPLEMENT RABBITMQ EXCHHANGE
func ImplementExchangePublisher() {
	conn, err := NewConnection()
	failOnError(err, "failed connection")
	s := NewRabbitMq(conn)
	// s.SendMessage(os.Getenv("EXCHANGE_NAME"), os.Getenv("QUEU_NAME"), os.Getenv("ROUTE_KEY"), message)
	s.SendMessage("EXCHANGE_NAME", "QUEUE_NAME", "ROUTE_KEY", "message")

}

func NewConnection() (*amqp.Connection, error) {

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

type RabbitMQ struct {
	connection *amqp.Connection
}

func NewRabbitMq(connection *amqp.Connection) *RabbitMQ {
	return &RabbitMQ{connection: connection}
}

func (rabbitmq *RabbitMQ) DeclareExchangeTopic(exchangeName, queuName, route string) {
	ch, err := rabbitmq.connection.Channel()
	failOnError(err, "Failed to open a channel")

	err = ch.ExchangeDeclare(
		exchangeName,
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		amqp.Table{
			"alternate-exchange": "my-ae",
		},
	)

	failOnError(err, "Failed to declare exchange")
	log.Printf("RabbitMQ: %s exchange created", exchangeName)
	_, err = ch.QueueDeclare(queuName, false, false, false, false, amqp.Table{
		"alternate-exchange": "my-ae",
	})
	failOnError(err, "Queu Declare")

	err = ch.QueueBind(queuName, route, exchangeName, false, nil)
	failOnError(err, "Queu Declare")

}

func (rabbitmq *RabbitMQ) SendMessage(exchangeName, queuName, route string, message interface{}) {
	ch, err := rabbitmq.connection.Channel()
	failOnError(err, "Failed to open a channel")

	js, _ := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	content := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(string(js)),
	}

	// declare exchange topic
	rabbitmq.DeclareExchangeTopic(exchangeName, queuName, route)

	err = ch.Publish(
		exchangeName,
		route,
		false,
		false,
		content,
	)

	failOnError(err, "Failed to publish content")

	log.Printf("RabbitMQ: published message to %s", exchangeName)
}
