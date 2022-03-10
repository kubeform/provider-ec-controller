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

// CreateDeploymentTemplateV2Reader is a Reader for the CreateDeploymentTemplateV2 structure.
type CreateDeploymentTemplateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeploymentTemplateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDeploymentTemplateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateDeploymentTemplateV2Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDeploymentTemplateV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDeploymentTemplateV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewCreateDeploymentTemplateV2RetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDeploymentTemplateV2OK creates a CreateDeploymentTemplateV2OK with default headers values
func NewCreateDeploymentTemplateV2OK() *CreateDeploymentTemplateV2OK {
	return &CreateDeploymentTemplateV2OK{}
}

/* CreateDeploymentTemplateV2OK describes a response with status code 200, with default header values.

The request was valid (used when validate_only is true).
*/
type CreateDeploymentTemplateV2OK struct {
	Payload *models.IDResponse
}

func (o *CreateDeploymentTemplateV2OK) Error() string {
	return fmt.Sprintf("[POST /deployments/templates][%d] createDeploymentTemplateV2OK  %+v", 200, o.Payload)
}
func (o *CreateDeploymentTemplateV2OK) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *CreateDeploymentTemplateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentTemplateV2Created creates a CreateDeploymentTemplateV2Created with default headers values
func NewCreateDeploymentTemplateV2Created() *CreateDeploymentTemplateV2Created {
	return &CreateDeploymentTemplateV2Created{}
}

/* CreateDeploymentTemplateV2Created describes a response with status code 201, with default header values.

The deployment definition was valid and the template has been created.
*/
type CreateDeploymentTemplateV2Created struct {
	Payload *models.IDResponse
}

func (o *CreateDeploymentTemplateV2Created) Error() string {
	return fmt.Sprintf("[POST /deployments/templates][%d] createDeploymentTemplateV2Created  %+v", 201, o.Payload)
}
func (o *CreateDeploymentTemplateV2Created) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *CreateDeploymentTemplateV2Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentTemplateV2BadRequest creates a CreateDeploymentTemplateV2BadRequest with default headers values
func NewCreateDeploymentTemplateV2BadRequest() *CreateDeploymentTemplateV2BadRequest {
	return &CreateDeploymentTemplateV2BadRequest{}
}

/* CreateDeploymentTemplateV2BadRequest describes a response with status code 400, with default header values.

The requested region is not supported. (code: `templates.region_not_found`)
*/
type CreateDeploymentTemplateV2BadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateDeploymentTemplateV2BadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments/templates][%d] createDeploymentTemplateV2BadRequest  %+v", 400, o.Payload)
}
func (o *CreateDeploymentTemplateV2BadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentTemplateV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDeploymentTemplateV2Unauthorized creates a CreateDeploymentTemplateV2Unauthorized with default headers values
func NewCreateDeploymentTemplateV2Unauthorized() *CreateDeploymentTemplateV2Unauthorized {
	return &CreateDeploymentTemplateV2Unauthorized{}
}

/* CreateDeploymentTemplateV2Unauthorized describes a response with status code 401, with default header values.

The user is not authorized to access requested region. (code: `templates.region_not_allowed`)
*/
type CreateDeploymentTemplateV2Unauthorized struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateDeploymentTemplateV2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /deployments/templates][%d] createDeploymentTemplateV2Unauthorized  %+v", 401, o.Payload)
}
func (o *CreateDeploymentTemplateV2Unauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentTemplateV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDeploymentTemplateV2RetryWith creates a CreateDeploymentTemplateV2RetryWith with default headers values
func NewCreateDeploymentTemplateV2RetryWith() *CreateDeploymentTemplateV2RetryWith {
	return &CreateDeploymentTemplateV2RetryWith{}
}

/* CreateDeploymentTemplateV2RetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type CreateDeploymentTemplateV2RetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateDeploymentTemplateV2RetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/templates][%d] createDeploymentTemplateV2RetryWith  %+v", 449, o.Payload)
}
func (o *CreateDeploymentTemplateV2RetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentTemplateV2RetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
