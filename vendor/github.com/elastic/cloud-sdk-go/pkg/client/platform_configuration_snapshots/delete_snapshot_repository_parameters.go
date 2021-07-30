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

package platform_configuration_snapshots

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

// NewDeleteSnapshotRepositoryParams creates a new DeleteSnapshotRepositoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSnapshotRepositoryParams() *DeleteSnapshotRepositoryParams {
	return &DeleteSnapshotRepositoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnapshotRepositoryParamsWithTimeout creates a new DeleteSnapshotRepositoryParams object
// with the ability to set a timeout on a request.
func NewDeleteSnapshotRepositoryParamsWithTimeout(timeout time.Duration) *DeleteSnapshotRepositoryParams {
	return &DeleteSnapshotRepositoryParams{
		timeout: timeout,
	}
}

// NewDeleteSnapshotRepositoryParamsWithContext creates a new DeleteSnapshotRepositoryParams object
// with the ability to set a context for a request.
func NewDeleteSnapshotRepositoryParamsWithContext(ctx context.Context) *DeleteSnapshotRepositoryParams {
	return &DeleteSnapshotRepositoryParams{
		Context: ctx,
	}
}

// NewDeleteSnapshotRepositoryParamsWithHTTPClient creates a new DeleteSnapshotRepositoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSnapshotRepositoryParamsWithHTTPClient(client *http.Client) *DeleteSnapshotRepositoryParams {
	return &DeleteSnapshotRepositoryParams{
		HTTPClient: client,
	}
}

/* DeleteSnapshotRepositoryParams contains all the parameters to send to the API endpoint
   for the delete snapshot repository operation.

   Typically these are written to a http.Request.
*/
type DeleteSnapshotRepositoryParams struct {

	/* CleanupDeployments.

	   Removes references to this snapshot repository configuration and disables snapshots on the clusters that were referencing this configuration. If a request is made to delete a repository configuration that has already been deleted and this parameter is set to true and clusters still reference the configuration, then the request will have the side effects of removing references and disabling snapshots for clusters that reference the previously deleted configuration.
	*/
	CleanupDeployments *bool

	/* RepositoryName.

	   The name of the snapshot repository configuration.
	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete snapshot repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSnapshotRepositoryParams) WithDefaults() *DeleteSnapshotRepositoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete snapshot repository params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSnapshotRepositoryParams) SetDefaults() {
	var (
		cleanupDeploymentsDefault = bool(false)
	)

	val := DeleteSnapshotRepositoryParams{
		CleanupDeployments: &cleanupDeploymentsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete snapshot repository params
func (o *DeleteSnapshotRepositoryParams) WithTimeout(timeout time.Duration) *DeleteSnapshotRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete snapshot repository params
func (o *DeleteSnapshotRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete snapshot repository params
func (o *DeleteSnapshotRepositoryParams) WithContext(ctx context.Context) *DeleteSnapshotRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete snapshot repository params
func (o *DeleteSnapshotRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete snapshot repository params
func (o *DeleteSnapshotRepositoryParams) WithHTTPClient(client *http.Client) *DeleteSnapshotRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete snapshot repository params
func (o *DeleteSnapshotRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCleanupDeployments adds the cleanupDeployments to the delete snapshot repository params
func (o *DeleteSnapshotRepositoryParams) WithCleanupDeployments(cleanupDeployments *bool) *DeleteSnapshotRepositoryParams {
	o.SetCleanupDeployments(cleanupDeployments)
	return o
}

// SetCleanupDeployments adds the cleanupDeployments to the delete snapshot repository params
func (o *DeleteSnapshotRepositoryParams) SetCleanupDeployments(cleanupDeployments *bool) {
	o.CleanupDeployments = cleanupDeployments
}

// WithRepositoryName adds the repositoryName to the delete snapshot repository params
func (o *DeleteSnapshotRepositoryParams) WithRepositoryName(repositoryName string) *DeleteSnapshotRepositoryParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the delete snapshot repository params
func (o *DeleteSnapshotRepositoryParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnapshotRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CleanupDeployments != nil {

		// query param cleanup_deployments
		var qrCleanupDeployments bool

		if o.CleanupDeployments != nil {
			qrCleanupDeployments = *o.CleanupDeployments
		}
		qCleanupDeployments := swag.FormatBool(qrCleanupDeployments)
		if qCleanupDeployments != "" {

			if err := r.SetQueryParam("cleanup_deployments", qCleanupDeployments); err != nil {
				return err
			}
		}
	}

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}