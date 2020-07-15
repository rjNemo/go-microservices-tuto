// Package classification Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// 	- application/json
//
// Produces:
// 	- application/json
// swagger:meta
package handlers

import "github.com/rjNemo/go-micro/products/models"

// list of products in the response.
// swagger:response productsResponse
type productsResponse struct {
	// All products in the datastore
	// in: body
	Body []models.Product
}

// product in the response.
// swagger:response productResponse
type productResponse struct {
	// One product in the datastore
	// in: body
	Body models.Product
}

// swagger:parameters deleteProduct updateProduct
type productIDParameter struct {
	// The ID of a product in the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// empty response
// swagger:response noContent
type productNoContent struct{}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponse struct {
	// Description of the error
	// in: body
	Body string
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidation struct {
	// Collection of the errors
	// in: body
	Body string
}
