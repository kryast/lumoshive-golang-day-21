package main

import (
	"day-21/handler"
	"day-21/middleware"
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
	serverMux.HandleFunc("POST /login", handler.UserLoginHandler)

	muxWithMiddleware := http.NewServeMux()
	muxWithMiddleware.HandleFunc("POST /create-task", handler.CreateTaskHandler)
	muxWithMiddleware.HandleFunc("POST /update-task", handler.UpdateTaskStatusHandler)
	muxWithMiddleware.HandleFunc("GET /read-task", handler.GetAllTasksHandler)

	authMiddleware := middleware.Middleware(muxWithMiddleware)
	roleMiddleware := middleware.RoleMiddleware(authMiddleware)

	serverMux.Handle("POST /create-task", roleMiddleware)
	serverMux.Handle("POST /update-task", roleMiddleware)
	serverMux.Handle("GET /read-task", roleMiddleware)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", serverMux)
}
