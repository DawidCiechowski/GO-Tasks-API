package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Task Struct
type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks []Task

// Get all tasks
func getTasks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks)
}

// Get a singular task
func getTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request) //Get all URL params
	for _, task := range tasks {
		if task.ID == params["id"] {
			json.NewEncoder(writer).Encode(task)
			return
		}
	}

	json.NewEncoder(writer).Encode(&Task{})
}

// Create new Task
func createTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var task Task

	_ = json.NewDecoder(request.Body).Decode(&task)
	id, _ := generateHex(100) //Mock Id
	task.ID = id
	tasks = append(tasks, task)
	json.NewEncoder(writer).Encode(task)
}

// Generate random hexadecimal value for ID
func generateHex(number int) (string, error) {
	bytes := make([]byte, number)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

// Update currently existing task
func updateTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, task := range tasks {
		if task.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Task
			_ = json.NewDecoder(request.Body).Decode(&task)
			task.ID = params["id"]
			tasks = append(tasks, task)
			json.NewEncoder(writer).Encode(task)
			return
		}
	}
}

// Get all tasks
func deleteTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, task := range tasks {
		if task.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode(tasks)
}
func main() {

	//Init Router
	router := mux.NewRouter()
	fmt.Println("Running...")
	//Mock tasks for testing purposes
	// TODO: Implement database
	tasks = append(tasks, Task{ID: "1", Name: "Finsh this program", Description: "You have to finish this program", Completed: false})

	// Routes and handlers
	router.HandleFunc("/api/v1/tasks", getTasks).Methods("GET")
	router.HandleFunc("/api/v1/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/api/v1/tasks", createTask).Methods("POST")
	router.HandleFunc("/api/v1/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/api/v1/tasks/{id}", deleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
