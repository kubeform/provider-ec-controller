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

// UpdateRunnerLoggingSettingsReader is a Reader for the UpdateRunnerLoggingSettings structure.
type UpdateRunnerLoggingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRunnerLoggingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRunnerLoggingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRunnerLoggingSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRunnerLoggingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRunnerLoggingSettingsOK creates a UpdateRunnerLoggingSettingsOK with default headers values
func NewUpdateRunnerLoggingSettingsOK() *UpdateRunnerLoggingSettingsOK {
	return &UpdateRunnerLoggingSettingsOK{}
}

/* UpdateRunnerLoggingSettingsOK describes a response with status code 200, with default header values.

The updated logging settings for the runner instance
*/
type UpdateRunnerLoggingSettingsOK struct {

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

func (o *UpdateRunnerLoggingSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/runners/{runner_id}/logging_settings][%d] updateRunnerLoggingSettingsOK  %+v", 200, o.Payload)
}
func (o *UpdateRunnerLoggingSettingsOK) GetPayload() *models.LoggingSettings {
	return o.Payload
}

func (o *UpdateRunnerLoggingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRunnerLoggingSettingsBadRequest creates a UpdateRunnerLoggingSettingsBadRequest with default headers values
func NewUpdateRunnerLoggingSettingsBadRequest() *UpdateRunnerLoggingSettingsBadRequest {
	return &UpdateRunnerLoggingSettingsBadRequest{}
}

/* UpdateRunnerLoggingSettingsBadRequest describes a response with status code 400, with default header values.

The update request is invalid. (code: `patch.request_malformed`)
*/
type UpdateRunnerLoggingSettingsBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateRunnerLoggingSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/runners/{runner_id}/logging_settings][%d] updateRunnerLoggingSettingsBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateRunnerLoggingSettingsBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateRunnerLoggingSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRunnerLoggingSettingsNotFound creates a UpdateRunnerLoggingSettingsNotFound with default headers values
func NewUpdateRunnerLoggingSettingsNotFound() *UpdateRunnerLoggingSettingsNotFound {
	return &UpdateRunnerLoggingSettingsNotFound{}
}

/* UpdateRunnerLoggingSettingsNotFound describes a response with status code 404, with default header values.

The logging settings for this runner were not found. (code: `runners.logging_settings.not_found`)
*/
type UpdateRunnerLoggingSettingsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateRunnerLoggingSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/runners/{runner_id}/logging_settings][%d] updateRunnerLoggingSettingsNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRunnerLoggingSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateRunnerLoggingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
