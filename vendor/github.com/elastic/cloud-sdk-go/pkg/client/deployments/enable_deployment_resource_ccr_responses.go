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

// EnableDeploymentResourceCcrReader is a Reader for the EnableDeploymentResourceCcr structure.
type EnableDeploymentResourceCcrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableDeploymentResourceCcrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableDeploymentResourceCcrOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewEnableDeploymentResourceCcrNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewEnableDeploymentResourceCcrRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnableDeploymentResourceCcrOK creates a EnableDeploymentResourceCcrOK with default headers values
func NewEnableDeploymentResourceCcrOK() *EnableDeploymentResourceCcrOK {
	return &EnableDeploymentResourceCcrOK{}
}

/* EnableDeploymentResourceCcrOK describes a response with status code 200, with default header values.

Standard response
*/
type EnableDeploymentResourceCcrOK struct {
	Payload *models.DeploymentResourceCommandResponse
}

func (o *EnableDeploymentResourceCcrOK) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_enable-ccr][%d] enableDeploymentResourceCcrOK  %+v", 200, o.Payload)
}
func (o *EnableDeploymentResourceCcrOK) GetPayload() *models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *EnableDeploymentResourceCcrOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableDeploymentResourceCcrNotFound creates a EnableDeploymentResourceCcrNotFound with default headers values
func NewEnableDeploymentResourceCcrNotFound() *EnableDeploymentResourceCcrNotFound {
	return &EnableDeploymentResourceCcrNotFound{}
}

/* EnableDeploymentResourceCcrNotFound describes a response with status code 404, with default header values.

 * The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type EnableDeploymentResourceCcrNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *EnableDeploymentResourceCcrNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_enable-ccr][%d] enableDeploymentResourceCcrNotFound  %+v", 404, o.Payload)
}
func (o *EnableDeploymentResourceCcrNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *EnableDeploymentResourceCcrNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEnableDeploymentResourceCcrRetryWith creates a EnableDeploymentResourceCcrRetryWith with default headers values
func NewEnableDeploymentResourceCcrRetryWith() *EnableDeploymentResourceCcrRetryWith {
	return &EnableDeploymentResourceCcrRetryWith{}
}

/* EnableDeploymentResourceCcrRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type EnableDeploymentResourceCcrRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *EnableDeploymentResourceCcrRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_enable-ccr][%d] enableDeploymentResourceCcrRetryWith  %+v", 449, o.Payload)
}
func (o *EnableDeploymentResourceCcrRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *EnableDeploymentResourceCcrRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
