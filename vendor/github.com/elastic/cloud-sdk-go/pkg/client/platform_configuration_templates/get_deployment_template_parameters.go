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

package platform_configuration_templates

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

// NewGetDeploymentTemplateParams creates a new GetDeploymentTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentTemplateParams() *GetDeploymentTemplateParams {
	return &GetDeploymentTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentTemplateParamsWithTimeout creates a new GetDeploymentTemplateParams object
// with the ability to set a timeout on a request.
func NewGetDeploymentTemplateParamsWithTimeout(timeout time.Duration) *GetDeploymentTemplateParams {
	return &GetDeploymentTemplateParams{
		timeout: timeout,
	}
}

// NewGetDeploymentTemplateParamsWithContext creates a new GetDeploymentTemplateParams object
// with the ability to set a context for a request.
func NewGetDeploymentTemplateParamsWithContext(ctx context.Context) *GetDeploymentTemplateParams {
	return &GetDeploymentTemplateParams{
		Context: ctx,
	}
}

// NewGetDeploymentTemplateParamsWithHTTPClient creates a new GetDeploymentTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentTemplateParamsWithHTTPClient(client *http.Client) *GetDeploymentTemplateParams {
	return &GetDeploymentTemplateParams{
		HTTPClient: client,
	}
}

/* GetDeploymentTemplateParams contains all the parameters to send to the API endpoint
   for the get deployment template operation.

   Typically these are written to a http.Request.
*/
type GetDeploymentTemplateParams struct {

	/* Format.

	   If 'cluster' is specified populates cluster_template in the response, if 'deployment' is specified populates deployment_template in the response

	   Default: "cluster"
	*/
	Format *string

	/* ShowInstanceConfigurations.

	   If true, will return details for each instance configuration referenced by the template.
	*/
	ShowInstanceConfigurations *bool

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

// WithDefaults hydrates default values in the get deployment template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentTemplateParams) WithDefaults() *GetDeploymentTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployment template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentTemplateParams) SetDefaults() {
	var (
		formatDefault = string("cluster")

		showInstanceConfigurationsDefault = bool(false)
	)

	val := GetDeploymentTemplateParams{
		Format:                     &formatDefault,
		ShowInstanceConfigurations: &showInstanceConfigurationsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get deployment template params
func (o *GetDeploymentTemplateParams) WithTimeout(timeout time.Duration) *GetDeploymentTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment template params
func (o *GetDeploymentTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment template params
func (o *GetDeploymentTemplateParams) WithContext(ctx context.Context) *GetDeploymentTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment template params
func (o *GetDeploymentTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment template params
func (o *GetDeploymentTemplateParams) WithHTTPClient(client *http.Client) *GetDeploymentTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment template params
func (o *GetDeploymentTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the get deployment template params
func (o *GetDeploymentTemplateParams) WithFormat(format *string) *GetDeploymentTemplateParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get deployment template params
func (o *GetDeploymentTemplateParams) SetFormat(format *string) {
	o.Format = format
}

// WithShowInstanceConfigurations adds the showInstanceConfigurations to the get deployment template params
func (o *GetDeploymentTemplateParams) WithShowInstanceConfigurations(showInstanceConfigurations *bool) *GetDeploymentTemplateParams {
	o.SetShowInstanceConfigurations(showInstanceConfigurations)
	return o
}

// SetShowInstanceConfigurations adds the showInstanceConfigurations to the get deployment template params
func (o *GetDeploymentTemplateParams) SetShowInstanceConfigurations(showInstanceConfigurations *bool) {
	o.ShowInstanceConfigurations = showInstanceConfigurations
}

// WithStackVersion adds the stackVersion to the get deployment template params
func (o *GetDeploymentTemplateParams) WithStackVersion(stackVersion *string) *GetDeploymentTemplateParams {
	o.SetStackVersion(stackVersion)
	return o
}

// SetStackVersion adds the stackVersion to the get deployment template params
func (o *GetDeploymentTemplateParams) SetStackVersion(stackVersion *string) {
	o.StackVersion = stackVersion
}

// WithTemplateID adds the templateID to the get deployment template params
func (o *GetDeploymentTemplateParams) WithTemplateID(templateID string) *GetDeploymentTemplateParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the get deployment template params
func (o *GetDeploymentTemplateParams) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Format != nil {

		// query param format
		var qrFormat string

		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {

			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
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
