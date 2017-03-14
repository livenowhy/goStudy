package main



import (
	"log"
	"github.com/streadway/amqp"
	"github.com/livenowhy/RabbitResponse/Tutorials/common"

	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)

}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)

}

func FibonacciRPC(n int) (res int, err error) {
	conn, err := amqp.Dial(common.HOST_URL)
	common.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()



	ch, err := conn.Channel()
	common.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",       // name
		false,    // durable
		false,    // delete when usused
		false,     // exclusive
		false,    // noWait
		nil,      // arguments
	)
	common.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,    // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	common.FailOnError(err, "Failed to register a consumer")

	corrId := RandomString(32)

	err = ch.Publish(
		"",            // exchange
		"rpc_queue",   // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			CorrelationId: corrId,
			ReplyTo: q.Name,
			Body: []byte(strconv.Itoa(n)),
		})

	common.FailOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			if corrId == d.CorrelationId {
				res, err = strconv.Atoi(string(d.Body))
				common.FailOnError(err, "Failed to convert body to integer")
				break

			}
		}
	}

	return
}





func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	//n := bodyFrom(os.Args)

	for i := 0; i <300; i++ {

		log.Printf(" [x] Requesting fib(%d)", i)
	res, err := FibonacciRPC(i)
	common.FailOnError(err, "Failed to handle RPC request")

	log.Printf(" [.] Got %d", res)

	}
}

func bodyFrom(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	common.FailOnError(err, "Failed to convert arg to integer")
	return n
}