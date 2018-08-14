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

// NewGetPeerKeyParams creates a new GetPeerKeyParams object
// with the default values initialized.
func NewGetPeerKeyParams() *GetPeerKeyParams {

	return &GetPeerKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPeerKeyParamsWithTimeout creates a new GetPeerKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPeerKeyParamsWithTimeout(timeout time.Duration) *GetPeerKeyParams {

	return &GetPeerKeyParams{

		timeout: timeout,
	}
}

// NewGetPeerKeyParamsWithContext creates a new GetPeerKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPeerKeyParamsWithContext(ctx context.Context) *GetPeerKeyParams {

	return &GetPeerKeyParams{

		Context: ctx,
	}
}

// NewGetPeerKeyParamsWithHTTPClient creates a new GetPeerKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPeerKeyParamsWithHTTPClient(client *http.Client) *GetPeerKeyParams {

	return &GetPeerKeyParams{
		HTTPClient: client,
	}
}

/*GetPeerKeyParams contains all the parameters to send to the API endpoint
for the get peer key operation typically these are written to a http.Request
*/
type GetPeerKeyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get peer key params
func (o *GetPeerKeyParams) WithTimeout(timeout time.Duration) *GetPeerKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get peer key params
func (o *GetPeerKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get peer key params
func (o *GetPeerKeyParams) WithContext(ctx context.Context) *GetPeerKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get peer key params
func (o *GetPeerKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get peer key params
func (o *GetPeerKeyParams) WithHTTPClient(client *http.Client) *GetPeerKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get peer key params
func (o *GetPeerKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPeerKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
