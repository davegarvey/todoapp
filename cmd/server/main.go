package main

import (
	"log"
	"net/http"
	"todo_app/internal/server"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/users/{userID}", server.GetUserHandler).Methods("GET")
	r.HandleFunc("/todos", server.GetTodosHandler).Methods("GET")
	r.HandleFunc("/todos", server.CreateTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id}", server.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todos/{id}", server.DeleteTodoHandler).Methods("DELETE")

	port := "9080"
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
