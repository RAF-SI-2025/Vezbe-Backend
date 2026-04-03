package main

import (
	"log"
	"net/http"

	"github.com/RAF-SI-2025/Vezbe-Backend/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", handler.GetAllUsers)
	mux.HandleFunc("POST /users", handler.CreateUser)
	mux.HandleFunc("GET /users/{id}", handler.GetUserByID)
	mux.HandleFunc("PUT /users/{id}", handler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", handler.DeleteUser)
	mux.HandleFunc("PUT /users/{id}/password", handler.ChangePassword)

	const addr = ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("Server started on http://localhost%s\n", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
