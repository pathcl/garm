// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewDeleteOrgParams creates a new DeleteOrgParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrgParams() *DeleteOrgParams {
	return &DeleteOrgParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrgParamsWithTimeout creates a new DeleteOrgParams object
// with the ability to set a timeout on a request.
func NewDeleteOrgParamsWithTimeout(timeout time.Duration) *DeleteOrgParams {
	return &DeleteOrgParams{
		timeout: timeout,
	}
}

// NewDeleteOrgParamsWithContext creates a new DeleteOrgParams object
// with the ability to set a context for a request.
func NewDeleteOrgParamsWithContext(ctx context.Context) *DeleteOrgParams {
	return &DeleteOrgParams{
		Context: ctx,
	}
}

// NewDeleteOrgParamsWithHTTPClient creates a new DeleteOrgParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrgParamsWithHTTPClient(client *http.Client) *DeleteOrgParams {
	return &DeleteOrgParams{
		HTTPClient: client,
	}
}

/*
DeleteOrgParams contains all the parameters to send to the API endpoint

	for the delete org operation.

	Typically these are written to a http.Request.
*/
type DeleteOrgParams struct {

	/* KeepWebhook.

	   If true and a webhook is installed for this organization, it will not be removed.
	*/
	KeepWebhook *bool

	/* OrgID.

	   ID of the organization to delete.
	*/
	OrgID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgParams) WithDefaults() *DeleteOrgParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete org params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrgParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete org params
func (o *DeleteOrgParams) WithTimeout(timeout time.Duration) *DeleteOrgParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete org params
func (o *DeleteOrgParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete org params
func (o *DeleteOrgParams) WithContext(ctx context.Context) *DeleteOrgParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete org params
func (o *DeleteOrgParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete org params
func (o *DeleteOrgParams) WithHTTPClient(client *http.Client) *DeleteOrgParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete org params
func (o *DeleteOrgParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeepWebhook adds the keepWebhook to the delete org params
func (o *DeleteOrgParams) WithKeepWebhook(keepWebhook *bool) *DeleteOrgParams {
	o.SetKeepWebhook(keepWebhook)
	return o
}

// SetKeepWebhook adds the keepWebhook to the delete org params
func (o *DeleteOrgParams) SetKeepWebhook(keepWebhook *bool) {
	o.KeepWebhook = keepWebhook
}

// WithOrgID adds the orgID to the delete org params
func (o *DeleteOrgParams) WithOrgID(orgID string) *DeleteOrgParams {
	o.SetOrgID(orgID)
	return o
}

// SetOrgID adds the orgId to the delete org params
func (o *DeleteOrgParams) SetOrgID(orgID string) {
	o.OrgID = orgID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrgParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.KeepWebhook != nil {

		// query param keepWebhook
		var qrKeepWebhook bool

		if o.KeepWebhook != nil {
			qrKeepWebhook = *o.KeepWebhook
		}
		qKeepWebhook := swag.FormatBool(qrKeepWebhook)
		if qKeepWebhook != "" {

			if err := r.SetQueryParam("keepWebhook", qKeepWebhook); err != nil {
				return err
			}
		}
	}

	// path param orgID
	if err := r.SetPathParam("orgID", o.OrgID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
