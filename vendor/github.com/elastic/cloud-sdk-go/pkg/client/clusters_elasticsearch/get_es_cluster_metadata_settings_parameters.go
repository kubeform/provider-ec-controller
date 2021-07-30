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

// NewGetEsClusterMetadataSettingsParams creates a new GetEsClusterMetadataSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEsClusterMetadataSettingsParams() *GetEsClusterMetadataSettingsParams {
	return &GetEsClusterMetadataSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEsClusterMetadataSettingsParamsWithTimeout creates a new GetEsClusterMetadataSettingsParams object
// with the ability to set a timeout on a request.
func NewGetEsClusterMetadataSettingsParamsWithTimeout(timeout time.Duration) *GetEsClusterMetadataSettingsParams {
	return &GetEsClusterMetadataSettingsParams{
		timeout: timeout,
	}
}

// NewGetEsClusterMetadataSettingsParamsWithContext creates a new GetEsClusterMetadataSettingsParams object
// with the ability to set a context for a request.
func NewGetEsClusterMetadataSettingsParamsWithContext(ctx context.Context) *GetEsClusterMetadataSettingsParams {
	return &GetEsClusterMetadataSettingsParams{
		Context: ctx,
	}
}

// NewGetEsClusterMetadataSettingsParamsWithHTTPClient creates a new GetEsClusterMetadataSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEsClusterMetadataSettingsParamsWithHTTPClient(client *http.Client) *GetEsClusterMetadataSettingsParams {
	return &GetEsClusterMetadataSettingsParams{
		HTTPClient: client,
	}
}

/* GetEsClusterMetadataSettingsParams contains all the parameters to send to the API endpoint
   for the get es cluster metadata settings operation.

   Typically these are written to a http.Request.
*/
type GetEsClusterMetadataSettingsParams struct {

	/* ClusterID.

	   Elasticsearch cluster identifier
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get es cluster metadata settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEsClusterMetadataSettingsParams) WithDefaults() *GetEsClusterMetadataSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get es cluster metadata settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEsClusterMetadataSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get es cluster metadata settings params
func (o *GetEsClusterMetadataSettingsParams) WithTimeout(timeout time.Duration) *GetEsClusterMetadataSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get es cluster metadata settings params
func (o *GetEsClusterMetadataSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get es cluster metadata settings params
func (o *GetEsClusterMetadataSettingsParams) WithContext(ctx context.Context) *GetEsClusterMetadataSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get es cluster metadata settings params
func (o *GetEsClusterMetadataSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get es cluster metadata settings params
func (o *GetEsClusterMetadataSettingsParams) WithHTTPClient(client *http.Client) *GetEsClusterMetadataSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get es cluster metadata settings params
func (o *GetEsClusterMetadataSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get es cluster metadata settings params
func (o *GetEsClusterMetadataSettingsParams) WithClusterID(clusterID string) *GetEsClusterMetadataSettingsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get es cluster metadata settings params
func (o *GetEsClusterMetadataSettingsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEsClusterMetadataSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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