package main

import (
	_ "embed"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"io"
	"log"
	"net/http"
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
			fmt.Fprintf(w, "Error reading message: %v", err)
			return
		}
		log.Println("Received message:", string(message))
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func SendMsgToMQ(msg string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
