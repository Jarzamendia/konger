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
	"github.com/jarzamendia/konger/kong"
)

func main() {

	kong.GetPlugins()

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	r.HandleFunc("/services", controllers.ServicesHandler).Methods("GET")
	r.HandleFunc("/servicesbyid/{ID}", controllers.ServicesIDHandler).Methods("GET")
	r.HandleFunc("/servicesbyname/{Name}", controllers.ServicesNameHandler).Methods("GET")
	r.HandleFunc("/consumers", controllers.ConsumersHandler).Methods("GET")
	r.HandleFunc("/consumersbyid/{ID}", controllers.ConsumersIDHandler).Methods("GET")
	r.HandleFunc("/consumersbyname/{Name}", controllers.ConsumersNameHandler).Methods("GET")
	r.HandleFunc("/routesbyserviceid/{ID}", controllers.RoutesIDHandler).Methods("GET")
	r.HandleFunc("/routesbyservicename/{Name}", controllers.RoutesNameHandler).Methods("GET")
	r.HandleFunc("/plugins", controllers.PluginsHandler).Methods("GET")

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
