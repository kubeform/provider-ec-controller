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

package deployments

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

// NewRestoreDeploymentResourceParams creates a new RestoreDeploymentResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoreDeploymentResourceParams() *RestoreDeploymentResourceParams {
	return &RestoreDeploymentResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreDeploymentResourceParamsWithTimeout creates a new RestoreDeploymentResourceParams object
// with the ability to set a timeout on a request.
func NewRestoreDeploymentResourceParamsWithTimeout(timeout time.Duration) *RestoreDeploymentResourceParams {
	return &RestoreDeploymentResourceParams{
		timeout: timeout,
	}
}

// NewRestoreDeploymentResourceParamsWithContext creates a new RestoreDeploymentResourceParams object
// with the ability to set a context for a request.
func NewRestoreDeploymentResourceParamsWithContext(ctx context.Context) *RestoreDeploymentResourceParams {
	return &RestoreDeploymentResourceParams{
		Context: ctx,
	}
}

// NewRestoreDeploymentResourceParamsWithHTTPClient creates a new RestoreDeploymentResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestoreDeploymentResourceParamsWithHTTPClient(client *http.Client) *RestoreDeploymentResourceParams {
	return &RestoreDeploymentResourceParams{
		HTTPClient: client,
	}
}

/* RestoreDeploymentResourceParams contains all the parameters to send to the API endpoint
   for the restore deployment resource operation.

   Typically these are written to a http.Request.
*/
type RestoreDeploymentResourceParams struct {

	/* DeploymentID.

	   Identifier for the Deployment
	*/
	DeploymentID string

	/* RefID.

	   User-specified RefId for the Resource
	*/
	RefID string

	/* ResourceKind.

	   The kind of resource
	*/
	ResourceKind string

	/* RestoreSnapshot.

	   Whether or not to restore a snapshot for those resources that allow it.
	*/
	RestoreSnapshot *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restore deployment resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreDeploymentResourceParams) WithDefaults() *RestoreDeploymentResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore deployment resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreDeploymentResourceParams) SetDefaults() {
	var (
		restoreSnapshotDefault = bool(false)
	)

	val := RestoreDeploymentResourceParams{
		RestoreSnapshot: &restoreSnapshotDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) WithTimeout(timeout time.Duration) *RestoreDeploymentResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) WithContext(ctx context.Context) *RestoreDeploymentResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) WithHTTPClient(client *http.Client) *RestoreDeploymentResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) WithDeploymentID(deploymentID string) *RestoreDeploymentResourceParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) WithRefID(refID string) *RestoreDeploymentResourceParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithResourceKind adds the resourceKind to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) WithResourceKind(resourceKind string) *RestoreDeploymentResourceParams {
	o.SetResourceKind(resourceKind)
	return o
}

// SetResourceKind adds the resourceKind to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) SetResourceKind(resourceKind string) {
	o.ResourceKind = resourceKind
}

// WithRestoreSnapshot adds the restoreSnapshot to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) WithRestoreSnapshot(restoreSnapshot *bool) *RestoreDeploymentResourceParams {
	o.SetRestoreSnapshot(restoreSnapshot)
	return o
}

// SetRestoreSnapshot adds the restoreSnapshot to the restore deployment resource params
func (o *RestoreDeploymentResourceParams) SetRestoreSnapshot(restoreSnapshot *bool) {
	o.RestoreSnapshot = restoreSnapshot
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreDeploymentResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	// path param ref_id
	if err := r.SetPathParam("ref_id", o.RefID); err != nil {
		return err
	}

	// path param resource_kind
	if err := r.SetPathParam("resource_kind", o.ResourceKind); err != nil {
		return err
	}

	if o.RestoreSnapshot != nil {

		// query param restore_snapshot
		var qrRestoreSnapshot bool

		if o.RestoreSnapshot != nil {
			qrRestoreSnapshot = *o.RestoreSnapshot
		}
		qRestoreSnapshot := swag.FormatBool(qrRestoreSnapshot)
		if qRestoreSnapshot != "" {

			if err := r.SetQueryParam("restore_snapshot", qRestoreSnapshot); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
