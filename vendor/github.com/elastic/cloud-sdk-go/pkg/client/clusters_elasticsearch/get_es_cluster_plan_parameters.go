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

// NewGetEsClusterPlanParams creates a new GetEsClusterPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEsClusterPlanParams() *GetEsClusterPlanParams {
	return &GetEsClusterPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEsClusterPlanParamsWithTimeout creates a new GetEsClusterPlanParams object
// with the ability to set a timeout on a request.
func NewGetEsClusterPlanParamsWithTimeout(timeout time.Duration) *GetEsClusterPlanParams {
	return &GetEsClusterPlanParams{
		timeout: timeout,
	}
}

// NewGetEsClusterPlanParamsWithContext creates a new GetEsClusterPlanParams object
// with the ability to set a context for a request.
func NewGetEsClusterPlanParamsWithContext(ctx context.Context) *GetEsClusterPlanParams {
	return &GetEsClusterPlanParams{
		Context: ctx,
	}
}

// NewGetEsClusterPlanParamsWithHTTPClient creates a new GetEsClusterPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEsClusterPlanParamsWithHTTPClient(client *http.Client) *GetEsClusterPlanParams {
	return &GetEsClusterPlanParams{
		HTTPClient: client,
	}
}

/* GetEsClusterPlanParams contains all the parameters to send to the API endpoint
   for the get es cluster plan operation.

   Typically these are written to a http.Request.
*/
type GetEsClusterPlanParams struct {

	/* ClusterID.

	   The Elasticsearch cluster identifier.
	*/
	ClusterID string

	/* ConvertLegacyPlans.

	   When `true`, converts the plans to the 2.0.x format. When `false`, uses the 1.x format. The default is `false`.
	*/
	ConvertLegacyPlans *bool

	/* EnrichWithTemplate.

	   When plans are shown, includes the missing elements from the applicable deployment template.
	*/
	EnrichWithTemplate *bool

	/* ShowPlanDefaults.

	   When plans are shown, includes the default values in the response. NOTE: This option results in large responses.
	*/
	ShowPlanDefaults *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get es cluster plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEsClusterPlanParams) WithDefaults() *GetEsClusterPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get es cluster plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEsClusterPlanParams) SetDefaults() {
	var (
		convertLegacyPlansDefault = bool(false)

		enrichWithTemplateDefault = bool(false)

		showPlanDefaultsDefault = bool(false)
	)

	val := GetEsClusterPlanParams{
		ConvertLegacyPlans: &convertLegacyPlansDefault,
		EnrichWithTemplate: &enrichWithTemplateDefault,
		ShowPlanDefaults:   &showPlanDefaultsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get es cluster plan params
func (o *GetEsClusterPlanParams) WithTimeout(timeout time.Duration) *GetEsClusterPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get es cluster plan params
func (o *GetEsClusterPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get es cluster plan params
func (o *GetEsClusterPlanParams) WithContext(ctx context.Context) *GetEsClusterPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get es cluster plan params
func (o *GetEsClusterPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get es cluster plan params
func (o *GetEsClusterPlanParams) WithHTTPClient(client *http.Client) *GetEsClusterPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get es cluster plan params
func (o *GetEsClusterPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get es cluster plan params
func (o *GetEsClusterPlanParams) WithClusterID(clusterID string) *GetEsClusterPlanParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get es cluster plan params
func (o *GetEsClusterPlanParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithConvertLegacyPlans adds the convertLegacyPlans to the get es cluster plan params
func (o *GetEsClusterPlanParams) WithConvertLegacyPlans(convertLegacyPlans *bool) *GetEsClusterPlanParams {
	o.SetConvertLegacyPlans(convertLegacyPlans)
	return o
}

// SetConvertLegacyPlans adds the convertLegacyPlans to the get es cluster plan params
func (o *GetEsClusterPlanParams) SetConvertLegacyPlans(convertLegacyPlans *bool) {
	o.ConvertLegacyPlans = convertLegacyPlans
}

// WithEnrichWithTemplate adds the enrichWithTemplate to the get es cluster plan params
func (o *GetEsClusterPlanParams) WithEnrichWithTemplate(enrichWithTemplate *bool) *GetEsClusterPlanParams {
	o.SetEnrichWithTemplate(enrichWithTemplate)
	return o
}

// SetEnrichWithTemplate adds the enrichWithTemplate to the get es cluster plan params
func (o *GetEsClusterPlanParams) SetEnrichWithTemplate(enrichWithTemplate *bool) {
	o.EnrichWithTemplate = enrichWithTemplate
}

// WithShowPlanDefaults adds the showPlanDefaults to the get es cluster plan params
func (o *GetEsClusterPlanParams) WithShowPlanDefaults(showPlanDefaults *bool) *GetEsClusterPlanParams {
	o.SetShowPlanDefaults(showPlanDefaults)
	return o
}

// SetShowPlanDefaults adds the showPlanDefaults to the get es cluster plan params
func (o *GetEsClusterPlanParams) SetShowPlanDefaults(showPlanDefaults *bool) {
	o.ShowPlanDefaults = showPlanDefaults
}

// WriteToRequest writes these params to a swagger request
func (o *GetEsClusterPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if o.ConvertLegacyPlans != nil {

		// query param convert_legacy_plans
		var qrConvertLegacyPlans bool

		if o.ConvertLegacyPlans != nil {
			qrConvertLegacyPlans = *o.ConvertLegacyPlans
		}
		qConvertLegacyPlans := swag.FormatBool(qrConvertLegacyPlans)
		if qConvertLegacyPlans != "" {

			if err := r.SetQueryParam("convert_legacy_plans", qConvertLegacyPlans); err != nil {
				return err
			}
		}
	}

	if o.EnrichWithTemplate != nil {

		// query param enrich_with_template
		var qrEnrichWithTemplate bool

		if o.EnrichWithTemplate != nil {
			qrEnrichWithTemplate = *o.EnrichWithTemplate
		}
		qEnrichWithTemplate := swag.FormatBool(qrEnrichWithTemplate)
		if qEnrichWithTemplate != "" {

			if err := r.SetQueryParam("enrich_with_template", qEnrichWithTemplate); err != nil {
				return err
			}
		}
	}

	if o.ShowPlanDefaults != nil {

		// query param show_plan_defaults
		var qrShowPlanDefaults bool

		if o.ShowPlanDefaults != nil {
			qrShowPlanDefaults = *o.ShowPlanDefaults
		}
		qShowPlanDefaults := swag.FormatBool(qrShowPlanDefaults)
		if qShowPlanDefaults != "" {

			if err := r.SetQueryParam("show_plan_defaults", qShowPlanDefaults); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}