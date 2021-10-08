// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/chrischenyc/microservices-in-go/ep-08/models"
)

// NewUpdateProductParams creates a new UpdateProductParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProductParams() *UpdateProductParams {
	return &UpdateProductParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProductParamsWithTimeout creates a new UpdateProductParams object
// with the ability to set a timeout on a request.
func NewUpdateProductParamsWithTimeout(timeout time.Duration) *UpdateProductParams {
	return &UpdateProductParams{
		timeout: timeout,
	}
}

// NewUpdateProductParamsWithContext creates a new UpdateProductParams object
// with the ability to set a context for a request.
func NewUpdateProductParamsWithContext(ctx context.Context) *UpdateProductParams {
	return &UpdateProductParams{
		Context: ctx,
	}
}

// NewUpdateProductParamsWithHTTPClient creates a new UpdateProductParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProductParamsWithHTTPClient(client *http.Client) *UpdateProductParams {
	return &UpdateProductParams{
		HTTPClient: client,
	}
}

/* UpdateProductParams contains all the parameters to send to the API endpoint
   for the update product operation.

   Typically these are written to a http.Request.
*/
type UpdateProductParams struct {

	/* Body.

	   Product data structure to Update or Create
	*/
	Body *models.Product

	/* ID.

	   The id of the product

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProductParams) WithDefaults() *UpdateProductParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update product params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProductParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update product params
func (o *UpdateProductParams) WithTimeout(timeout time.Duration) *UpdateProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update product params
func (o *UpdateProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update product params
func (o *UpdateProductParams) WithContext(ctx context.Context) *UpdateProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update product params
func (o *UpdateProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update product params
func (o *UpdateProductParams) WithHTTPClient(client *http.Client) *UpdateProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update product params
func (o *UpdateProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update product params
func (o *UpdateProductParams) WithBody(body *models.Product) *UpdateProductParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update product params
func (o *UpdateProductParams) SetBody(body *models.Product) {
	o.Body = body
}

// WithID adds the id to the update product params
func (o *UpdateProductParams) WithID(id int64) *UpdateProductParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update product params
func (o *UpdateProductParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
