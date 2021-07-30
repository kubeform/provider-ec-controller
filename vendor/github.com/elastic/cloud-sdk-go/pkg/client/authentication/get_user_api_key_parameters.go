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

package authentication

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

// NewGetUserAPIKeyParams creates a new GetUserAPIKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserAPIKeyParams() *GetUserAPIKeyParams {
	return &GetUserAPIKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserAPIKeyParamsWithTimeout creates a new GetUserAPIKeyParams object
// with the ability to set a timeout on a request.
func NewGetUserAPIKeyParamsWithTimeout(timeout time.Duration) *GetUserAPIKeyParams {
	return &GetUserAPIKeyParams{
		timeout: timeout,
	}
}

// NewGetUserAPIKeyParamsWithContext creates a new GetUserAPIKeyParams object
// with the ability to set a context for a request.
func NewGetUserAPIKeyParamsWithContext(ctx context.Context) *GetUserAPIKeyParams {
	return &GetUserAPIKeyParams{
		Context: ctx,
	}
}

// NewGetUserAPIKeyParamsWithHTTPClient creates a new GetUserAPIKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserAPIKeyParamsWithHTTPClient(client *http.Client) *GetUserAPIKeyParams {
	return &GetUserAPIKeyParams{
		HTTPClient: client,
	}
}

/* GetUserAPIKeyParams contains all the parameters to send to the API endpoint
   for the get user api key operation.

   Typically these are written to a http.Request.
*/
type GetUserAPIKeyParams struct {

	/* APIKeyID.

	   The API Key ID.
	*/
	APIKeyID string

	/* UserID.

	   The user ID.
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user api key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAPIKeyParams) WithDefaults() *GetUserAPIKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user api key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAPIKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user api key params
func (o *GetUserAPIKeyParams) WithTimeout(timeout time.Duration) *GetUserAPIKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user api key params
func (o *GetUserAPIKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user api key params
func (o *GetUserAPIKeyParams) WithContext(ctx context.Context) *GetUserAPIKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user api key params
func (o *GetUserAPIKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user api key params
func (o *GetUserAPIKeyParams) WithHTTPClient(client *http.Client) *GetUserAPIKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user api key params
func (o *GetUserAPIKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKeyID adds the aPIKeyID to the get user api key params
func (o *GetUserAPIKeyParams) WithAPIKeyID(aPIKeyID string) *GetUserAPIKeyParams {
	o.SetAPIKeyID(aPIKeyID)
	return o
}

// SetAPIKeyID adds the apiKeyId to the get user api key params
func (o *GetUserAPIKeyParams) SetAPIKeyID(aPIKeyID string) {
	o.APIKeyID = aPIKeyID
}

// WithUserID adds the userID to the get user api key params
func (o *GetUserAPIKeyParams) WithUserID(userID string) *GetUserAPIKeyParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user api key params
func (o *GetUserAPIKeyParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserAPIKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param api_key_id
	if err := r.SetPathParam("api_key_id", o.APIKeyID); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
