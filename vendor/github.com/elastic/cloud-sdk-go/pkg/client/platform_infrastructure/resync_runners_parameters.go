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
)

// NewResyncRunnersParams creates a new ResyncRunnersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResyncRunnersParams() *ResyncRunnersParams {
	return &ResyncRunnersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResyncRunnersParamsWithTimeout creates a new ResyncRunnersParams object
// with the ability to set a timeout on a request.
func NewResyncRunnersParamsWithTimeout(timeout time.Duration) *ResyncRunnersParams {
	return &ResyncRunnersParams{
		timeout: timeout,
	}
}

// NewResyncRunnersParamsWithContext creates a new ResyncRunnersParams object
// with the ability to set a context for a request.
func NewResyncRunnersParamsWithContext(ctx context.Context) *ResyncRunnersParams {
	return &ResyncRunnersParams{
		Context: ctx,
	}
}

// NewResyncRunnersParamsWithHTTPClient creates a new ResyncRunnersParams object
// with the ability to set a custom HTTPClient for a request.
func NewResyncRunnersParamsWithHTTPClient(client *http.Client) *ResyncRunnersParams {
	return &ResyncRunnersParams{
		HTTPClient: client,
	}
}

/* ResyncRunnersParams contains all the parameters to send to the API endpoint
   for the resync runners operation.

   Typically these are written to a http.Request.
*/
type ResyncRunnersParams struct {

	/* SkipMatchingVersion.

	   When true, skips the document indexing when the version matches the in-memory copy.

	   Default: true
	*/
	SkipMatchingVersion *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resync runners params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResyncRunnersParams) WithDefaults() *ResyncRunnersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resync runners params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResyncRunnersParams) SetDefaults() {
	var (
		skipMatchingVersionDefault = bool(true)
	)

	val := ResyncRunnersParams{
		SkipMatchingVersion: &skipMatchingVersionDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the resync runners params
func (o *ResyncRunnersParams) WithTimeout(timeout time.Duration) *ResyncRunnersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resync runners params
func (o *ResyncRunnersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resync runners params
func (o *ResyncRunnersParams) WithContext(ctx context.Context) *ResyncRunnersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resync runners params
func (o *ResyncRunnersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resync runners params
func (o *ResyncRunnersParams) WithHTTPClient(client *http.Client) *ResyncRunnersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resync runners params
func (o *ResyncRunnersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkipMatchingVersion adds the skipMatchingVersion to the resync runners params
func (o *ResyncRunnersParams) WithSkipMatchingVersion(skipMatchingVersion *bool) *ResyncRunnersParams {
	o.SetSkipMatchingVersion(skipMatchingVersion)
	return o
}

// SetSkipMatchingVersion adds the skipMatchingVersion to the resync runners params
func (o *ResyncRunnersParams) SetSkipMatchingVersion(skipMatchingVersion *bool) {
	o.SkipMatchingVersion = skipMatchingVersion
}

// WriteToRequest writes these params to a swagger request
func (o *ResyncRunnersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SkipMatchingVersion != nil {

		// query param skip_matching_version
		var qrSkipMatchingVersion bool

		if o.SkipMatchingVersion != nil {
			qrSkipMatchingVersion = *o.SkipMatchingVersion
		}
		qSkipMatchingVersion := swag.FormatBool(qrSkipMatchingVersion)
		if qSkipMatchingVersion != "" {

			if err := r.SetQueryParam("skip_matching_version", qSkipMatchingVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
