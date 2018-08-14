// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/models"
)

// PostNameRevokeTxReader is a Reader for the PostNameRevokeTx structure.
type PostNameRevokeTxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNameRevokeTxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostNameRevokeTxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostNameRevokeTxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostNameRevokeTxNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostNameRevokeTxOK creates a PostNameRevokeTxOK with default headers values
func NewPostNameRevokeTxOK() *PostNameRevokeTxOK {
	return &PostNameRevokeTxOK{}
}

/*PostNameRevokeTxOK handles this case with default header values.

successful operation
*/
type PostNameRevokeTxOK struct {
	Payload *models.NameHash
}

func (o *PostNameRevokeTxOK) Error() string {
	return fmt.Sprintf("[POST /name-revoke-tx][%d] postNameRevokeTxOK  %+v", 200, o.Payload)
}

func (o *PostNameRevokeTxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NameHash)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNameRevokeTxBadRequest creates a PostNameRevokeTxBadRequest with default headers values
func NewPostNameRevokeTxBadRequest() *PostNameRevokeTxBadRequest {
	return &PostNameRevokeTxBadRequest{}
}

/*PostNameRevokeTxBadRequest handles this case with default header values.

Name revoke transaction validation error
*/
type PostNameRevokeTxBadRequest struct {
	Payload *models.Error
}

func (o *PostNameRevokeTxBadRequest) Error() string {
	return fmt.Sprintf("[POST /name-revoke-tx][%d] postNameRevokeTxBadRequest  %+v", 400, o.Payload)
}

func (o *PostNameRevokeTxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNameRevokeTxNotFound creates a PostNameRevokeTxNotFound with default headers values
func NewPostNameRevokeTxNotFound() *PostNameRevokeTxNotFound {
	return &PostNameRevokeTxNotFound{}
}

/*PostNameRevokeTxNotFound handles this case with default header values.

Account not found error
*/
type PostNameRevokeTxNotFound struct {
	Payload *models.Error
}

func (o *PostNameRevokeTxNotFound) Error() string {
	return fmt.Sprintf("[POST /name-revoke-tx][%d] postNameRevokeTxNotFound  %+v", 404, o.Payload)
}

func (o *PostNameRevokeTxNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
