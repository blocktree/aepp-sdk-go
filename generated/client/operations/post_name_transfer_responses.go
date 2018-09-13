// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/generated/models"
)

// PostNameTransferReader is a Reader for the PostNameTransfer structure.
type PostNameTransferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNameTransferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostNameTransferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostNameTransferBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostNameTransferNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostNameTransferOK creates a PostNameTransferOK with default headers values
func NewPostNameTransferOK() *PostNameTransferOK {
	return &PostNameTransferOK{}
}

/*PostNameTransferOK handles this case with default header values.

Successful operation
*/
type PostNameTransferOK struct {
	Payload *models.UnsignedTx
}

func (o *PostNameTransferOK) Error() string {
	return fmt.Sprintf("[POST /debug/names/transfer][%d] postNameTransferOK  %+v", 200, o.Payload)
}

func (o *PostNameTransferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnsignedTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNameTransferBadRequest creates a PostNameTransferBadRequest with default headers values
func NewPostNameTransferBadRequest() *PostNameTransferBadRequest {
	return &PostNameTransferBadRequest{}
}

/*PostNameTransferBadRequest handles this case with default header values.

Invalid transaction
*/
type PostNameTransferBadRequest struct {
	Payload *models.Error
}

func (o *PostNameTransferBadRequest) Error() string {
	return fmt.Sprintf("[POST /debug/names/transfer][%d] postNameTransferBadRequest  %+v", 400, o.Payload)
}

func (o *PostNameTransferBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNameTransferNotFound creates a PostNameTransferNotFound with default headers values
func NewPostNameTransferNotFound() *PostNameTransferNotFound {
	return &PostNameTransferNotFound{}
}

/*PostNameTransferNotFound handles this case with default header values.

Account or name not found
*/
type PostNameTransferNotFound struct {
	Payload *models.Error
}

func (o *PostNameTransferNotFound) Error() string {
	return fmt.Sprintf("[POST /debug/names/transfer][%d] postNameTransferNotFound  %+v", 404, o.Payload)
}

func (o *PostNameTransferNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
