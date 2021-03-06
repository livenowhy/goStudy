package main


import (
	"log"
	"github.com/streadway/amqp"
	"github.com/livenowhy/RabbitResponse/Tutorials/common"
	"os"
)

func main() {
	conn, err := amqp.Dial(common.HOST_URL)
	common.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   //  name
		"fanout", //  type
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil, // arguments
	)
	common.FailOnError(err, "Failed to declare an exchange")

	body := common.BodyFrom(os.Args)
	err = ch.Publish(
		"logs", // exchange
		"",   // routing key
		false, // mandatory
		false, // immediate 立即
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})

	common.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
