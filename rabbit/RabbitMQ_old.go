package main
//
//// queue 与 exchange 绑定
//func (rs *RabbitMQ) ExchangeQueueBind(queue_name, routing_key, exchange_name string) error {
//	return rs.Ch.QueueBind(queue_name, routing_key, exchange_name, false, nil)
//}
//
//// 注册一个Consume
//func (rs *RabbitMQ) RegisterConsume(queue_name string) (<-chan amqp.Delivery, error) {
//	msgs, err := rs.Ch.Consume(
//		queue_name, // queue
//		"",         // consumer
//		true,       // autoAck
//		false,      // exclusive
//		false,      // noLocal
//		false,      // noWait
//		nil,        // args
//
//	)
//	return msgs, err
//}
//
//// 声明一个exchange; 并声明一个与之绑定的队列,返回队列名
//func (rs *RabbitMQ) DeliveryDeclare(exchange_name string) (<-chan amqp.Delivery, error) {
//
//	// 声明一个用于广播的exchange
//	err := rs.FanoutExchangeDeclare(exchange_name)
//	common.FailOnError(err, "Failed to declare an exchange")
//
//	// 声明一个随机命名的队列
//	q, err := rs.QueueDeclareExclusive()
//	common.FailOnError(err, "Failed to declare a queue")
//
//	// queue 与 exchange 绑定
//	err = rs.ExchangeQueueBind(q.Name, "", exchange_name)
//	common.FailOnError(err, "Failed to bind a queue")
//
//	// 注册 Consume
//	msgs, err := rs.RegisterConsume(q.Name)
//	common.FailOnError(err, "Failed to register a consumer")
//
//	return msgs, err
//}
//
//// 接受广播信息
//func (rs *RabbitMQ) Receive_broadcast_msg(exchange_name string) {
//
//	msgs, err := rs.DeliveryDeclare(exchange_name)
//
//	if err != nil {
//		common.FailOnError(err, "Failed to register a consumer")
//		return
//	}
//	forever := make(chan bool)
//	go func() {
//		for d := range msgs {
//			log.Printf(" [x] %s", d.Body)
//		}
//	}()
//
//	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
//	<-forever
//}
//
//// 声明一个广播模式的 Exchange 和 queue
//func (rs *RabbitMQ) BroadCastDeclare(exchange_name string) {
//	rs.RabbitMQDial()
//	rs.FanoutExchangeDeclare("exchange_test")
//}
//
//// 发送广播信息
//func (rs *RabbitMQ) Send_broadcast_msg(exchange_name string, body []byte) {
//
//	err := rs.Ch.Publish(
//		exchange_name, // exchange
//		"",            // routing key
//		false,         // mandatory
//		false,         // immediate 立即
//		amqp.Publishing{
//			ContentType: "application/json",
//			Body:        body,
//		})
//
//	// []byte(`{"asa": "sdsd"}`)
//	common.FailOnError(err, "Failed to publish a message")
//	if err == nil {
//		log.Println("send is ok : " + string(body))
//	}
//}
//
//
//
