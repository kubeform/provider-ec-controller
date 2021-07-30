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

package clusters_elasticsearch

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

// NewGenerateEsClusterDiagnosticsParams creates a new GenerateEsClusterDiagnosticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateEsClusterDiagnosticsParams() *GenerateEsClusterDiagnosticsParams {
	return &GenerateEsClusterDiagnosticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateEsClusterDiagnosticsParamsWithTimeout creates a new GenerateEsClusterDiagnosticsParams object
// with the ability to set a timeout on a request.
func NewGenerateEsClusterDiagnosticsParamsWithTimeout(timeout time.Duration) *GenerateEsClusterDiagnosticsParams {
	return &GenerateEsClusterDiagnosticsParams{
		timeout: timeout,
	}
}

// NewGenerateEsClusterDiagnosticsParamsWithContext creates a new GenerateEsClusterDiagnosticsParams object
// with the ability to set a context for a request.
func NewGenerateEsClusterDiagnosticsParamsWithContext(ctx context.Context) *GenerateEsClusterDiagnosticsParams {
	return &GenerateEsClusterDiagnosticsParams{
		Context: ctx,
	}
}

// NewGenerateEsClusterDiagnosticsParamsWithHTTPClient creates a new GenerateEsClusterDiagnosticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateEsClusterDiagnosticsParamsWithHTTPClient(client *http.Client) *GenerateEsClusterDiagnosticsParams {
	return &GenerateEsClusterDiagnosticsParams{
		HTTPClient: client,
	}
}

/* GenerateEsClusterDiagnosticsParams contains all the parameters to send to the API endpoint
   for the generate es cluster diagnostics operation.

   Typically these are written to a http.Request.
*/
type GenerateEsClusterDiagnosticsParams struct {

	/* ClusterID.

	   The Elasticsearch cluster identifier.
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate es cluster diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateEsClusterDiagnosticsParams) WithDefaults() *GenerateEsClusterDiagnosticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate es cluster diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateEsClusterDiagnosticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generate es cluster diagnostics params
func (o *GenerateEsClusterDiagnosticsParams) WithTimeout(timeout time.Duration) *GenerateEsClusterDiagnosticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate es cluster diagnostics params
func (o *GenerateEsClusterDiagnosticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate es cluster diagnostics params
func (o *GenerateEsClusterDiagnosticsParams) WithContext(ctx context.Context) *GenerateEsClusterDiagnosticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate es cluster diagnostics params
func (o *GenerateEsClusterDiagnosticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate es cluster diagnostics params
func (o *GenerateEsClusterDiagnosticsParams) WithHTTPClient(client *http.Client) *GenerateEsClusterDiagnosticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate es cluster diagnostics params
func (o *GenerateEsClusterDiagnosticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the generate es cluster diagnostics params
func (o *GenerateEsClusterDiagnosticsParams) WithClusterID(clusterID string) *GenerateEsClusterDiagnosticsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the generate es cluster diagnostics params
func (o *GenerateEsClusterDiagnosticsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateEsClusterDiagnosticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
