package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/chrischenyc/microservices-in-go/ep-05/data"
	"github.com/gorilla/mux"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{
		logger,
	}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unexpected error", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(ContextKeyProduct{}).(*data.Product)

	data.AddProduct(product)

	rw.WriteHeader(http.StatusCreated)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "invalid product id", http.StatusBadRequest)
		return
	}

	product := r.Context().Value(ContextKeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, product)

	if err == data.ErrProductNotFound {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "unknown error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "invalid product id", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "unknown error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusAccepted)
}

// Gorilla Mux middleware that
// a) deserializes HTTP request body into a Product struct object
// b) set the object in request context
// c) discontinue HTTP handling if it can't deserialize request body
type ContextKeyProduct struct{}

func (p *Products) RequireProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := &data.Product{}
		err := product.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "cannot parse request body", http.StatusBadRequest)
			return
		}

		newContext := context.WithValue(r.Context(), ContextKeyProduct{}, product)
		newRequest := r.Clone(newContext)

		next.ServeHTTP(rw, newRequest)
	})
}
