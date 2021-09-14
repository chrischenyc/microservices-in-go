package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HelloHandler struct {
	logger *log.Logger
}

func NewHelloHandler(logger *log.Logger) *HelloHandler {
	return &HelloHandler{
		logger,
	}
}

// implements http.Handler interface
func (hello *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	hello.logger.Printf("method: %s", r.Method)
	hello.logger.Printf("path: %s", r.URL.Path)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "someting went wrong", http.StatusBadRequest)
		return
	}
	hello.logger.Printf("body: %s", body)

	fmt.Fprintf(rw, "requested with\nmethod: %s\npath: %s\nbody: %s\n", r.Method, r.URL.Path, body)
}
