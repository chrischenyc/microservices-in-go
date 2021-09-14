package handlers

import (
	"log"
	"net/http"
)

type GoodbyeHandler struct {
	logger *log.Logger
}

func NewGoodbyeHandler(logger *log.Logger) *GoodbyeHandler {
	return &GoodbyeHandler{
		logger,
	}
}

// implements http.Handler interface
func (goodbye *GoodbyeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Goodbye!\n"))
}
