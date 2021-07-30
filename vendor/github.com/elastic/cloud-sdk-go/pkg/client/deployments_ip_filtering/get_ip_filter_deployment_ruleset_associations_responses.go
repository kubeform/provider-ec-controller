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

package deployments_ip_filtering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetIPFilterDeploymentRulesetAssociationsReader is a Reader for the GetIPFilterDeploymentRulesetAssociations structure.
type GetIPFilterDeploymentRulesetAssociationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPFilterDeploymentRulesetAssociationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIPFilterDeploymentRulesetAssociationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetIPFilterDeploymentRulesetAssociationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIPFilterDeploymentRulesetAssociationsOK creates a GetIPFilterDeploymentRulesetAssociationsOK with default headers values
func NewGetIPFilterDeploymentRulesetAssociationsOK() *GetIPFilterDeploymentRulesetAssociationsOK {
	return &GetIPFilterDeploymentRulesetAssociationsOK{}
}

/* GetIPFilterDeploymentRulesetAssociationsOK describes a response with status code 200, with default header values.

Rulesets in the deployment were successfully returned
*/
type GetIPFilterDeploymentRulesetAssociationsOK struct {
	Payload *models.IPFilteringSettings
}

func (o *GetIPFilterDeploymentRulesetAssociationsOK) Error() string {
	return fmt.Sprintf("[GET /deployments/ip-filtering/associations/{association_type}/{associated_entity_id}/rulesets][%d] getIpFilterDeploymentRulesetAssociationsOK  %+v", 200, o.Payload)
}
func (o *GetIPFilterDeploymentRulesetAssociationsOK) GetPayload() *models.IPFilteringSettings {
	return o.Payload
}

func (o *GetIPFilterDeploymentRulesetAssociationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPFilteringSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPFilterDeploymentRulesetAssociationsInternalServerError creates a GetIPFilterDeploymentRulesetAssociationsInternalServerError with default headers values
func NewGetIPFilterDeploymentRulesetAssociationsInternalServerError() *GetIPFilterDeploymentRulesetAssociationsInternalServerError {
	return &GetIPFilterDeploymentRulesetAssociationsInternalServerError{}
}

/* GetIPFilterDeploymentRulesetAssociationsInternalServerError describes a response with status code 500, with default header values.

Request execution failed (code: 'ip_filtering.request_execution_failed')
*/
type GetIPFilterDeploymentRulesetAssociationsInternalServerError struct {
	Payload *models.BasicFailedReply
}

func (o *GetIPFilterDeploymentRulesetAssociationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/ip-filtering/associations/{association_type}/{associated_entity_id}/rulesets][%d] getIpFilterDeploymentRulesetAssociationsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetIPFilterDeploymentRulesetAssociationsInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetIPFilterDeploymentRulesetAssociationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
