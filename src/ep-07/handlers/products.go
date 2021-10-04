// Package classification of Product API
//
// Schemes: http
// BashPath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/chrischenyc/microservices-in-go/ep-07/data"
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

// swagger:response productsResponse
type productsResponseWrapper struct {
	// in: body
	Body []data.Product
}

// swagger:parameters delete update
type productIDParameterWrapper struct {
	// The id of the product
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:route GET /products products list
// Returns a list of products
//
// Responses:
// 	200: productsResponse
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unexpected error", http.StatusInternalServerError)
	}
}

// swagger:route POST /products products create
// Create a product
//
// Responses:
// 	202: noContent
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(ContextKeyProduct{}).(*data.Product)

	data.AddProduct(product)

	rw.WriteHeader(http.StatusCreated)
}

// swagger:route PUT /products/{id} products update
// Update a product by its id
//
// Responses:
// 	202: noContent
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

// swagger:route DELETE /products/{id} products delete
// Delete a product by its id
//
// Responses:
// 	202: noContent
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

		// validate
		err = product.Validate()
		if err != nil {
			msg := fmt.Sprintf("Error validating product:\n%s", err.Error())
			p.logger.Println(msg)
			http.Error(rw, msg, http.StatusBadRequest)
			return
		}

		newContext := context.WithValue(r.Context(), ContextKeyProduct{}, product)
		newRequest := r.Clone(newContext)

		next.ServeHTTP(rw, newRequest)
	})
}
