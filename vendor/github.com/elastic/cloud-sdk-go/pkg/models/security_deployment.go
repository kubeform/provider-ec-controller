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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SecurityDeployment The Elasticsearch security deployment.
//
// swagger:model SecurityDeployment
type SecurityDeployment struct {

	// The identifier for the security deployment cluster
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// True if a pending plan is in progress
	// Required: true
	HasPendingPlan *bool `json:"has_pending_plan"`

	// True if the security cluster is currently enabled
	// Required: true
	IsEnabled *bool `json:"is_enabled"`

	// True if the cluster is healthy
	// Required: true
	IsHealthy *bool `json:"is_healthy"`

	// The name of the security deployment cluster
	// Required: true
	Name *string `json:"name"`

	// The current status of the cluster
	// Required: true
	// Enum: [initializing stopping stopped rebooting restarting reconfiguring started]
	Status *string `json:"status"`

	// The version of the Elasticsearch cluster
	Version string `json:"version,omitempty"`
}

// Validate validates this security deployment
func (m *SecurityDeployment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasPendingPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityDeployment) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *SecurityDeployment) validateHasPendingPlan(formats strfmt.Registry) error {

	if err := validate.Required("has_pending_plan", "body", m.HasPendingPlan); err != nil {
		return err
	}

	return nil
}

func (m *SecurityDeployment) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("is_enabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

func (m *SecurityDeployment) validateIsHealthy(formats strfmt.Registry) error {

	if err := validate.Required("is_healthy", "body", m.IsHealthy); err != nil {
		return err
	}

	return nil
}

func (m *SecurityDeployment) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var securityDeploymentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["initializing","stopping","stopped","rebooting","restarting","reconfiguring","started"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityDeploymentTypeStatusPropEnum = append(securityDeploymentTypeStatusPropEnum, v)
	}
}

const (

	// SecurityDeploymentStatusInitializing captures enum value "initializing"
	SecurityDeploymentStatusInitializing string = "initializing"

	// SecurityDeploymentStatusStopping captures enum value "stopping"
	SecurityDeploymentStatusStopping string = "stopping"

	// SecurityDeploymentStatusStopped captures enum value "stopped"
	SecurityDeploymentStatusStopped string = "stopped"

	// SecurityDeploymentStatusRebooting captures enum value "rebooting"
	SecurityDeploymentStatusRebooting string = "rebooting"

	// SecurityDeploymentStatusRestarting captures enum value "restarting"
	SecurityDeploymentStatusRestarting string = "restarting"

	// SecurityDeploymentStatusReconfiguring captures enum value "reconfiguring"
	SecurityDeploymentStatusReconfiguring string = "reconfiguring"

	// SecurityDeploymentStatusStarted captures enum value "started"
	SecurityDeploymentStatusStarted string = "started"
)

// prop value enum
func (m *SecurityDeployment) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityDeploymentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecurityDeployment) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this security deployment based on context it is used
func (m *SecurityDeployment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecurityDeployment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityDeployment) UnmarshalBinary(b []byte) error {
	var res SecurityDeployment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
