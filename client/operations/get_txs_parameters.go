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

// NewGetTxsParams creates a new GetTxsParams object
// with the default values initialized.
func NewGetTxsParams() *GetTxsParams {

	return &GetTxsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTxsParamsWithTimeout creates a new GetTxsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTxsParamsWithTimeout(timeout time.Duration) *GetTxsParams {

	return &GetTxsParams{

		timeout: timeout,
	}
}

// NewGetTxsParamsWithContext creates a new GetTxsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTxsParamsWithContext(ctx context.Context) *GetTxsParams {

	return &GetTxsParams{

		Context: ctx,
	}
}

// NewGetTxsParamsWithHTTPClient creates a new GetTxsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTxsParamsWithHTTPClient(client *http.Client) *GetTxsParams {

	return &GetTxsParams{
		HTTPClient: client,
	}
}

/*GetTxsParams contains all the parameters to send to the API endpoint
for the get txs operation typically these are written to a http.Request
*/
type GetTxsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get txs params
func (o *GetTxsParams) WithTimeout(timeout time.Duration) *GetTxsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get txs params
func (o *GetTxsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get txs params
func (o *GetTxsParams) WithContext(ctx context.Context) *GetTxsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get txs params
func (o *GetTxsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get txs params
func (o *GetTxsParams) WithHTTPClient(client *http.Client) *GetTxsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get txs params
func (o *GetTxsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTxsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
