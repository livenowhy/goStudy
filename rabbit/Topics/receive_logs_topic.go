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
		"logs_topic",   // name
		"topic",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	common.FailOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",        // name
		false,     // durable
		false,     // delete when usused
		true,      // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_direct", s)
		err = ch.QueueBind(
			q.Name,   // queue name
			s,        // routing key
			"logs_topic",  // exchange
			false,
			nil)
		common.FailOnError(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name,     // queue
		"",         // consumer
		true,       // auto ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)
	common.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<- forever
}