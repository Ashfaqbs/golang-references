package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Person represents a simple data model for JSON input/output.
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// In-memory storage for Person records.
var (
	people []Person
	nextID = 1
)

// statusHandler returns a simple JSON status.
func statusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"status": "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		log.Println("error encoding status response:", err)
	}
}

// listPeopleHandler returns all people as JSON.
func listPeopleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println("error encoding people list:", err)
	}
}

// createPersonRequest represents the expected JSON body for creating a person.
type createPersonRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// createPersonHandler reads JSON body, creates a new Person, and returns it.
func createPersonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createPersonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Age <= 0 {
		http.Error(w, "name and age must be provided and valid", http.StatusBadRequest)
		return
	}

	person := Person{
		ID:   nextID,
		Name: req.Name,
		Age:  req.Age,
	}
	nextID++
	people = append(people, person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Println("error encoding created person:", err)
	}
}

func main() {
	// Preload some in-memory data.
	people = append(people, Person{ID: nextID, Name: "Alice", Age: 30})
	nextID++
	people = append(people, Person{ID: nextID, Name: "Bob", Age: 25})
	nextID++

	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listPeopleHandler(w, r)
		case http.MethodPost:
			createPersonHandler(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := 8080
	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("server error:", err)
	}
}
