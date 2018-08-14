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

// GetPeerKeyReader is a Reader for the GetPeerKey structure.
type GetPeerKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPeerKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPeerKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetPeerKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPeerKeyOK creates a GetPeerKeyOK with default headers values
func NewGetPeerKeyOK() *GetPeerKeyOK {
	return &GetPeerKeyOK{}
}

/*GetPeerKeyOK handles this case with default header values.

Node's peer public key
*/
type GetPeerKeyOK struct {
	Payload *models.InlineResponse2004
}

func (o *GetPeerKeyOK) Error() string {
	return fmt.Sprintf("[GET /peer/key][%d] getPeerKeyOK  %+v", 200, o.Payload)
}

func (o *GetPeerKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InlineResponse2004)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeerKeyNotFound creates a GetPeerKeyNotFound with default headers values
func NewGetPeerKeyNotFound() *GetPeerKeyNotFound {
	return &GetPeerKeyNotFound{}
}

/*GetPeerKeyNotFound handles this case with default header values.

No key pair
*/
type GetPeerKeyNotFound struct {
	Payload *models.Error
}

func (o *GetPeerKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /peer/key][%d] getPeerKeyNotFound  %+v", 404, o.Payload)
}

func (o *GetPeerKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
