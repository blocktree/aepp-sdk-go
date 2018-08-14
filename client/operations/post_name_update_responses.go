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

// PostNameUpdateReader is a Reader for the PostNameUpdate structure.
type PostNameUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNameUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostNameUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostNameUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostNameUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostNameUpdateOK creates a PostNameUpdateOK with default headers values
func NewPostNameUpdateOK() *PostNameUpdateOK {
	return &PostNameUpdateOK{}
}

/*PostNameUpdateOK handles this case with default header values.

Successful operation
*/
type PostNameUpdateOK struct {
	Payload *models.UnsignedTx
}

func (o *PostNameUpdateOK) Error() string {
	return fmt.Sprintf("[POST /tx/name/update][%d] postNameUpdateOK  %+v", 200, o.Payload)
}

func (o *PostNameUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnsignedTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNameUpdateBadRequest creates a PostNameUpdateBadRequest with default headers values
func NewPostNameUpdateBadRequest() *PostNameUpdateBadRequest {
	return &PostNameUpdateBadRequest{}
}

/*PostNameUpdateBadRequest handles this case with default header values.

Invalid transaction
*/
type PostNameUpdateBadRequest struct {
	Payload *models.Error
}

func (o *PostNameUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /tx/name/update][%d] postNameUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PostNameUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNameUpdateNotFound creates a PostNameUpdateNotFound with default headers values
func NewPostNameUpdateNotFound() *PostNameUpdateNotFound {
	return &PostNameUpdateNotFound{}
}

/*PostNameUpdateNotFound handles this case with default header values.

Account or oracle not found
*/
type PostNameUpdateNotFound struct {
	Payload *models.Error
}

func (o *PostNameUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /tx/name/update][%d] postNameUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PostNameUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
