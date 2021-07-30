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

// CrossClusterSearchInfo The cross-cluster search settings and status for the Elasticsearch cluster.
//
// swagger:model CrossClusterSearchInfo
type CrossClusterSearchInfo struct {

	// Flag signaling health issues when at least one remote has an incompatible version with this cluster
	// Required: true
	Healthy *bool `json:"healthy"`

	// The list of remote clusters this cluster can access using cross-cluster search
	// Required: true
	RemoteClusters []*RemoteClusterInfo `json:"remote_clusters"`
}

// Validate validates this cross cluster search info
func (m *CrossClusterSearchInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrossClusterSearchInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *CrossClusterSearchInfo) validateRemoteClusters(formats strfmt.Registry) error {

	if err := validate.Required("remote_clusters", "body", m.RemoteClusters); err != nil {
		return err
	}

	for i := 0; i < len(m.RemoteClusters); i++ {
		if swag.IsZero(m.RemoteClusters[i]) { // not required
			continue
		}

		if m.RemoteClusters[i] != nil {
			if err := m.RemoteClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remote_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cross cluster search info based on the context it is used
func (m *CrossClusterSearchInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRemoteClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrossClusterSearchInfo) contextValidateRemoteClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RemoteClusters); i++ {

		if m.RemoteClusters[i] != nil {
			if err := m.RemoteClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remote_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CrossClusterSearchInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrossClusterSearchInfo) UnmarshalBinary(b []byte) error {
	var res CrossClusterSearchInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}