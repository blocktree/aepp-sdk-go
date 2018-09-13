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

// PostNameClaimReader is a Reader for the PostNameClaim structure.
type PostNameClaimReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNameClaimReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostNameClaimOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostNameClaimBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostNameClaimNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostNameClaimOK creates a PostNameClaimOK with default headers values
func NewPostNameClaimOK() *PostNameClaimOK {
	return &PostNameClaimOK{}
}

/*PostNameClaimOK handles this case with default header values.

Successful operation
*/
type PostNameClaimOK struct {
	Payload *models.UnsignedTx
}

func (o *PostNameClaimOK) Error() string {
	return fmt.Sprintf("[POST /debug/names/claim][%d] postNameClaimOK  %+v", 200, o.Payload)
}

func (o *PostNameClaimOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnsignedTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNameClaimBadRequest creates a PostNameClaimBadRequest with default headers values
func NewPostNameClaimBadRequest() *PostNameClaimBadRequest {
	return &PostNameClaimBadRequest{}
}

/*PostNameClaimBadRequest handles this case with default header values.

Invalid transaction
*/
type PostNameClaimBadRequest struct {
	Payload *models.Error
}

func (o *PostNameClaimBadRequest) Error() string {
	return fmt.Sprintf("[POST /debug/names/claim][%d] postNameClaimBadRequest  %+v", 400, o.Payload)
}

func (o *PostNameClaimBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNameClaimNotFound creates a PostNameClaimNotFound with default headers values
func NewPostNameClaimNotFound() *PostNameClaimNotFound {
	return &PostNameClaimNotFound{}
}

/*PostNameClaimNotFound handles this case with default header values.

Account or name not found
*/
type PostNameClaimNotFound struct {
	Payload *models.Error
}

func (o *PostNameClaimNotFound) Error() string {
	return fmt.Sprintf("[POST /debug/names/claim][%d] postNameClaimNotFound  %+v", 404, o.Payload)
}

func (o *PostNameClaimNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
