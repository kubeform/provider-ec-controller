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

// RunnersSummary Summarized information about the runners.
//
// swagger:model RunnersSummary
type RunnersSummary struct {

	// Total capacity of connected runners in megabytes
	// Required: true
	ConnectedCapacity *int32 `json:"connected_capacity"`

	// Number of connected runners
	// Required: true
	ConnectedRunners *int32 `json:"connected_runners"`

	// Total number of containers
	// Required: true
	ContainersCount *int32 `json:"containers_count"`

	// Whether all runners are healthy
	// Required: true
	Healthy *bool `json:"healthy"`

	// Number of healthy runners
	// Required: true
	HealthyRunners *int32 `json:"healthy_runners"`

	// Maximum capacity available in one runner in megabytes
	// Required: true
	MaxAvailableCapacity *int32 `json:"max_available_capacity"`

	// Total number of runners
	// Required: true
	TotalRunners *int32 `json:"total_runners"`
}

// Validate validates this runners summary
func (m *RunnersSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectedCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectedRunners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainersCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthyRunners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxAvailableCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalRunners(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunnersSummary) validateConnectedCapacity(formats strfmt.Registry) error {

	if err := validate.Required("connected_capacity", "body", m.ConnectedCapacity); err != nil {
		return err
	}

	return nil
}

func (m *RunnersSummary) validateConnectedRunners(formats strfmt.Registry) error {

	if err := validate.Required("connected_runners", "body", m.ConnectedRunners); err != nil {
		return err
	}

	return nil
}

func (m *RunnersSummary) validateContainersCount(formats strfmt.Registry) error {

	if err := validate.Required("containers_count", "body", m.ContainersCount); err != nil {
		return err
	}

	return nil
}

func (m *RunnersSummary) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *RunnersSummary) validateHealthyRunners(formats strfmt.Registry) error {

	if err := validate.Required("healthy_runners", "body", m.HealthyRunners); err != nil {
		return err
	}

	return nil
}

func (m *RunnersSummary) validateMaxAvailableCapacity(formats strfmt.Registry) error {

	if err := validate.Required("max_available_capacity", "body", m.MaxAvailableCapacity); err != nil {
		return err
	}

	return nil
}

func (m *RunnersSummary) validateTotalRunners(formats strfmt.Registry) error {

	if err := validate.Required("total_runners", "body", m.TotalRunners); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this runners summary based on context it is used
func (m *RunnersSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RunnersSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunnersSummary) UnmarshalBinary(b []byte) error {
	var res RunnersSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
