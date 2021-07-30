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

// NewUpdateEsClusterPlanParams creates a new UpdateEsClusterPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateEsClusterPlanParams() *UpdateEsClusterPlanParams {
	return &UpdateEsClusterPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEsClusterPlanParamsWithTimeout creates a new UpdateEsClusterPlanParams object
// with the ability to set a timeout on a request.
func NewUpdateEsClusterPlanParamsWithTimeout(timeout time.Duration) *UpdateEsClusterPlanParams {
	return &UpdateEsClusterPlanParams{
		timeout: timeout,
	}
}

// NewUpdateEsClusterPlanParamsWithContext creates a new UpdateEsClusterPlanParams object
// with the ability to set a context for a request.
func NewUpdateEsClusterPlanParamsWithContext(ctx context.Context) *UpdateEsClusterPlanParams {
	return &UpdateEsClusterPlanParams{
		Context: ctx,
	}
}

// NewUpdateEsClusterPlanParamsWithHTTPClient creates a new UpdateEsClusterPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateEsClusterPlanParamsWithHTTPClient(client *http.Client) *UpdateEsClusterPlanParams {
	return &UpdateEsClusterPlanParams{
		HTTPClient: client,
	}
}

/* UpdateEsClusterPlanParams contains all the parameters to send to the API endpoint
   for the update es cluster plan operation.

   Typically these are written to a http.Request.
*/
type UpdateEsClusterPlanParams struct {

	/* Body.

	   The update plan definition
	*/
	Body *models.ElasticsearchClusterPlan

	/* ClusterID.

	   The Elasticsearch cluster identifier.
	*/
	ClusterID string

	/* ValidateOnly.

	   When `true`, validates the cluster definition without performing the update.
	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update es cluster plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEsClusterPlanParams) WithDefaults() *UpdateEsClusterPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update es cluster plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateEsClusterPlanParams) SetDefaults() {
	var (
		validateOnlyDefault = bool(false)
	)

	val := UpdateEsClusterPlanParams{
		ValidateOnly: &validateOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) WithTimeout(timeout time.Duration) *UpdateEsClusterPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) WithContext(ctx context.Context) *UpdateEsClusterPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) WithHTTPClient(client *http.Client) *UpdateEsClusterPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) WithBody(body *models.ElasticsearchClusterPlan) *UpdateEsClusterPlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) SetBody(body *models.ElasticsearchClusterPlan) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) WithClusterID(clusterID string) *UpdateEsClusterPlanParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithValidateOnly adds the validateOnly to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) WithValidateOnly(validateOnly *bool) *UpdateEsClusterPlanParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the update es cluster plan params
func (o *UpdateEsClusterPlanParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEsClusterPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
