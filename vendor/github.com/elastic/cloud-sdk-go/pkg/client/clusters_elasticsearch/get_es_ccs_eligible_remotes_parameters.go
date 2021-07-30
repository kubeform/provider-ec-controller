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

// NewGetEsCcsEligibleRemotesParams creates a new GetEsCcsEligibleRemotesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEsCcsEligibleRemotesParams() *GetEsCcsEligibleRemotesParams {
	return &GetEsCcsEligibleRemotesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEsCcsEligibleRemotesParamsWithTimeout creates a new GetEsCcsEligibleRemotesParams object
// with the ability to set a timeout on a request.
func NewGetEsCcsEligibleRemotesParamsWithTimeout(timeout time.Duration) *GetEsCcsEligibleRemotesParams {
	return &GetEsCcsEligibleRemotesParams{
		timeout: timeout,
	}
}

// NewGetEsCcsEligibleRemotesParamsWithContext creates a new GetEsCcsEligibleRemotesParams object
// with the ability to set a context for a request.
func NewGetEsCcsEligibleRemotesParamsWithContext(ctx context.Context) *GetEsCcsEligibleRemotesParams {
	return &GetEsCcsEligibleRemotesParams{
		Context: ctx,
	}
}

// NewGetEsCcsEligibleRemotesParamsWithHTTPClient creates a new GetEsCcsEligibleRemotesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEsCcsEligibleRemotesParamsWithHTTPClient(client *http.Client) *GetEsCcsEligibleRemotesParams {
	return &GetEsCcsEligibleRemotesParams{
		HTTPClient: client,
	}
}

/* GetEsCcsEligibleRemotesParams contains all the parameters to send to the API endpoint
   for the get es ccs eligible remotes operation.

   Typically these are written to a http.Request.
*/
type GetEsCcsEligibleRemotesParams struct {

	/* OwnerID.

	   (Optional) Returns only clusters filtered by the provided owner id.
	*/
	OwnerID *string

	/* Q.

	   (Optional) Cluster name or id prefix to filters the candidates.
	*/
	Q *string

	/* Size.

	   (Optional) Maximum number of clusters to include in the response.
	*/
	Size *int64

	/* Version.

	   The version of the cross-cluster search cluster that will link with the remote candidates.
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get es ccs eligible remotes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEsCcsEligibleRemotesParams) WithDefaults() *GetEsCcsEligibleRemotesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get es ccs eligible remotes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEsCcsEligibleRemotesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) WithTimeout(timeout time.Duration) *GetEsCcsEligibleRemotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) WithContext(ctx context.Context) *GetEsCcsEligibleRemotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) WithHTTPClient(client *http.Client) *GetEsCcsEligibleRemotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwnerID adds the ownerID to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) WithOwnerID(ownerID *string) *GetEsCcsEligibleRemotesParams {
	o.SetOwnerID(ownerID)
	return o
}

// SetOwnerID adds the ownerId to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) SetOwnerID(ownerID *string) {
	o.OwnerID = ownerID
}

// WithQ adds the q to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) WithQ(q *string) *GetEsCcsEligibleRemotesParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) SetQ(q *string) {
	o.Q = q
}

// WithSize adds the size to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) WithSize(size *int64) *GetEsCcsEligibleRemotesParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) SetSize(size *int64) {
	o.Size = size
}

// WithVersion adds the version to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) WithVersion(version string) *GetEsCcsEligibleRemotesParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get es ccs eligible remotes params
func (o *GetEsCcsEligibleRemotesParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetEsCcsEligibleRemotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OwnerID != nil {

		// query param owner_id
		var qrOwnerID string

		if o.OwnerID != nil {
			qrOwnerID = *o.OwnerID
		}
		qOwnerID := qrOwnerID
		if qOwnerID != "" {

			if err := r.SetQueryParam("owner_id", qOwnerID); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int64

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {

		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
