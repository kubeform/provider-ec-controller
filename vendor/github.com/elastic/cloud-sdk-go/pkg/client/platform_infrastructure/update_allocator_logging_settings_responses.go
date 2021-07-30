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

package platform_infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateAllocatorLoggingSettingsReader is a Reader for the UpdateAllocatorLoggingSettings structure.
type UpdateAllocatorLoggingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAllocatorLoggingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAllocatorLoggingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAllocatorLoggingSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAllocatorLoggingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAllocatorLoggingSettingsOK creates a UpdateAllocatorLoggingSettingsOK with default headers values
func NewUpdateAllocatorLoggingSettingsOK() *UpdateAllocatorLoggingSettingsOK {
	return &UpdateAllocatorLoggingSettingsOK{}
}

/* UpdateAllocatorLoggingSettingsOK describes a response with status code 200, with default header values.

The updated logging settings for the allocator instance
*/
type UpdateAllocatorLoggingSettingsOK struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.LoggingSettings
}

func (o *UpdateAllocatorLoggingSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/allocators/{allocator_id}/logging_settings][%d] updateAllocatorLoggingSettingsOK  %+v", 200, o.Payload)
}
func (o *UpdateAllocatorLoggingSettingsOK) GetPayload() *models.LoggingSettings {
	return o.Payload
}

func (o *UpdateAllocatorLoggingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.LoggingSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAllocatorLoggingSettingsBadRequest creates a UpdateAllocatorLoggingSettingsBadRequest with default headers values
func NewUpdateAllocatorLoggingSettingsBadRequest() *UpdateAllocatorLoggingSettingsBadRequest {
	return &UpdateAllocatorLoggingSettingsBadRequest{}
}

/* UpdateAllocatorLoggingSettingsBadRequest describes a response with status code 400, with default header values.

The update request is invalid. (code: `patch.request_malformed`)
*/
type UpdateAllocatorLoggingSettingsBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateAllocatorLoggingSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/allocators/{allocator_id}/logging_settings][%d] updateAllocatorLoggingSettingsBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateAllocatorLoggingSettingsBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateAllocatorLoggingSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateAllocatorLoggingSettingsNotFound creates a UpdateAllocatorLoggingSettingsNotFound with default headers values
func NewUpdateAllocatorLoggingSettingsNotFound() *UpdateAllocatorLoggingSettingsNotFound {
	return &UpdateAllocatorLoggingSettingsNotFound{}
}

/* UpdateAllocatorLoggingSettingsNotFound describes a response with status code 404, with default header values.

The logging settings for this allocator were not found. (code: `allocators.logging_settings.not_found`)
*/
type UpdateAllocatorLoggingSettingsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateAllocatorLoggingSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/allocators/{allocator_id}/logging_settings][%d] updateAllocatorLoggingSettingsNotFound  %+v", 404, o.Payload)
}
func (o *UpdateAllocatorLoggingSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateAllocatorLoggingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
