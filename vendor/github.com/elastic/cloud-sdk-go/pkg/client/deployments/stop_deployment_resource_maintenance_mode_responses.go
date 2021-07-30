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

// StopDeploymentResourceMaintenanceModeReader is a Reader for the StopDeploymentResourceMaintenanceMode structure.
type StopDeploymentResourceMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopDeploymentResourceMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStopDeploymentResourceMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStopDeploymentResourceMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopDeploymentResourceMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStopDeploymentResourceMaintenanceModeRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopDeploymentResourceMaintenanceModeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopDeploymentResourceMaintenanceModeAccepted creates a StopDeploymentResourceMaintenanceModeAccepted with default headers values
func NewStopDeploymentResourceMaintenanceModeAccepted() *StopDeploymentResourceMaintenanceModeAccepted {
	return &StopDeploymentResourceMaintenanceModeAccepted{}
}

/* StopDeploymentResourceMaintenanceModeAccepted describes a response with status code 202, with default header values.

The stop maintenance mode command was issued successfully.
*/
type StopDeploymentResourceMaintenanceModeAccepted struct {
	Payload *models.DeploymentResourceCommandResponse
}

func (o *StopDeploymentResourceMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/{instance_ids}/maintenance-mode/_stop][%d] stopDeploymentResourceMaintenanceModeAccepted  %+v", 202, o.Payload)
}
func (o *StopDeploymentResourceMaintenanceModeAccepted) GetPayload() *models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *StopDeploymentResourceMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopDeploymentResourceMaintenanceModeForbidden creates a StopDeploymentResourceMaintenanceModeForbidden with default headers values
func NewStopDeploymentResourceMaintenanceModeForbidden() *StopDeploymentResourceMaintenanceModeForbidden {
	return &StopDeploymentResourceMaintenanceModeForbidden{}
}

/* StopDeploymentResourceMaintenanceModeForbidden describes a response with status code 403, with default header values.

The stop maintenance mode command was prohibited for the given Resource. (code: `deployments.instance_update_prohibited_error`)
*/
type StopDeploymentResourceMaintenanceModeForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/{instance_ids}/maintenance-mode/_stop][%d] stopDeploymentResourceMaintenanceModeForbidden  %+v", 403, o.Payload)
}
func (o *StopDeploymentResourceMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopDeploymentResourceMaintenanceModeNotFound creates a StopDeploymentResourceMaintenanceModeNotFound with default headers values
func NewStopDeploymentResourceMaintenanceModeNotFound() *StopDeploymentResourceMaintenanceModeNotFound {
	return &StopDeploymentResourceMaintenanceModeNotFound{}
}

/* StopDeploymentResourceMaintenanceModeNotFound describes a response with status code 404, with default header values.

 * The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
* One or more instances of the given resource type are missing. (code: `deployments.instances_missing_on_update_error`)
*/
type StopDeploymentResourceMaintenanceModeNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/{instance_ids}/maintenance-mode/_stop][%d] stopDeploymentResourceMaintenanceModeNotFound  %+v", 404, o.Payload)
}
func (o *StopDeploymentResourceMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopDeploymentResourceMaintenanceModeRetryWith creates a StopDeploymentResourceMaintenanceModeRetryWith with default headers values
func NewStopDeploymentResourceMaintenanceModeRetryWith() *StopDeploymentResourceMaintenanceModeRetryWith {
	return &StopDeploymentResourceMaintenanceModeRetryWith{}
}

/* StopDeploymentResourceMaintenanceModeRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StopDeploymentResourceMaintenanceModeRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceMaintenanceModeRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/{instance_ids}/maintenance-mode/_stop][%d] stopDeploymentResourceMaintenanceModeRetryWith  %+v", 449, o.Payload)
}
func (o *StopDeploymentResourceMaintenanceModeRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceMaintenanceModeRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopDeploymentResourceMaintenanceModeInternalServerError creates a StopDeploymentResourceMaintenanceModeInternalServerError with default headers values
func NewStopDeploymentResourceMaintenanceModeInternalServerError() *StopDeploymentResourceMaintenanceModeInternalServerError {
	return &StopDeploymentResourceMaintenanceModeInternalServerError{}
}

/* StopDeploymentResourceMaintenanceModeInternalServerError describes a response with status code 500, with default header values.

A Resource that was previously stored no longer exists. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type StopDeploymentResourceMaintenanceModeInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceMaintenanceModeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/{instance_ids}/maintenance-mode/_stop][%d] stopDeploymentResourceMaintenanceModeInternalServerError  %+v", 500, o.Payload)
}
func (o *StopDeploymentResourceMaintenanceModeInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceMaintenanceModeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
