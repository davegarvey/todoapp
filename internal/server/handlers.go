package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"todo_app/internal/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var todos = make(map[string]auth.Todo)

func checkAuthZ(req auth.AuthZENRequest) (bool, error) {
	log.Printf("checkAuthZ called with request: %+v\n", req)
	pdpURL := os.Getenv("AUTHZEN_PDP_URL")
	if pdpURL == "" {
		return false, fmt.Errorf("AUTHZEN_PDP_URL not set")
	}

	jsonBody, err := json.Marshal(req)
	if err != nil {
		return false, err
	}

	httpReq, err := http.NewRequest("POST", pdpURL+"/access/v1/evaluation", bytes.NewBuffer(jsonBody))
	if err != nil {
		return false, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var authzResp auth.AuthZENResponse
	if err := json.NewDecoder(resp.Body).Decode(&authzResp); err != nil {
		return false, err
	}

	log.Printf("Authorization response: %+v\n", authzResp)
	return authzResp.Decision, nil
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUserHandler called")
	vars := mux.Vars(r)
	userID := vars["userID"]
	log.Printf("User ID: %s\n", userID)
	subject, err := extractSubjectFromJWT(r)
	if err != nil {
		log.Printf("Error extracting subject from JWT: %v\n", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	log.Printf("Subject extracted from JWT: %s\n", subject)

	authzReq := auth.AuthZENRequest{
		Subject: auth.Subject{
			Type: "user",
			ID:   subject,
		},
		Action: auth.Action{
			Name: "can_read_user",
		},
		Resource: auth.Resource{
			Type: "user",
			ID:   userID,
		},
	}

	log.Printf("Authorization request: %+v\n", authzReq)
	allowed, err := checkAuthZ(authzReq)
	if err != nil {
		log.Printf("Authorization error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !allowed {
		log.Println("Unauthorized access")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Mock user response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":      userID,
		"name":    "Mock User",
		"email":   userID,
		"picture": "https://example.com/avatar.jpg",
	})
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetTodosHandler called")
	subject, err := extractSubjectFromJWT(r)
	if err != nil {
		log.Printf("Error extracting subject from JWT: %v\n", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	log.Printf("Subject extracted from JWT: %s\n", subject)

	authzReq := auth.AuthZENRequest{
		Subject: auth.Subject{
			Type: "user",
			ID:   subject,
		},
		Action: auth.Action{
			Name: "can_read_todos",
		},
		Resource: auth.Resource{
			Type: "todo",
			ID:   "todo-1",
		},
	}

	log.Printf("Authorization request: %+v\n", authzReq)
	allowed, err := checkAuthZ(authzReq)
	if err != nil {
		log.Printf("Authorization error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !allowed {
		log.Println("Unauthorized access")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	todosList := make([]auth.Todo, 0, len(todos))
	for _, todo := range todos {
		todosList = append(todosList, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todosList)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateTodoHandler called")
	subject, err := extractSubjectFromJWT(r)
	if err != nil {
		log.Printf("Error extracting subject from JWT: %v\n", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	log.Printf("Subject extracted from JWT: %s\n", subject)

	authzReq := auth.AuthZENRequest{
		Subject: auth.Subject{
			Type: "user",
			ID:   subject,
		},
		Action: auth.Action{
			Name: "can_create_todo",
		},
		Resource: auth.Resource{
			Type: "todo",
			ID:   "todo-1",
		},
	}

	log.Printf("Authorization request: %+v\n", authzReq)
	allowed, err := checkAuthZ(authzReq)
	if err != nil {
		log.Printf("Authorization error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !allowed {
		log.Println("Unauthorized access")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var todo auth.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.ID = uuid.New().String()
	todo.OwnerID = subject
	todos[todo.ID] = todo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateTodoHandler called")
	vars := mux.Vars(r)
	todoID := vars["id"]
	log.Printf("Todo ID: %s\n", todoID)
	subject, err := extractSubjectFromJWT(r)
	if err != nil {
		log.Printf("Error extracting subject from JWT: %v\n", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	log.Printf("Subject extracted from JWT: %s\n", subject)

	todo, exists := todos[todoID]
	if !exists {
		log.Println("Todo not found")
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	authzReq := auth.AuthZENRequest{
		Subject: auth.Subject{
			Type: "user",
			ID:   subject,
		},
		Action: auth.Action{
			Name: "can_update_todo",
		},
		Resource: auth.Resource{
			Type: "todo",
			ID:   todoID,
			Properties: auth.Properties{
				OwnerID: todo.OwnerID,
			},
		},
	}

	log.Printf("Authorization request: %+v\n", authzReq)
	allowed, err := checkAuthZ(authzReq)
	if err != nil {
		log.Printf("Authorization error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !allowed {
		log.Println("Unauthorized access")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var updatedTodo auth.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTodo.ID = todoID
	updatedTodo.OwnerID = todo.OwnerID
	todos[todoID] = updatedTodo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteTodoHandler called")

	vars := mux.Vars(r)
	todoID := vars["id"]
	log.Printf("Todo ID: %s\n", todoID)

	subject, err := extractSubjectFromJWT(r)
	if err != nil {
		log.Printf("Error extracting subject from JWT: %v\n", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	log.Printf("Subject extracted from JWT: %s\n", subject)

	todo, exists := todos[todoID]
	if !exists {
		log.Println("Todo not found")
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	authzReq := auth.AuthZENRequest{
		Subject: auth.Subject{
			Type: "user",
			ID:   subject,
		},
		Action: auth.Action{
			Name: "can_delete_todo",
		},
		Resource: auth.Resource{
			Type: "todo",
			ID:   todoID,
			Properties: auth.Properties{
				OwnerID: todo.OwnerID,
			},
		},
	}

	log.Printf("Authorization request: %+v\n", authzReq)
	allowed, err := checkAuthZ(authzReq)
	if err != nil {
		log.Printf("Authorization error: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !allowed {
		log.Println("Unauthorized access")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	delete(todos, todoID)
	w.WriteHeader(http.StatusNoContent)
}

// Helper function to extract subject from JWT
func extractSubjectFromJWT(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return "", fmt.Errorf("invalid authorization header")
	}

	tokenString := authHeader[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your-256-bit-secret"), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		subject, ok := claims["sub"].(string)
		if !ok {
			return "", fmt.Errorf("subject claim not found")
		}
		return subject, nil
	}

	return "", fmt.Errorf("invalid token")
}
