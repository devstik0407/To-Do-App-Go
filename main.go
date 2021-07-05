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
	// log.Print(ctx.Err())

	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Failed:%+v", err)
	} else {
		log.Print("Server Exited Properly")
	}
}
