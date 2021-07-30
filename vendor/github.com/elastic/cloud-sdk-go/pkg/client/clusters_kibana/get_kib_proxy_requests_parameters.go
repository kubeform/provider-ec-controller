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
)

// NewGetKibProxyRequestsParams creates a new GetKibProxyRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKibProxyRequestsParams() *GetKibProxyRequestsParams {
	return &GetKibProxyRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKibProxyRequestsParamsWithTimeout creates a new GetKibProxyRequestsParams object
// with the ability to set a timeout on a request.
func NewGetKibProxyRequestsParamsWithTimeout(timeout time.Duration) *GetKibProxyRequestsParams {
	return &GetKibProxyRequestsParams{
		timeout: timeout,
	}
}

// NewGetKibProxyRequestsParamsWithContext creates a new GetKibProxyRequestsParams object
// with the ability to set a context for a request.
func NewGetKibProxyRequestsParamsWithContext(ctx context.Context) *GetKibProxyRequestsParams {
	return &GetKibProxyRequestsParams{
		Context: ctx,
	}
}

// NewGetKibProxyRequestsParamsWithHTTPClient creates a new GetKibProxyRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKibProxyRequestsParamsWithHTTPClient(client *http.Client) *GetKibProxyRequestsParams {
	return &GetKibProxyRequestsParams{
		HTTPClient: client,
	}
}

/* GetKibProxyRequestsParams contains all the parameters to send to the API endpoint
   for the get kib proxy requests operation.

   Typically these are written to a http.Request.
*/
type GetKibProxyRequestsParams struct {

	/* XManagementRequest.

	   When set to `true`, includes the X-Management-Request header value.
	*/
	XManagementRequest string

	/* ClusterID.

	   The Kibana deployment identifier
	*/
	ClusterID string

	/* KibanaPath.

	   The URL part to proxy to the Kibana cluster. Example: /api/spaces/space or /api/upgrade_assistant/status
	*/
	KibanaPath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kib proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKibProxyRequestsParams) WithDefaults() *GetKibProxyRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kib proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKibProxyRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) WithTimeout(timeout time.Duration) *GetKibProxyRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) WithContext(ctx context.Context) *GetKibProxyRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) WithHTTPClient(client *http.Client) *GetKibProxyRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXManagementRequest adds the xManagementRequest to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) WithXManagementRequest(xManagementRequest string) *GetKibProxyRequestsParams {
	o.SetXManagementRequest(xManagementRequest)
	return o
}

// SetXManagementRequest adds the xManagementRequest to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) SetXManagementRequest(xManagementRequest string) {
	o.XManagementRequest = xManagementRequest
}

// WithClusterID adds the clusterID to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) WithClusterID(clusterID string) *GetKibProxyRequestsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithKibanaPath adds the kibanaPath to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) WithKibanaPath(kibanaPath string) *GetKibProxyRequestsParams {
	o.SetKibanaPath(kibanaPath)
	return o
}

// SetKibanaPath adds the kibanaPath to the get kib proxy requests params
func (o *GetKibProxyRequestsParams) SetKibanaPath(kibanaPath string) {
	o.KibanaPath = kibanaPath
}

// WriteToRequest writes these params to a swagger request
func (o *GetKibProxyRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Management-Request
	if err := r.SetHeaderParam("X-Management-Request", o.XManagementRequest); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param kibana_path
	if err := r.SetPathParam("kibana_path", o.KibanaPath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}