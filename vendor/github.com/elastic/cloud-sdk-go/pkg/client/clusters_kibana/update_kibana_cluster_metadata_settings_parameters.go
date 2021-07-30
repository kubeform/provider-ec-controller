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

package clusters_kibana

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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewUpdateKibanaClusterMetadataSettingsParams creates a new UpdateKibanaClusterMetadataSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateKibanaClusterMetadataSettingsParams() *UpdateKibanaClusterMetadataSettingsParams {
	return &UpdateKibanaClusterMetadataSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateKibanaClusterMetadataSettingsParamsWithTimeout creates a new UpdateKibanaClusterMetadataSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateKibanaClusterMetadataSettingsParamsWithTimeout(timeout time.Duration) *UpdateKibanaClusterMetadataSettingsParams {
	return &UpdateKibanaClusterMetadataSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateKibanaClusterMetadataSettingsParamsWithContext creates a new UpdateKibanaClusterMetadataSettingsParams object
// with the ability to set a context for a request.
func NewUpdateKibanaClusterMetadataSettingsParamsWithContext(ctx context.Context) *UpdateKibanaClusterMetadataSettingsParams {
	return &UpdateKibanaClusterMetadataSettingsParams{
		Context: ctx,
	}
}

// NewUpdateKibanaClusterMetadataSettingsParamsWithHTTPClient creates a new UpdateKibanaClusterMetadataSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateKibanaClusterMetadataSettingsParamsWithHTTPClient(client *http.Client) *UpdateKibanaClusterMetadataSettingsParams {
	return &UpdateKibanaClusterMetadataSettingsParams{
		HTTPClient: client,
	}
}

/* UpdateKibanaClusterMetadataSettingsParams contains all the parameters to send to the API endpoint
   for the update kibana cluster metadata settings operation.

   Typically these are written to a http.Request.
*/
type UpdateKibanaClusterMetadataSettingsParams struct {

	/* Body.

	   The cluster settings including updated values
	*/
	Body *models.ClusterMetadataSettings

	/* ClusterID.

	   The Kibana deployment identifier.
	*/
	ClusterID string

	/* Version.

	   Checks for conflicts against the metadata version, then returns the value in the `x-cloud-resource-version` header.
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update kibana cluster metadata settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateKibanaClusterMetadataSettingsParams) WithDefaults() *UpdateKibanaClusterMetadataSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update kibana cluster metadata settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateKibanaClusterMetadataSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) WithTimeout(timeout time.Duration) *UpdateKibanaClusterMetadataSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) WithContext(ctx context.Context) *UpdateKibanaClusterMetadataSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) WithHTTPClient(client *http.Client) *UpdateKibanaClusterMetadataSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) WithBody(body *models.ClusterMetadataSettings) *UpdateKibanaClusterMetadataSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) SetBody(body *models.ClusterMetadataSettings) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) WithClusterID(clusterID string) *UpdateKibanaClusterMetadataSettingsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithVersion adds the version to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) WithVersion(version *int64) *UpdateKibanaClusterMetadataSettingsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update kibana cluster metadata settings params
func (o *UpdateKibanaClusterMetadataSettingsParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateKibanaClusterMetadataSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
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