// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new endpoints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for endpoints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateGithubEndpoint(params *CreateGithubEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGithubEndpointOK, error)

	DeleteGithubEndpoint(params *DeleteGithubEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	GetGithubEndpoint(params *GetGithubEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGithubEndpointOK, error)

	ListGithubEndpoints(params *ListGithubEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGithubEndpointsOK, error)

	UpdateGithubEndpoint(params *UpdateGithubEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGithubEndpointOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateGithubEndpoint creates a git hub endpoint
*/
func (a *Client) CreateGithubEndpoint(params *CreateGithubEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGithubEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGithubEndpointParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateGithubEndpoint",
		Method:             "POST",
		PathPattern:        "/github/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateGithubEndpointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateGithubEndpointOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateGithubEndpointDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteGithubEndpoint deletes a git hub endpoint
*/
func (a *Client) DeleteGithubEndpoint(params *DeleteGithubEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGithubEndpointParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteGithubEndpoint",
		Method:             "DELETE",
		PathPattern:        "/github/endpoints/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteGithubEndpointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
GetGithubEndpoint gets a git hub endpoint
*/
func (a *Client) GetGithubEndpoint(params *GetGithubEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGithubEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGithubEndpointParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGithubEndpoint",
		Method:             "GET",
		PathPattern:        "/github/endpoints/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGithubEndpointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGithubEndpointOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetGithubEndpointDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGithubEndpoints lists all git hub endpoints
*/
func (a *Client) ListGithubEndpoints(params *ListGithubEndpointsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGithubEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGithubEndpointsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListGithubEndpoints",
		Method:             "GET",
		PathPattern:        "/github/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListGithubEndpointsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGithubEndpointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGithubEndpointsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateGithubEndpoint updates a git hub endpoint
*/
func (a *Client) UpdateGithubEndpoint(params *UpdateGithubEndpointParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGithubEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGithubEndpointParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateGithubEndpoint",
		Method:             "PUT",
		PathPattern:        "/github/endpoints/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGithubEndpointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateGithubEndpointOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateGithubEndpointDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
