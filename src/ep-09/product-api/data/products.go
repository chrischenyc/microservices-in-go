package data

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// swagger:model
type Product struct {
	// the id of the product
	//
	// required: true
	// min: 1
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	SKU         string `json:"sku" validate:"required,sku"`
	// product price in cents
	// required: true
	// min: 0
	Price     int    `json:"price,omitempty" validate:"gt=0"` // in cents
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"-"`
}

type Products []*Product

func (products Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(products)
}

func (product *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(product)
}

func (product *Product) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(product)
}

func (p *Product) Validate() error {
	productValidator := validator.New()
	productValidator.RegisterValidation("sku", func(fl validator.FieldLevel) bool {
		// sku format abc-def-ghi
		regx := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
		matches := regx.FindAllString(fl.Field().String(), -1)

		return len(matches) == 1
	})
	return productValidator.Struct(p)
}

func GetProducts() Products {
	return productsInMemory
}

func GetProductByID(id int) *Product {
	for _, product := range productsInMemory {
		if product.ID == id {
			return product
		}
	}

	return nil
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
