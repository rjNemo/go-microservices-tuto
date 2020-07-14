package data

import (
	"encoding/json"
	"fmt"
	"io"
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

// FromJSON read JSON data to create a new product
func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

// Products is the collection of products.
// It encapsulates data access logic
type Products []*Product

// AllProducts returns all existing products
func AllProducts() Products {
	return productList
}

// AddProduct add a Product to the dataStore
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct edits a Product identified by its id
func UpdateProduct(id int, p *Product) error {
	idx, _, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[idx] = p
	return nil
}

// ErrorProductNotFound its thrown when the product is not found
var ErrorProductNotFound = fmt.Errorf("Product not found")

// findProduct retrieves a product via its unique identifier
func findProduct(id int) (int, *Product, error) {
	for i, p := range productList {
		if p.ID == id {
			return i, p, nil
		}
	}
	return 0, nil, ErrorProductNotFound
}

// getNextID handle ID creation
func getNextID() int {
	return productList[len(productList)-1].ID + 1
}

// ToJSON returns all existing product in JSON format
func (p *Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p) // more efficient in memory and time than Marshal
}
