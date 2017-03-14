package main

import (
	"log"
	"github.com/streadway/amqp"
	"github.com/livenowhy/RabbitResponse/Tutorials/common"
)

func main() {

	conn, err := amqp.Dial(common.HOST_URL)
	common.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"gohello",  // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil, // arguments
	)
	common.FailOnError(err,  "Failed to declare a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"", // consumer 消费者
		true, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, //args
	)

	common.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}