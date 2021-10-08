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

	"github.com/chrischenyc/microservices-in-go/handlers"
	"github.com/go-openapi/runtime/middleware"
	gorilla_handlers "github.com/gorilla/handlers"
)

func main() {
	apiLogger := log.New(os.Stdout, "[api] ", log.LstdFlags)

	root := mux.NewRouter()

	// API routes
	productsRouter := root.PathPrefix("/products").Subrouter()
	configProductsRouter(productsRouter, apiLogger)

	// API docs routes
	root.Handle("/swagger.yaml", http.FileServer(http.Dir("./"))).Methods(http.MethodGet)
	root.Handle("/docs", middleware.SwaggerUI(middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}, nil)).Methods(http.MethodGet)

	// middleware
	root.Use(func(next http.Handler) http.Handler {
		return gorilla_handlers.CombinedLoggingHandler(os.Stdout, next)
	})
	corsHandler := gorilla_handlers.CORS(gorilla_handlers.AllowedOrigins([]string{"*"}))(root)
	compressHandler := gorilla_handlers.CompressHandler(corsHandler)

	// serve
	server := http.Server{
		Addr:         ":3000",
		Handler:      compressHandler,
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

func configProductsRouter(productsRouter *mux.Router, logger *log.Logger) {
	productsHandler := handlers.NewProducts(logger)

	productsRouter.HandleFunc("", productsHandler.GetProducts).Methods(http.MethodGet)

	// ??? how to easily use middleware for certain path/method? do I have to create a sub router?
	// ??? can't I just use r.HandleFunc("/path", handler).Methods(http.MethodPut).use(middleware)?
	productsPostRouter := productsRouter.Methods(http.MethodPost).Subrouter()
	productsPostRouter.Use(productsHandler.RequireProduct)
	productsPostRouter.HandleFunc("", productsHandler.AddProducts)

	// ??? how to easily use middleware for certain path/method? do I have to create a sub router?
	// ??? can't I just use r.HandleFunc("/path", handler).Methods(http.MethodPut).use(middleware)?
	productsPutRouter := productsRouter.Methods(http.MethodPut).Subrouter()
	productsPutRouter.Use(productsHandler.RequireProduct)
	productsPutRouter.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProduct).Methods(http.MethodPut)

	productsRouter.HandleFunc("/{id:[0-9]+}", productsHandler.DeleteProduct).Methods(http.MethodDelete)
}
