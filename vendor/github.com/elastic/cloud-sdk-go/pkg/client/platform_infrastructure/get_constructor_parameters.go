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

// NewGetConstructorParams creates a new GetConstructorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConstructorParams() *GetConstructorParams {
	return &GetConstructorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConstructorParamsWithTimeout creates a new GetConstructorParams object
// with the ability to set a timeout on a request.
func NewGetConstructorParamsWithTimeout(timeout time.Duration) *GetConstructorParams {
	return &GetConstructorParams{
		timeout: timeout,
	}
}

// NewGetConstructorParamsWithContext creates a new GetConstructorParams object
// with the ability to set a context for a request.
func NewGetConstructorParamsWithContext(ctx context.Context) *GetConstructorParams {
	return &GetConstructorParams{
		Context: ctx,
	}
}

// NewGetConstructorParamsWithHTTPClient creates a new GetConstructorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConstructorParamsWithHTTPClient(client *http.Client) *GetConstructorParams {
	return &GetConstructorParams{
		HTTPClient: client,
	}
}

/* GetConstructorParams contains all the parameters to send to the API endpoint
   for the get constructor operation.

   Typically these are written to a http.Request.
*/
type GetConstructorParams struct {

	/* ConstructorID.

	   Identifier for the constructor
	*/
	ConstructorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get constructor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConstructorParams) WithDefaults() *GetConstructorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get constructor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConstructorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get constructor params
func (o *GetConstructorParams) WithTimeout(timeout time.Duration) *GetConstructorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get constructor params
func (o *GetConstructorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get constructor params
func (o *GetConstructorParams) WithContext(ctx context.Context) *GetConstructorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get constructor params
func (o *GetConstructorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get constructor params
func (o *GetConstructorParams) WithHTTPClient(client *http.Client) *GetConstructorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get constructor params
func (o *GetConstructorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstructorID adds the constructorID to the get constructor params
func (o *GetConstructorParams) WithConstructorID(constructorID string) *GetConstructorParams {
	o.SetConstructorID(constructorID)
	return o
}

// SetConstructorID adds the constructorId to the get constructor params
func (o *GetConstructorParams) SetConstructorID(constructorID string) {
	o.ConstructorID = constructorID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConstructorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param constructor_id
	if err := r.SetPathParam("constructor_id", o.ConstructorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}