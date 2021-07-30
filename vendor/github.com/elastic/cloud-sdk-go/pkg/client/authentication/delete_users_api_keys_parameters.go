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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewDeleteUsersAPIKeysParams creates a new DeleteUsersAPIKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUsersAPIKeysParams() *DeleteUsersAPIKeysParams {
	return &DeleteUsersAPIKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsersAPIKeysParamsWithTimeout creates a new DeleteUsersAPIKeysParams object
// with the ability to set a timeout on a request.
func NewDeleteUsersAPIKeysParamsWithTimeout(timeout time.Duration) *DeleteUsersAPIKeysParams {
	return &DeleteUsersAPIKeysParams{
		timeout: timeout,
	}
}

// NewDeleteUsersAPIKeysParamsWithContext creates a new DeleteUsersAPIKeysParams object
// with the ability to set a context for a request.
func NewDeleteUsersAPIKeysParamsWithContext(ctx context.Context) *DeleteUsersAPIKeysParams {
	return &DeleteUsersAPIKeysParams{
		Context: ctx,
	}
}

// NewDeleteUsersAPIKeysParamsWithHTTPClient creates a new DeleteUsersAPIKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUsersAPIKeysParamsWithHTTPClient(client *http.Client) *DeleteUsersAPIKeysParams {
	return &DeleteUsersAPIKeysParams{
		HTTPClient: client,
	}
}

/* DeleteUsersAPIKeysParams contains all the parameters to send to the API endpoint
   for the delete users api keys operation.

   Typically these are written to a http.Request.
*/
type DeleteUsersAPIKeysParams struct {

	/* Body.

	   The request to delete API keys.
	*/
	Body *models.DeleteUsersAPIKeysRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete users api keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsersAPIKeysParams) WithDefaults() *DeleteUsersAPIKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete users api keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsersAPIKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete users api keys params
func (o *DeleteUsersAPIKeysParams) WithTimeout(timeout time.Duration) *DeleteUsersAPIKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete users api keys params
func (o *DeleteUsersAPIKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete users api keys params
func (o *DeleteUsersAPIKeysParams) WithContext(ctx context.Context) *DeleteUsersAPIKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete users api keys params
func (o *DeleteUsersAPIKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete users api keys params
func (o *DeleteUsersAPIKeysParams) WithHTTPClient(client *http.Client) *DeleteUsersAPIKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete users api keys params
func (o *DeleteUsersAPIKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete users api keys params
func (o *DeleteUsersAPIKeysParams) WithBody(body *models.DeleteUsersAPIKeysRequest) *DeleteUsersAPIKeysParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete users api keys params
func (o *DeleteUsersAPIKeysParams) SetBody(body *models.DeleteUsersAPIKeysRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsersAPIKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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