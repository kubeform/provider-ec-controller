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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpgradeApmReader is a Reader for the UpgradeApm structure.
type UpgradeApmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeApmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewUpgradeApmAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpgradeApmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpgradeApmRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpgradeApmAccepted creates a UpgradeApmAccepted with default headers values
func NewUpgradeApmAccepted() *UpgradeApmAccepted {
	return &UpgradeApmAccepted{}
}

/* UpgradeApmAccepted describes a response with status code 202, with default header values.

The upgrade command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type UpgradeApmAccepted struct {
	Payload *models.ClusterUpgradeInfo
}

func (o *UpgradeApmAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_upgrade][%d] upgradeApmAccepted  %+v", 202, o.Payload)
}
func (o *UpgradeApmAccepted) GetPayload() *models.ClusterUpgradeInfo {
	return o.Payload
}

func (o *UpgradeApmAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterUpgradeInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpgradeApmNotFound creates a UpgradeApmNotFound with default headers values
func NewUpgradeApmNotFound() *UpgradeApmNotFound {
	return &UpgradeApmNotFound{}
}

/* UpgradeApmNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type UpgradeApmNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpgradeApmNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_upgrade][%d] upgradeApmNotFound  %+v", 404, o.Payload)
}
func (o *UpgradeApmNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpgradeApmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpgradeApmRetryWith creates a UpgradeApmRetryWith with default headers values
func NewUpgradeApmRetryWith() *UpgradeApmRetryWith {
	return &UpgradeApmRetryWith{}
}

/* UpgradeApmRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type UpgradeApmRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpgradeApmRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_upgrade][%d] upgradeApmRetryWith  %+v", 449, o.Payload)
}
func (o *UpgradeApmRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpgradeApmRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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