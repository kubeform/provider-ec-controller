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

// MoveClustersCommandResponse Information about the move operations for the Elasticsearch clusters, multiple Kibana instances, and multiple APM Servers, prior to the move.
//
// swagger:model MoveClustersCommandResponse
type MoveClustersCommandResponse struct {

	// Detailed information about the clusters that failed to move.
	// Required: true
	Failures *MoveClustersDetails `json:"failures"`

	// Detailed information about the clusters being moved off the allocator.
	// Required: true
	Moves *MoveClustersDetails `json:"moves"`
}

// Validate validates this move clusters command response
func (m *MoveClustersCommandResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoves(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveClustersCommandResponse) validateFailures(formats strfmt.Registry) error {

	if err := validate.Required("failures", "body", m.Failures); err != nil {
		return err
	}

	if m.Failures != nil {
		if err := m.Failures.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failures")
			}
			return err
		}
	}

	return nil
}

func (m *MoveClustersCommandResponse) validateMoves(formats strfmt.Registry) error {

	if err := validate.Required("moves", "body", m.Moves); err != nil {
		return err
	}

	if m.Moves != nil {
		if err := m.Moves.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moves")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this move clusters command response based on the context it is used
func (m *MoveClustersCommandResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFailures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMoves(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveClustersCommandResponse) contextValidateFailures(ctx context.Context, formats strfmt.Registry) error {

	if m.Failures != nil {
		if err := m.Failures.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failures")
			}
			return err
		}
	}

	return nil
}

func (m *MoveClustersCommandResponse) contextValidateMoves(ctx context.Context, formats strfmt.Registry) error {

	if m.Moves != nil {
		if err := m.Moves.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moves")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveClustersCommandResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveClustersCommandResponse) UnmarshalBinary(b []byte) error {
	var res MoveClustersCommandResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
