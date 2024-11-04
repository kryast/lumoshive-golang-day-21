package main

import (
	"day-21/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("POST /login", handler.LoginHandler)
	log.Printf("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
