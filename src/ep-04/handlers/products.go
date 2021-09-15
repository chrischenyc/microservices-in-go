package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/chrischenyc/microservices-in-go/ep-04/data"
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
		p.getProducts(rw, r)

	case http.MethodPost:
		p.addProducts(rw, r)

	// PUT /products/:id
	case http.MethodPut:
		regex := regexp.MustCompile(`/products/([0-9]+)`)
		matches := regex.FindAllStringSubmatch(r.URL.Path, -1)
		if len(matches) != 1 ||
			len(matches[0]) != 2 {
			http.Error(rw, "invalid url query parameter, use PUT /products/:id", http.StatusBadRequest)
			return
		}
		productIDString := matches[0][1]
		productID, _ := strconv.Atoi(productIDString)

		p.updateProduct(productID, rw, r)

		// DELETE /products/:id
	case http.MethodDelete:
		regex := regexp.MustCompile(`/products/([0-9]+)`)
		matches := regex.FindAllStringSubmatch(r.URL.Path, -1)
		if len(matches) != 1 ||
			len(matches[0]) != 2 {
			http.Error(rw, "invalid url query parameter, use PUT /products/:id", http.StatusBadRequest)
			return
		}
		productIDString := matches[0][1]
		productID, _ := strconv.Atoi(productIDString)

		p.deleteProduct(productID, rw, r)

	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "unexpected error", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	product := &data.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "cannot parse request body", http.StatusBadRequest)
		return
	}

	data.AddProduct(product)

	rw.WriteHeader(http.StatusCreated)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	product := &data.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "cannot parse request body", http.StatusBadRequest)
		return
	}

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

func (p *Products) deleteProduct(id int, rw http.ResponseWriter, r *http.Request) {
	err := data.DeleteProduct(id)

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
