package handlers

import (
	"net/http"

	"github.com/chrischenyc/microservices-in-go/data"
)

// swagger:route POST /products products createProduct
// Create a product
//
// Responses:
// 	202: noContent
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(ContextKeyProduct{}).(*data.Product)

	data.AddProduct(product)

	rw.WriteHeader(http.StatusCreated)
}
