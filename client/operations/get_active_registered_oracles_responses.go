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

// GetActiveRegisteredOraclesReader is a Reader for the GetActiveRegisteredOracles structure.
type GetActiveRegisteredOraclesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActiveRegisteredOraclesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetActiveRegisteredOraclesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetActiveRegisteredOraclesOK creates a GetActiveRegisteredOraclesOK with default headers values
func NewGetActiveRegisteredOraclesOK() *GetActiveRegisteredOraclesOK {
	return &GetActiveRegisteredOraclesOK{}
}

/*GetActiveRegisteredOraclesOK handles this case with default header values.

Active registered oracles
*/
type GetActiveRegisteredOraclesOK struct {
	Payload models.RegisteredOracles
}

func (o *GetActiveRegisteredOraclesOK) Error() string {
	return fmt.Sprintf("[GET /oracles][%d] getActiveRegisteredOraclesOK  %+v", 200, o.Payload)
}

func (o *GetActiveRegisteredOraclesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
