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

package platform_configuration_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetDeploymentTemplatesReader is a Reader for the GetDeploymentTemplates structure.
type GetDeploymentTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeploymentTemplatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentTemplatesOK creates a GetDeploymentTemplatesOK with default headers values
func NewGetDeploymentTemplatesOK() *GetDeploymentTemplatesOK {
	return &GetDeploymentTemplatesOK{}
}

/* GetDeploymentTemplatesOK describes a response with status code 200, with default header values.

The deployment templates were returned successfully.
*/
type GetDeploymentTemplatesOK struct {
	Payload []*models.DeploymentTemplateInfo
}

func (o *GetDeploymentTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /platform/configuration/templates/deployments][%d] getDeploymentTemplatesOK  %+v", 200, o.Payload)
}
func (o *GetDeploymentTemplatesOK) GetPayload() []*models.DeploymentTemplateInfo {
	return o.Payload
}

func (o *GetDeploymentTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentTemplatesBadRequest creates a GetDeploymentTemplatesBadRequest with default headers values
func NewGetDeploymentTemplatesBadRequest() *GetDeploymentTemplatesBadRequest {
	return &GetDeploymentTemplatesBadRequest{}
}

/* GetDeploymentTemplatesBadRequest describes a response with status code 400, with default header values.

The template is not compatible with the [cluster] format. (code: `deployment.migration_invalid`)
*/
type GetDeploymentTemplatesBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentTemplatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /platform/configuration/templates/deployments][%d] getDeploymentTemplatesBadRequest  %+v", 400, o.Payload)
}
func (o *GetDeploymentTemplatesBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentTemplatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}