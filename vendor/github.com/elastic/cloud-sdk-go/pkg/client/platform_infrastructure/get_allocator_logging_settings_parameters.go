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

// NewGetAllocatorLoggingSettingsParams creates a new GetAllocatorLoggingSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllocatorLoggingSettingsParams() *GetAllocatorLoggingSettingsParams {
	return &GetAllocatorLoggingSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllocatorLoggingSettingsParamsWithTimeout creates a new GetAllocatorLoggingSettingsParams object
// with the ability to set a timeout on a request.
func NewGetAllocatorLoggingSettingsParamsWithTimeout(timeout time.Duration) *GetAllocatorLoggingSettingsParams {
	return &GetAllocatorLoggingSettingsParams{
		timeout: timeout,
	}
}

// NewGetAllocatorLoggingSettingsParamsWithContext creates a new GetAllocatorLoggingSettingsParams object
// with the ability to set a context for a request.
func NewGetAllocatorLoggingSettingsParamsWithContext(ctx context.Context) *GetAllocatorLoggingSettingsParams {
	return &GetAllocatorLoggingSettingsParams{
		Context: ctx,
	}
}

// NewGetAllocatorLoggingSettingsParamsWithHTTPClient creates a new GetAllocatorLoggingSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllocatorLoggingSettingsParamsWithHTTPClient(client *http.Client) *GetAllocatorLoggingSettingsParams {
	return &GetAllocatorLoggingSettingsParams{
		HTTPClient: client,
	}
}

/* GetAllocatorLoggingSettingsParams contains all the parameters to send to the API endpoint
   for the get allocator logging settings operation.

   Typically these are written to a http.Request.
*/
type GetAllocatorLoggingSettingsParams struct {

	/* AllocatorID.

	   The allocator identifier.
	*/
	AllocatorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get allocator logging settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllocatorLoggingSettingsParams) WithDefaults() *GetAllocatorLoggingSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get allocator logging settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllocatorLoggingSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get allocator logging settings params
func (o *GetAllocatorLoggingSettingsParams) WithTimeout(timeout time.Duration) *GetAllocatorLoggingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get allocator logging settings params
func (o *GetAllocatorLoggingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get allocator logging settings params
func (o *GetAllocatorLoggingSettingsParams) WithContext(ctx context.Context) *GetAllocatorLoggingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get allocator logging settings params
func (o *GetAllocatorLoggingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get allocator logging settings params
func (o *GetAllocatorLoggingSettingsParams) WithHTTPClient(client *http.Client) *GetAllocatorLoggingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get allocator logging settings params
func (o *GetAllocatorLoggingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatorID adds the allocatorID to the get allocator logging settings params
func (o *GetAllocatorLoggingSettingsParams) WithAllocatorID(allocatorID string) *GetAllocatorLoggingSettingsParams {
	o.SetAllocatorID(allocatorID)
	return o
}

// SetAllocatorID adds the allocatorId to the get allocator logging settings params
func (o *GetAllocatorLoggingSettingsParams) SetAllocatorID(allocatorID string) {
	o.AllocatorID = allocatorID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllocatorLoggingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param allocator_id
	if err := r.SetPathParam("allocator_id", o.AllocatorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
