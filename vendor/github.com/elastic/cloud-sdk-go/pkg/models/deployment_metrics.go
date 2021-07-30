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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentMetrics Metrics information for a deployment
//
// swagger:model DeploymentMetrics
type DeploymentMetrics struct {

	// Whether the deployment metrics are healthy or not
	// Required: true
	Healthy *bool `json:"healthy"`

	// Metrics health issues for the deployment
	Issues []*ObservabilityIssue `json:"issues"`

	// The URLs to view this deployment's metrics in Kibana
	Urls map[string]string `json:"urls,omitempty"`
}

// Validate validates this deployment metrics
func (m *DeploymentMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentMetrics) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentMetrics) validateIssues(formats strfmt.Registry) error {
	if swag.IsZero(m.Issues) { // not required
		return nil
	}

	for i := 0; i < len(m.Issues); i++ {
		if swag.IsZero(m.Issues[i]) { // not required
			continue
		}

		if m.Issues[i] != nil {
			if err := m.Issues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("issues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this deployment metrics based on the context it is used
func (m *DeploymentMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIssues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentMetrics) contextValidateIssues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Issues); i++ {

		if m.Issues[i] != nil {
			if err := m.Issues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("issues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentMetrics) UnmarshalBinary(b []byte) error {
	var res DeploymentMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}