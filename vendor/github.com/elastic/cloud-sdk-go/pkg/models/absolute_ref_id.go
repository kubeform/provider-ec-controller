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

// AbsoluteRefID A reference to a specific resource of a deployment
//
// swagger:model AbsoluteRefId
type AbsoluteRefID struct {

	// The deployment id
	// Required: true
	DeploymentID *string `json:"deployment_id"`

	// The reference id of the resource in the given deployment
	// Required: true
	RefID *string `json:"ref_id"`
}

// Validate validates this absolute ref Id
func (m *AbsoluteRefID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AbsoluteRefID) validateDeploymentID(formats strfmt.Registry) error {

	if err := validate.Required("deployment_id", "body", m.DeploymentID); err != nil {
		return err
	}

	return nil
}

func (m *AbsoluteRefID) validateRefID(formats strfmt.Registry) error {

	if err := validate.Required("ref_id", "body", m.RefID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this absolute ref Id based on context it is used
func (m *AbsoluteRefID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AbsoluteRefID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AbsoluteRefID) UnmarshalBinary(b []byte) error {
	var res AbsoluteRefID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
