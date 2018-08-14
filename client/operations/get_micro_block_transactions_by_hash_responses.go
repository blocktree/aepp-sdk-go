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

// GetMicroBlockTransactionsByHashReader is a Reader for the GetMicroBlockTransactionsByHash structure.
type GetMicroBlockTransactionsByHashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMicroBlockTransactionsByHashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMicroBlockTransactionsByHashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetMicroBlockTransactionsByHashBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetMicroBlockTransactionsByHashNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMicroBlockTransactionsByHashOK creates a GetMicroBlockTransactionsByHashOK with default headers values
func NewGetMicroBlockTransactionsByHashOK() *GetMicroBlockTransactionsByHashOK {
	return &GetMicroBlockTransactionsByHashOK{}
}

/*GetMicroBlockTransactionsByHashOK handles this case with default header values.

Successful operation
*/
type GetMicroBlockTransactionsByHashOK struct {
	Payload *models.JSONTxs
}

func (o *GetMicroBlockTransactionsByHashOK) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/transactions][%d] getMicroBlockTransactionsByHashOK  %+v", 200, o.Payload)
}

func (o *GetMicroBlockTransactionsByHashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONTxs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicroBlockTransactionsByHashBadRequest creates a GetMicroBlockTransactionsByHashBadRequest with default headers values
func NewGetMicroBlockTransactionsByHashBadRequest() *GetMicroBlockTransactionsByHashBadRequest {
	return &GetMicroBlockTransactionsByHashBadRequest{}
}

/*GetMicroBlockTransactionsByHashBadRequest handles this case with default header values.

Invalid hash
*/
type GetMicroBlockTransactionsByHashBadRequest struct {
	Payload *models.Error
}

func (o *GetMicroBlockTransactionsByHashBadRequest) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/transactions][%d] getMicroBlockTransactionsByHashBadRequest  %+v", 400, o.Payload)
}

func (o *GetMicroBlockTransactionsByHashBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMicroBlockTransactionsByHashNotFound creates a GetMicroBlockTransactionsByHashNotFound with default headers values
func NewGetMicroBlockTransactionsByHashNotFound() *GetMicroBlockTransactionsByHashNotFound {
	return &GetMicroBlockTransactionsByHashNotFound{}
}

/*GetMicroBlockTransactionsByHashNotFound handles this case with default header values.

Block not found
*/
type GetMicroBlockTransactionsByHashNotFound struct {
	Payload *models.Error
}

func (o *GetMicroBlockTransactionsByHashNotFound) Error() string {
	return fmt.Sprintf("[GET /micro-blocks/hash/{hash}/transactions][%d] getMicroBlockTransactionsByHashNotFound  %+v", 404, o.Payload)
}

func (o *GetMicroBlockTransactionsByHashNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
