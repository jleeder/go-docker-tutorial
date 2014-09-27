package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Application starting...")

	server := NewAppServer()
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}
}
