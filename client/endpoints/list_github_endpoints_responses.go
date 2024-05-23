// Code generated by go-swagger; DO NOT EDIT.

package endpoints

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

// ListGithubEndpointsReader is a Reader for the ListGithubEndpoints structure.
type ListGithubEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGithubEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGithubEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGithubEndpointsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGithubEndpointsOK creates a ListGithubEndpointsOK with default headers values
func NewListGithubEndpointsOK() *ListGithubEndpointsOK {
	return &ListGithubEndpointsOK{}
}

/*
ListGithubEndpointsOK describes a response with status code 200, with default header values.

GithubEndpoints
*/
type ListGithubEndpointsOK struct {
	Payload garm_params.GithubEndpoints
}

// IsSuccess returns true when this list github endpoints o k response has a 2xx status code
func (o *ListGithubEndpointsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list github endpoints o k response has a 3xx status code
func (o *ListGithubEndpointsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list github endpoints o k response has a 4xx status code
func (o *ListGithubEndpointsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list github endpoints o k response has a 5xx status code
func (o *ListGithubEndpointsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list github endpoints o k response a status code equal to that given
func (o *ListGithubEndpointsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list github endpoints o k response
func (o *ListGithubEndpointsOK) Code() int {
	return 200
}

func (o *ListGithubEndpointsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /github/endpoints][%d] listGithubEndpointsOK %s", 200, payload)
}

func (o *ListGithubEndpointsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /github/endpoints][%d] listGithubEndpointsOK %s", 200, payload)
}

func (o *ListGithubEndpointsOK) GetPayload() garm_params.GithubEndpoints {
	return o.Payload
}

func (o *ListGithubEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGithubEndpointsDefault creates a ListGithubEndpointsDefault with default headers values
func NewListGithubEndpointsDefault(code int) *ListGithubEndpointsDefault {
	return &ListGithubEndpointsDefault{
		_statusCode: code,
	}
}

/*
ListGithubEndpointsDefault describes a response with status code -1, with default header values.

APIErrorResponse
*/
type ListGithubEndpointsDefault struct {
	_statusCode int

	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this list github endpoints default response has a 2xx status code
func (o *ListGithubEndpointsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list github endpoints default response has a 3xx status code
func (o *ListGithubEndpointsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list github endpoints default response has a 4xx status code
func (o *ListGithubEndpointsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list github endpoints default response has a 5xx status code
func (o *ListGithubEndpointsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list github endpoints default response a status code equal to that given
func (o *ListGithubEndpointsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list github endpoints default response
func (o *ListGithubEndpointsDefault) Code() int {
	return o._statusCode
}

func (o *ListGithubEndpointsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /github/endpoints][%d] ListGithubEndpoints default %s", o._statusCode, payload)
}

func (o *ListGithubEndpointsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /github/endpoints][%d] ListGithubEndpoints default %s", o._statusCode, payload)
}

func (o *ListGithubEndpointsDefault) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *ListGithubEndpointsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
