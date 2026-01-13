package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type webRequest struct {
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
	Action    string `json:"action"`
}

type webResponse struct {
	Allowed bool   `json:"allowed"`
	Message string `json:"message,omitempty"`
}

var allowedNamespace = []string{"prod", "staging", "dev"}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var req webRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	allowed := false
	for _, ns := range allowedNamespace {
		if ns == req.Namespace {
			allowed = true
			break
		}
	}

	resp := webResponse{Allowed: allowed,}

	if !allowed {
		resp.Message = fmt.Sprintf("namespace '%s' not allowed", req.Namespace)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func webhookHandlerGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Webhook endpoint is active",
		"method":  "GET",
		"path":    r.URL.Path,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/webhook", webhookHandler).Methods("POST")
	r.HandleFunc("/webhook", webhookHandlerGet).Methods("GET")
	http.ListenAndServe(":8080", r)
}
