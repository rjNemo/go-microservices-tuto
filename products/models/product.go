package models

import (
	"encoding/json"
	"io"
	"regexp"

	"github.com/go-playground/validator"
)

// Product defines the structure of a product
type Product struct {
	ID          int     `json:"id"` //TODO: use uuid
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"` // TODO: use int
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// FromJSON read JSON data to create a new product
func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

// Validate checks object validity
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-efgh-ijklm
	re := regexp.MustCompile("[a-z]+-[a-z]+-[a-z]+")
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}
