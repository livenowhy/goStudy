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

	// 声明一个队列
	q, err := ch.QueueDeclare(
		common.WORK_QUEUE_NAME,  // name
		true, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil, // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")

	body := common.BodyFrom(os.Args)
	err = ch.Publish(
		"", // exchange
		q.Name, // routing key
		false, // mandatory  强制的
		false, // immediate  立即
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body: []byte(body),
		})

	log.Printf(" [x] Sent %s", body)
	common.FailOnError(err, "Failed to publish a message")
}