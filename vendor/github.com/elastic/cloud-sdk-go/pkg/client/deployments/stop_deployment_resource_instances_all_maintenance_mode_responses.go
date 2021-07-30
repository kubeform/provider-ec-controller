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

// StopDeploymentResourceInstancesAllMaintenanceModeReader is a Reader for the StopDeploymentResourceInstancesAllMaintenanceMode structure.
type StopDeploymentResourceInstancesAllMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopDeploymentResourceInstancesAllMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStopDeploymentResourceInstancesAllMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStopDeploymentResourceInstancesAllMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopDeploymentResourceInstancesAllMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStopDeploymentResourceInstancesAllMaintenanceModeRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopDeploymentResourceInstancesAllMaintenanceModeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopDeploymentResourceInstancesAllMaintenanceModeAccepted creates a StopDeploymentResourceInstancesAllMaintenanceModeAccepted with default headers values
func NewStopDeploymentResourceInstancesAllMaintenanceModeAccepted() *StopDeploymentResourceInstancesAllMaintenanceModeAccepted {
	return &StopDeploymentResourceInstancesAllMaintenanceModeAccepted{}
}

/* StopDeploymentResourceInstancesAllMaintenanceModeAccepted describes a response with status code 202, with default header values.

The stop maintenance mode command was issued successfully.
*/
type StopDeploymentResourceInstancesAllMaintenanceModeAccepted struct {
	Payload *models.DeploymentResourceCommandResponse
}

func (o *StopDeploymentResourceInstancesAllMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_stop][%d] stopDeploymentResourceInstancesAllMaintenanceModeAccepted  %+v", 202, o.Payload)
}
func (o *StopDeploymentResourceInstancesAllMaintenanceModeAccepted) GetPayload() *models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *StopDeploymentResourceInstancesAllMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopDeploymentResourceInstancesAllMaintenanceModeForbidden creates a StopDeploymentResourceInstancesAllMaintenanceModeForbidden with default headers values
func NewStopDeploymentResourceInstancesAllMaintenanceModeForbidden() *StopDeploymentResourceInstancesAllMaintenanceModeForbidden {
	return &StopDeploymentResourceInstancesAllMaintenanceModeForbidden{}
}

/* StopDeploymentResourceInstancesAllMaintenanceModeForbidden describes a response with status code 403, with default header values.

The stop maintenance mode command was prohibited for the given Resource. (code: `deployments.instance_update_prohibited_error`)
*/
type StopDeploymentResourceInstancesAllMaintenanceModeForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceInstancesAllMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_stop][%d] stopDeploymentResourceInstancesAllMaintenanceModeForbidden  %+v", 403, o.Payload)
}
func (o *StopDeploymentResourceInstancesAllMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceInstancesAllMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopDeploymentResourceInstancesAllMaintenanceModeNotFound creates a StopDeploymentResourceInstancesAllMaintenanceModeNotFound with default headers values
func NewStopDeploymentResourceInstancesAllMaintenanceModeNotFound() *StopDeploymentResourceInstancesAllMaintenanceModeNotFound {
	return &StopDeploymentResourceInstancesAllMaintenanceModeNotFound{}
}

/* StopDeploymentResourceInstancesAllMaintenanceModeNotFound describes a response with status code 404, with default header values.

 * The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
* One or more instances of the given resource type are missing. (code: `deployments.instances_missing_on_update_error`)
*/
type StopDeploymentResourceInstancesAllMaintenanceModeNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceInstancesAllMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_stop][%d] stopDeploymentResourceInstancesAllMaintenanceModeNotFound  %+v", 404, o.Payload)
}
func (o *StopDeploymentResourceInstancesAllMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceInstancesAllMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopDeploymentResourceInstancesAllMaintenanceModeRetryWith creates a StopDeploymentResourceInstancesAllMaintenanceModeRetryWith with default headers values
func NewStopDeploymentResourceInstancesAllMaintenanceModeRetryWith() *StopDeploymentResourceInstancesAllMaintenanceModeRetryWith {
	return &StopDeploymentResourceInstancesAllMaintenanceModeRetryWith{}
}

/* StopDeploymentResourceInstancesAllMaintenanceModeRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StopDeploymentResourceInstancesAllMaintenanceModeRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceInstancesAllMaintenanceModeRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_stop][%d] stopDeploymentResourceInstancesAllMaintenanceModeRetryWith  %+v", 449, o.Payload)
}
func (o *StopDeploymentResourceInstancesAllMaintenanceModeRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceInstancesAllMaintenanceModeRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopDeploymentResourceInstancesAllMaintenanceModeInternalServerError creates a StopDeploymentResourceInstancesAllMaintenanceModeInternalServerError with default headers values
func NewStopDeploymentResourceInstancesAllMaintenanceModeInternalServerError() *StopDeploymentResourceInstancesAllMaintenanceModeInternalServerError {
	return &StopDeploymentResourceInstancesAllMaintenanceModeInternalServerError{}
}

/* StopDeploymentResourceInstancesAllMaintenanceModeInternalServerError describes a response with status code 500, with default header values.

A Resource that was previously stored no longer exists. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type StopDeploymentResourceInstancesAllMaintenanceModeInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopDeploymentResourceInstancesAllMaintenanceModeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_stop][%d] stopDeploymentResourceInstancesAllMaintenanceModeInternalServerError  %+v", 500, o.Payload)
}
func (o *StopDeploymentResourceInstancesAllMaintenanceModeInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceInstancesAllMaintenanceModeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
