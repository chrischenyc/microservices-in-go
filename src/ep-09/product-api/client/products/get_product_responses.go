// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chrischenyc/microservices-in-go/models"
)

// GetProductReader is a Reader for the GetProduct structure.
type GetProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProductInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProductOK creates a GetProductOK with default headers values
func NewGetProductOK() *GetProductOK {
	return &GetProductOK{}
}

/* GetProductOK describes a response with status code 200, with default header values.

GetProductOK get product o k
*/
type GetProductOK struct {
	Payload *models.Product
}

func (o *GetProductOK) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductOK  %+v", 200, o.Payload)
}
func (o *GetProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *GetProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductNotFound creates a GetProductNotFound with default headers values
func NewGetProductNotFound() *GetProductNotFound {
	return &GetProductNotFound{}
}

/* GetProductNotFound describes a response with status code 404, with default header values.

GetProductNotFound get product not found
*/
type GetProductNotFound struct {
	Payload *models.GenericError
}

func (o *GetProductNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductNotFound  %+v", 404, o.Payload)
}
func (o *GetProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductInternalServerError creates a GetProductInternalServerError with default headers values
func NewGetProductInternalServerError() *GetProductInternalServerError {
	return &GetProductInternalServerError{}
}

/* GetProductInternalServerError describes a response with status code 500, with default header values.

GetProductInternalServerError get product internal server error
*/
type GetProductInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetProductInternalServerError) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] getProductInternalServerError  %+v", 500, o.Payload)
}
func (o *GetProductInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetProductInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
