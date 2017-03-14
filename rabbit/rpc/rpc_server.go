package main

import (
	"log"
	"github.com/streadway/amqp"
	"github.com/livenowhy/RabbitResponse/Tutorials/common"
	"strconv"
	"time"
)


func fib(n int) int {


	time.Sleep(time.Second)

	return 34
}

func main() {
	conn, err := amqp.Dial(common.HOST_URL)
	common.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue_ee",   // name
		false,         // durable
		false,         // delete when usused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")
	err = ch.Qos(
		1,       // prefetch count
		0,       // prefetch size
		false,   // global
	)

	msgs, err := ch.Consume(
		q.Name,    // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	common.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			common.FailOnError(err, "Failed to convert body to integer")

			log.Printf(" [.] fib(%d)", n)
			response := fib(n)
			err = ch.Publish(
				"",         // exchange
				d.ReplyTo,  // routing key
				false,      // mandatory
				false,      // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					CorrelationId: d.CorrelationId,
					Body: []byte(strconv.Itoa(response)),
				})
			common.FailOnError(err, "Failed to publish a message")
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}