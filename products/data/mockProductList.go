package data

import (
	"time"

	"github.com/rjNemo/go-micro/products/models"
)

// dummy persistence layer
var productList = []*models.Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Prothy Milky Coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short Strong Coffee without Milk",
		Price:       1.99,
		SKU:         "efg456",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
}
