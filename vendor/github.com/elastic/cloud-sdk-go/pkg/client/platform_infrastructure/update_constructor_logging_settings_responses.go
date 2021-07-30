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

// UpdateConstructorLoggingSettingsReader is a Reader for the UpdateConstructorLoggingSettings structure.
type UpdateConstructorLoggingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConstructorLoggingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateConstructorLoggingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateConstructorLoggingSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConstructorLoggingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateConstructorLoggingSettingsOK creates a UpdateConstructorLoggingSettingsOK with default headers values
func NewUpdateConstructorLoggingSettingsOK() *UpdateConstructorLoggingSettingsOK {
	return &UpdateConstructorLoggingSettingsOK{}
}

/* UpdateConstructorLoggingSettingsOK describes a response with status code 200, with default header values.

The updated logging settings for the constructor instance
*/
type UpdateConstructorLoggingSettingsOK struct {

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

func (o *UpdateConstructorLoggingSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] updateConstructorLoggingSettingsOK  %+v", 200, o.Payload)
}
func (o *UpdateConstructorLoggingSettingsOK) GetPayload() *models.LoggingSettings {
	return o.Payload
}

func (o *UpdateConstructorLoggingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateConstructorLoggingSettingsBadRequest creates a UpdateConstructorLoggingSettingsBadRequest with default headers values
func NewUpdateConstructorLoggingSettingsBadRequest() *UpdateConstructorLoggingSettingsBadRequest {
	return &UpdateConstructorLoggingSettingsBadRequest{}
}

/* UpdateConstructorLoggingSettingsBadRequest describes a response with status code 400, with default header values.

The update request is invalid. (code: `patch.request_malformed`)
*/
type UpdateConstructorLoggingSettingsBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateConstructorLoggingSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] updateConstructorLoggingSettingsBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateConstructorLoggingSettingsBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateConstructorLoggingSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateConstructorLoggingSettingsNotFound creates a UpdateConstructorLoggingSettingsNotFound with default headers values
func NewUpdateConstructorLoggingSettingsNotFound() *UpdateConstructorLoggingSettingsNotFound {
	return &UpdateConstructorLoggingSettingsNotFound{}
}

/* UpdateConstructorLoggingSettingsNotFound describes a response with status code 404, with default header values.

The logging settings for this constructor were not found. (code: `constructors.logging_settings.not_found`)
*/
type UpdateConstructorLoggingSettingsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateConstructorLoggingSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] updateConstructorLoggingSettingsNotFound  %+v", 404, o.Payload)
}
func (o *UpdateConstructorLoggingSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateConstructorLoggingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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