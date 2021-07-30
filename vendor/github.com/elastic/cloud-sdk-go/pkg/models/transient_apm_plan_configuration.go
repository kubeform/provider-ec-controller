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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransientApmPlanConfiguration Defines the configuration parameters that control how the plan is applied. For example, the Elasticsearch cluster topology and APM Server settings.
//
// swagger:model TransientApmPlanConfiguration
type TransientApmPlanConfiguration struct {

	// plan configuration
	PlanConfiguration *ApmPlanControlConfiguration `json:"plan_configuration,omitempty"`

	// strategy
	Strategy *PlanStrategy `json:"strategy,omitempty"`
}

// Validate validates this transient apm plan configuration
func (m *TransientApmPlanConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlanConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransientApmPlanConfiguration) validatePlanConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.PlanConfiguration) { // not required
		return nil
	}

	if m.PlanConfiguration != nil {
		if err := m.PlanConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *TransientApmPlanConfiguration) validateStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.Strategy) { // not required
		return nil
	}

	if m.Strategy != nil {
		if err := m.Strategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("strategy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this transient apm plan configuration based on the context it is used
func (m *TransientApmPlanConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlanConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransientApmPlanConfiguration) contextValidatePlanConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.PlanConfiguration != nil {
		if err := m.PlanConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *TransientApmPlanConfiguration) contextValidateStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.Strategy != nil {
		if err := m.Strategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("strategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransientApmPlanConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransientApmPlanConfiguration) UnmarshalBinary(b []byte) error {
	var res TransientApmPlanConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
