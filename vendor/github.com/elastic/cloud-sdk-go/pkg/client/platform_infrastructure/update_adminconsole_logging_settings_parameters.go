// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package platform_infrastructure

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
)

// NewUpdateAdminconsoleLoggingSettingsParams creates a new UpdateAdminconsoleLoggingSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAdminconsoleLoggingSettingsParams() *UpdateAdminconsoleLoggingSettingsParams {
	return &UpdateAdminconsoleLoggingSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAdminconsoleLoggingSettingsParamsWithTimeout creates a new UpdateAdminconsoleLoggingSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateAdminconsoleLoggingSettingsParamsWithTimeout(timeout time.Duration) *UpdateAdminconsoleLoggingSettingsParams {
	return &UpdateAdminconsoleLoggingSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateAdminconsoleLoggingSettingsParamsWithContext creates a new UpdateAdminconsoleLoggingSettingsParams object
// with the ability to set a context for a request.
func NewUpdateAdminconsoleLoggingSettingsParamsWithContext(ctx context.Context) *UpdateAdminconsoleLoggingSettingsParams {
	return &UpdateAdminconsoleLoggingSettingsParams{
		Context: ctx,
	}
}

// NewUpdateAdminconsoleLoggingSettingsParamsWithHTTPClient creates a new UpdateAdminconsoleLoggingSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAdminconsoleLoggingSettingsParamsWithHTTPClient(client *http.Client) *UpdateAdminconsoleLoggingSettingsParams {
	return &UpdateAdminconsoleLoggingSettingsParams{
		HTTPClient: client,
	}
}

/* UpdateAdminconsoleLoggingSettingsParams contains all the parameters to send to the API endpoint
   for the update adminconsole logging settings operation.

   Typically these are written to a http.Request.
*/
type UpdateAdminconsoleLoggingSettingsParams struct {

	/* AdminconsoleID.

	   The identifier for the adminconsole instance
	*/
	AdminconsoleID string

	/* Body.

	   The logging settings to update
	*/
	Body string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update adminconsole logging settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAdminconsoleLoggingSettingsParams) WithDefaults() *UpdateAdminconsoleLoggingSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update adminconsole logging settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAdminconsoleLoggingSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update adminconsole logging settings params
func (o *UpdateAdminconsoleLoggingSettingsParams) WithTimeout(timeout time.Duration) *UpdateAdminconsoleLoggingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update adminconsole logging settings params
func (o *UpdateAdminconsoleLoggingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update adminconsole logging settings params
func (o *UpdateAdminconsoleLoggingSettingsParams) WithContext(ctx context.Context) *UpdateAdminconsoleLoggingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update adminconsole logging settings params
func (o *UpdateAdminconsoleLoggingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update adminconsole logging settings params
func (o *UpdateAdminconsoleLoggingSettingsParams) WithHTTPClient(client *http.Client) *UpdateAdminconsoleLoggingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update adminconsole logging settings params
func (o *UpdateAdminconsoleLoggingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdminconsoleID adds the adminconsoleID to the update adminconsole logging settings params
func (o *UpdateAdminconsoleLoggingSettingsParams) WithAdminconsoleID(adminconsoleID string) *UpdateAdminconsoleLoggingSettingsParams {
	o.SetAdminconsoleID(adminconsoleID)
	return o
}

// SetAdminconsoleID adds the adminconsoleId to the update adminconsole logging settings params
func (o *UpdateAdminconsoleLoggingSettingsParams) SetAdminconsoleID(adminconsoleID string) {
	o.AdminconsoleID = adminconsoleID
}

// WithBody adds the body to the update adminconsole logging settings params
func (o *UpdateAdminconsoleLoggingSettingsParams) WithBody(body string) *UpdateAdminconsoleLoggingSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update adminconsole logging settings params
func (o *UpdateAdminconsoleLoggingSettingsParams) SetBody(body string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAdminconsoleLoggingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param adminconsole_id
	if err := r.SetPathParam("adminconsole_id", o.AdminconsoleID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}