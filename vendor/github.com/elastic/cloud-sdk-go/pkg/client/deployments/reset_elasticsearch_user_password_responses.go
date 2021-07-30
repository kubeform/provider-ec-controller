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

// ResetElasticsearchUserPasswordReader is a Reader for the ResetElasticsearchUserPassword structure.
type ResetElasticsearchUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetElasticsearchUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetElasticsearchUserPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewResetElasticsearchUserPasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewResetElasticsearchUserPasswordRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResetElasticsearchUserPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResetElasticsearchUserPasswordOK creates a ResetElasticsearchUserPasswordOK with default headers values
func NewResetElasticsearchUserPasswordOK() *ResetElasticsearchUserPasswordOK {
	return &ResetElasticsearchUserPasswordOK{}
}

/* ResetElasticsearchUserPasswordOK describes a response with status code 200, with default header values.

The password reset was out carried successfully
*/
type ResetElasticsearchUserPasswordOK struct {
	Payload *models.ElasticsearchElasticUserPasswordResetResponse
}

func (o *ResetElasticsearchUserPasswordOK) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_reset-password][%d] resetElasticsearchUserPasswordOK  %+v", 200, o.Payload)
}
func (o *ResetElasticsearchUserPasswordOK) GetPayload() *models.ElasticsearchElasticUserPasswordResetResponse {
	return o.Payload
}

func (o *ResetElasticsearchUserPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElasticsearchElasticUserPasswordResetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetElasticsearchUserPasswordNotFound creates a ResetElasticsearchUserPasswordNotFound with default headers values
func NewResetElasticsearchUserPasswordNotFound() *ResetElasticsearchUserPasswordNotFound {
	return &ResetElasticsearchUserPasswordNotFound{}
}

/* ResetElasticsearchUserPasswordNotFound describes a response with status code 404, with default header values.

 * The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type ResetElasticsearchUserPasswordNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResetElasticsearchUserPasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_reset-password][%d] resetElasticsearchUserPasswordNotFound  %+v", 404, o.Payload)
}
func (o *ResetElasticsearchUserPasswordNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResetElasticsearchUserPasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewResetElasticsearchUserPasswordRetryWith creates a ResetElasticsearchUserPasswordRetryWith with default headers values
func NewResetElasticsearchUserPasswordRetryWith() *ResetElasticsearchUserPasswordRetryWith {
	return &ResetElasticsearchUserPasswordRetryWith{}
}

/* ResetElasticsearchUserPasswordRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ResetElasticsearchUserPasswordRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResetElasticsearchUserPasswordRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_reset-password][%d] resetElasticsearchUserPasswordRetryWith  %+v", 449, o.Payload)
}
func (o *ResetElasticsearchUserPasswordRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResetElasticsearchUserPasswordRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewResetElasticsearchUserPasswordInternalServerError creates a ResetElasticsearchUserPasswordInternalServerError with default headers values
func NewResetElasticsearchUserPasswordInternalServerError() *ResetElasticsearchUserPasswordInternalServerError {
	return &ResetElasticsearchUserPasswordInternalServerError{}
}

/* ResetElasticsearchUserPasswordInternalServerError describes a response with status code 500, with default header values.

Failed to reset the 'elastic' user's password. (code: `deployments.elasticsearch.password_reset_error`)
*/
type ResetElasticsearchUserPasswordInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResetElasticsearchUserPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_reset-password][%d] resetElasticsearchUserPasswordInternalServerError  %+v", 500, o.Payload)
}
func (o *ResetElasticsearchUserPasswordInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResetElasticsearchUserPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
