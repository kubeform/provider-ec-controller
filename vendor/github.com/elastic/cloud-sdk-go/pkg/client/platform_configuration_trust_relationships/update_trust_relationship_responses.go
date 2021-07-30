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

package platform_configuration_trust_relationships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateTrustRelationshipReader is a Reader for the UpdateTrustRelationship structure.
type UpdateTrustRelationshipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTrustRelationshipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTrustRelationshipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTrustRelationshipBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateTrustRelationshipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTrustRelationshipNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTrustRelationshipOK creates a UpdateTrustRelationshipOK with default headers values
func NewUpdateTrustRelationshipOK() *UpdateTrustRelationshipOK {
	return &UpdateTrustRelationshipOK{}
}

/* UpdateTrustRelationshipOK describes a response with status code 200, with default header values.

The request was valid and the trust relationship was updated.
*/
type UpdateTrustRelationshipOK struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.TrustRelationshipUpdateResponse
}

func (o *UpdateTrustRelationshipOK) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/trust-relationships/{trust_relationship_id}][%d] updateTrustRelationshipOK  %+v", 200, o.Payload)
}
func (o *UpdateTrustRelationshipOK) GetPayload() *models.TrustRelationshipUpdateResponse {
	return o.Payload
}

func (o *UpdateTrustRelationshipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-resource-created
	hdrXCloudResourceCreated := response.GetHeader("x-cloud-resource-created")

	if hdrXCloudResourceCreated != "" {
		o.XCloudResourceCreated = hdrXCloudResourceCreated
	}

	// hydrates response header x-cloud-resource-last-modified
	hdrXCloudResourceLastModified := response.GetHeader("x-cloud-resource-last-modified")

	if hdrXCloudResourceLastModified != "" {
		o.XCloudResourceLastModified = hdrXCloudResourceLastModified
	}

	// hydrates response header x-cloud-resource-version
	hdrXCloudResourceVersion := response.GetHeader("x-cloud-resource-version")

	if hdrXCloudResourceVersion != "" {
		o.XCloudResourceVersion = hdrXCloudResourceVersion
	}

	o.Payload = new(models.TrustRelationshipUpdateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTrustRelationshipBadRequest creates a UpdateTrustRelationshipBadRequest with default headers values
func NewUpdateTrustRelationshipBadRequest() *UpdateTrustRelationshipBadRequest {
	return &UpdateTrustRelationshipBadRequest{}
}

/* UpdateTrustRelationshipBadRequest describes a response with status code 400, with default header values.

The trust relationship request had errors.
*/
type UpdateTrustRelationshipBadRequest struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateTrustRelationshipBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/trust-relationships/{trust_relationship_id}][%d] updateTrustRelationshipBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateTrustRelationshipBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateTrustRelationshipBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTrustRelationshipUnauthorized creates a UpdateTrustRelationshipUnauthorized with default headers values
func NewUpdateTrustRelationshipUnauthorized() *UpdateTrustRelationshipUnauthorized {
	return &UpdateTrustRelationshipUnauthorized{}
}

/* UpdateTrustRelationshipUnauthorized describes a response with status code 401, with default header values.

You are not authorized to perform this action.
*/
type UpdateTrustRelationshipUnauthorized struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateTrustRelationshipUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/trust-relationships/{trust_relationship_id}][%d] updateTrustRelationshipUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateTrustRelationshipUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateTrustRelationshipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTrustRelationshipNotFound creates a UpdateTrustRelationshipNotFound with default headers values
func NewUpdateTrustRelationshipNotFound() *UpdateTrustRelationshipNotFound {
	return &UpdateTrustRelationshipNotFound{}
}

/* UpdateTrustRelationshipNotFound describes a response with status code 404, with default header values.

The trust relationship specified by {trust_relationship_id} cannot be found. (code: `trust_relationships.not_found`)
*/
type UpdateTrustRelationshipNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateTrustRelationshipNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/trust-relationships/{trust_relationship_id}][%d] updateTrustRelationshipNotFound  %+v", 404, o.Payload)
}
func (o *UpdateTrustRelationshipNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateTrustRelationshipNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
