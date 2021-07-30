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

package deployment_templates

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

// NewSetDeploymentTemplateV2Params creates a new SetDeploymentTemplateV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetDeploymentTemplateV2Params() *SetDeploymentTemplateV2Params {
	return &SetDeploymentTemplateV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetDeploymentTemplateV2ParamsWithTimeout creates a new SetDeploymentTemplateV2Params object
// with the ability to set a timeout on a request.
func NewSetDeploymentTemplateV2ParamsWithTimeout(timeout time.Duration) *SetDeploymentTemplateV2Params {
	return &SetDeploymentTemplateV2Params{
		timeout: timeout,
	}
}

// NewSetDeploymentTemplateV2ParamsWithContext creates a new SetDeploymentTemplateV2Params object
// with the ability to set a context for a request.
func NewSetDeploymentTemplateV2ParamsWithContext(ctx context.Context) *SetDeploymentTemplateV2Params {
	return &SetDeploymentTemplateV2Params{
		Context: ctx,
	}
}

// NewSetDeploymentTemplateV2ParamsWithHTTPClient creates a new SetDeploymentTemplateV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewSetDeploymentTemplateV2ParamsWithHTTPClient(client *http.Client) *SetDeploymentTemplateV2Params {
	return &SetDeploymentTemplateV2Params{
		HTTPClient: client,
	}
}

/* SetDeploymentTemplateV2Params contains all the parameters to send to the API endpoint
   for the set deployment template v2 operation.

   Typically these are written to a http.Request.
*/
type SetDeploymentTemplateV2Params struct {

	/* Body.

	   The deployment template definition.
	*/
	Body *models.DeploymentTemplateRequestBody

	/* CreateOnly.

	   If true, will fail if the deployment template already exists at the given id
	*/
	CreateOnly *bool

	/* Region.

	   Region of the deployment template
	*/
	Region string

	/* TemplateID.

	   The identifier for the deployment template.
	*/
	TemplateID string

	/* Version.

	   If specified, checks for conflicts against the version of the template (returned in 'x-cloud-resource-version' of the GET request)
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set deployment template v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDeploymentTemplateV2Params) WithDefaults() *SetDeploymentTemplateV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set deployment template v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDeploymentTemplateV2Params) SetDefaults() {
	var (
		createOnlyDefault = bool(false)
	)

	val := SetDeploymentTemplateV2Params{
		CreateOnly: &createOnlyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) WithTimeout(timeout time.Duration) *SetDeploymentTemplateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) WithContext(ctx context.Context) *SetDeploymentTemplateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) WithHTTPClient(client *http.Client) *SetDeploymentTemplateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) WithBody(body *models.DeploymentTemplateRequestBody) *SetDeploymentTemplateV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) SetBody(body *models.DeploymentTemplateRequestBody) {
	o.Body = body
}

// WithCreateOnly adds the createOnly to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) WithCreateOnly(createOnly *bool) *SetDeploymentTemplateV2Params {
	o.SetCreateOnly(createOnly)
	return o
}

// SetCreateOnly adds the createOnly to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) SetCreateOnly(createOnly *bool) {
	o.CreateOnly = createOnly
}

// WithRegion adds the region to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) WithRegion(region string) *SetDeploymentTemplateV2Params {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) SetRegion(region string) {
	o.Region = region
}

// WithTemplateID adds the templateID to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) WithTemplateID(templateID string) *SetDeploymentTemplateV2Params {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WithVersion adds the version to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) WithVersion(version *string) *SetDeploymentTemplateV2Params {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the set deployment template v2 params
func (o *SetDeploymentTemplateV2Params) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SetDeploymentTemplateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CreateOnly != nil {

		// query param create_only
		var qrCreateOnly bool

		if o.CreateOnly != nil {
			qrCreateOnly = *o.CreateOnly
		}
		qCreateOnly := swag.FormatBool(qrCreateOnly)
		if qCreateOnly != "" {

			if err := r.SetQueryParam("create_only", qCreateOnly); err != nil {
				return err
			}
		}
	}

	// query param region
	qrRegion := o.Region
	qRegion := qrRegion
	if qRegion != "" {

		if err := r.SetQueryParam("region", qRegion); err != nil {
			return err
		}
	}

	// path param template_id
	if err := r.SetPathParam("template_id", o.TemplateID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
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
