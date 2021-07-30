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

package platform_configuration_trust_relationships

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

// NewCreateTrustRelationshipParams creates a new CreateTrustRelationshipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateTrustRelationshipParams() *CreateTrustRelationshipParams {
	return &CreateTrustRelationshipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTrustRelationshipParamsWithTimeout creates a new CreateTrustRelationshipParams object
// with the ability to set a timeout on a request.
func NewCreateTrustRelationshipParamsWithTimeout(timeout time.Duration) *CreateTrustRelationshipParams {
	return &CreateTrustRelationshipParams{
		timeout: timeout,
	}
}

// NewCreateTrustRelationshipParamsWithContext creates a new CreateTrustRelationshipParams object
// with the ability to set a context for a request.
func NewCreateTrustRelationshipParamsWithContext(ctx context.Context) *CreateTrustRelationshipParams {
	return &CreateTrustRelationshipParams{
		Context: ctx,
	}
}

// NewCreateTrustRelationshipParamsWithHTTPClient creates a new CreateTrustRelationshipParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateTrustRelationshipParamsWithHTTPClient(client *http.Client) *CreateTrustRelationshipParams {
	return &CreateTrustRelationshipParams{
		HTTPClient: client,
	}
}

/* CreateTrustRelationshipParams contains all the parameters to send to the API endpoint
   for the create trust relationship operation.

   Typically these are written to a http.Request.
*/
type CreateTrustRelationshipParams struct {

	/* Body.

	   The trust relationship definition
	*/
	Body *models.TrustRelationshipCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create trust relationship params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTrustRelationshipParams) WithDefaults() *CreateTrustRelationshipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create trust relationship params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateTrustRelationshipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create trust relationship params
func (o *CreateTrustRelationshipParams) WithTimeout(timeout time.Duration) *CreateTrustRelationshipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create trust relationship params
func (o *CreateTrustRelationshipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create trust relationship params
func (o *CreateTrustRelationshipParams) WithContext(ctx context.Context) *CreateTrustRelationshipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create trust relationship params
func (o *CreateTrustRelationshipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create trust relationship params
func (o *CreateTrustRelationshipParams) WithHTTPClient(client *http.Client) *CreateTrustRelationshipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create trust relationship params
func (o *CreateTrustRelationshipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create trust relationship params
func (o *CreateTrustRelationshipParams) WithBody(body *models.TrustRelationshipCreateRequest) *CreateTrustRelationshipParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create trust relationship params
func (o *CreateTrustRelationshipParams) SetBody(body *models.TrustRelationshipCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTrustRelationshipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
