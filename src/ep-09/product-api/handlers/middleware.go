package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/chrischenyc/microservices-in-go/data"
)

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
