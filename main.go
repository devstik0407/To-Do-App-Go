package main

import (
	"fmt"
	"net/http"
	"todo/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Welcome to homepage")
	})

	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")

	r.HandleFunc("/todos", handlers.CreateTaskList).Methods("POST")

	r.HandleFunc("/todos/{listId}", handlers.AddTask).Methods("POST")

	r.HandleFunc("/todos/{listId}", handlers.DeleteTaskList).Methods("DELETE")

	r.HandleFunc("/todos/{listId}/{taskId}", handlers.DeleteTask).Methods("DELETE")

	r.HandleFunc("/todos/{listId}/{taskId}", handlers.UpdateTask).Methods("PUT")

	http.ListenAndServe(":8081", r)
}
