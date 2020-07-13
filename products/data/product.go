package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure of a product
type Product struct {
	ID          int     `json:"id"` //TODO: use uuid
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"` // TODO: use int
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is the collection of products.
// It encapsulates data access logic
type Products []*Product

// All returns all existing products
func AllProducts() Products {
	return productList
}

// ToJSON returns all existing product in JSON format
func (p *Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p) // more efficient in memory and time than Marshal
}

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
