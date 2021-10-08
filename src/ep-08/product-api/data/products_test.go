package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductValidator(t *testing.T) {
	t.Run("it should not return error for Procut with all required fields", func(t *testing.T) {
		p := &Product{
			Name:  "demo",
			Price: 1,
			SKU:   "food-veg-banana",
		}

		err := p.Validate()

		assert.NoError(t, err)
	})

	t.Run("it should return error for Procut without name", func(t *testing.T) {
		p := &Product{}

		err := p.Validate()

		assert.Error(t, err)
	})

	t.Run("it should return error for Procut with invalid sku", func(t *testing.T) {
		p := &Product{
			Name:  "demo",
			Price: 1,
			SKU:   "food-veg-123",
		}

		err := p.Validate()

		assert.Error(t, err)
	})

}
