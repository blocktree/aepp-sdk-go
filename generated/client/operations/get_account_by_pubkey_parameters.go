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

// NewGetAccountByPubkeyParams creates a new GetAccountByPubkeyParams object
// with the default values initialized.
func NewGetAccountByPubkeyParams() *GetAccountByPubkeyParams {
	var ()
	return &GetAccountByPubkeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountByPubkeyParamsWithTimeout creates a new GetAccountByPubkeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountByPubkeyParamsWithTimeout(timeout time.Duration) *GetAccountByPubkeyParams {
	var ()
	return &GetAccountByPubkeyParams{

		timeout: timeout,
	}
}

// NewGetAccountByPubkeyParamsWithContext creates a new GetAccountByPubkeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountByPubkeyParamsWithContext(ctx context.Context) *GetAccountByPubkeyParams {
	var ()
	return &GetAccountByPubkeyParams{

		Context: ctx,
	}
}

// NewGetAccountByPubkeyParamsWithHTTPClient creates a new GetAccountByPubkeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountByPubkeyParamsWithHTTPClient(client *http.Client) *GetAccountByPubkeyParams {
	var ()
	return &GetAccountByPubkeyParams{
		HTTPClient: client,
	}
}

/*GetAccountByPubkeyParams contains all the parameters to send to the API endpoint
for the get account by pubkey operation typically these are written to a http.Request
*/
type GetAccountByPubkeyParams struct {

	/*Pubkey
	  The public key of the account

	*/
	Pubkey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get account by pubkey params
func (o *GetAccountByPubkeyParams) WithTimeout(timeout time.Duration) *GetAccountByPubkeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account by pubkey params
func (o *GetAccountByPubkeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account by pubkey params
func (o *GetAccountByPubkeyParams) WithContext(ctx context.Context) *GetAccountByPubkeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account by pubkey params
func (o *GetAccountByPubkeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account by pubkey params
func (o *GetAccountByPubkeyParams) WithHTTPClient(client *http.Client) *GetAccountByPubkeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account by pubkey params
func (o *GetAccountByPubkeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPubkey adds the pubkey to the get account by pubkey params
func (o *GetAccountByPubkeyParams) WithPubkey(pubkey string) *GetAccountByPubkeyParams {
	o.SetPubkey(pubkey)
	return o
}

// SetPubkey adds the pubkey to the get account by pubkey params
func (o *GetAccountByPubkeyParams) SetPubkey(pubkey string) {
	o.Pubkey = pubkey
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountByPubkeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pubkey
	if err := r.SetPathParam("pubkey", o.Pubkey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
