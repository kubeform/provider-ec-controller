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

// KibanaClustersInfo Information about a set of Kibana instances.
//
// swagger:model KibanaClustersInfo
type KibanaClustersInfo struct {

	// kibana clusters
	// Required: true
	KibanaClusters []*KibanaClusterInfo `json:"kibana_clusters"`

	// If a query is supplied, then the total number of clusters that matched
	MatchCount int32 `json:"match_count,omitempty"`

	// The number of clusters actually returned
	// Required: true
	ReturnCount *int32 `json:"return_count"`
}

// Validate validates this kibana clusters info
func (m *KibanaClustersInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKibanaClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KibanaClustersInfo) validateKibanaClusters(formats strfmt.Registry) error {

	if err := validate.Required("kibana_clusters", "body", m.KibanaClusters); err != nil {
		return err
	}

	for i := 0; i < len(m.KibanaClusters); i++ {
		if swag.IsZero(m.KibanaClusters[i]) { // not required
			continue
		}

		if m.KibanaClusters[i] != nil {
			if err := m.KibanaClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kibana_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KibanaClustersInfo) validateReturnCount(formats strfmt.Registry) error {

	if err := validate.Required("return_count", "body", m.ReturnCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this kibana clusters info based on the context it is used
func (m *KibanaClustersInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKibanaClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KibanaClustersInfo) contextValidateKibanaClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.KibanaClusters); i++ {

		if m.KibanaClusters[i] != nil {
			if err := m.KibanaClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kibana_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KibanaClustersInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KibanaClustersInfo) UnmarshalBinary(b []byte) error {
	var res KibanaClustersInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
