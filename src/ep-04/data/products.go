package data

import (
	"encoding/json"
	"errors"
	"io"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SKU         string `json:"sku"`
	Price       int    `json:"price,omitempty"` // in cents
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"-"`
}

type Products []*Product

func (products Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(products)
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func GetProducts() Products {
	return productsInMemory
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productsInMemory = append(productsInMemory, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProductWithId(id)

	if err != nil {
		return err
	}

	p.ID = id
	productsInMemory[pos] = p
	return nil
}

func DeleteProduct(id int) error {
	_, pos, err := findProductWithId(id)

	if err != nil {
		return err
	}

	productsInMemory = append(productsInMemory[:pos], productsInMemory[pos+1:]...)

	return nil
}

func getNextID() int {
	lastProduct := productsInMemory[len(productsInMemory)-1]
	return lastProduct.ID + 1
}

var ErrProductNotFound = errors.New("product not found")

func findProductWithId(id int) (*Product, int, error) {
	for i, v := range productsInMemory {
		if v.ID == id {
			return v, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

var productsInMemory = []*Product{
	{
		ID:          1,
		Name:        "Sobe - Lizard Fuel",
		Description: "Networked 3rd generation emulation",
		SKU:         "WAU32AFD0FN874428",
		CreatedAt:   "2021-01-04T12:34:56",
		UpdatedAt:   "2021-01-04T12:34:56",
	},
	{
		ID:          2,
		Name:        "Cherries - Bing, Canned",
		Description: "Extended hybrid hierarchy",
		Price:       100,
		SKU:         "JM1GJ1T6XF1611223",
		CreatedAt:   "2021-03-04T12:34:56",
		UpdatedAt:   "2021-03-04T12:34:56",
		DeletedAt:   "2021-04-04T12:34:56",
	},
	{
		ID:          3,
		Name:        "Syrup - Monin, Amaretta",
		Description: "Adaptive global contingency",
		Price:       200,
		SKU:         "19UUA65515A343666",
		CreatedAt:   "2021-05-04T12:34:56",
		UpdatedAt:   "2021-06-04T12:34:56",
	},
}
