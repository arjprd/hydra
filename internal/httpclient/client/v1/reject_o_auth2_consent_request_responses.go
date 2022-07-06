// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/client/v1"
	"github.com/ory/hydra/internal/httpclient/models"
)

// RejectOAuth2ConsentRequestReader is a Reader for the RejectOAuth2ConsentRequest structure.
type RejectOAuth2ConsentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RejectOAuth2ConsentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRejectOAuth2ConsentRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRejectOAuth2ConsentRequestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRejectOAuth2ConsentRequestOK creates a RejectOAuth2ConsentRequestOK with default headers values
func NewRejectOAuth2ConsentRequestOK() *RejectOAuth2ConsentRequestOK {
	return &RejectOAuth2ConsentRequestOK{}
}

/* RejectOAuth2ConsentRequestOK describes a response with status code 200, with default header values.

handledOAuth2ConsentRequest
*/
type RejectOAuth2ConsentRequestOK struct {
	Payload *models.HandledOAuth2ConsentRequest
}

func (o *RejectOAuth2ConsentRequestOK) Error() string {
	return fmt.Sprintf("[PUT /admin/oauth2/auth/requests/consent/reject][%d] rejectOAuth2ConsentRequestOK  %+v", 200, o.Payload)
}
func (o *RejectOAuth2ConsentRequestOK) GetPayload() *models.HandledOAuth2ConsentRequest {
	return o.Payload
}

func (o *RejectOAuth2ConsentRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandledOAuth2ConsentRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectOAuth2ConsentRequestDefault creates a RejectOAuth2ConsentRequestDefault with default headers values
func NewRejectOAuth2ConsentRequestDefault(code int) *RejectOAuth2ConsentRequestDefault {
	return &RejectOAuth2ConsentRequestDefault{
		_statusCode: code,
	}
}

/* RejectOAuth2ConsentRequestDefault describes a response with status code -1, with default header values.

oAuth2ApiError
*/
type RejectOAuth2ConsentRequestDefault struct {
	_statusCode int

	Payload *models.OAuth2APIError
}

// Code gets the status code for the reject o auth2 consent request default response
func (o *RejectOAuth2ConsentRequestDefault) Code() int {
	return o._statusCode
}

func (o *RejectOAuth2ConsentRequestDefault) Error() string {
	return fmt.Sprintf("[PUT /admin/oauth2/auth/requests/consent/reject][%d] rejectOAuth2ConsentRequest default  %+v", o._statusCode, o.Payload)
}
func (o *RejectOAuth2ConsentRequestDefault) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *RejectOAuth2ConsentRequestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}