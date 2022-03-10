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
)

// NewGetDeploymentTemplateV2Params creates a new GetDeploymentTemplateV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentTemplateV2Params() *GetDeploymentTemplateV2Params {
	return &GetDeploymentTemplateV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentTemplateV2ParamsWithTimeout creates a new GetDeploymentTemplateV2Params object
// with the ability to set a timeout on a request.
func NewGetDeploymentTemplateV2ParamsWithTimeout(timeout time.Duration) *GetDeploymentTemplateV2Params {
	return &GetDeploymentTemplateV2Params{
		timeout: timeout,
	}
}

// NewGetDeploymentTemplateV2ParamsWithContext creates a new GetDeploymentTemplateV2Params object
// with the ability to set a context for a request.
func NewGetDeploymentTemplateV2ParamsWithContext(ctx context.Context) *GetDeploymentTemplateV2Params {
	return &GetDeploymentTemplateV2Params{
		Context: ctx,
	}
}

// NewGetDeploymentTemplateV2ParamsWithHTTPClient creates a new GetDeploymentTemplateV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentTemplateV2ParamsWithHTTPClient(client *http.Client) *GetDeploymentTemplateV2Params {
	return &GetDeploymentTemplateV2Params{
		HTTPClient: client,
	}
}

/* GetDeploymentTemplateV2Params contains all the parameters to send to the API endpoint
   for the get deployment template v2 operation.

   Typically these are written to a http.Request.
*/
type GetDeploymentTemplateV2Params struct {

	/* Region.

	   Region of the deployment template
	*/
	Region string

	/* ShowInstanceConfigurations.

	   If true, will return details for each instance configuration referenced by the template.

	   Default: true
	*/
	ShowInstanceConfigurations *bool

	/* ShowMaxZones.

	   If true, will populate the max_zones field in the instance configurations. Only relevant if show_instance_configurations=true.
	*/
	ShowMaxZones *bool

	/* StackVersion.

	   If present, it will cause the returned deployment template to be adapted to return only the elements allowed in that version.
	*/
	StackVersion *string

	/* TemplateID.

	   The identifier for the deployment template.
	*/
	TemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployment template v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentTemplateV2Params) WithDefaults() *GetDeploymentTemplateV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployment template v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentTemplateV2Params) SetDefaults() {
	var (
		showInstanceConfigurationsDefault = bool(true)

		showMaxZonesDefault = bool(false)
	)

	val := GetDeploymentTemplateV2Params{
		ShowInstanceConfigurations: &showInstanceConfigurationsDefault,
		ShowMaxZones:               &showMaxZonesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) WithTimeout(timeout time.Duration) *GetDeploymentTemplateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) WithContext(ctx context.Context) *GetDeploymentTemplateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) WithHTTPClient(client *http.Client) *GetDeploymentTemplateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegion adds the region to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) WithRegion(region string) *GetDeploymentTemplateV2Params {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) SetRegion(region string) {
	o.Region = region
}

// WithShowInstanceConfigurations adds the showInstanceConfigurations to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) WithShowInstanceConfigurations(showInstanceConfigurations *bool) *GetDeploymentTemplateV2Params {
	o.SetShowInstanceConfigurations(showInstanceConfigurations)
	return o
}

// SetShowInstanceConfigurations adds the showInstanceConfigurations to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) SetShowInstanceConfigurations(showInstanceConfigurations *bool) {
	o.ShowInstanceConfigurations = showInstanceConfigurations
}

// WithShowMaxZones adds the showMaxZones to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) WithShowMaxZones(showMaxZones *bool) *GetDeploymentTemplateV2Params {
	o.SetShowMaxZones(showMaxZones)
	return o
}

// SetShowMaxZones adds the showMaxZones to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) SetShowMaxZones(showMaxZones *bool) {
	o.ShowMaxZones = showMaxZones
}

// WithStackVersion adds the stackVersion to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) WithStackVersion(stackVersion *string) *GetDeploymentTemplateV2Params {
	o.SetStackVersion(stackVersion)
	return o
}

// SetStackVersion adds the stackVersion to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) SetStackVersion(stackVersion *string) {
	o.StackVersion = stackVersion
}

// WithTemplateID adds the templateID to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) WithTemplateID(templateID string) *GetDeploymentTemplateV2Params {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the get deployment template v2 params
func (o *GetDeploymentTemplateV2Params) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentTemplateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param region
	qrRegion := o.Region
	qRegion := qrRegion
	if qRegion != "" {

		if err := r.SetQueryParam("region", qRegion); err != nil {
			return err
		}
	}

	if o.ShowInstanceConfigurations != nil {

		// query param show_instance_configurations
		var qrShowInstanceConfigurations bool

		if o.ShowInstanceConfigurations != nil {
			qrShowInstanceConfigurations = *o.ShowInstanceConfigurations
		}
		qShowInstanceConfigurations := swag.FormatBool(qrShowInstanceConfigurations)
		if qShowInstanceConfigurations != "" {

			if err := r.SetQueryParam("show_instance_configurations", qShowInstanceConfigurations); err != nil {
				return err
			}
		}
	}

	if o.ShowMaxZones != nil {

		// query param show_max_zones
		var qrShowMaxZones bool

		if o.ShowMaxZones != nil {
			qrShowMaxZones = *o.ShowMaxZones
		}
		qShowMaxZones := swag.FormatBool(qrShowMaxZones)
		if qShowMaxZones != "" {

			if err := r.SetQueryParam("show_max_zones", qShowMaxZones); err != nil {
				return err
			}
		}
	}

	if o.StackVersion != nil {

		// query param stack_version
		var qrStackVersion string

		if o.StackVersion != nil {
			qrStackVersion = *o.StackVersion
		}
		qStackVersion := qrStackVersion
		if qStackVersion != "" {

			if err := r.SetQueryParam("stack_version", qStackVersion); err != nil {
				return err
			}
		}
	}

	// path param template_id
	if err := r.SetPathParam("template_id", o.TemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
