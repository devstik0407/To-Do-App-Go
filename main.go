package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"todo/handlers"
	mongostore "todo/mongo-store"
	"todo/todos"

	"github.com/gorilla/mux"
)

var (
	dataStore   mongostore.MongoDB
	todoService todos.Service
	handler     handlers.Handler
)

func init() {
	ctx, _ := context.WithCancel(context.Background())
	dataStore = mongostore.Connect(ctx)
	todoService = todos.Service{
		DataStore: dataStore,
	}
	handler = handlers.Handler{
		Service: todoService,
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Welcome to homepage")
	})
	r.HandleFunc("/todos", handler.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handler.CreateTaskList).Methods("POST")
	r.HandleFunc("/todos/{listId}", handler.AddTask).Methods("POST")
	r.HandleFunc("/todos/{listId}", handler.DeleteTaskList).Methods("DELETE")
	r.HandleFunc("/todos/{listId}/{taskId}", handler.DeleteTask).Methods("DELETE")
	r.HandleFunc("/todos/{listId}/{taskId}", handler.UpdateTask).Methods("PUT")

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")

	<-ctx.Done()
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Failed:%+v", err)
	} else {
		log.Print("Server Exited Properly")
	}
}
