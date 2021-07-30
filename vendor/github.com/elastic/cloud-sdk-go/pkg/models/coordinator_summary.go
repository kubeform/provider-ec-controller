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

// CoordinatorSummary Summarized information about a coordinator.
//
// swagger:model CoordinatorSummary
type CoordinatorSummary struct {

	// Attributes of this coordinator
	// Required: true
	Attributes map[string]string `json:"attributes"`

	// Client port of this coordinator
	// Required: true
	ClientPort *int32 `json:"client_port"`

	// Election port of this coordinator
	// Required: true
	ElectionPort *int32 `json:"election_port"`

	// Leader port of this coordinator
	// Required: true
	LeaderPort *int32 `json:"leader_port"`

	// Name of this coordinator
	// Required: true
	Name *string `json:"name"`

	// Public hostname of this coordinator
	// Required: true
	PublicHostname *string `json:"public_hostname"`
}

// Validate validates this coordinator summary
func (m *CoordinatorSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElectionPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeaderPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *CoordinatorSummary) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	return nil
}

func (m *CoordinatorSummary) validateClientPort(formats strfmt.Registry) error {

	if err := validate.Required("client_port", "body", m.ClientPort); err != nil {
		return err
	}

	return nil
}

func (m *CoordinatorSummary) validateElectionPort(formats strfmt.Registry) error {

	if err := validate.Required("election_port", "body", m.ElectionPort); err != nil {
		return err
	}

	return nil
}

func (m *CoordinatorSummary) validateLeaderPort(formats strfmt.Registry) error {

	if err := validate.Required("leader_port", "body", m.LeaderPort); err != nil {
		return err
	}

	return nil
}

func (m *CoordinatorSummary) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CoordinatorSummary) validatePublicHostname(formats strfmt.Registry) error {

	if err := validate.Required("public_hostname", "body", m.PublicHostname); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this coordinator summary based on context it is used
func (m *CoordinatorSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoordinatorSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoordinatorSummary) UnmarshalBinary(b []byte) error {
	var res CoordinatorSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}