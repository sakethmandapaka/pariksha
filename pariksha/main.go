package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"Title"`
	Completed bool
}

//initilize tasks var as  a slice struct
var tasks []Task

// Get all Tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Get single Task
func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

// Add new Task
func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)

}

// Update Task
func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Task
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			tasks = append(tasks, task)
			json.NewEncoder(w).Encode(task)
			return
		}
	}

}

// Delete Task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tasks)

}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	tasks = append(tasks, Task{ID: "1", Title: "test task 1", Completed: true})
	tasks = append(tasks, Task{ID: "2", Title: "test task 2", Completed: false})
	tasks = append(tasks, Task{ID: "3", Title: "test task 3", Completed: true})

	// Route handles & endpoints
	r.HandleFunc("/GET/v1/tasks", getTasks).Methods("GET")
	r.HandleFunc("/GET/v1/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/POST/v1/tasks", createTask).Methods("POST")
	r.HandleFunc("/PUT/v1/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/DELETE/v1/tasks/{id}", deleteTask).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
