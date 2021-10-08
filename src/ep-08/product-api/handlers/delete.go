package handlers

import (
	"net/http"
	"strconv"

	"github.com/chrischenyc/microservices-in-go/ep-08/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
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
