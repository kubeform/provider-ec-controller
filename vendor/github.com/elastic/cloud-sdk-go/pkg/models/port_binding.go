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

// PortBinding Value used to bind to a port on the host.
//
// swagger:model PortBinding
type PortBinding struct {

	// IP to bind to on the host. I.e {@code 0.0.0.0}
	// Example: 0.0.0.0
	HostIP string `json:"host_ip,omitempty"`

	// Port as observed by the host.
	// Required: true
	HostPort *string `json:"host_port"`
}

// Validate validates this port binding
func (m *PortBinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostPort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortBinding) validateHostPort(formats strfmt.Registry) error {

	if err := validate.Required("host_port", "body", m.HostPort); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this port binding based on context it is used
func (m *PortBinding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortBinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortBinding) UnmarshalBinary(b []byte) error {
	var res PortBinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
