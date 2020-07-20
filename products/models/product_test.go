package models

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "Frappe",
		Price: 1.1,
		SKU:   "a-adfg-fdds",
	}

	err := p.Validate()

	if err != nil {
		t.Error(err)
	}
}
