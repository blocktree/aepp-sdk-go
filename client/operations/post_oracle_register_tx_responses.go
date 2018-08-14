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

// PostOracleRegisterTxReader is a Reader for the PostOracleRegisterTx structure.
type PostOracleRegisterTxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOracleRegisterTxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostOracleRegisterTxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostOracleRegisterTxNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostOracleRegisterTxOK creates a PostOracleRegisterTxOK with default headers values
func NewPostOracleRegisterTxOK() *PostOracleRegisterTxOK {
	return &PostOracleRegisterTxOK{}
}

/*PostOracleRegisterTxOK handles this case with default header values.

successful operation
*/
type PostOracleRegisterTxOK struct {
	Payload *models.OracleRegisterResponse
}

func (o *PostOracleRegisterTxOK) Error() string {
	return fmt.Sprintf("[POST /oracle-register-tx][%d] postOracleRegisterTxOK  %+v", 200, o.Payload)
}

func (o *PostOracleRegisterTxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OracleRegisterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOracleRegisterTxNotFound creates a PostOracleRegisterTxNotFound with default headers values
func NewPostOracleRegisterTxNotFound() *PostOracleRegisterTxNotFound {
	return &PostOracleRegisterTxNotFound{}
}

/*PostOracleRegisterTxNotFound handles this case with default header values.

Oracle register transaction validation error
*/
type PostOracleRegisterTxNotFound struct {
	Payload *models.Error
}

func (o *PostOracleRegisterTxNotFound) Error() string {
	return fmt.Sprintf("[POST /oracle-register-tx][%d] postOracleRegisterTxNotFound  %+v", 404, o.Payload)
}

func (o *PostOracleRegisterTxNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
