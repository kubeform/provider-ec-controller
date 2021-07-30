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

// ProxyInfo Information about a proxy.
//
// swagger:model ProxyInfo
type ProxyInfo struct {

	// Allocation information by type.
	// Required: true
	Allocations []*ProxyAllocationInfo `json:"allocations"`

	// Specifies if the proxy is healthy
	// Required: true
	Healthy *bool `json:"healthy"`

	// IP of the host the proxy runs on
	// Required: true
	HostIP *string `json:"host_ip"`

	// Arbitrary metadata associated with the proxy
	// Required: true
	Metadata interface{} `json:"metadata"`

	// The proxy identifier
	// Required: true
	ProxyID *string `json:"proxy_id"`

	// Public hostname of the host the proxy runs on
	// Required: true
	PublicHostname *string `json:"public_hostname"`

	// Identifier of the runner for the proxy
	RunnerID string `json:"runner_id,omitempty"`

	// The zone.
	Zone string `json:"zone,omitempty"`
}

// Validate validates this proxy info
func (m *ProxyInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicHostname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxyInfo) validateAllocations(formats strfmt.Registry) error {

	if err := validate.Required("allocations", "body", m.Allocations); err != nil {
		return err
	}

	for i := 0; i < len(m.Allocations); i++ {
		if swag.IsZero(m.Allocations[i]) { // not required
			continue
		}

		if m.Allocations[i] != nil {
			if err := m.Allocations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProxyInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *ProxyInfo) validateHostIP(formats strfmt.Registry) error {

	if err := validate.Required("host_ip", "body", m.HostIP); err != nil {
		return err
	}

	return nil
}

func (m *ProxyInfo) validateMetadata(formats strfmt.Registry) error {

	if m.Metadata == nil {
		return errors.Required("metadata", "body", nil)
	}

	return nil
}

func (m *ProxyInfo) validateProxyID(formats strfmt.Registry) error {

	if err := validate.Required("proxy_id", "body", m.ProxyID); err != nil {
		return err
	}

	return nil
}

func (m *ProxyInfo) validatePublicHostname(formats strfmt.Registry) error {

	if err := validate.Required("public_hostname", "body", m.PublicHostname); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this proxy info based on the context it is used
func (m *ProxyInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxyInfo) contextValidateAllocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Allocations); i++ {

		if m.Allocations[i] != nil {
			if err := m.Allocations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxyInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxyInfo) UnmarshalBinary(b []byte) error {
	var res ProxyInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
