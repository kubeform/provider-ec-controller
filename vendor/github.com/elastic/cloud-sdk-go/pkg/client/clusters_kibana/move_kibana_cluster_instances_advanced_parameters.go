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

// NewMoveKibanaClusterInstancesAdvancedParams creates a new MoveKibanaClusterInstancesAdvancedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMoveKibanaClusterInstancesAdvancedParams() *MoveKibanaClusterInstancesAdvancedParams {
	return &MoveKibanaClusterInstancesAdvancedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMoveKibanaClusterInstancesAdvancedParamsWithTimeout creates a new MoveKibanaClusterInstancesAdvancedParams object
// with the ability to set a timeout on a request.
func NewMoveKibanaClusterInstancesAdvancedParamsWithTimeout(timeout time.Duration) *MoveKibanaClusterInstancesAdvancedParams {
	return &MoveKibanaClusterInstancesAdvancedParams{
		timeout: timeout,
	}
}

// NewMoveKibanaClusterInstancesAdvancedParamsWithContext creates a new MoveKibanaClusterInstancesAdvancedParams object
// with the ability to set a context for a request.
func NewMoveKibanaClusterInstancesAdvancedParamsWithContext(ctx context.Context) *MoveKibanaClusterInstancesAdvancedParams {
	return &MoveKibanaClusterInstancesAdvancedParams{
		Context: ctx,
	}
}

// NewMoveKibanaClusterInstancesAdvancedParamsWithHTTPClient creates a new MoveKibanaClusterInstancesAdvancedParams object
// with the ability to set a custom HTTPClient for a request.
func NewMoveKibanaClusterInstancesAdvancedParamsWithHTTPClient(client *http.Client) *MoveKibanaClusterInstancesAdvancedParams {
	return &MoveKibanaClusterInstancesAdvancedParams{
		HTTPClient: client,
	}
}

/* MoveKibanaClusterInstancesAdvancedParams contains all the parameters to send to the API endpoint
   for the move kibana cluster instances advanced operation.

   Typically these are written to a http.Request.
*/
type MoveKibanaClusterInstancesAdvancedParams struct {

	/* Body.

	   Overrides defaults for the move, including setting the configuration of instances specified in the path
	*/
	Body *models.TransientKibanaPlanConfiguration

	/* ClusterID.

	   The Kibana deployment identifier.
	*/
	ClusterID string

	/* ForceUpdate.

	   When `true`, cancels and overwrites the pending plans, or treats the instance as an error.
	*/
	ForceUpdate *bool

	/* ValidateOnly.

	   When `true`, validates the move request, then returns the calculated plan without applying the plan.
	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the move kibana cluster instances advanced params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MoveKibanaClusterInstancesAdvancedParams) WithDefaults() *MoveKibanaClusterInstancesAdvancedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the move kibana cluster instances advanced params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MoveKibanaClusterInstancesAdvancedParams) SetDefaults() {
	var (
		forceUpdateDefault = bool(false)

		validateOnlyDefault = bool(false)
	)

	val := MoveKibanaClusterInstancesAdvancedParams{
		ForceUpdate:  &forceUpdateDefault,
		ValidateOnly: &validateOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) WithTimeout(timeout time.Duration) *MoveKibanaClusterInstancesAdvancedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) WithContext(ctx context.Context) *MoveKibanaClusterInstancesAdvancedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) WithHTTPClient(client *http.Client) *MoveKibanaClusterInstancesAdvancedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) WithBody(body *models.TransientKibanaPlanConfiguration) *MoveKibanaClusterInstancesAdvancedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) SetBody(body *models.TransientKibanaPlanConfiguration) {
	o.Body = body
}

// WithClusterID adds the clusterID to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) WithClusterID(clusterID string) *MoveKibanaClusterInstancesAdvancedParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithForceUpdate adds the forceUpdate to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) WithForceUpdate(forceUpdate *bool) *MoveKibanaClusterInstancesAdvancedParams {
	o.SetForceUpdate(forceUpdate)
	return o
}

// SetForceUpdate adds the forceUpdate to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) SetForceUpdate(forceUpdate *bool) {
	o.ForceUpdate = forceUpdate
}

// WithValidateOnly adds the validateOnly to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) WithValidateOnly(validateOnly *bool) *MoveKibanaClusterInstancesAdvancedParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the move kibana cluster instances advanced params
func (o *MoveKibanaClusterInstancesAdvancedParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *MoveKibanaClusterInstancesAdvancedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ForceUpdate != nil {

		// query param force_update
		var qrForceUpdate bool

		if o.ForceUpdate != nil {
			qrForceUpdate = *o.ForceUpdate
		}
		qForceUpdate := swag.FormatBool(qrForceUpdate)
		if qForceUpdate != "" {

			if err := r.SetQueryParam("force_update", qForceUpdate); err != nil {
				return err
			}
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