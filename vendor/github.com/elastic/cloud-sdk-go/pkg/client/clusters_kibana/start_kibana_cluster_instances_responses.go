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

package clusters_kibana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// StartKibanaClusterInstancesReader is a Reader for the StartKibanaClusterInstances structure.
type StartKibanaClusterInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartKibanaClusterInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartKibanaClusterInstancesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartKibanaClusterInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartKibanaClusterInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStartKibanaClusterInstancesRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartKibanaClusterInstancesAccepted creates a StartKibanaClusterInstancesAccepted with default headers values
func NewStartKibanaClusterInstancesAccepted() *StartKibanaClusterInstancesAccepted {
	return &StartKibanaClusterInstancesAccepted{}
}

/* StartKibanaClusterInstancesAccepted describes a response with status code 202, with default header values.

The start command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type StartKibanaClusterInstancesAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StartKibanaClusterInstancesAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/{instance_ids}/_start][%d] startKibanaClusterInstancesAccepted  %+v", 202, o.Payload)
}
func (o *StartKibanaClusterInstancesAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *StartKibanaClusterInstancesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartKibanaClusterInstancesForbidden creates a StartKibanaClusterInstancesForbidden with default headers values
func NewStartKibanaClusterInstancesForbidden() *StartKibanaClusterInstancesForbidden {
	return &StartKibanaClusterInstancesForbidden{}
}

/* StartKibanaClusterInstancesForbidden describes a response with status code 403, with default header values.

The start command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StartKibanaClusterInstancesForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartKibanaClusterInstancesForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/{instance_ids}/_start][%d] startKibanaClusterInstancesForbidden  %+v", 403, o.Payload)
}
func (o *StartKibanaClusterInstancesForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartKibanaClusterInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartKibanaClusterInstancesNotFound creates a StartKibanaClusterInstancesNotFound with default headers values
func NewStartKibanaClusterInstancesNotFound() *StartKibanaClusterInstancesNotFound {
	return &StartKibanaClusterInstancesNotFound{}
}

/* StartKibanaClusterInstancesNotFound describes a response with status code 404, with default header values.

 * The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
* One or more of the instances specified at {instance_ids} could not be found. (code: `clusters.instances_not_found`)
*/
type StartKibanaClusterInstancesNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartKibanaClusterInstancesNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/{instance_ids}/_start][%d] startKibanaClusterInstancesNotFound  %+v", 404, o.Payload)
}
func (o *StartKibanaClusterInstancesNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartKibanaClusterInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartKibanaClusterInstancesRetryWith creates a StartKibanaClusterInstancesRetryWith with default headers values
func NewStartKibanaClusterInstancesRetryWith() *StartKibanaClusterInstancesRetryWith {
	return &StartKibanaClusterInstancesRetryWith{}
}

/* StartKibanaClusterInstancesRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StartKibanaClusterInstancesRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartKibanaClusterInstancesRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/{instance_ids}/_start][%d] startKibanaClusterInstancesRetryWith  %+v", 449, o.Payload)
}
func (o *StartKibanaClusterInstancesRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartKibanaClusterInstancesRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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