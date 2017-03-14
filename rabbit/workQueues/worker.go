package main

import (
	"log"
	"github.com/streadway/amqp"
	"github.com/livenowhy/RabbitResponse/Tutorials/common"
	"bytes"
	"time"
)

func main() {

	conn, err := amqp.Dial(common.HOST_URL)
	common.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		common.WORK_QUEUE_NAME,  // name
		true,   // durable, 持久的
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil, // arguments
	)
	common.FailOnError(err,  "Failed to declare a queue")


	err = ch.Qos(
		1,   // prefetch count
		0,  // prefetch size
		false, // global
	)
	common.FailOnError(err, "Failed to set QoS")


	msgs, err := ch.Consume(
		q.Name, // queue
		"", // consumer 消费者
		false, // auto-ack
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

			d.Ack(false)
			dot_count := bytes.Count(d.Body, []byte("."))
			log.Println(dot_count)
			t := time.Duration(dot_count)

			time.Sleep(t * time.Second)
			log.Printf("Done")
			log.Println(t * time.Second)

		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}