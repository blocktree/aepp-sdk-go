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

// NewGetContractCallFromTxParams creates a new GetContractCallFromTxParams object
// with the default values initialized.
func NewGetContractCallFromTxParams() *GetContractCallFromTxParams {
	var ()
	return &GetContractCallFromTxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetContractCallFromTxParamsWithTimeout creates a new GetContractCallFromTxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetContractCallFromTxParamsWithTimeout(timeout time.Duration) *GetContractCallFromTxParams {
	var ()
	return &GetContractCallFromTxParams{

		timeout: timeout,
	}
}

// NewGetContractCallFromTxParamsWithContext creates a new GetContractCallFromTxParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetContractCallFromTxParamsWithContext(ctx context.Context) *GetContractCallFromTxParams {
	var ()
	return &GetContractCallFromTxParams{

		Context: ctx,
	}
}

// NewGetContractCallFromTxParamsWithHTTPClient creates a new GetContractCallFromTxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetContractCallFromTxParamsWithHTTPClient(client *http.Client) *GetContractCallFromTxParams {
	var ()
	return &GetContractCallFromTxParams{
		HTTPClient: client,
	}
}

/*GetContractCallFromTxParams contains all the parameters to send to the API endpoint
for the get contract call from tx operation typically these are written to a http.Request
*/
type GetContractCallFromTxParams struct {

	/*TxHash
	  Hash of the transaction

	*/
	TxHash string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get contract call from tx params
func (o *GetContractCallFromTxParams) WithTimeout(timeout time.Duration) *GetContractCallFromTxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contract call from tx params
func (o *GetContractCallFromTxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contract call from tx params
func (o *GetContractCallFromTxParams) WithContext(ctx context.Context) *GetContractCallFromTxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contract call from tx params
func (o *GetContractCallFromTxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contract call from tx params
func (o *GetContractCallFromTxParams) WithHTTPClient(client *http.Client) *GetContractCallFromTxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contract call from tx params
func (o *GetContractCallFromTxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTxHash adds the txHash to the get contract call from tx params
func (o *GetContractCallFromTxParams) WithTxHash(txHash string) *GetContractCallFromTxParams {
	o.SetTxHash(txHash)
	return o
}

// SetTxHash adds the txHash to the get contract call from tx params
func (o *GetContractCallFromTxParams) SetTxHash(txHash string) {
	o.TxHash = txHash
}

// WriteToRequest writes these params to a swagger request
func (o *GetContractCallFromTxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tx_hash
	if err := r.SetPathParam("tx_hash", o.TxHash); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
