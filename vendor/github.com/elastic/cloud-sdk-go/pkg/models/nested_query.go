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

// NestedQuery A query that matches nested objects.
//
// swagger:model NestedQuery
type NestedQuery struct {

	// The path to the nested object.
	// Required: true
	Path *string `json:"path"`

	// The actual query to execute on the nested objects.
	// Required: true
	Query *QueryContainer `json:"query"`

	// Allows to specify how inner children matching affects score of the parent. Refer to the Elasticsearch documentation for details.
	// Enum: [avg sum min max none]
	ScoreMode string `json:"score_mode,omitempty"`
}

// Validate validates this nested query
func (m *NestedQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScoreMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedQuery) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *NestedQuery) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

var nestedQueryTypeScoreModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["avg","sum","min","max","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nestedQueryTypeScoreModePropEnum = append(nestedQueryTypeScoreModePropEnum, v)
	}
}

const (

	// NestedQueryScoreModeAvg captures enum value "avg"
	NestedQueryScoreModeAvg string = "avg"

	// NestedQueryScoreModeSum captures enum value "sum"
	NestedQueryScoreModeSum string = "sum"

	// NestedQueryScoreModeMin captures enum value "min"
	NestedQueryScoreModeMin string = "min"

	// NestedQueryScoreModeMax captures enum value "max"
	NestedQueryScoreModeMax string = "max"

	// NestedQueryScoreModeNone captures enum value "none"
	NestedQueryScoreModeNone string = "none"
)

// prop value enum
func (m *NestedQuery) validateScoreModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nestedQueryTypeScoreModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NestedQuery) validateScoreMode(formats strfmt.Registry) error {
	if swag.IsZero(m.ScoreMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateScoreModeEnum("score_mode", "body", m.ScoreMode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nested query based on the context it is used
func (m *NestedQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedQuery) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	if m.Query != nil {
		if err := m.Query.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedQuery) UnmarshalBinary(b []byte) error {
	var res NestedQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}