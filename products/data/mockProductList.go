package data

import "time"

// dummy persistence layer
var productList = []*Product{
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
