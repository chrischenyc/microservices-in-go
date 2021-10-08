// NOTE: types defined here are for swagger doc generation

package handlers

import "github.com/chrischenyc/microservices-in-go/data"

// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products
	// in: body
	Body []data.Product
}

// swagger:response productResponse
type productResponseWrapper struct {
	// Product detail
	// in: body
	Body data.Product
}

// swagger:response noContent
type noContentWrapper struct {
}

// swagger:response errorResponse
type errorResponseWrapper struct {
	// error description
	// in: body
	Body GenericError
}

type GenericError struct {
	Message   string `json:"message"`
	ErrorCode string `json:"code"`
}

// swagger:parameters deleteProduct updateProduct getProduct
type productIDParameterWrapper struct {
	// The id of the product
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters updateProduct createProduct
type productParameterWrapper struct {
	// Product data structure to Update or Create
	// in: body
	// required: true
	Body data.Product
}
