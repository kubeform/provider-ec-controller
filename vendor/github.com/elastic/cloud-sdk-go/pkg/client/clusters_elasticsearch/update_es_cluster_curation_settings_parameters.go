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
	"github.com/go-openapi/swag"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewUpdateEsClusterCurationSettingsParams creates a new UpdateEsClusterCurationSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateEsClusterCurationSettingsParams() *UpdateEsClusterCurationSettingsParams {
	return &UpdateEsClusterCurationSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEsClusterCurationSettingsParamsWithTimeout creates a new UpdateEsClusterCurationSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateEsClusterCurationSettingsParamsWithTimeout(timeout time.Duration) *UpdateEsClusterCurationSettingsParams {
	return &UpdateEsClusterCurationSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateEsClusterCurationSettingsParamsWithContext creates a new UpdateEsClusterCurationSettingsParams object
// with the ability to set a context for a request.
func NewUpdateEsClusterCurationSettingsParamsWithContext(ctx context.Context) *UpdateEsClusterCurationSettingsParams {
	return &UpdateEsClusterCurationSettingsParams{
		Context: ctx,
	}
}

// NewUpdateEsClusterCurationSettingsParamsWithHTTPClient creates a new UpdateEsClusterCurationSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateEsClusterCurationSettingsParamsWithHTTPClient(client *http.Client) *UpdateEsClusterCurationSettingsParams {
	return &UpdateEsClusterCurationSettingsParams{
		HTTPClient: client,
	}
}

/* UpdateEsClusterCurationSettingsParams contains all the parameters to send to the API endpoint
   for the update es cluster curation settings operation.

   Typically these are written to a http.Request.
*/
type UpdateEsClusterCurationSettingsParams struct {

	/* Body.

	   The cluster curation settings including updated values
	*/
	Body *models.ClusterCurationSettings

	/* ClusterID.

	   The Elasticsearch cluster identifier.
	*/
	ClusterID string

	/* Version.

	   If specified then checks for conflicts against the version of the cluster curation settings (returned in 'x-cloud-resource-version' of the GET request)
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update es cluster curation settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEsClusterCurationSettingsParams) WithDefaults() *UpdateEsClusterCurationSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update es cluster curation settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEsClusterCurationSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) WithTimeout(timeout time.Duration) *UpdateEsClusterCurationSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) WithContext(ctx context.Context) *UpdateEsClusterCurationSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) WithHTTPClient(client *http.Client) *UpdateEsClusterCurationSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) WithBody(body *models.ClusterCurationSettings) *UpdateEsClusterCurationSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) SetBody(body *models.ClusterCurationSettings) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) WithClusterID(clusterID string) *UpdateEsClusterCurationSettingsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithVersion adds the version to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) WithVersion(version *int64) *UpdateEsClusterCurationSettingsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update es cluster curation settings params
func (o *UpdateEsClusterCurationSettingsParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEsClusterCurationSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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