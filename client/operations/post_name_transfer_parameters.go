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

// NewPostNameTransferParams creates a new PostNameTransferParams object
// with the default values initialized.
func NewPostNameTransferParams() *PostNameTransferParams {
	var ()
	return &PostNameTransferParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNameTransferParamsWithTimeout creates a new PostNameTransferParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNameTransferParamsWithTimeout(timeout time.Duration) *PostNameTransferParams {
	var ()
	return &PostNameTransferParams{

		timeout: timeout,
	}
}

// NewPostNameTransferParamsWithContext creates a new PostNameTransferParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNameTransferParamsWithContext(ctx context.Context) *PostNameTransferParams {
	var ()
	return &PostNameTransferParams{

		Context: ctx,
	}
}

// NewPostNameTransferParamsWithHTTPClient creates a new PostNameTransferParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNameTransferParamsWithHTTPClient(client *http.Client) *PostNameTransferParams {
	var ()
	return &PostNameTransferParams{
		HTTPClient: client,
	}
}

/*PostNameTransferParams contains all the parameters to send to the API endpoint
for the post name transfer operation typically these are written to a http.Request
*/
type PostNameTransferParams struct {

	/*Body*/
	Body *models.NameTransferTx

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post name transfer params
func (o *PostNameTransferParams) WithTimeout(timeout time.Duration) *PostNameTransferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post name transfer params
func (o *PostNameTransferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post name transfer params
func (o *PostNameTransferParams) WithContext(ctx context.Context) *PostNameTransferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post name transfer params
func (o *PostNameTransferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post name transfer params
func (o *PostNameTransferParams) WithHTTPClient(client *http.Client) *PostNameTransferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post name transfer params
func (o *PostNameTransferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post name transfer params
func (o *PostNameTransferParams) WithBody(body *models.NameTransferTx) *PostNameTransferParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post name transfer params
func (o *PostNameTransferParams) SetBody(body *models.NameTransferTx) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostNameTransferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
