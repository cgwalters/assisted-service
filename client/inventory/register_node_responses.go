// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/models"
)

// RegisterNodeReader is a Reader for the RegisterNode structure.
type RegisterNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRegisterNodeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegisterNodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegisterNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegisterNodeCreated creates a RegisterNodeCreated with default headers values
func NewRegisterNodeCreated() *RegisterNodeCreated {
	return &RegisterNodeCreated{}
}

/*RegisterNodeCreated handles this case with default header values.

Registered node
*/
type RegisterNodeCreated struct {
	Payload *models.Node
}

func (o *RegisterNodeCreated) Error() string {
	return fmt.Sprintf("[POST /nodes][%d] registerNodeCreated  %+v", 201, o.Payload)
}

func (o *RegisterNodeCreated) GetPayload() *models.Node {
	return o.Payload
}

func (o *RegisterNodeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterNodeBadRequest creates a RegisterNodeBadRequest with default headers values
func NewRegisterNodeBadRequest() *RegisterNodeBadRequest {
	return &RegisterNodeBadRequest{}
}

/*RegisterNodeBadRequest handles this case with default header values.

Invalid input
*/
type RegisterNodeBadRequest struct {
}

func (o *RegisterNodeBadRequest) Error() string {
	return fmt.Sprintf("[POST /nodes][%d] registerNodeBadRequest ", 400)
}

func (o *RegisterNodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterNodeInternalServerError creates a RegisterNodeInternalServerError with default headers values
func NewRegisterNodeInternalServerError() *RegisterNodeInternalServerError {
	return &RegisterNodeInternalServerError{}
}

/*RegisterNodeInternalServerError handles this case with default header values.

Internal server error
*/
type RegisterNodeInternalServerError struct {
}

func (o *RegisterNodeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /nodes][%d] registerNodeInternalServerError ", 500)
}

func (o *RegisterNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
