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

// GetCommitmentIDReader is a Reader for the GetCommitmentID structure.
type GetCommitmentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCommitmentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCommitmentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetCommitmentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCommitmentIDOK creates a GetCommitmentIDOK with default headers values
func NewGetCommitmentIDOK() *GetCommitmentIDOK {
	return &GetCommitmentIDOK{}
}

/*GetCommitmentIDOK handles this case with default header values.

Successful operation
*/
type GetCommitmentIDOK struct {
	Payload *models.CommitmentID
}

func (o *GetCommitmentIDOK) Error() string {
	return fmt.Sprintf("[GET /debug/names/commitment-id][%d] getCommitmentIdOK  %+v", 200, o.Payload)
}

func (o *GetCommitmentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommitmentID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCommitmentIDBadRequest creates a GetCommitmentIDBadRequest with default headers values
func NewGetCommitmentIDBadRequest() *GetCommitmentIDBadRequest {
	return &GetCommitmentIDBadRequest{}
}

/*GetCommitmentIDBadRequest handles this case with default header values.

Invalid name
*/
type GetCommitmentIDBadRequest struct {
	Payload *models.Error
}

func (o *GetCommitmentIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /debug/names/commitment-id][%d] getCommitmentIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetCommitmentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
