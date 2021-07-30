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

// ExternalTrustRelationship The trust relationship with external entities (remote environments, remote accounts...).
//
// swagger:model ExternalTrustRelationship
type ExternalTrustRelationship struct {

	// If true, all clusters in this external entity will be trusted and the `trust_allowlist` is ignored.
	// Required: true
	TrustAll *bool `json:"trust_all"`

	// The list of clusters to trust. Only used when `trust_all` is false.
	TrustAllowlist []string `json:"trust_allowlist"`

	// the ID of the external trust relationship
	// Required: true
	TrustRelationshipID *string `json:"trust_relationship_id"`
}

// Validate validates this external trust relationship
func (m *ExternalTrustRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrustAll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrustRelationshipID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalTrustRelationship) validateTrustAll(formats strfmt.Registry) error {

	if err := validate.Required("trust_all", "body", m.TrustAll); err != nil {
		return err
	}

	return nil
}

func (m *ExternalTrustRelationship) validateTrustRelationshipID(formats strfmt.Registry) error {

	if err := validate.Required("trust_relationship_id", "body", m.TrustRelationshipID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this external trust relationship based on context it is used
func (m *ExternalTrustRelationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalTrustRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalTrustRelationship) UnmarshalBinary(b []byte) error {
	var res ExternalTrustRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
