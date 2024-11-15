// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateProductReader is a Reader for the UpdateProduct structure.
type UpdateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateProductNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateProductUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateProductNoContent creates a UpdateProductNoContent with default headers values
func NewUpdateProductNoContent() *UpdateProductNoContent {
	return &UpdateProductNoContent{}
}

/*UpdateProductNoContent handles this case with default header values.

empty response
*/
type UpdateProductNoContent struct {
}

func (o *UpdateProductNoContent) Error() string {
	return fmt.Sprintf("[PUT /products/{id}][%d] updateProductNoContent ", 204)
}

func (o *UpdateProductNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductNotFound creates a UpdateProductNotFound with default headers values
func NewUpdateProductNotFound() *UpdateProductNotFound {
	return &UpdateProductNotFound{}
}

/*UpdateProductNotFound handles this case with default header values.

Generic error message returned as a string
*/
type UpdateProductNotFound struct {
}

func (o *UpdateProductNotFound) Error() string {
	return fmt.Sprintf("[PUT /products/{id}][%d] updateProductNotFound ", 404)
}

func (o *UpdateProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductUnprocessableEntity creates a UpdateProductUnprocessableEntity with default headers values
func NewUpdateProductUnprocessableEntity() *UpdateProductUnprocessableEntity {
	return &UpdateProductUnprocessableEntity{}
}

/*UpdateProductUnprocessableEntity handles this case with default header values.

Validation errors defined as an array of strings
*/
type UpdateProductUnprocessableEntity struct {
}

func (o *UpdateProductUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /products/{id}][%d] updateProductUnprocessableEntity ", 422)
}

func (o *UpdateProductUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
