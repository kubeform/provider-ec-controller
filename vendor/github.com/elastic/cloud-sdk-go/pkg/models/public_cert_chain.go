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

// PublicCertChain The public portion of the certificate chain that contains the PEM encoded server certificate, intermediate certificates, and the CA certificate. NOTE: The private key, normally included in certificate chains, is omitted.
//
// swagger:model PublicCertChain
type PublicCertChain struct {

	// The list of PEM encoded X509 certificates that make up the certificate chain
	// Required: true
	Chain []string `json:"chain"`
}

// Validate validates this public cert chain
func (m *PublicCertChain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicCertChain) validateChain(formats strfmt.Registry) error {

	if err := validate.Required("chain", "body", m.Chain); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this public cert chain based on context it is used
func (m *PublicCertChain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublicCertChain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicCertChain) UnmarshalBinary(b []byte) error {
	var res PublicCertChain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
