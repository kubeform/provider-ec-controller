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
	"github.com/go-openapi/validate"
)

// IndexPattern An index pattern described indicating how it has to be migrated to ILM.
//
// swagger:model IndexPattern
type IndexPattern struct {

	// Index pattern to which the ILM policy will be applied.
	// Required: true
	IndexPattern *string `json:"index_pattern"`

	// Defines the Elasticsearch node attributes for the warm element of the topology
	NodeAttributes map[string]string `json:"node_attributes,omitempty"`

	// Name of the policy to create
	// Required: true
	PolicyName *string `json:"policy_name"`
}

// Validate validates this index pattern
func (m *IndexPattern) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndexPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IndexPattern) validateIndexPattern(formats strfmt.Registry) error {

	if err := validate.Required("index_pattern", "body", m.IndexPattern); err != nil {
		return err
	}

	return nil
}

func (m *IndexPattern) validatePolicyName(formats strfmt.Registry) error {

	if err := validate.Required("policy_name", "body", m.PolicyName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this index pattern based on context it is used
func (m *IndexPattern) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IndexPattern) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndexPattern) UnmarshalBinary(b []byte) error {
	var res IndexPattern
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
