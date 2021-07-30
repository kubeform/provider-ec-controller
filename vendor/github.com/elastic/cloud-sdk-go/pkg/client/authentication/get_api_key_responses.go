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

// GetAPIKeyReader is a Reader for the GetAPIKey structure.
type GetAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAPIKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIKeyOK creates a GetAPIKeyOK with default headers values
func NewGetAPIKeyOK() *GetAPIKeyOK {
	return &GetAPIKeyOK{}
}

/* GetAPIKeyOK describes a response with status code 200, with default header values.

The API key metadata is retrieved.
*/
type GetAPIKeyOK struct {
	Payload *models.APIKeyResponse
}

func (o *GetAPIKeyOK) Error() string {
	return fmt.Sprintf("[GET /users/auth/keys/{api_key_id}][%d] getApiKeyOK  %+v", 200, o.Payload)
}
func (o *GetAPIKeyOK) GetPayload() *models.APIKeyResponse {
	return o.Payload
}

func (o *GetAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIKeyNotFound creates a GetAPIKeyNotFound with default headers values
func NewGetAPIKeyNotFound() *GetAPIKeyNotFound {
	return &GetAPIKeyNotFound{}
}

/* GetAPIKeyNotFound describes a response with status code 404, with default header values.

The {api_key_id} can't be found. (code: `api_keys.key_not_found`)
*/
type GetAPIKeyNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetAPIKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /users/auth/keys/{api_key_id}][%d] getApiKeyNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIKeyNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetAPIKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
