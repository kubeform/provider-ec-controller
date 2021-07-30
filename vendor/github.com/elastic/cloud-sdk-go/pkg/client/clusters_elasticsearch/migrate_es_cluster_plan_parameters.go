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
)

// NewMigrateEsClusterPlanParams creates a new MigrateEsClusterPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMigrateEsClusterPlanParams() *MigrateEsClusterPlanParams {
	return &MigrateEsClusterPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMigrateEsClusterPlanParamsWithTimeout creates a new MigrateEsClusterPlanParams object
// with the ability to set a timeout on a request.
func NewMigrateEsClusterPlanParamsWithTimeout(timeout time.Duration) *MigrateEsClusterPlanParams {
	return &MigrateEsClusterPlanParams{
		timeout: timeout,
	}
}

// NewMigrateEsClusterPlanParamsWithContext creates a new MigrateEsClusterPlanParams object
// with the ability to set a context for a request.
func NewMigrateEsClusterPlanParamsWithContext(ctx context.Context) *MigrateEsClusterPlanParams {
	return &MigrateEsClusterPlanParams{
		Context: ctx,
	}
}

// NewMigrateEsClusterPlanParamsWithHTTPClient creates a new MigrateEsClusterPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewMigrateEsClusterPlanParamsWithHTTPClient(client *http.Client) *MigrateEsClusterPlanParams {
	return &MigrateEsClusterPlanParams{
		HTTPClient: client,
	}
}

/* MigrateEsClusterPlanParams contains all the parameters to send to the API endpoint
   for the migrate es cluster plan operation.

   Typically these are written to a http.Request.
*/
type MigrateEsClusterPlanParams struct {

	/* ClusterID.

	   The Elasticsearch cluster identifier.
	*/
	ClusterID string

	/* Template.

	   The ID of the deployment template to migrate to
	*/
	Template string

	/* ValidateOnly.

	   When true, validates the cluster definition, but does not perform the update
	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the migrate es cluster plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MigrateEsClusterPlanParams) WithDefaults() *MigrateEsClusterPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the migrate es cluster plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MigrateEsClusterPlanParams) SetDefaults() {
	var (
		validateOnlyDefault = bool(false)
	)

	val := MigrateEsClusterPlanParams{
		ValidateOnly: &validateOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) WithTimeout(timeout time.Duration) *MigrateEsClusterPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) WithContext(ctx context.Context) *MigrateEsClusterPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) WithHTTPClient(client *http.Client) *MigrateEsClusterPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) WithClusterID(clusterID string) *MigrateEsClusterPlanParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithTemplate adds the template to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) WithTemplate(template string) *MigrateEsClusterPlanParams {
	o.SetTemplate(template)
	return o
}

// SetTemplate adds the template to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) SetTemplate(template string) {
	o.Template = template
}

// WithValidateOnly adds the validateOnly to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) WithValidateOnly(validateOnly *bool) *MigrateEsClusterPlanParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the migrate es cluster plan params
func (o *MigrateEsClusterPlanParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *MigrateEsClusterPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// query param template
	qrTemplate := o.Template
	qTemplate := qrTemplate
	if qTemplate != "" {

		if err := r.SetQueryParam("template", qTemplate); err != nil {
			return err
		}
	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool

		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {

			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}