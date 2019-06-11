package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jarzamendia/konger/controllers"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")

	r.HandleFunc("/service/all", controllers.ServicesHandler).Methods("GET")
	r.HandleFunc("/service/id", controllers.ServicesIDHandler).Methods("POST")
	r.HandleFunc("/service/name", controllers.ServicesNameHandler).Methods("POST")

	r.HandleFunc("/consumer/all", controllers.ConsumersHandler).Methods("GET")
	r.HandleFunc("/consumer/id", controllers.ConsumersIDHandler).Methods("POST")
	r.HandleFunc("/consumer/name", controllers.ConsumersNameHandler).Methods("POST")
	r.HandleFunc("/consumer/create", controllers.ConsumersCreateNameHandler).Methods("POST")

	r.HandleFunc("/route/id", controllers.RoutesIDHandler).Methods("POST")
	r.HandleFunc("/route/name", controllers.RoutesNameHandler).Methods("POST")

	r.HandleFunc("/plugin/all", controllers.PluginsHandler).Methods("GET")

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8081",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start Server
	go func() {

		log.Println("Starting Server")

		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)

}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
