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

package deployments_traffic_filter

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

// NewDeleteTrafficFilterRulesetAssociationParams creates a new DeleteTrafficFilterRulesetAssociationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTrafficFilterRulesetAssociationParams() *DeleteTrafficFilterRulesetAssociationParams {
	return &DeleteTrafficFilterRulesetAssociationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTrafficFilterRulesetAssociationParamsWithTimeout creates a new DeleteTrafficFilterRulesetAssociationParams object
// with the ability to set a timeout on a request.
func NewDeleteTrafficFilterRulesetAssociationParamsWithTimeout(timeout time.Duration) *DeleteTrafficFilterRulesetAssociationParams {
	return &DeleteTrafficFilterRulesetAssociationParams{
		timeout: timeout,
	}
}

// NewDeleteTrafficFilterRulesetAssociationParamsWithContext creates a new DeleteTrafficFilterRulesetAssociationParams object
// with the ability to set a context for a request.
func NewDeleteTrafficFilterRulesetAssociationParamsWithContext(ctx context.Context) *DeleteTrafficFilterRulesetAssociationParams {
	return &DeleteTrafficFilterRulesetAssociationParams{
		Context: ctx,
	}
}

// NewDeleteTrafficFilterRulesetAssociationParamsWithHTTPClient creates a new DeleteTrafficFilterRulesetAssociationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTrafficFilterRulesetAssociationParamsWithHTTPClient(client *http.Client) *DeleteTrafficFilterRulesetAssociationParams {
	return &DeleteTrafficFilterRulesetAssociationParams{
		HTTPClient: client,
	}
}

/* DeleteTrafficFilterRulesetAssociationParams contains all the parameters to send to the API endpoint
   for the delete traffic filter ruleset association operation.

   Typically these are written to a http.Request.
*/
type DeleteTrafficFilterRulesetAssociationParams struct {

	/* AssociatedEntityID.

	   Associated entity ID
	*/
	AssociatedEntityID string

	/* AssociationType.

	   Association type
	*/
	AssociationType string

	/* RulesetID.

	   The mandatory ruleset ID.
	*/
	RulesetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete traffic filter ruleset association params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTrafficFilterRulesetAssociationParams) WithDefaults() *DeleteTrafficFilterRulesetAssociationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete traffic filter ruleset association params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTrafficFilterRulesetAssociationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) WithTimeout(timeout time.Duration) *DeleteTrafficFilterRulesetAssociationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) WithContext(ctx context.Context) *DeleteTrafficFilterRulesetAssociationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) WithHTTPClient(client *http.Client) *DeleteTrafficFilterRulesetAssociationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssociatedEntityID adds the associatedEntityID to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) WithAssociatedEntityID(associatedEntityID string) *DeleteTrafficFilterRulesetAssociationParams {
	o.SetAssociatedEntityID(associatedEntityID)
	return o
}

// SetAssociatedEntityID adds the associatedEntityId to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) SetAssociatedEntityID(associatedEntityID string) {
	o.AssociatedEntityID = associatedEntityID
}

// WithAssociationType adds the associationType to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) WithAssociationType(associationType string) *DeleteTrafficFilterRulesetAssociationParams {
	o.SetAssociationType(associationType)
	return o
}

// SetAssociationType adds the associationType to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) SetAssociationType(associationType string) {
	o.AssociationType = associationType
}

// WithRulesetID adds the rulesetID to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) WithRulesetID(rulesetID string) *DeleteTrafficFilterRulesetAssociationParams {
	o.SetRulesetID(rulesetID)
	return o
}

// SetRulesetID adds the rulesetId to the delete traffic filter ruleset association params
func (o *DeleteTrafficFilterRulesetAssociationParams) SetRulesetID(rulesetID string) {
	o.RulesetID = rulesetID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTrafficFilterRulesetAssociationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param associated_entity_id
	if err := r.SetPathParam("associated_entity_id", o.AssociatedEntityID); err != nil {
		return err
	}

	// path param association_type
	if err := r.SetPathParam("association_type", o.AssociationType); err != nil {
		return err
	}

	// path param ruleset_id
	if err := r.SetPathParam("ruleset_id", o.RulesetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
