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

	models "github.com/aeternity/aepp-sdk-go/models"
)

// NewPostNameClaimTxParams creates a new PostNameClaimTxParams object
// with the default values initialized.
func NewPostNameClaimTxParams() *PostNameClaimTxParams {
	var ()
	return &PostNameClaimTxParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNameClaimTxParamsWithTimeout creates a new PostNameClaimTxParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNameClaimTxParamsWithTimeout(timeout time.Duration) *PostNameClaimTxParams {
	var ()
	return &PostNameClaimTxParams{

		timeout: timeout,
	}
}

// NewPostNameClaimTxParamsWithContext creates a new PostNameClaimTxParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNameClaimTxParamsWithContext(ctx context.Context) *PostNameClaimTxParams {
	var ()
	return &PostNameClaimTxParams{

		Context: ctx,
	}
}

// NewPostNameClaimTxParamsWithHTTPClient creates a new PostNameClaimTxParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNameClaimTxParamsWithHTTPClient(client *http.Client) *PostNameClaimTxParams {
	var ()
	return &PostNameClaimTxParams{
		HTTPClient: client,
	}
}

/*PostNameClaimTxParams contains all the parameters to send to the API endpoint
for the post name claim tx operation typically these are written to a http.Request
*/
type PostNameClaimTxParams struct {

	/*Body
	  Creates new name claim transaction

	*/
	Body *models.NameClaimTx

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post name claim tx params
func (o *PostNameClaimTxParams) WithTimeout(timeout time.Duration) *PostNameClaimTxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post name claim tx params
func (o *PostNameClaimTxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post name claim tx params
func (o *PostNameClaimTxParams) WithContext(ctx context.Context) *PostNameClaimTxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post name claim tx params
func (o *PostNameClaimTxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post name claim tx params
func (o *PostNameClaimTxParams) WithHTTPClient(client *http.Client) *PostNameClaimTxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post name claim tx params
func (o *PostNameClaimTxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post name claim tx params
func (o *PostNameClaimTxParams) WithBody(body *models.NameClaimTx) *PostNameClaimTxParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post name claim tx params
func (o *PostNameClaimTxParams) SetBody(body *models.NameClaimTx) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostNameClaimTxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
