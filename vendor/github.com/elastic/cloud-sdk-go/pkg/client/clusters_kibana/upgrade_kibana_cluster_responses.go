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

// UpgradeKibanaClusterReader is a Reader for the UpgradeKibanaCluster structure.
type UpgradeKibanaClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeKibanaClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUpgradeKibanaClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpgradeKibanaClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpgradeKibanaClusterRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpgradeKibanaClusterAccepted creates a UpgradeKibanaClusterAccepted with default headers values
func NewUpgradeKibanaClusterAccepted() *UpgradeKibanaClusterAccepted {
	return &UpgradeKibanaClusterAccepted{}
}

/* UpgradeKibanaClusterAccepted describes a response with status code 202, with default header values.

The upgrade command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type UpgradeKibanaClusterAccepted struct {
	Payload *models.ClusterUpgradeInfo
}

func (o *UpgradeKibanaClusterAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/_upgrade][%d] upgradeKibanaClusterAccepted  %+v", 202, o.Payload)
}
func (o *UpgradeKibanaClusterAccepted) GetPayload() *models.ClusterUpgradeInfo {
	return o.Payload
}

func (o *UpgradeKibanaClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterUpgradeInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpgradeKibanaClusterNotFound creates a UpgradeKibanaClusterNotFound with default headers values
func NewUpgradeKibanaClusterNotFound() *UpgradeKibanaClusterNotFound {
	return &UpgradeKibanaClusterNotFound{}
}

/* UpgradeKibanaClusterNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type UpgradeKibanaClusterNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpgradeKibanaClusterNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/_upgrade][%d] upgradeKibanaClusterNotFound  %+v", 404, o.Payload)
}
func (o *UpgradeKibanaClusterNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpgradeKibanaClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpgradeKibanaClusterRetryWith creates a UpgradeKibanaClusterRetryWith with default headers values
func NewUpgradeKibanaClusterRetryWith() *UpgradeKibanaClusterRetryWith {
	return &UpgradeKibanaClusterRetryWith{}
}

/* UpgradeKibanaClusterRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type UpgradeKibanaClusterRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpgradeKibanaClusterRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/_upgrade][%d] upgradeKibanaClusterRetryWith  %+v", 449, o.Payload)
}
func (o *UpgradeKibanaClusterRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpgradeKibanaClusterRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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