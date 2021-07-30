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

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteUserAPIKeyReader is a Reader for the DeleteUserAPIKey structure.
type DeleteUserAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteUserAPIKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewDeleteUserAPIKeyRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserAPIKeyOK creates a DeleteUserAPIKeyOK with default headers values
func NewDeleteUserAPIKeyOK() *DeleteUserAPIKeyOK {
	return &DeleteUserAPIKeyOK{}
}

/* DeleteUserAPIKeyOK describes a response with status code 200, with default header values.

The API key is deleted.
*/
type DeleteUserAPIKeyOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteUserAPIKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/auth/keys/{api_key_id}][%d] deleteUserApiKeyOK  %+v", 200, o.Payload)
}
func (o *DeleteUserAPIKeyOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteUserAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserAPIKeyNotFound creates a DeleteUserAPIKeyNotFound with default headers values
func NewDeleteUserAPIKeyNotFound() *DeleteUserAPIKeyNotFound {
	return &DeleteUserAPIKeyNotFound{}
}

/* DeleteUserAPIKeyNotFound describes a response with status code 404, with default header values.

The {api_key_id} can't be found. (code: `api_keys.key_not_found`)
*/
type DeleteUserAPIKeyNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteUserAPIKeyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/auth/keys/{api_key_id}][%d] deleteUserApiKeyNotFound  %+v", 404, o.Payload)
}
func (o *DeleteUserAPIKeyNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteUserAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserAPIKeyRetryWith creates a DeleteUserAPIKeyRetryWith with default headers values
func NewDeleteUserAPIKeyRetryWith() *DeleteUserAPIKeyRetryWith {
	return &DeleteUserAPIKeyRetryWith{}
}

/* DeleteUserAPIKeyRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type DeleteUserAPIKeyRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteUserAPIKeyRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/auth/keys/{api_key_id}][%d] deleteUserApiKeyRetryWith  %+v", 449, o.Payload)
}
func (o *DeleteUserAPIKeyRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteUserAPIKeyRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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