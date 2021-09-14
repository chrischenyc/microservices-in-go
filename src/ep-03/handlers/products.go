package handlers

import (
	"log"
	"net/http"

	"github.com/chrischenyc/microservices-in-go/ep-03/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{
		logger,
	}
}

// implement http.Handler interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		getProducts(rw, r)

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.WriteJSON(rw)

	if err != nil {
		http.Error(rw, "unexpected error", http.StatusInternalServerError)
	}
}
