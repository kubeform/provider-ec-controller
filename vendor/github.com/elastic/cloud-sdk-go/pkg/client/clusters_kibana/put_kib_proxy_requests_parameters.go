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

// NewPutKibProxyRequestsParams creates a new PutKibProxyRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutKibProxyRequestsParams() *PutKibProxyRequestsParams {
	return &PutKibProxyRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutKibProxyRequestsParamsWithTimeout creates a new PutKibProxyRequestsParams object
// with the ability to set a timeout on a request.
func NewPutKibProxyRequestsParamsWithTimeout(timeout time.Duration) *PutKibProxyRequestsParams {
	return &PutKibProxyRequestsParams{
		timeout: timeout,
	}
}

// NewPutKibProxyRequestsParamsWithContext creates a new PutKibProxyRequestsParams object
// with the ability to set a context for a request.
func NewPutKibProxyRequestsParamsWithContext(ctx context.Context) *PutKibProxyRequestsParams {
	return &PutKibProxyRequestsParams{
		Context: ctx,
	}
}

// NewPutKibProxyRequestsParamsWithHTTPClient creates a new PutKibProxyRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutKibProxyRequestsParamsWithHTTPClient(client *http.Client) *PutKibProxyRequestsParams {
	return &PutKibProxyRequestsParams{
		HTTPClient: client,
	}
}

/* PutKibProxyRequestsParams contains all the parameters to send to the API endpoint
   for the put kib proxy requests operation.

   Typically these are written to a http.Request.
*/
type PutKibProxyRequestsParams struct {

	/* XManagementRequest.

	   When set to `true`, includes the X-Management-Request header value.
	*/
	XManagementRequest string

	/* Body.

	   The JSON payload that is used to proxy the Kibana deployment.
	*/
	Body string

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

// WithDefaults hydrates default values in the put kib proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutKibProxyRequestsParams) WithDefaults() *PutKibProxyRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put kib proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutKibProxyRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) WithTimeout(timeout time.Duration) *PutKibProxyRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) WithContext(ctx context.Context) *PutKibProxyRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) WithHTTPClient(client *http.Client) *PutKibProxyRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXManagementRequest adds the xManagementRequest to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) WithXManagementRequest(xManagementRequest string) *PutKibProxyRequestsParams {
	o.SetXManagementRequest(xManagementRequest)
	return o
}

// SetXManagementRequest adds the xManagementRequest to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) SetXManagementRequest(xManagementRequest string) {
	o.XManagementRequest = xManagementRequest
}

// WithBody adds the body to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) WithBody(body string) *PutKibProxyRequestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) SetBody(body string) {
	o.Body = body
}

// WithClusterID adds the clusterID to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) WithClusterID(clusterID string) *PutKibProxyRequestsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithKibanaPath adds the kibanaPath to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) WithKibanaPath(kibanaPath string) *PutKibProxyRequestsParams {
	o.SetKibanaPath(kibanaPath)
	return o
}

// SetKibanaPath adds the kibanaPath to the put kib proxy requests params
func (o *PutKibProxyRequestsParams) SetKibanaPath(kibanaPath string) {
	o.KibanaPath = kibanaPath
}

// WriteToRequest writes these params to a swagger request
func (o *PutKibProxyRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Management-Request
	if err := r.SetHeaderParam("X-Management-Request", o.XManagementRequest); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
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