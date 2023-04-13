package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/NeowayLabs/wabbit"
	"github.com/NeowayLabs/wabbit/amqp"
	// "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var (
	AmqpDial = amqp.Dial
	mb       wabbit.Conn
)

type EmailInvitationUser struct {
	Name           string `json:"name"`
	StaffName      string `json:"staff_name"`
	Email          string `json:"email"`
	EncryptedEmail string `json:"encrypted"`
	TeamName       string `json:"team_name"`
	MerchantName   string `json:"merchant_name"`
	MerchantImage  string `json:"merchant_image"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	// mbUrl := os.Getenv("MB_URL")
	mb, err = AmqpDial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Succees load queue.")

	go listenServiceUserInvitation(mb)

	// q, err := ch.QueueDeclare(
	// 	"hello", // name
	// 	false,   // durable
	// 	false,   // delete when unused
	// 	false,   // exclusive
	// 	false,   // no-wait
	// 	nil,     // arguments
	// )
	// failOnError(err, "Failed to declare a queue")

	// msgs, err := ch.Consume(
	// 	q.Name,                  // queue
	// 	"reqSendUserInvitation", // consumer
	// 	true,                    // auto-ack
	// 	false,                   // exclusive
	// 	false,                   // no-local
	// 	false,                   // no-wait
	// 	nil,                     // args
	// )
	// failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
func listenQueue(mb wabbit.Conn, client string, queueName string) (string, wabbit.Channel, <-chan wabbit.Delivery) {
	// if os.Getenv("ENABLE_LISTEN_EMAIL") == "false" {
	// 	return "nil", nil, nil
	// }
	ch, err := mb.Channel()
	if err != nil {
		fmt.Println(err)
		return "nil", nil, nil
	}

	q, _ := ch.QueueDeclare(
		queueName, // name
		wabbit.Option{
			"durable":   false,
			"delete":    false,
			"exclusive": false,
			"noWait":    false,
		}, // arguments
	)
	if err != nil {
		fmt.Println("error", err.Error())
		return "nil", nil, nil
	}

	msgs, err := ch.Consume(
		q.Name(), // queue
		"",       // consumer
		wabbit.Option{
			"autoAck":   true,
			"exclusive": false,
			"noLocal":   false,
			"noWait":    false,
		}, //args
	)
	if err != nil {
		return "nil", nil, nil
	}

	return "emailHttpDelivery", ch, msgs
}

func listenServiceUserInvitation(mb wabbit.Conn) {
	_, ch, msgs := listenQueue(mb, "client", "reqSendUserInvitation")
	if ch == nil {
		return
	}
	defer ch.Close()

	fmt.Println("Start listen request email user invitation")

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			var Invitation EmailInvitationUser
			json.Unmarshal(d.Body(), &Invitation)
			fmt.Println(d.Body(), "reqSendUserInvitation", Invitation)
		}
	}()

	<-forever
}
