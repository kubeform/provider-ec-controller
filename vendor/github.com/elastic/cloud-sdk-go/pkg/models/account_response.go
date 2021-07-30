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

// AccountResponse An account is the entity that owns deployments, and it is accessed by users. Accounts are isolated from each other.
//
// swagger:model AccountResponse
type AccountResponse struct {

	// The account's identifier
	// Required: true
	ID *string `json:"id"`

	// Settings related to the level of trust of the clusters in this account
	Trust *AccountTrustSettings `json:"trust,omitempty"`
}

// Validate validates this account response
func (m *AccountResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrust(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AccountResponse) validateTrust(formats strfmt.Registry) error {
	if swag.IsZero(m.Trust) { // not required
		return nil
	}

	if m.Trust != nil {
		if err := m.Trust.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trust")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this account response based on the context it is used
func (m *AccountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTrust(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountResponse) contextValidateTrust(ctx context.Context, formats strfmt.Registry) error {

	if m.Trust != nil {
		if err := m.Trust.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trust")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountResponse) UnmarshalBinary(b []byte) error {
	var res AccountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
