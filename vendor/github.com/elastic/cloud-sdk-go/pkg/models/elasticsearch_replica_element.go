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

// ElasticsearchReplicaElement Information about the unavailable replicas. NOTE: Unlike shards, unavailable replicas indicate a loss of redundancy rather than a loss of availability.
//
// swagger:model ElasticsearchReplicaElement
type ElasticsearchReplicaElement struct {

	// The Elastic Cloud name/id of the instance (container)
	// Required: true
	InstanceName *string `json:"instance_name"`

	// The number of unavailable replicas on this instance
	// Required: true
	ReplicaCount *int32 `json:"replica_count"`
}

// Validate validates this elasticsearch replica element
func (m *ElasticsearchReplicaElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchReplicaElement) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instance_name", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchReplicaElement) validateReplicaCount(formats strfmt.Registry) error {

	if err := validate.Required("replica_count", "body", m.ReplicaCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this elasticsearch replica element based on context it is used
func (m *ElasticsearchReplicaElement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchReplicaElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchReplicaElement) UnmarshalBinary(b []byte) error {
	var res ElasticsearchReplicaElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
