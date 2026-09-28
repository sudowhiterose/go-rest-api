package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{
	{ID: 1, Title: "GO", Done: false},
	{ID: 2, Title: "jerk off", Done: true},
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var newTask Task
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newTask.ID = len(tasks) + 1
		tasks = append(tasks, newTask)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)

	}
}

func main() {
	http.HandleFunc("GET /tasks", handleTasks)
	http.HandleFunc("POST /tasks", handleTasks)

	println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
