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

// ClusterCrudResponse The response to an Elasticsearch cluster or Kibana instance CRUD (create/update-plan) request.
//
// swagger:model ClusterCrudResponse
type ClusterCrudResponse struct {

	// If the endpoint is called with URL param 'validate_only=true', then this contains advanced debug info (the internal plan representation) for the APM that was created along with the submitted ES plan.
	Apm *ApmCrudResponse `json:"apm,omitempty"`

	// For an operation creating or updating an APM, the Id of that cluster
	ApmID string `json:"apm_id,omitempty"`

	// An encoded string that provides other Elastic services with the necessary information to connect to this Elasticsearch and Kibana
	CloudID string `json:"cloud_id,omitempty"`

	// credentials
	Credentials *ClusterCredentials `json:"credentials,omitempty"`

	// If the endpoint is called with URL param 'validate_only=true', then this contains advanced debug info (the internal plan representation)
	Diagnostics interface{} `json:"diagnostics,omitempty"`

	// For an operation creating or updating an Elasticsearch cluster, the Id of that cluster
	ElasticsearchClusterID string `json:"elasticsearch_cluster_id,omitempty"`

	// If the endpoint is called with URL param 'validate_only=true', then this contains advanced debug info (the internal plan representation) for the Kibana that was created along with the submitted ES plan.
	Kibana *ClusterCrudResponse `json:"kibana,omitempty"`

	// For an operation creating or updating a Kibana cluster, the Id of that cluster
	KibanaClusterID string `json:"kibana_cluster_id,omitempty"`
}

// Validate validates this cluster crud response
func (m *ClusterCrudResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKibana(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCrudResponse) validateApm(formats strfmt.Registry) error {
	if swag.IsZero(m.Apm) { // not required
		return nil
	}

	if m.Apm != nil {
		if err := m.Apm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apm")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCrudResponse) validateCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCrudResponse) validateKibana(formats strfmt.Registry) error {
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

// ContextValidate validate this cluster crud response based on the context it is used
func (m *ClusterCrudResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKibana(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCrudResponse) contextValidateApm(ctx context.Context, formats strfmt.Registry) error {

	if m.Apm != nil {
		if err := m.Apm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apm")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCrudResponse) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {
		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCrudResponse) contextValidateKibana(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ClusterCrudResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCrudResponse) UnmarshalBinary(b []byte) error {
	var res ClusterCrudResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
