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

package accounts

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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewUpdateCurrentAccountParams creates a new UpdateCurrentAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCurrentAccountParams() *UpdateCurrentAccountParams {
	return &UpdateCurrentAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCurrentAccountParamsWithTimeout creates a new UpdateCurrentAccountParams object
// with the ability to set a timeout on a request.
func NewUpdateCurrentAccountParamsWithTimeout(timeout time.Duration) *UpdateCurrentAccountParams {
	return &UpdateCurrentAccountParams{
		timeout: timeout,
	}
}

// NewUpdateCurrentAccountParamsWithContext creates a new UpdateCurrentAccountParams object
// with the ability to set a context for a request.
func NewUpdateCurrentAccountParamsWithContext(ctx context.Context) *UpdateCurrentAccountParams {
	return &UpdateCurrentAccountParams{
		Context: ctx,
	}
}

// NewUpdateCurrentAccountParamsWithHTTPClient creates a new UpdateCurrentAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCurrentAccountParamsWithHTTPClient(client *http.Client) *UpdateCurrentAccountParams {
	return &UpdateCurrentAccountParams{
		HTTPClient: client,
	}
}

/* UpdateCurrentAccountParams contains all the parameters to send to the API endpoint
   for the update current account operation.

   Typically these are written to a http.Request.
*/
type UpdateCurrentAccountParams struct {

	/* Body.

	   the current account
	*/
	Body *models.AccountUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update current account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCurrentAccountParams) WithDefaults() *UpdateCurrentAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update current account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCurrentAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update current account params
func (o *UpdateCurrentAccountParams) WithTimeout(timeout time.Duration) *UpdateCurrentAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update current account params
func (o *UpdateCurrentAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update current account params
func (o *UpdateCurrentAccountParams) WithContext(ctx context.Context) *UpdateCurrentAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update current account params
func (o *UpdateCurrentAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update current account params
func (o *UpdateCurrentAccountParams) WithHTTPClient(client *http.Client) *UpdateCurrentAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update current account params
func (o *UpdateCurrentAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update current account params
func (o *UpdateCurrentAccountParams) WithBody(body *models.AccountUpdateRequest) *UpdateCurrentAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update current account params
func (o *UpdateCurrentAccountParams) SetBody(body *models.AccountUpdateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCurrentAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
