package models

import (
	"encoding/json"
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
