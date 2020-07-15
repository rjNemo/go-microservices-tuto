package models

import (
	"encoding/json"
	"io"
	"regexp"

	"github.com/go-playground/validator"
)

// Product defines the structure of a product
// swagger:model
type Product struct {
	// the id for this product
	//
	// required: true
	// min: 1
	ID int `json:"id"` //TODO: use uuid
	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`
	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`
	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"` // TODO: use int
	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU       string `json:"sku" validate:"required,sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

// FromJSON read JSON data to create a new product
func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

// ToJSON convert product to JSON
func (p *Product) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
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
