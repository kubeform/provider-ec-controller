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

// DiscreteSizes Instance sizes that are supported by the Elasticsearch instance, Kibana instance, or APM Server configuration.
//
// swagger:model DiscreteSizes
type DiscreteSizes struct {

	// The default size
	// Required: true
	DefaultSize *int32 `json:"default_size"`

	// The unit that each size represents
	// Required: true
	// Enum: [memory storage]
	Resource *string `json:"resource"`

	// List of supported sizes
	// Required: true
	Sizes []int32 `json:"sizes"`
}

// Validate validates this discrete sizes
func (m *DiscreteSizes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiscreteSizes) validateDefaultSize(formats strfmt.Registry) error {

	if err := validate.Required("default_size", "body", m.DefaultSize); err != nil {
		return err
	}

	return nil
}

var discreteSizesTypeResourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["memory","storage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		discreteSizesTypeResourcePropEnum = append(discreteSizesTypeResourcePropEnum, v)
	}
}

const (

	// DiscreteSizesResourceMemory captures enum value "memory"
	DiscreteSizesResourceMemory string = "memory"

	// DiscreteSizesResourceStorage captures enum value "storage"
	DiscreteSizesResourceStorage string = "storage"
)

// prop value enum
func (m *DiscreteSizes) validateResourceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, discreteSizesTypeResourcePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DiscreteSizes) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	// value enum
	if err := m.validateResourceEnum("resource", "body", *m.Resource); err != nil {
		return err
	}

	return nil
}

func (m *DiscreteSizes) validateSizes(formats strfmt.Registry) error {

	if err := validate.Required("sizes", "body", m.Sizes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this discrete sizes based on context it is used
func (m *DiscreteSizes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DiscreteSizes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiscreteSizes) UnmarshalBinary(b []byte) error {
	var res DiscreteSizes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
