// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apiserver_params "github.com/cloudbase/garm/apiserver/params"
	garm_params "github.com/cloudbase/garm/params"
)

// LoginReader is a Reader for the Login structure.
type LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLoginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /auth/login] Login", response, response.Code())
	}
}

// NewLoginOK creates a LoginOK with default headers values
func NewLoginOK() *LoginOK {
	return &LoginOK{}
}

/*
LoginOK describes a response with status code 200, with default header values.

JWTResponse
*/
type LoginOK struct {
	Payload garm_params.JWTResponse
}

// IsSuccess returns true when this login o k response has a 2xx status code
func (o *LoginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this login o k response has a 3xx status code
func (o *LoginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login o k response has a 4xx status code
func (o *LoginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this login o k response has a 5xx status code
func (o *LoginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this login o k response a status code equal to that given
func (o *LoginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the login o k response
func (o *LoginOK) Code() int {
	return 200
}

func (o *LoginOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /auth/login][%d] loginOK %s", 200, payload)
}

func (o *LoginOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /auth/login][%d] loginOK %s", 200, payload)
}

func (o *LoginOK) GetPayload() garm_params.JWTResponse {
	return o.Payload
}

func (o *LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginBadRequest creates a LoginBadRequest with default headers values
func NewLoginBadRequest() *LoginBadRequest {
	return &LoginBadRequest{}
}

/*
LoginBadRequest describes a response with status code 400, with default header values.

APIErrorResponse
*/
type LoginBadRequest struct {
	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this login bad request response has a 2xx status code
func (o *LoginBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this login bad request response has a 3xx status code
func (o *LoginBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login bad request response has a 4xx status code
func (o *LoginBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this login bad request response has a 5xx status code
func (o *LoginBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this login bad request response a status code equal to that given
func (o *LoginBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the login bad request response
func (o *LoginBadRequest) Code() int {
	return 400
}

func (o *LoginBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /auth/login][%d] loginBadRequest %s", 400, payload)
}

func (o *LoginBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /auth/login][%d] loginBadRequest %s", 400, payload)
}

func (o *LoginBadRequest) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *LoginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
