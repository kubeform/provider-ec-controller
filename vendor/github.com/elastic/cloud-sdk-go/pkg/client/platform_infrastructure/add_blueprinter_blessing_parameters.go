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
	"github.com/go-openapi/swag"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewAddBlueprinterBlessingParams creates a new AddBlueprinterBlessingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddBlueprinterBlessingParams() *AddBlueprinterBlessingParams {
	return &AddBlueprinterBlessingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddBlueprinterBlessingParamsWithTimeout creates a new AddBlueprinterBlessingParams object
// with the ability to set a timeout on a request.
func NewAddBlueprinterBlessingParamsWithTimeout(timeout time.Duration) *AddBlueprinterBlessingParams {
	return &AddBlueprinterBlessingParams{
		timeout: timeout,
	}
}

// NewAddBlueprinterBlessingParamsWithContext creates a new AddBlueprinterBlessingParams object
// with the ability to set a context for a request.
func NewAddBlueprinterBlessingParamsWithContext(ctx context.Context) *AddBlueprinterBlessingParams {
	return &AddBlueprinterBlessingParams{
		Context: ctx,
	}
}

// NewAddBlueprinterBlessingParamsWithHTTPClient creates a new AddBlueprinterBlessingParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddBlueprinterBlessingParamsWithHTTPClient(client *http.Client) *AddBlueprinterBlessingParams {
	return &AddBlueprinterBlessingParams{
		HTTPClient: client,
	}
}

/* AddBlueprinterBlessingParams contains all the parameters to send to the API endpoint
   for the add blueprinter blessing operation.

   Typically these are written to a http.Request.
*/
type AddBlueprinterBlessingParams struct {

	/* BlueprinterRoleID.

	   User-specified Blueprinter role ID.
	*/
	BlueprinterRoleID string

	/* Body.

	   The blessing to add.
	*/
	Body *models.Blessing

	/* RunnerID.

	   Runner ID for a blessing associated with a role.
	*/
	RunnerID string

	/* Version.

	   When specified, checks for conflicts against the version stored in the persistent store (returned in 'x-cloud-resource-version' of the GET request)
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add blueprinter blessing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBlueprinterBlessingParams) WithDefaults() *AddBlueprinterBlessingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add blueprinter blessing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddBlueprinterBlessingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) WithTimeout(timeout time.Duration) *AddBlueprinterBlessingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) WithContext(ctx context.Context) *AddBlueprinterBlessingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) WithHTTPClient(client *http.Client) *AddBlueprinterBlessingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlueprinterRoleID adds the blueprinterRoleID to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) WithBlueprinterRoleID(blueprinterRoleID string) *AddBlueprinterBlessingParams {
	o.SetBlueprinterRoleID(blueprinterRoleID)
	return o
}

// SetBlueprinterRoleID adds the blueprinterRoleId to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) SetBlueprinterRoleID(blueprinterRoleID string) {
	o.BlueprinterRoleID = blueprinterRoleID
}

// WithBody adds the body to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) WithBody(body *models.Blessing) *AddBlueprinterBlessingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) SetBody(body *models.Blessing) {
	o.Body = body
}

// WithRunnerID adds the runnerID to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) WithRunnerID(runnerID string) *AddBlueprinterBlessingParams {
	o.SetRunnerID(runnerID)
	return o
}

// SetRunnerID adds the runnerId to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) SetRunnerID(runnerID string) {
	o.RunnerID = runnerID
}

// WithVersion adds the version to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) WithVersion(version *int64) *AddBlueprinterBlessingParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the add blueprinter blessing params
func (o *AddBlueprinterBlessingParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *AddBlueprinterBlessingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param blueprinter_role_id
	if err := r.SetPathParam("blueprinter_role_id", o.BlueprinterRoleID); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param runner_id
	if err := r.SetPathParam("runner_id", o.RunnerID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
