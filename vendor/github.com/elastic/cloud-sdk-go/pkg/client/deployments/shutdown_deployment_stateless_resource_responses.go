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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// ShutdownDeploymentStatelessResourceReader is a Reader for the ShutdownDeploymentStatelessResource structure.
type ShutdownDeploymentStatelessResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShutdownDeploymentStatelessResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShutdownDeploymentStatelessResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShutdownDeploymentStatelessResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShutdownDeploymentStatelessResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewShutdownDeploymentStatelessResourceRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShutdownDeploymentStatelessResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShutdownDeploymentStatelessResourceOK creates a ShutdownDeploymentStatelessResourceOK with default headers values
func NewShutdownDeploymentStatelessResourceOK() *ShutdownDeploymentStatelessResourceOK {
	return &ShutdownDeploymentStatelessResourceOK{}
}

/* ShutdownDeploymentStatelessResourceOK describes a response with status code 200, with default header values.

Standard response
*/
type ShutdownDeploymentStatelessResourceOK struct {
	Payload *models.DeploymentResourceCommandResponse
}

func (o *ShutdownDeploymentStatelessResourceOK) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_shutdown][%d] shutdownDeploymentStatelessResourceOK  %+v", 200, o.Payload)
}
func (o *ShutdownDeploymentStatelessResourceOK) GetPayload() *models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *ShutdownDeploymentStatelessResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShutdownDeploymentStatelessResourceBadRequest creates a ShutdownDeploymentStatelessResourceBadRequest with default headers values
func NewShutdownDeploymentStatelessResourceBadRequest() *ShutdownDeploymentStatelessResourceBadRequest {
	return &ShutdownDeploymentStatelessResourceBadRequest{}
}

/* ShutdownDeploymentStatelessResourceBadRequest describes a response with status code 400, with default header values.

Parameter is restricted and can only be set by a Platform administrator. (code: `deployments.restricted_parameter`)
*/
type ShutdownDeploymentStatelessResourceBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ShutdownDeploymentStatelessResourceBadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_shutdown][%d] shutdownDeploymentStatelessResourceBadRequest  %+v", 400, o.Payload)
}
func (o *ShutdownDeploymentStatelessResourceBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ShutdownDeploymentStatelessResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewShutdownDeploymentStatelessResourceNotFound creates a ShutdownDeploymentStatelessResourceNotFound with default headers values
func NewShutdownDeploymentStatelessResourceNotFound() *ShutdownDeploymentStatelessResourceNotFound {
	return &ShutdownDeploymentStatelessResourceNotFound{}
}

/* ShutdownDeploymentStatelessResourceNotFound describes a response with status code 404, with default header values.

 * The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type ShutdownDeploymentStatelessResourceNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ShutdownDeploymentStatelessResourceNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_shutdown][%d] shutdownDeploymentStatelessResourceNotFound  %+v", 404, o.Payload)
}
func (o *ShutdownDeploymentStatelessResourceNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ShutdownDeploymentStatelessResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewShutdownDeploymentStatelessResourceRetryWith creates a ShutdownDeploymentStatelessResourceRetryWith with default headers values
func NewShutdownDeploymentStatelessResourceRetryWith() *ShutdownDeploymentStatelessResourceRetryWith {
	return &ShutdownDeploymentStatelessResourceRetryWith{}
}

/* ShutdownDeploymentStatelessResourceRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ShutdownDeploymentStatelessResourceRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ShutdownDeploymentStatelessResourceRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_shutdown][%d] shutdownDeploymentStatelessResourceRetryWith  %+v", 449, o.Payload)
}
func (o *ShutdownDeploymentStatelessResourceRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ShutdownDeploymentStatelessResourceRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewShutdownDeploymentStatelessResourceInternalServerError creates a ShutdownDeploymentStatelessResourceInternalServerError with default headers values
func NewShutdownDeploymentStatelessResourceInternalServerError() *ShutdownDeploymentStatelessResourceInternalServerError {
	return &ShutdownDeploymentStatelessResourceInternalServerError{}
}

/* ShutdownDeploymentStatelessResourceInternalServerError describes a response with status code 500, with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type ShutdownDeploymentStatelessResourceInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ShutdownDeploymentStatelessResourceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_shutdown][%d] shutdownDeploymentStatelessResourceInternalServerError  %+v", 500, o.Payload)
}
func (o *ShutdownDeploymentStatelessResourceInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ShutdownDeploymentStatelessResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
