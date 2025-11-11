package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

const (
	Listen = ":8082"
)

//go:embed frontend/static/index.html
var html string

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
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
	fmt.Fprintf(w, "Hello, World!")
}
