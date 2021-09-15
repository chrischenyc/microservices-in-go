package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chrischenyc/microservices-in-go/ep-04/handlers"
)

func main() {
	apiLogger := log.New(os.Stdout, "[api] ", log.LstdFlags)

	serveMux := http.NewServeMux()

	serveMux.Handle("/products/", handlers.NewProducts(apiLogger))

	server := http.Server{
		Addr:         ":3000",
		Handler:      serveMux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// listening to port in a go routine, so it won't block the rest
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			apiLogger.Fatal(err)
		}
	}()

	osSignalChan := make(chan os.Signal, 1)
	// use syscall.SIGTERM instead of os.Kill: https://github.com/braintree/manners/issues/45
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGTERM)

	sig := <-osSignalChan
	apiLogger.Print("received OS termination signal, gracefully shut down", sig)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(timeoutContext)
}
