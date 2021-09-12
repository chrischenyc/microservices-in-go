package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chrischenyc/microservices-in-go/ep-02/handlers"
)

func main() {
	apiLogger := log.New(os.Stdout, "api", log.LstdFlags)

	serveMux := http.NewServeMux()

	serveMux.Handle("/", handlers.NewHelloHandler(apiLogger))
	serveMux.Handle("/goodbye", handlers.NewGoodbyeHandler(apiLogger))

	http.ListenAndServe(":3000", serveMux)
}
