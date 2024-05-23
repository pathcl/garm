// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// ListOrgInstancesReader is a Reader for the ListOrgInstances structure.
type ListOrgInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrgInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrgInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOrgInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOrgInstancesOK creates a ListOrgInstancesOK with default headers values
func NewListOrgInstancesOK() *ListOrgInstancesOK {
	return &ListOrgInstancesOK{}
}

/*
ListOrgInstancesOK describes a response with status code 200, with default header values.

Instances
*/
type ListOrgInstancesOK struct {
	Payload garm_params.Instances
}

// IsSuccess returns true when this list org instances o k response has a 2xx status code
func (o *ListOrgInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list org instances o k response has a 3xx status code
func (o *ListOrgInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list org instances o k response has a 4xx status code
func (o *ListOrgInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list org instances o k response has a 5xx status code
func (o *ListOrgInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list org instances o k response a status code equal to that given
func (o *ListOrgInstancesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list org instances o k response
func (o *ListOrgInstancesOK) Code() int {
	return 200
}

func (o *ListOrgInstancesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{orgID}/instances][%d] listOrgInstancesOK %s", 200, payload)
}

func (o *ListOrgInstancesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{orgID}/instances][%d] listOrgInstancesOK %s", 200, payload)
}

func (o *ListOrgInstancesOK) GetPayload() garm_params.Instances {
	return o.Payload
}

func (o *ListOrgInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrgInstancesDefault creates a ListOrgInstancesDefault with default headers values
func NewListOrgInstancesDefault(code int) *ListOrgInstancesDefault {
	return &ListOrgInstancesDefault{
		_statusCode: code,
	}
}

/*
ListOrgInstancesDefault describes a response with status code -1, with default header values.

APIErrorResponse
*/
type ListOrgInstancesDefault struct {
	_statusCode int

	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this list org instances default response has a 2xx status code
func (o *ListOrgInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list org instances default response has a 3xx status code
func (o *ListOrgInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list org instances default response has a 4xx status code
func (o *ListOrgInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list org instances default response has a 5xx status code
func (o *ListOrgInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list org instances default response a status code equal to that given
func (o *ListOrgInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list org instances default response
func (o *ListOrgInstancesDefault) Code() int {
	return o._statusCode
}

func (o *ListOrgInstancesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{orgID}/instances][%d] ListOrgInstances default %s", o._statusCode, payload)
}

func (o *ListOrgInstancesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /organizations/{orgID}/instances][%d] ListOrgInstances default %s", o._statusCode, payload)
}

func (o *ListOrgInstancesDefault) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *ListOrgInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
