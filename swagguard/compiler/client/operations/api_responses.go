// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/swagguard/compiler/models"
)

// APIReader is a Reader for the API structure.
type APIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAPIOK creates a APIOK with default headers values
func NewAPIOK() *APIOK {
	return &APIOK{}
}

/*APIOK handles this case with default header values.

API description
*/
type APIOK struct {
	Payload models.API
}

func (o *APIOK) Error() string {
	return fmt.Sprintf("[GET /api][%d] apiOK  %+v", 200, o.Payload)
}

func (o *APIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIForbidden creates a APIForbidden with default headers values
func NewAPIForbidden() *APIForbidden {
	return &APIForbidden{}
}

/*APIForbidden handles this case with default header values.

Error
*/
type APIForbidden struct {
	Payload *models.Error
}

func (o *APIForbidden) Error() string {
	return fmt.Sprintf("[GET /api][%d] apiForbidden  %+v", 403, o.Payload)
}

func (o *APIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
