package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	Listen = ":8081"
)

func main() {

	http.HandleFunc("/send", SendMessage)

	fmt.Printf("Starting server at port %s ", Listen)
	err := http.ListenAndServe(Listen, nil)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
