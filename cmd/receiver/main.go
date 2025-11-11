package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	receiver()
}
func receiver() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sendservice", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Panic("Failed to declare a queue", err)
	}

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Panic("Failed to register a consumer", err)
	}

	var exit chan struct{}

	f := func(d amqp.Delivery) {
		log.Printf("Received a message from server: %s", d.Body)
	}

	for d := range messages {
		go f(d)
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-exit
}
