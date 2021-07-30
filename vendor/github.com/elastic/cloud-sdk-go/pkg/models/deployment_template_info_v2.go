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

// DeploymentTemplateInfoV2 Deployment template detailed information
//
// swagger:model DeploymentTemplateInfoV2
type DeploymentTemplateInfoV2 struct {

	// The body of the deployment template to use for creating a deployment.
	// Required: true
	DeploymentTemplate *DeploymentCreateRequest `json:"deployment_template"`

	// An optional description for the template.
	Description string `json:"description,omitempty"`

	// Whether or not the deployment template is hidden by default.
	Hidden *bool `json:"hidden,omitempty"`

	// The unique identifier for the template.
	// Required: true
	ID *string `json:"id"`

	// List of instance configurations used in the cluster template.
	// Required: true
	InstanceConfigurations []*InstanceConfigurationInfo `json:"instance_configurations"`

	// The Kibana Deeplink for this type of deployment.
	KibanaDeeplink []*KibanaDeeplink `json:"kibana_deeplink"`

	// Optional arbitrary metadata to associate with this template.
	Metadata []*MetadataItem `json:"metadata"`

	// Minimum stack version required by this template, if any.
	MinVersion string `json:"min_version,omitempty"`

	// A human readable name for the template.
	// Required: true
	Name *string `json:"name"`

	// Determines the order in which this template should be returned when listed. Templates are returned in ascending order. If not specified, then the template willbe appended to the end of the list.
	Order *int32 `json:"order,omitempty"`

	// Information describing the source that created or modified the template.
	Source *ChangeSourceInfo `json:"source,omitempty"`

	// Whether or not if this is system owned template.
	SystemOwned *bool `json:"system_owned,omitempty"`

	// Provider and version agnostic template identifier used for grouping related template types.
	TemplateCategoryID string `json:"template_category_id,omitempty"`
}

// Validate validates this deployment template info v2
func (m *DeploymentTemplateInfoV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceConfigurations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKibanaDeeplink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentTemplateInfoV2) validateDeploymentTemplate(formats strfmt.Registry) error {

	if err := validate.Required("deployment_template", "body", m.DeploymentTemplate); err != nil {
		return err
	}

	if m.DeploymentTemplate != nil {
		if err := m.DeploymentTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment_template")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentTemplateInfoV2) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentTemplateInfoV2) validateInstanceConfigurations(formats strfmt.Registry) error {

	if err := validate.Required("instance_configurations", "body", m.InstanceConfigurations); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceConfigurations); i++ {
		if swag.IsZero(m.InstanceConfigurations[i]) { // not required
			continue
		}

		if m.InstanceConfigurations[i] != nil {
			if err := m.InstanceConfigurations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instance_configurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentTemplateInfoV2) validateKibanaDeeplink(formats strfmt.Registry) error {
	if swag.IsZero(m.KibanaDeeplink) { // not required
		return nil
	}

	for i := 0; i < len(m.KibanaDeeplink); i++ {
		if swag.IsZero(m.KibanaDeeplink[i]) { // not required
			continue
		}

		if m.KibanaDeeplink[i] != nil {
			if err := m.KibanaDeeplink[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kibana_deeplink" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentTemplateInfoV2) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	for i := 0; i < len(m.Metadata); i++ {
		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentTemplateInfoV2) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentTemplateInfoV2) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this deployment template info v2 based on the context it is used
func (m *DeploymentTemplateInfoV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeploymentTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstanceConfigurations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKibanaDeeplink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentTemplateInfoV2) contextValidateDeploymentTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.DeploymentTemplate != nil {
		if err := m.DeploymentTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment_template")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentTemplateInfoV2) contextValidateInstanceConfigurations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstanceConfigurations); i++ {

		if m.InstanceConfigurations[i] != nil {
			if err := m.InstanceConfigurations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instance_configurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentTemplateInfoV2) contextValidateKibanaDeeplink(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.KibanaDeeplink); i++ {

		if m.KibanaDeeplink[i] != nil {
			if err := m.KibanaDeeplink[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kibana_deeplink" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentTemplateInfoV2) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metadata); i++ {

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentTemplateInfoV2) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentTemplateInfoV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentTemplateInfoV2) UnmarshalBinary(b []byte) error {
	var res DeploymentTemplateInfoV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}