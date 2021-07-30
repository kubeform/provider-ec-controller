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

// NewDeleteRunnerLoggingSettingsParams creates a new DeleteRunnerLoggingSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRunnerLoggingSettingsParams() *DeleteRunnerLoggingSettingsParams {
	return &DeleteRunnerLoggingSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRunnerLoggingSettingsParamsWithTimeout creates a new DeleteRunnerLoggingSettingsParams object
// with the ability to set a timeout on a request.
func NewDeleteRunnerLoggingSettingsParamsWithTimeout(timeout time.Duration) *DeleteRunnerLoggingSettingsParams {
	return &DeleteRunnerLoggingSettingsParams{
		timeout: timeout,
	}
}

// NewDeleteRunnerLoggingSettingsParamsWithContext creates a new DeleteRunnerLoggingSettingsParams object
// with the ability to set a context for a request.
func NewDeleteRunnerLoggingSettingsParamsWithContext(ctx context.Context) *DeleteRunnerLoggingSettingsParams {
	return &DeleteRunnerLoggingSettingsParams{
		Context: ctx,
	}
}

// NewDeleteRunnerLoggingSettingsParamsWithHTTPClient creates a new DeleteRunnerLoggingSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRunnerLoggingSettingsParamsWithHTTPClient(client *http.Client) *DeleteRunnerLoggingSettingsParams {
	return &DeleteRunnerLoggingSettingsParams{
		HTTPClient: client,
	}
}

/* DeleteRunnerLoggingSettingsParams contains all the parameters to send to the API endpoint
   for the delete runner logging settings operation.

   Typically these are written to a http.Request.
*/
type DeleteRunnerLoggingSettingsParams struct {

	/* RunnerID.

	   The identifier for the runner
	*/
	RunnerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete runner logging settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRunnerLoggingSettingsParams) WithDefaults() *DeleteRunnerLoggingSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete runner logging settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRunnerLoggingSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete runner logging settings params
func (o *DeleteRunnerLoggingSettingsParams) WithTimeout(timeout time.Duration) *DeleteRunnerLoggingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete runner logging settings params
func (o *DeleteRunnerLoggingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete runner logging settings params
func (o *DeleteRunnerLoggingSettingsParams) WithContext(ctx context.Context) *DeleteRunnerLoggingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete runner logging settings params
func (o *DeleteRunnerLoggingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete runner logging settings params
func (o *DeleteRunnerLoggingSettingsParams) WithHTTPClient(client *http.Client) *DeleteRunnerLoggingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete runner logging settings params
func (o *DeleteRunnerLoggingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRunnerID adds the runnerID to the delete runner logging settings params
func (o *DeleteRunnerLoggingSettingsParams) WithRunnerID(runnerID string) *DeleteRunnerLoggingSettingsParams {
	o.SetRunnerID(runnerID)
	return o
}

// SetRunnerID adds the runnerId to the delete runner logging settings params
func (o *DeleteRunnerLoggingSettingsParams) SetRunnerID(runnerID string) {
	o.RunnerID = runnerID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRunnerLoggingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param runner_id
	if err := r.SetPathParam("runner_id", o.RunnerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
