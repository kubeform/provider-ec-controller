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
)

// KibanaClusterTopologyElement The topology of the Kibana nodes, including the number, capacity, and type of nodes, and where they can be allocated.
//
// swagger:model KibanaClusterTopologyElement
type KibanaClusterTopologyElement struct {

	// Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the id of an existing instance configuration.
	InstanceConfigurationID string `json:"instance_configuration_id,omitempty"`

	// kibana
	Kibana *KibanaConfiguration `json:"kibana,omitempty"`

	// The memory capacity in MB for each node of this type built in each zone.
	MemoryPerNode int32 `json:"memory_per_node,omitempty"`

	// The number of nodes of this type that are allocated within each zone (i.e. total capacity per zone = `node_count_per_zone` * `memory_per_node` in MB).
	NodeCountPerZone int32 `json:"node_count_per_zone,omitempty"`

	// size
	Size *TopologySize `json:"size,omitempty"`

	// number of zones in which nodes will be placed
	ZoneCount int32 `json:"zone_count,omitempty"`
}

// Validate validates this kibana cluster topology element
func (m *KibanaClusterTopologyElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKibana(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KibanaClusterTopologyElement) validateKibana(formats strfmt.Registry) error {
	if swag.IsZero(m.Kibana) { // not required
		return nil
	}

	if m.Kibana != nil {
		if err := m.Kibana.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kibana")
			}
			return err
		}
	}

	return nil
}

func (m *KibanaClusterTopologyElement) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kibana cluster topology element based on the context it is used
func (m *KibanaClusterTopologyElement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKibana(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KibanaClusterTopologyElement) contextValidateKibana(ctx context.Context, formats strfmt.Registry) error {

	if m.Kibana != nil {
		if err := m.Kibana.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kibana")
			}
			return err
		}
	}

	return nil
}

func (m *KibanaClusterTopologyElement) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if m.Size != nil {
		if err := m.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KibanaClusterTopologyElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KibanaClusterTopologyElement) UnmarshalBinary(b []byte) error {
	var res KibanaClusterTopologyElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
