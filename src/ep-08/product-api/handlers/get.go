package handlers

import (
	"net/http"
	"strconv"

	"github.com/chrischenyc/microservices-in-go/ep-08/data"
	"github.com/gorilla/mux"
)

// swagger:route GET /products products listProducts
// Return a list of products from the datastore
//
// Responses:
// 	200: productsResponse
//  500: errorResponse
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	products := data.GetProducts()

	err := products.ToJSON(rw)
	if err != nil {
		p.logger.Println("[ERROR] serializing products", err)
		http.Error(rw, "unexpected error", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products getProduct
// Return product details
//
// Responses:
// 	200: productResponse
//	404: errorResponse
//	500: errorResponse
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "invalid product id", http.StatusBadRequest)
		return
	}

	product := data.GetProductByID(id)

	if product == nil {
		http.Error(rw, "not found", http.StatusNotFound)
		return
	}

	err = product.ToJSON(rw)
	if err != nil {
		p.logger.Println("[ERROR] serializing product", err)
		http.Error(rw, "unexpected error", http.StatusInternalServerError)
	}
}
