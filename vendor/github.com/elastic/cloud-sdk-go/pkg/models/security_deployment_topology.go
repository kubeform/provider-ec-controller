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

// SecurityDeploymentTopology The cluster topology for the Elasticsearch security deployment.
//
// swagger:model SecurityDeploymentTopology
type SecurityDeploymentTopology struct {

	// The security deployment topology size
	Size *TopologySize `json:"size,omitempty"`

	// The number of zones in which data nodes will be placed
	ZoneCount int32 `json:"zone_count,omitempty"`
}

// Validate validates this security deployment topology
func (m *SecurityDeploymentTopology) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityDeploymentTopology) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security deployment topology based on the context it is used
func (m *SecurityDeploymentTopology) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityDeploymentTopology) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if m.Size != nil {
		if err := m.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityDeploymentTopology) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityDeploymentTopology) UnmarshalBinary(b []byte) error {
	var res SecurityDeploymentTopology
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
