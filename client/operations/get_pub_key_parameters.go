// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPubKeyParams creates a new GetPubKeyParams object
// with the default values initialized.
func NewGetPubKeyParams() *GetPubKeyParams {

	return &GetPubKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPubKeyParamsWithTimeout creates a new GetPubKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPubKeyParamsWithTimeout(timeout time.Duration) *GetPubKeyParams {

	return &GetPubKeyParams{

		timeout: timeout,
	}
}

// NewGetPubKeyParamsWithContext creates a new GetPubKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPubKeyParamsWithContext(ctx context.Context) *GetPubKeyParams {

	return &GetPubKeyParams{

		Context: ctx,
	}
}

// NewGetPubKeyParamsWithHTTPClient creates a new GetPubKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPubKeyParamsWithHTTPClient(client *http.Client) *GetPubKeyParams {

	return &GetPubKeyParams{
		HTTPClient: client,
	}
}

/*GetPubKeyParams contains all the parameters to send to the API endpoint
for the get pub key operation typically these are written to a http.Request
*/
type GetPubKeyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pub key params
func (o *GetPubKeyParams) WithTimeout(timeout time.Duration) *GetPubKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pub key params
func (o *GetPubKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pub key params
func (o *GetPubKeyParams) WithContext(ctx context.Context) *GetPubKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pub key params
func (o *GetPubKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pub key params
func (o *GetPubKeyParams) WithHTTPClient(client *http.Client) *GetPubKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pub key params
func (o *GetPubKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPubKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
