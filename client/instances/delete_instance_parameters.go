// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteInstanceParams creates a new DeleteInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteInstanceParams() *DeleteInstanceParams {
	return &DeleteInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstanceParamsWithTimeout creates a new DeleteInstanceParams object
// with the ability to set a timeout on a request.
func NewDeleteInstanceParamsWithTimeout(timeout time.Duration) *DeleteInstanceParams {
	return &DeleteInstanceParams{
		timeout: timeout,
	}
}

// NewDeleteInstanceParamsWithContext creates a new DeleteInstanceParams object
// with the ability to set a context for a request.
func NewDeleteInstanceParamsWithContext(ctx context.Context) *DeleteInstanceParams {
	return &DeleteInstanceParams{
		Context: ctx,
	}
}

// NewDeleteInstanceParamsWithHTTPClient creates a new DeleteInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteInstanceParamsWithHTTPClient(client *http.Client) *DeleteInstanceParams {
	return &DeleteInstanceParams{
		HTTPClient: client,
	}
}

/*
DeleteInstanceParams contains all the parameters to send to the API endpoint

	for the delete instance operation.

	Typically these are written to a http.Request.
*/
type DeleteInstanceParams struct {

	/* BypassGHUnauthorized.

	   If true GARM will ignore unauthorized errors returned by GitHub when removing a runner. This is useful if you want to clean up runners and your credentials have expired.
	*/
	BypassGHUnauthorized *bool

	/* ForceRemove.

	   If true GARM will ignore any provider error when removing the runner and will continue to remove the runner from github and the GARM database.
	*/
	ForceRemove *bool

	/* InstanceName.

	   Runner instance name.
	*/
	InstanceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInstanceParams) WithDefaults() *DeleteInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete instance params
func (o *DeleteInstanceParams) WithTimeout(timeout time.Duration) *DeleteInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instance params
func (o *DeleteInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instance params
func (o *DeleteInstanceParams) WithContext(ctx context.Context) *DeleteInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instance params
func (o *DeleteInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instance params
func (o *DeleteInstanceParams) WithHTTPClient(client *http.Client) *DeleteInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instance params
func (o *DeleteInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBypassGHUnauthorized adds the bypassGHUnauthorized to the delete instance params
func (o *DeleteInstanceParams) WithBypassGHUnauthorized(bypassGHUnauthorized *bool) *DeleteInstanceParams {
	o.SetBypassGHUnauthorized(bypassGHUnauthorized)
	return o
}

// SetBypassGHUnauthorized adds the bypassGHUnauthorized to the delete instance params
func (o *DeleteInstanceParams) SetBypassGHUnauthorized(bypassGHUnauthorized *bool) {
	o.BypassGHUnauthorized = bypassGHUnauthorized
}

// WithForceRemove adds the forceRemove to the delete instance params
func (o *DeleteInstanceParams) WithForceRemove(forceRemove *bool) *DeleteInstanceParams {
	o.SetForceRemove(forceRemove)
	return o
}

// SetForceRemove adds the forceRemove to the delete instance params
func (o *DeleteInstanceParams) SetForceRemove(forceRemove *bool) {
	o.ForceRemove = forceRemove
}

// WithInstanceName adds the instanceName to the delete instance params
func (o *DeleteInstanceParams) WithInstanceName(instanceName string) *DeleteInstanceParams {
	o.SetInstanceName(instanceName)
	return o
}

// SetInstanceName adds the instanceName to the delete instance params
func (o *DeleteInstanceParams) SetInstanceName(instanceName string) {
	o.InstanceName = instanceName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BypassGHUnauthorized != nil {

		// query param bypassGHUnauthorized
		var qrBypassGHUnauthorized bool

		if o.BypassGHUnauthorized != nil {
			qrBypassGHUnauthorized = *o.BypassGHUnauthorized
		}
		qBypassGHUnauthorized := swag.FormatBool(qrBypassGHUnauthorized)
		if qBypassGHUnauthorized != "" {

			if err := r.SetQueryParam("bypassGHUnauthorized", qBypassGHUnauthorized); err != nil {
				return err
			}
		}
	}

	if o.ForceRemove != nil {

		// query param forceRemove
		var qrForceRemove bool

		if o.ForceRemove != nil {
			qrForceRemove = *o.ForceRemove
		}
		qForceRemove := swag.FormatBool(qrForceRemove)
		if qForceRemove != "" {

			if err := r.SetQueryParam("forceRemove", qForceRemove); err != nil {
				return err
			}
		}
	}

	// path param instanceName
	if err := r.SetPathParam("instanceName", o.InstanceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
