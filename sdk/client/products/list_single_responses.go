// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"coffe-api/sdk/models"
)

// ListSingleReader is a Reader for the ListSingle structure.
type ListSingleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSingleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSingleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListSingleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSingleOK creates a ListSingleOK with default headers values
func NewListSingleOK() *ListSingleOK {
	return &ListSingleOK{}
}

/*ListSingleOK handles this case with default header values.

Data structure representing a single product
*/
type ListSingleOK struct {
	Payload *models.Product
}

func (o *ListSingleOK) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] listSingleOK  %+v", 200, o.Payload)
}

func (o *ListSingleOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *ListSingleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSingleNotFound creates a ListSingleNotFound with default headers values
func NewListSingleNotFound() *ListSingleNotFound {
	return &ListSingleNotFound{}
}

/*ListSingleNotFound handles this case with default header values.

Generic error message returned as a string
*/
type ListSingleNotFound struct {
	Payload *models.GenericError
}

func (o *ListSingleNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] listSingleNotFound  %+v", 404, o.Payload)
}

func (o *ListSingleNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *ListSingleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
