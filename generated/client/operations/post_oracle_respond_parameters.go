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

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// NewPostOracleRespondParams creates a new PostOracleRespondParams object
// with the default values initialized.
func NewPostOracleRespondParams() *PostOracleRespondParams {
	var ()
	return &PostOracleRespondParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOracleRespondParamsWithTimeout creates a new PostOracleRespondParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOracleRespondParamsWithTimeout(timeout time.Duration) *PostOracleRespondParams {
	var ()
	return &PostOracleRespondParams{

		timeout: timeout,
	}
}

// NewPostOracleRespondParamsWithContext creates a new PostOracleRespondParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOracleRespondParamsWithContext(ctx context.Context) *PostOracleRespondParams {
	var ()
	return &PostOracleRespondParams{

		Context: ctx,
	}
}

// NewPostOracleRespondParamsWithHTTPClient creates a new PostOracleRespondParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOracleRespondParamsWithHTTPClient(client *http.Client) *PostOracleRespondParams {
	var ()
	return &PostOracleRespondParams{
		HTTPClient: client,
	}
}

/*PostOracleRespondParams contains all the parameters to send to the API endpoint
for the post oracle respond operation typically these are written to a http.Request
*/
type PostOracleRespondParams struct {

	/*Body*/
	Body *models.OracleRespondTx

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post oracle respond params
func (o *PostOracleRespondParams) WithTimeout(timeout time.Duration) *PostOracleRespondParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post oracle respond params
func (o *PostOracleRespondParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post oracle respond params
func (o *PostOracleRespondParams) WithContext(ctx context.Context) *PostOracleRespondParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post oracle respond params
func (o *PostOracleRespondParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post oracle respond params
func (o *PostOracleRespondParams) WithHTTPClient(client *http.Client) *PostOracleRespondParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post oracle respond params
func (o *PostOracleRespondParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post oracle respond params
func (o *PostOracleRespondParams) WithBody(body *models.OracleRespondTx) *PostOracleRespondParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post oracle respond params
func (o *PostOracleRespondParams) SetBody(body *models.OracleRespondTx) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostOracleRespondParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
