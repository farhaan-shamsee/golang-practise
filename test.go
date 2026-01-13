package main

import (
	"encoding/json"

	"log"
	"net/http"
	"sync"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{Name: "John Doe", Email: "john@example.com"},
	{Name: "Jane Smith", Email: "jane@example.com"},
	{Name: "Bob Johnson", Email: "bob@example.com"},
}

var mut = sync.RWMutex{}

func main() {

	r := http.NewServeMux()

	r.HandleFunc("GET /", handleHome)
	r.HandleFunc("GET /users", handleGetUsers)
	r.HandleFunc("POST /user", handleCreateSingleUser)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Application failed: %s", err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to user home"))
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application-json")
	mut.Lock()
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}
	mut.Unlock()
}

func handleCreateSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var newUser = User{}
	json.NewDecoder(r.Body).Decode(&newUser)
}

