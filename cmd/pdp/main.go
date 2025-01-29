package main

import (
	"encoding/json"
	"log"
	"net/http"
	"todo_app/internal/auth"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/access/v1/evaluation", evaluationHandler).Methods("POST")
	r.HandleFunc("/access/v1/evaluations", evaluationsHandler).Methods("POST")

	port := "8081"
	log.Printf("PDP Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func evaluationHandler(w http.ResponseWriter, r *http.Request) {
	var req auth.AuthZENRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decision := auth.EvaluatePolicy(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth.AuthZENResponse{Decision: decision})
}

func evaluationsHandler(w http.ResponseWriter, r *http.Request) {
	var req auth.BatchEvaluationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var responses []auth.AuthZENResponse
	for _, resource := range req.Evaluations {
		evalReq := auth.AuthZENRequest{
			Subject:  req.Subject,
			Action:   req.Action,
			Resource: resource,
			Context:  req.Context,
		}
		decision := auth.EvaluatePolicy(evalReq)
		responses = append(responses, auth.AuthZENResponse{Decision: decision})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth.BatchEvaluationResponse{Evaluations: responses})
}
