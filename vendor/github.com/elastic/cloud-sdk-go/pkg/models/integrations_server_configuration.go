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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IntegrationsServerConfiguration The configuration options for the Integrations Server.
//
// swagger:model IntegrationsServerConfiguration
type IntegrationsServerConfiguration struct {

	// A docker URI that allows overriding of the default docker image specified for this version
	DockerImage string `json:"docker_image,omitempty"`

	// The mode the Integrations Server is operating in.
	// Enum: [standalone managed]
	Mode string `json:"mode,omitempty"`

	// system settings
	SystemSettings *IntegrationsServerSystemSettings `json:"system_settings,omitempty"`

	// An arbitrary JSON object allowing (non-admin) cluster owners to set their parameters (only one of this and 'user_settings_yaml' is allowed), provided the parameters are on the allowlist and not on the denylist. (This field together with 'user_settings_override*' and 'system_settings' defines the total set of Integrations Server settings)
	UserSettingsJSON interface{} `json:"user_settings_json,omitempty"`

	// An arbitrary JSON object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_yaml' is allowed), ie in addition to the documented 'system_settings'. (This field together with 'system_settings' and 'user_settings*' defines the total set of Integrations Server settings)
	UserSettingsOverrideJSON interface{} `json:"user_settings_override_json,omitempty"`

	// An arbitrary YAML object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_json' is allowed), ie in addition to the documented 'system_settings'. (This field together with 'system_settings' and 'user_settings*' defines the total set of Integrations Server settings)
	UserSettingsOverrideYaml string `json:"user_settings_override_yaml,omitempty"`

	// An arbitrary YAML object allowing (non-admin) cluster owners to set their parameters (only one of this and 'user_settings_json' is allowed), provided the parameters are on the allowlist and not on the denylist. (These field together with 'user_settings_override*' and 'system_settings' defines the total set of Integrations Server settings)
	UserSettingsYaml string `json:"user_settings_yaml,omitempty"`

	// The version of the Integrations Server cluster (must be one of the ECE supported versions, and won't work unless it matches the Integrations Server version. Leave blank to auto-detect version.)
	Version string `json:"version,omitempty"`
}

// Validate validates this integrations server configuration
func (m *IntegrationsServerConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var integrationsServerConfigurationTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["standalone","managed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		integrationsServerConfigurationTypeModePropEnum = append(integrationsServerConfigurationTypeModePropEnum, v)
	}
}

const (

	// IntegrationsServerConfigurationModeStandalone captures enum value "standalone"
	IntegrationsServerConfigurationModeStandalone string = "standalone"

	// IntegrationsServerConfigurationModeManaged captures enum value "managed"
	IntegrationsServerConfigurationModeManaged string = "managed"
)

// prop value enum
func (m *IntegrationsServerConfiguration) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, integrationsServerConfigurationTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IntegrationsServerConfiguration) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationsServerConfiguration) validateSystemSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemSettings) { // not required
		return nil
	}

	if m.SystemSettings != nil {
		if err := m.SystemSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system_settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this integrations server configuration based on the context it is used
func (m *IntegrationsServerConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSystemSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsServerConfiguration) contextValidateSystemSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemSettings != nil {
		if err := m.SystemSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system_settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsServerConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsServerConfiguration) UnmarshalBinary(b []byte) error {
	var res IntegrationsServerConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
