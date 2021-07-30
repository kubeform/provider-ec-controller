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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetEsClusterInstancesSettingsOverridesReader is a Reader for the SetEsClusterInstancesSettingsOverrides structure.
type SetEsClusterInstancesSettingsOverridesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetEsClusterInstancesSettingsOverridesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetEsClusterInstancesSettingsOverridesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSetEsClusterInstancesSettingsOverridesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewSetEsClusterInstancesSettingsOverridesRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetEsClusterInstancesSettingsOverridesOK creates a SetEsClusterInstancesSettingsOverridesOK with default headers values
func NewSetEsClusterInstancesSettingsOverridesOK() *SetEsClusterInstancesSettingsOverridesOK {
	return &SetEsClusterInstancesSettingsOverridesOK{}
}

/* SetEsClusterInstancesSettingsOverridesOK describes a response with status code 200, with default header values.

Returns the updated settings overrides for the specified instances
*/
type SetEsClusterInstancesSettingsOverridesOK struct {
	Payload *models.ElasticsearchClusterInstanceSettingsOverrides
}

func (o *SetEsClusterInstancesSettingsOverridesOK) Error() string {
	return fmt.Sprintf("[PUT /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/settings][%d] setEsClusterInstancesSettingsOverridesOK  %+v", 200, o.Payload)
}
func (o *SetEsClusterInstancesSettingsOverridesOK) GetPayload() *models.ElasticsearchClusterInstanceSettingsOverrides {
	return o.Payload
}

func (o *SetEsClusterInstancesSettingsOverridesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElasticsearchClusterInstanceSettingsOverrides)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetEsClusterInstancesSettingsOverridesNotFound creates a SetEsClusterInstancesSettingsOverridesNotFound with default headers values
func NewSetEsClusterInstancesSettingsOverridesNotFound() *SetEsClusterInstancesSettingsOverridesNotFound {
	return &SetEsClusterInstancesSettingsOverridesNotFound{}
}

/* SetEsClusterInstancesSettingsOverridesNotFound describes a response with status code 404, with default header values.

One or more of the instances specified at {instance_ids} could not be found. (code: `clusters.instances_not_found`)
*/
type SetEsClusterInstancesSettingsOverridesNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetEsClusterInstancesSettingsOverridesNotFound) Error() string {
	return fmt.Sprintf("[PUT /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/settings][%d] setEsClusterInstancesSettingsOverridesNotFound  %+v", 404, o.Payload)
}
func (o *SetEsClusterInstancesSettingsOverridesNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetEsClusterInstancesSettingsOverridesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetEsClusterInstancesSettingsOverridesRetryWith creates a SetEsClusterInstancesSettingsOverridesRetryWith with default headers values
func NewSetEsClusterInstancesSettingsOverridesRetryWith() *SetEsClusterInstancesSettingsOverridesRetryWith {
	return &SetEsClusterInstancesSettingsOverridesRetryWith{}
}

/* SetEsClusterInstancesSettingsOverridesRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type SetEsClusterInstancesSettingsOverridesRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetEsClusterInstancesSettingsOverridesRetryWith) Error() string {
	return fmt.Sprintf("[PUT /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/settings][%d] setEsClusterInstancesSettingsOverridesRetryWith  %+v", 449, o.Payload)
}
func (o *SetEsClusterInstancesSettingsOverridesRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetEsClusterInstancesSettingsOverridesRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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