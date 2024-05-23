// Code generated by go-swagger; DO NOT EDIT.

package metrics_token

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

// GetMetricsTokenReader is a Reader for the GetMetricsToken structure.
type GetMetricsTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricsTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetricsTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMetricsTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /metrics-token] GetMetricsToken", response, response.Code())
	}
}

// NewGetMetricsTokenOK creates a GetMetricsTokenOK with default headers values
func NewGetMetricsTokenOK() *GetMetricsTokenOK {
	return &GetMetricsTokenOK{}
}

/*
GetMetricsTokenOK describes a response with status code 200, with default header values.

JWTResponse
*/
type GetMetricsTokenOK struct {
	Payload garm_params.JWTResponse
}

// IsSuccess returns true when this get metrics token o k response has a 2xx status code
func (o *GetMetricsTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get metrics token o k response has a 3xx status code
func (o *GetMetricsTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metrics token o k response has a 4xx status code
func (o *GetMetricsTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get metrics token o k response has a 5xx status code
func (o *GetMetricsTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get metrics token o k response a status code equal to that given
func (o *GetMetricsTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get metrics token o k response
func (o *GetMetricsTokenOK) Code() int {
	return 200
}

func (o *GetMetricsTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /metrics-token][%d] getMetricsTokenOK %s", 200, payload)
}

func (o *GetMetricsTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /metrics-token][%d] getMetricsTokenOK %s", 200, payload)
}

func (o *GetMetricsTokenOK) GetPayload() garm_params.JWTResponse {
	return o.Payload
}

func (o *GetMetricsTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsTokenUnauthorized creates a GetMetricsTokenUnauthorized with default headers values
func NewGetMetricsTokenUnauthorized() *GetMetricsTokenUnauthorized {
	return &GetMetricsTokenUnauthorized{}
}

/*
GetMetricsTokenUnauthorized describes a response with status code 401, with default header values.

APIErrorResponse
*/
type GetMetricsTokenUnauthorized struct {
	Payload apiserver_params.APIErrorResponse
}

// IsSuccess returns true when this get metrics token unauthorized response has a 2xx status code
func (o *GetMetricsTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get metrics token unauthorized response has a 3xx status code
func (o *GetMetricsTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metrics token unauthorized response has a 4xx status code
func (o *GetMetricsTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get metrics token unauthorized response has a 5xx status code
func (o *GetMetricsTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get metrics token unauthorized response a status code equal to that given
func (o *GetMetricsTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get metrics token unauthorized response
func (o *GetMetricsTokenUnauthorized) Code() int {
	return 401
}

func (o *GetMetricsTokenUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /metrics-token][%d] getMetricsTokenUnauthorized %s", 401, payload)
}

func (o *GetMetricsTokenUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /metrics-token][%d] getMetricsTokenUnauthorized %s", 401, payload)
}

func (o *GetMetricsTokenUnauthorized) GetPayload() apiserver_params.APIErrorResponse {
	return o.Payload
}

func (o *GetMetricsTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
