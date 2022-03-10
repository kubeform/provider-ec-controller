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

// GetDeploymentEsResourceKeystoreReader is a Reader for the GetDeploymentEsResourceKeystore structure.
type GetDeploymentEsResourceKeystoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentEsResourceKeystoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentEsResourceKeystoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeploymentEsResourceKeystoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewGetDeploymentEsResourceKeystoreRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeploymentEsResourceKeystoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentEsResourceKeystoreOK creates a GetDeploymentEsResourceKeystoreOK with default headers values
func NewGetDeploymentEsResourceKeystoreOK() *GetDeploymentEsResourceKeystoreOK {
	return &GetDeploymentEsResourceKeystoreOK{}
}

/* GetDeploymentEsResourceKeystoreOK describes a response with status code 200, with default header values.

The contents of the Elasticsearch keystore
*/
type GetDeploymentEsResourceKeystoreOK struct {
	Payload *models.KeystoreContents
}

func (o *GetDeploymentEsResourceKeystoreOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] getDeploymentEsResourceKeystoreOK  %+v", 200, o.Payload)
}
func (o *GetDeploymentEsResourceKeystoreOK) GetPayload() *models.KeystoreContents {
	return o.Payload
}

func (o *GetDeploymentEsResourceKeystoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeystoreContents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentEsResourceKeystoreNotFound creates a GetDeploymentEsResourceKeystoreNotFound with default headers values
func NewGetDeploymentEsResourceKeystoreNotFound() *GetDeploymentEsResourceKeystoreNotFound {
	return &GetDeploymentEsResourceKeystoreNotFound{}
}

/* GetDeploymentEsResourceKeystoreNotFound describes a response with status code 404, with default header values.

 * The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type GetDeploymentEsResourceKeystoreNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentEsResourceKeystoreNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] getDeploymentEsResourceKeystoreNotFound  %+v", 404, o.Payload)
}
func (o *GetDeploymentEsResourceKeystoreNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEsResourceKeystoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeploymentEsResourceKeystoreRetryWith creates a GetDeploymentEsResourceKeystoreRetryWith with default headers values
func NewGetDeploymentEsResourceKeystoreRetryWith() *GetDeploymentEsResourceKeystoreRetryWith {
	return &GetDeploymentEsResourceKeystoreRetryWith{}
}

/* GetDeploymentEsResourceKeystoreRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type GetDeploymentEsResourceKeystoreRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentEsResourceKeystoreRetryWith) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] getDeploymentEsResourceKeystoreRetryWith  %+v", 449, o.Payload)
}
func (o *GetDeploymentEsResourceKeystoreRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEsResourceKeystoreRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeploymentEsResourceKeystoreInternalServerError creates a GetDeploymentEsResourceKeystoreInternalServerError with default headers values
func NewGetDeploymentEsResourceKeystoreInternalServerError() *GetDeploymentEsResourceKeystoreInternalServerError {
	return &GetDeploymentEsResourceKeystoreInternalServerError{}
}

/* GetDeploymentEsResourceKeystoreInternalServerError describes a response with status code 500, with default header values.

We have failed you. (code: `deployments.metadata_internal_error`)
*/
type GetDeploymentEsResourceKeystoreInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentEsResourceKeystoreInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] getDeploymentEsResourceKeystoreInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDeploymentEsResourceKeystoreInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEsResourceKeystoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
