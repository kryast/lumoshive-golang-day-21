package main

import (
	"day-21/handler"
	"fmt"
	"net/http"
)

func main() {

	// http.HandleFunc("POST /login", handler.LoginHandler)
	// log.Printf("Starting server on :8080")
	// http.ListenAndServe(":8080", nil)

	// http.HandleFunc("POST /login", handler.LoginHandler)
	// http.HandleFunc("GET /login", handler.GetAdminByIDHandler)
	// log.Printf("Starting server on :8080")
	// http.ListenAndServe(":8080", nil)

	serverMux := http.NewServeMux()

	serverMux.HandleFunc("POST /create", handler.CreateUserHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", serverMux)
}
