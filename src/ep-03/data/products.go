package data

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SKU         string `json:"sku"`
	Price       int32  `json:"price,omitempty"` // in cents
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"-"`
}

type Products []*Product

func (products Products) WriteJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(products)

	return err
}

func GetProducts() Products {
	return productsInMemory
}

var productsInMemory = []*Product{
	{
		ID:          "d5ce1cfc-a82a-4300-9655-eac0bde41f71",
		Name:        "Sobe - Lizard Fuel",
		Description: "Networked 3rd generation emulation",
		SKU:         "WAU32AFD0FN874428",
		CreatedAt:   "2021-01-04T12:34:56",
		UpdatedAt:   "2021-01-04T12:34:56",
	},
	{
		ID:          "15823ea7-0a3f-4cdf-af26-2b119c113938",
		Name:        "Cherries - Bing, Canned",
		Description: "Extended hybrid hierarchy",
		Price:       100,
		SKU:         "JM1GJ1T6XF1611223",
		CreatedAt:   "2021-03-04T12:34:56",
		UpdatedAt:   "2021-03-04T12:34:56",
		DeletedAt:   "2021-04-04T12:34:56",
	},
	{
		ID:          "4e983def-db40-4334-9ade-a35d7bd1ab14",
		Name:        "Syrup - Monin, Amaretta",
		Description: "Adaptive global contingency",
		Price:       200,
		SKU:         "19UUA65515A343666",
		CreatedAt:   "2021-05-04T12:34:56",
		UpdatedAt:   "2021-06-04T12:34:56",
	},
}
