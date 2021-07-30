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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApmInfo The overview information for the APM Server.
//
// swagger:model ApmInfo
type ApmInfo struct {

	// The id of the deployment that this APM Server belongs to.
	DeploymentID string `json:"deployment_id,omitempty"`

	// elasticsearch cluster
	// Required: true
	ElasticsearchCluster *TargetElasticsearchCluster `json:"elasticsearch_cluster"`

	// External resources related to the APM
	// Required: true
	// Unique: true
	ExternalLinks []*ExternalHyperlink `json:"external_links"`

	// Whether the APM is healthy or not (one or more of the info subsections will have healthy: false)
	// Required: true
	Healthy *bool `json:"healthy"`

	// The id of the APM
	// Required: true
	ID *string `json:"id"`

	// A map of application-specific operations (which map to 'operationId's in the Swagger API) to metadata about that operation
	Links map[string]Hyperlink `json:"links,omitempty"`

	// metadata
	Metadata *ClusterMetadataInfo `json:"metadata,omitempty"`

	// The name of the APM
	// Required: true
	Name *string `json:"name"`

	// plan info
	// Required: true
	PlanInfo *ApmPlansInfo `json:"plan_info"`

	// The region that this APM belongs to. Only populated in SaaS or federated ECE.
	Region string `json:"region,omitempty"`

	// The cluster metadata settings for the APM
	Settings *ApmSettings `json:"settings,omitempty"`

	// APM status
	// Required: true
	// Enum: [initializing stopping stopped rebooting restarting reconfiguring started]
	Status *string `json:"status"`

	// topology
	// Required: true
	Topology *ClusterTopologyInfo `json:"topology"`
}

// Validate validates this apm info
func (m *ApmInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElasticsearchCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopology(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApmInfo) validateElasticsearchCluster(formats strfmt.Registry) error {

	if err := validate.Required("elasticsearch_cluster", "body", m.ElasticsearchCluster); err != nil {
		return err
	}

	if m.ElasticsearchCluster != nil {
		if err := m.ElasticsearchCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elasticsearch_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ApmInfo) validateExternalLinks(formats strfmt.Registry) error {

	if err := validate.Required("external_links", "body", m.ExternalLinks); err != nil {
		return err
	}

	if err := validate.UniqueItems("external_links", "body", m.ExternalLinks); err != nil {
		return err
	}

	for i := 0; i < len(m.ExternalLinks); i++ {
		if swag.IsZero(m.ExternalLinks[i]) { // not required
			continue
		}

		if m.ExternalLinks[i] != nil {
			if err := m.ExternalLinks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApmInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *ApmInfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ApmInfo) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for k := range m.Links {

		if err := validate.Required("links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ApmInfo) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ApmInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ApmInfo) validatePlanInfo(formats strfmt.Registry) error {

	if err := validate.Required("plan_info", "body", m.PlanInfo); err != nil {
		return err
	}

	if m.PlanInfo != nil {
		if err := m.PlanInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan_info")
			}
			return err
		}
	}

	return nil
}

func (m *ApmInfo) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

var apmInfoTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["initializing","stopping","stopped","rebooting","restarting","reconfiguring","started"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apmInfoTypeStatusPropEnum = append(apmInfoTypeStatusPropEnum, v)
	}
}

const (

	// ApmInfoStatusInitializing captures enum value "initializing"
	ApmInfoStatusInitializing string = "initializing"

	// ApmInfoStatusStopping captures enum value "stopping"
	ApmInfoStatusStopping string = "stopping"

	// ApmInfoStatusStopped captures enum value "stopped"
	ApmInfoStatusStopped string = "stopped"

	// ApmInfoStatusRebooting captures enum value "rebooting"
	ApmInfoStatusRebooting string = "rebooting"

	// ApmInfoStatusRestarting captures enum value "restarting"
	ApmInfoStatusRestarting string = "restarting"

	// ApmInfoStatusReconfiguring captures enum value "reconfiguring"
	ApmInfoStatusReconfiguring string = "reconfiguring"

	// ApmInfoStatusStarted captures enum value "started"
	ApmInfoStatusStarted string = "started"
)

// prop value enum
func (m *ApmInfo) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, apmInfoTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApmInfo) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ApmInfo) validateTopology(formats strfmt.Registry) error {

	if err := validate.Required("topology", "body", m.Topology); err != nil {
		return err
	}

	if m.Topology != nil {
		if err := m.Topology.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topology")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this apm info based on the context it is used
func (m *ApmInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElasticsearchCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlanInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopology(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApmInfo) contextValidateElasticsearchCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.ElasticsearchCluster != nil {
		if err := m.ElasticsearchCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elasticsearch_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ApmInfo) contextValidateExternalLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalLinks); i++ {

		if m.ExternalLinks[i] != nil {
			if err := m.ExternalLinks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApmInfo) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Links {

		if val, ok := m.Links[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ApmInfo) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ApmInfo) contextValidatePlanInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.PlanInfo != nil {
		if err := m.PlanInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan_info")
			}
			return err
		}
	}

	return nil
}

func (m *ApmInfo) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.Settings != nil {
		if err := m.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *ApmInfo) contextValidateTopology(ctx context.Context, formats strfmt.Registry) error {

	if m.Topology != nil {
		if err := m.Topology.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topology")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApmInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApmInfo) UnmarshalBinary(b []byte) error {
	var res ApmInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}