package data

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/rjNemo/go-micro/products/models"
)

// Products is the collection of products.
// It encapsulates data access logic
type Products []*models.Product

// ToJSON returns all existing product in JSON format
func (p *Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p) // more efficient in memory and time than Marshal
}

// AllProducts returns all existing products
func AllProducts() Products {
	return productList
}

// AddProduct add a Product to the dataStore
func AddProduct(p *models.Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct edits a Product identified by its id
func UpdateProduct(id int, p *models.Product) error {
	idx, _, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[idx] = p
	return nil
}

// DeleteProduct removes a Product idntified by its id
func DeleteProduct(id int) error {
	idx, _, err := findProduct(id)
	if err != nil {
		return err
	}
	productList = append(productList[0:idx], productList[idx+1:]...)
	return nil
}

// ErrorProductNotFound its thrown when the product is not found
var ErrorProductNotFound = fmt.Errorf("Product not found")

// findProduct retrieves a product via its unique identifier
func findProduct(id int) (int, *models.Product, error) {
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
