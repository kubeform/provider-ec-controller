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

// NewStartAllocatorMaintenanceModeParams creates a new StartAllocatorMaintenanceModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartAllocatorMaintenanceModeParams() *StartAllocatorMaintenanceModeParams {
	return &StartAllocatorMaintenanceModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartAllocatorMaintenanceModeParamsWithTimeout creates a new StartAllocatorMaintenanceModeParams object
// with the ability to set a timeout on a request.
func NewStartAllocatorMaintenanceModeParamsWithTimeout(timeout time.Duration) *StartAllocatorMaintenanceModeParams {
	return &StartAllocatorMaintenanceModeParams{
		timeout: timeout,
	}
}

// NewStartAllocatorMaintenanceModeParamsWithContext creates a new StartAllocatorMaintenanceModeParams object
// with the ability to set a context for a request.
func NewStartAllocatorMaintenanceModeParamsWithContext(ctx context.Context) *StartAllocatorMaintenanceModeParams {
	return &StartAllocatorMaintenanceModeParams{
		Context: ctx,
	}
}

// NewStartAllocatorMaintenanceModeParamsWithHTTPClient creates a new StartAllocatorMaintenanceModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartAllocatorMaintenanceModeParamsWithHTTPClient(client *http.Client) *StartAllocatorMaintenanceModeParams {
	return &StartAllocatorMaintenanceModeParams{
		HTTPClient: client,
	}
}

/* StartAllocatorMaintenanceModeParams contains all the parameters to send to the API endpoint
   for the start allocator maintenance mode operation.

   Typically these are written to a http.Request.
*/
type StartAllocatorMaintenanceModeParams struct {

	/* AllocatorID.

	   The allocator identifier.
	*/
	AllocatorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start allocator maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartAllocatorMaintenanceModeParams) WithDefaults() *StartAllocatorMaintenanceModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start allocator maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartAllocatorMaintenanceModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start allocator maintenance mode params
func (o *StartAllocatorMaintenanceModeParams) WithTimeout(timeout time.Duration) *StartAllocatorMaintenanceModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start allocator maintenance mode params
func (o *StartAllocatorMaintenanceModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start allocator maintenance mode params
func (o *StartAllocatorMaintenanceModeParams) WithContext(ctx context.Context) *StartAllocatorMaintenanceModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start allocator maintenance mode params
func (o *StartAllocatorMaintenanceModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start allocator maintenance mode params
func (o *StartAllocatorMaintenanceModeParams) WithHTTPClient(client *http.Client) *StartAllocatorMaintenanceModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start allocator maintenance mode params
func (o *StartAllocatorMaintenanceModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatorID adds the allocatorID to the start allocator maintenance mode params
func (o *StartAllocatorMaintenanceModeParams) WithAllocatorID(allocatorID string) *StartAllocatorMaintenanceModeParams {
	o.SetAllocatorID(allocatorID)
	return o
}

// SetAllocatorID adds the allocatorId to the start allocator maintenance mode params
func (o *StartAllocatorMaintenanceModeParams) SetAllocatorID(allocatorID string) {
	o.AllocatorID = allocatorID
}

// WriteToRequest writes these params to a swagger request
func (o *StartAllocatorMaintenanceModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
