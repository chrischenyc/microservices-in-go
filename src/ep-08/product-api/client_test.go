package main

import (
	"testing"

	"github.com/chrischenyc/microservices-in-go/ep-08/client"
	"github.com/chrischenyc/microservices-in-go/ep-08/client/products"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestAPIClient(t *testing.T) {
	apiClient := client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{
		Host: "localhost:3000",
	})

	listResult, err := apiClient.Products.ListProducts(products.NewListProductsParams())
	assert.NoError(t, err)
	assert.Len(t, listResult.Payload, 3)

}
