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

// NewResyncConstructorParams creates a new ResyncConstructorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResyncConstructorParams() *ResyncConstructorParams {
	return &ResyncConstructorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResyncConstructorParamsWithTimeout creates a new ResyncConstructorParams object
// with the ability to set a timeout on a request.
func NewResyncConstructorParamsWithTimeout(timeout time.Duration) *ResyncConstructorParams {
	return &ResyncConstructorParams{
		timeout: timeout,
	}
}

// NewResyncConstructorParamsWithContext creates a new ResyncConstructorParams object
// with the ability to set a context for a request.
func NewResyncConstructorParamsWithContext(ctx context.Context) *ResyncConstructorParams {
	return &ResyncConstructorParams{
		Context: ctx,
	}
}

// NewResyncConstructorParamsWithHTTPClient creates a new ResyncConstructorParams object
// with the ability to set a custom HTTPClient for a request.
func NewResyncConstructorParamsWithHTTPClient(client *http.Client) *ResyncConstructorParams {
	return &ResyncConstructorParams{
		HTTPClient: client,
	}
}

/* ResyncConstructorParams contains all the parameters to send to the API endpoint
   for the resync constructor operation.

   Typically these are written to a http.Request.
*/
type ResyncConstructorParams struct {

	/* ConstructorID.

	   Identifier for the constructor
	*/
	ConstructorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resync constructor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResyncConstructorParams) WithDefaults() *ResyncConstructorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resync constructor params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResyncConstructorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resync constructor params
func (o *ResyncConstructorParams) WithTimeout(timeout time.Duration) *ResyncConstructorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resync constructor params
func (o *ResyncConstructorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resync constructor params
func (o *ResyncConstructorParams) WithContext(ctx context.Context) *ResyncConstructorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resync constructor params
func (o *ResyncConstructorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resync constructor params
func (o *ResyncConstructorParams) WithHTTPClient(client *http.Client) *ResyncConstructorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resync constructor params
func (o *ResyncConstructorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConstructorID adds the constructorID to the resync constructor params
func (o *ResyncConstructorParams) WithConstructorID(constructorID string) *ResyncConstructorParams {
	o.SetConstructorID(constructorID)
	return o
}

// SetConstructorID adds the constructorId to the resync constructor params
func (o *ResyncConstructorParams) SetConstructorID(constructorID string) {
	o.ConstructorID = constructorID
}

// WriteToRequest writes these params to a swagger request
func (o *ResyncConstructorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
