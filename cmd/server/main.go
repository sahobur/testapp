package main

import (
	"context"
	_ "embed"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	Listen = ":8082"
)

//go:embed index.html
var html string

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, html)
	})
	http.HandleFunc("/send", SendMessage)

	fmt.Printf("Starting server at port %s ", Listen)
	err := http.ListenAndServe(Listen, nil)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		message, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading message:", err)
			fmt.Fprintf(w, "Error reading message: %v", err)
			return
		}
		log.Println("Received message:", string(message))

		err = SendMsgToMQ(string(message))
		if err != nil {
			log.Println("Error sending message:", err)
			fmt.Fprintf(w, "Error sending message: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func SendMsgToMQ(msg string) error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return fmt.Errorf("fail to connect to rabbit %v", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("fail to open channel %v", err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sendservice",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("fail to declare queue %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		return fmt.Errorf("fail to publish message %v", err)
	}
	log.Printf(" [x] Sent %s\n", msg)

	return nil
}
