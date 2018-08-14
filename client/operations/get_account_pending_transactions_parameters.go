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

// NewGetAccountPendingTransactionsParams creates a new GetAccountPendingTransactionsParams object
// with the default values initialized.
func NewGetAccountPendingTransactionsParams() *GetAccountPendingTransactionsParams {
	var ()
	return &GetAccountPendingTransactionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountPendingTransactionsParamsWithTimeout creates a new GetAccountPendingTransactionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountPendingTransactionsParamsWithTimeout(timeout time.Duration) *GetAccountPendingTransactionsParams {
	var ()
	return &GetAccountPendingTransactionsParams{

		timeout: timeout,
	}
}

// NewGetAccountPendingTransactionsParamsWithContext creates a new GetAccountPendingTransactionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountPendingTransactionsParamsWithContext(ctx context.Context) *GetAccountPendingTransactionsParams {
	var ()
	return &GetAccountPendingTransactionsParams{

		Context: ctx,
	}
}

// NewGetAccountPendingTransactionsParamsWithHTTPClient creates a new GetAccountPendingTransactionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountPendingTransactionsParamsWithHTTPClient(client *http.Client) *GetAccountPendingTransactionsParams {
	var ()
	return &GetAccountPendingTransactionsParams{
		HTTPClient: client,
	}
}

/*GetAccountPendingTransactionsParams contains all the parameters to send to the API endpoint
for the get account pending transactions operation typically these are written to a http.Request
*/
type GetAccountPendingTransactionsParams struct {

	/*AccountPubkey
	  Account pubkey to search for in mempool

	*/
	AccountPubkey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account pending transactions params
func (o *GetAccountPendingTransactionsParams) WithTimeout(timeout time.Duration) *GetAccountPendingTransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account pending transactions params
func (o *GetAccountPendingTransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account pending transactions params
func (o *GetAccountPendingTransactionsParams) WithContext(ctx context.Context) *GetAccountPendingTransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account pending transactions params
func (o *GetAccountPendingTransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account pending transactions params
func (o *GetAccountPendingTransactionsParams) WithHTTPClient(client *http.Client) *GetAccountPendingTransactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account pending transactions params
func (o *GetAccountPendingTransactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountPubkey adds the accountPubkey to the get account pending transactions params
func (o *GetAccountPendingTransactionsParams) WithAccountPubkey(accountPubkey string) *GetAccountPendingTransactionsParams {
	o.SetAccountPubkey(accountPubkey)
	return o
}

// SetAccountPubkey adds the accountPubkey to the get account pending transactions params
func (o *GetAccountPendingTransactionsParams) SetAccountPubkey(accountPubkey string) {
	o.AccountPubkey = accountPubkey
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountPendingTransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_pubkey
	if err := r.SetPathParam("account_pubkey", o.AccountPubkey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
