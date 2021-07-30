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

package deployment_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetDeploymentTemplatesV2Reader is a Reader for the GetDeploymentTemplatesV2 structure.
type GetDeploymentTemplatesV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentTemplatesV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentTemplatesV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeploymentTemplatesV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDeploymentTemplatesV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentTemplatesV2OK creates a GetDeploymentTemplatesV2OK with default headers values
func NewGetDeploymentTemplatesV2OK() *GetDeploymentTemplatesV2OK {
	return &GetDeploymentTemplatesV2OK{}
}

/* GetDeploymentTemplatesV2OK describes a response with status code 200, with default header values.

The deployment templates were returned successfully.
*/
type GetDeploymentTemplatesV2OK struct {
	Payload []*models.DeploymentTemplateInfoV2
}

func (o *GetDeploymentTemplatesV2OK) Error() string {
	return fmt.Sprintf("[GET /deployments/templates][%d] getDeploymentTemplatesV2OK  %+v", 200, o.Payload)
}
func (o *GetDeploymentTemplatesV2OK) GetPayload() []*models.DeploymentTemplateInfoV2 {
	return o.Payload
}

func (o *GetDeploymentTemplatesV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentTemplatesV2BadRequest creates a GetDeploymentTemplatesV2BadRequest with default headers values
func NewGetDeploymentTemplatesV2BadRequest() *GetDeploymentTemplatesV2BadRequest {
	return &GetDeploymentTemplatesV2BadRequest{}
}

/* GetDeploymentTemplatesV2BadRequest describes a response with status code 400, with default header values.

The requested region was not found. (code: `templates.region_not_found`)
*/
type GetDeploymentTemplatesV2BadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentTemplatesV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /deployments/templates][%d] getDeploymentTemplatesV2BadRequest  %+v", 400, o.Payload)
}
func (o *GetDeploymentTemplatesV2BadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentTemplatesV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeploymentTemplatesV2Unauthorized creates a GetDeploymentTemplatesV2Unauthorized with default headers values
func NewGetDeploymentTemplatesV2Unauthorized() *GetDeploymentTemplatesV2Unauthorized {
	return &GetDeploymentTemplatesV2Unauthorized{}
}

/* GetDeploymentTemplatesV2Unauthorized describes a response with status code 401, with default header values.

The user is not authorized to access requested region. (code: `templates.region_not_allowed`)
*/
type GetDeploymentTemplatesV2Unauthorized struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentTemplatesV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployments/templates][%d] getDeploymentTemplatesV2Unauthorized  %+v", 401, o.Payload)
}
func (o *GetDeploymentTemplatesV2Unauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentTemplatesV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
