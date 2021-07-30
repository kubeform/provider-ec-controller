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

// ElasticsearchMasterInfo Information about the master nodes in the Elasticsearch cluster.
//
// swagger:model ElasticsearchMasterInfo
type ElasticsearchMasterInfo struct {

	// Whether the master situation in the cluster is healthy (ie is the number of masters != 1), or do any instances have no master
	// Required: true
	Healthy *bool `json:"healthy"`

	// A list of any instances with no master
	// Required: true
	InstancesWithNoMaster []string `json:"instances_with_no_master"`

	// masters
	// Required: true
	Masters []*ElasticsearchMasterElement `json:"masters"`
}

// Validate validates this elasticsearch master info
func (m *ElasticsearchMasterInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstancesWithNoMaster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchMasterInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchMasterInfo) validateInstancesWithNoMaster(formats strfmt.Registry) error {

	if err := validate.Required("instances_with_no_master", "body", m.InstancesWithNoMaster); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchMasterInfo) validateMasters(formats strfmt.Registry) error {

	if err := validate.Required("masters", "body", m.Masters); err != nil {
		return err
	}

	for i := 0; i < len(m.Masters); i++ {
		if swag.IsZero(m.Masters[i]) { // not required
			continue
		}

		if m.Masters[i] != nil {
			if err := m.Masters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("masters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this elasticsearch master info based on the context it is used
func (m *ElasticsearchMasterInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMasters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchMasterInfo) contextValidateMasters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Masters); i++ {

		if m.Masters[i] != nil {
			if err := m.Masters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("masters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchMasterInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchMasterInfo) UnmarshalBinary(b []byte) error {
	var res ElasticsearchMasterInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
