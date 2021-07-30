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

// SamlSpSettings The configuration for the Elasticsearch security SAML Service Provider.
//
// swagger:model SamlSpSettings
type SamlSpSettings struct {

	// The URL of the Assertion Consumer service
	// Required: true
	Acs *string `json:"acs"`

	// The Entity ID to use for this SAML Service Provider. This should be entered as a URI.
	// Required: true
	EntityID *string `json:"entity_id"`

	// The URL of the Single Logout service
	// Required: true
	Logout *string `json:"logout"`
}

// Validate validates this saml sp settings
func (m *SamlSpSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogout(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SamlSpSettings) validateAcs(formats strfmt.Registry) error {

	if err := validate.Required("acs", "body", m.Acs); err != nil {
		return err
	}

	return nil
}

func (m *SamlSpSettings) validateEntityID(formats strfmt.Registry) error {

	if err := validate.Required("entity_id", "body", m.EntityID); err != nil {
		return err
	}

	return nil
}

func (m *SamlSpSettings) validateLogout(formats strfmt.Registry) error {

	if err := validate.Required("logout", "body", m.Logout); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this saml sp settings based on context it is used
func (m *SamlSpSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SamlSpSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SamlSpSettings) UnmarshalBinary(b []byte) error {
	var res SamlSpSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}