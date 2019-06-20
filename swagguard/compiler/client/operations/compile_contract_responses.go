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

// CompileContractReader is a Reader for the CompileContract structure.
type CompileContractReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompileContractReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCompileContractOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCompileContractBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCompileContractForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompileContractOK creates a CompileContractOK with default headers values
func NewCompileContractOK() *CompileContractOK {
	return &CompileContractOK{}
}

/*CompileContractOK handles this case with default header values.

Byte code response
*/
type CompileContractOK struct {
	Payload *models.ByteCode
}

func (o *CompileContractOK) Error() string {
	return fmt.Sprintf("[POST /compile][%d] compileContractOK  %+v", 200, o.Payload)
}

func (o *CompileContractOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ByteCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompileContractBadRequest creates a CompileContractBadRequest with default headers values
func NewCompileContractBadRequest() *CompileContractBadRequest {
	return &CompileContractBadRequest{}
}

/*CompileContractBadRequest handles this case with default header values.

Invalid data
*/
type CompileContractBadRequest struct {
	Payload *models.Error
}

func (o *CompileContractBadRequest) Error() string {
	return fmt.Sprintf("[POST /compile][%d] compileContractBadRequest  %+v", 400, o.Payload)
}

func (o *CompileContractBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompileContractForbidden creates a CompileContractForbidden with default headers values
func NewCompileContractForbidden() *CompileContractForbidden {
	return &CompileContractForbidden{}
}

/*CompileContractForbidden handles this case with default header values.

Invalid contract
*/
type CompileContractForbidden struct {
	Payload *models.Error
}

func (o *CompileContractForbidden) Error() string {
	return fmt.Sprintf("[POST /compile][%d] compileContractForbidden  %+v", 403, o.Payload)
}

func (o *CompileContractForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
