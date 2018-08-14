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

// NewPostContractCallParams creates a new PostContractCallParams object
// with the default values initialized.
func NewPostContractCallParams() *PostContractCallParams {
	var ()
	return &PostContractCallParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostContractCallParamsWithTimeout creates a new PostContractCallParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostContractCallParamsWithTimeout(timeout time.Duration) *PostContractCallParams {
	var ()
	return &PostContractCallParams{

		timeout: timeout,
	}
}

// NewPostContractCallParamsWithContext creates a new PostContractCallParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostContractCallParamsWithContext(ctx context.Context) *PostContractCallParams {
	var ()
	return &PostContractCallParams{

		Context: ctx,
	}
}

// NewPostContractCallParamsWithHTTPClient creates a new PostContractCallParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostContractCallParamsWithHTTPClient(client *http.Client) *PostContractCallParams {
	var ()
	return &PostContractCallParams{
		HTTPClient: client,
	}
}

/*PostContractCallParams contains all the parameters to send to the API endpoint
for the post contract call operation typically these are written to a http.Request
*/
type PostContractCallParams struct {

	/*Body*/
	Body *models.ContractCallData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post contract call params
func (o *PostContractCallParams) WithTimeout(timeout time.Duration) *PostContractCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post contract call params
func (o *PostContractCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post contract call params
func (o *PostContractCallParams) WithContext(ctx context.Context) *PostContractCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post contract call params
func (o *PostContractCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post contract call params
func (o *PostContractCallParams) WithHTTPClient(client *http.Client) *PostContractCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post contract call params
func (o *PostContractCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post contract call params
func (o *PostContractCallParams) WithBody(body *models.ContractCallData) *PostContractCallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post contract call params
func (o *PostContractCallParams) SetBody(body *models.ContractCallData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostContractCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
