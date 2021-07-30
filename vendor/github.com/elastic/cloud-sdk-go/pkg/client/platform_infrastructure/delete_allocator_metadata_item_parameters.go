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

// NewDeleteAllocatorMetadataItemParams creates a new DeleteAllocatorMetadataItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAllocatorMetadataItemParams() *DeleteAllocatorMetadataItemParams {
	return &DeleteAllocatorMetadataItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAllocatorMetadataItemParamsWithTimeout creates a new DeleteAllocatorMetadataItemParams object
// with the ability to set a timeout on a request.
func NewDeleteAllocatorMetadataItemParamsWithTimeout(timeout time.Duration) *DeleteAllocatorMetadataItemParams {
	return &DeleteAllocatorMetadataItemParams{
		timeout: timeout,
	}
}

// NewDeleteAllocatorMetadataItemParamsWithContext creates a new DeleteAllocatorMetadataItemParams object
// with the ability to set a context for a request.
func NewDeleteAllocatorMetadataItemParamsWithContext(ctx context.Context) *DeleteAllocatorMetadataItemParams {
	return &DeleteAllocatorMetadataItemParams{
		Context: ctx,
	}
}

// NewDeleteAllocatorMetadataItemParamsWithHTTPClient creates a new DeleteAllocatorMetadataItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAllocatorMetadataItemParamsWithHTTPClient(client *http.Client) *DeleteAllocatorMetadataItemParams {
	return &DeleteAllocatorMetadataItemParams{
		HTTPClient: client,
	}
}

/* DeleteAllocatorMetadataItemParams contains all the parameters to send to the API endpoint
   for the delete allocator metadata item operation.

   Typically these are written to a http.Request.
*/
type DeleteAllocatorMetadataItemParams struct {

	/* AllocatorID.

	   The allocator identifier.
	*/
	AllocatorID string

	/* Key.

	   The metadata item key.
	*/
	Key string

	/* Version.

	   Checks for conflicts against the metadata version, then returns the value in the `x-cloud-resource-version` header.
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete allocator metadata item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllocatorMetadataItemParams) WithDefaults() *DeleteAllocatorMetadataItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete allocator metadata item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllocatorMetadataItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) WithTimeout(timeout time.Duration) *DeleteAllocatorMetadataItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) WithContext(ctx context.Context) *DeleteAllocatorMetadataItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) WithHTTPClient(client *http.Client) *DeleteAllocatorMetadataItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatorID adds the allocatorID to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) WithAllocatorID(allocatorID string) *DeleteAllocatorMetadataItemParams {
	o.SetAllocatorID(allocatorID)
	return o
}

// SetAllocatorID adds the allocatorId to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) SetAllocatorID(allocatorID string) {
	o.AllocatorID = allocatorID
}

// WithKey adds the key to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) WithKey(key string) *DeleteAllocatorMetadataItemParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) SetKey(key string) {
	o.Key = key
}

// WithVersion adds the version to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) WithVersion(version *string) *DeleteAllocatorMetadataItemParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete allocator metadata item params
func (o *DeleteAllocatorMetadataItemParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAllocatorMetadataItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param allocator_id
	if err := r.SetPathParam("allocator_id", o.AllocatorID); err != nil {
		return err
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
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
