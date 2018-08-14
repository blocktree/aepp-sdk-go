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

// NewPostSpendParams creates a new PostSpendParams object
// with the default values initialized.
func NewPostSpendParams() *PostSpendParams {
	var ()
	return &PostSpendParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSpendParamsWithTimeout creates a new PostSpendParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSpendParamsWithTimeout(timeout time.Duration) *PostSpendParams {
	var ()
	return &PostSpendParams{

		timeout: timeout,
	}
}

// NewPostSpendParamsWithContext creates a new PostSpendParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSpendParamsWithContext(ctx context.Context) *PostSpendParams {
	var ()
	return &PostSpendParams{

		Context: ctx,
	}
}

// NewPostSpendParamsWithHTTPClient creates a new PostSpendParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSpendParamsWithHTTPClient(client *http.Client) *PostSpendParams {
	var ()
	return &PostSpendParams{
		HTTPClient: client,
	}
}

/*PostSpendParams contains all the parameters to send to the API endpoint
for the post spend operation typically these are written to a http.Request
*/
type PostSpendParams struct {

	/*Body
	  A spend transaction

	*/
	Body *models.SpendTx

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post spend params
func (o *PostSpendParams) WithTimeout(timeout time.Duration) *PostSpendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post spend params
func (o *PostSpendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post spend params
func (o *PostSpendParams) WithContext(ctx context.Context) *PostSpendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post spend params
func (o *PostSpendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post spend params
func (o *PostSpendParams) WithHTTPClient(client *http.Client) *PostSpendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post spend params
func (o *PostSpendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post spend params
func (o *PostSpendParams) WithBody(body *models.SpendTx) *PostSpendParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post spend params
func (o *PostSpendParams) SetBody(body *models.SpendTx) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostSpendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
