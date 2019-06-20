// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/aeternity/aepp-sdk-go/swagguard/node/models"
)

// GetTopBlockReader is a Reader for the GetTopBlock structure.
type GetTopBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTopBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTopBlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetTopBlockNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTopBlockOK creates a GetTopBlockOK with default headers values
func NewGetTopBlockOK() *GetTopBlockOK {
	return &GetTopBlockOK{}
}

/*GetTopBlockOK handles this case with default header values.

Successful operation
*/
type GetTopBlockOK struct {
	Payload *models.KeyBlockOrMicroBlockHeader
}

func (o *GetTopBlockOK) Error() string {
	return fmt.Sprintf("[GET /blocks/top][%d] getTopBlockOK  %+v", 200, o.Payload)
}

func (o *GetTopBlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeyBlockOrMicroBlockHeader)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTopBlockNotFound creates a GetTopBlockNotFound with default headers values
func NewGetTopBlockNotFound() *GetTopBlockNotFound {
	return &GetTopBlockNotFound{}
}

/*GetTopBlockNotFound handles this case with default header values.

Block not found
*/
type GetTopBlockNotFound struct {
	Payload *models.Error
}

func (o *GetTopBlockNotFound) Error() string {
	return fmt.Sprintf("[GET /blocks/top][%d] getTopBlockNotFound  %+v", 404, o.Payload)
}

func (o *GetTopBlockNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}