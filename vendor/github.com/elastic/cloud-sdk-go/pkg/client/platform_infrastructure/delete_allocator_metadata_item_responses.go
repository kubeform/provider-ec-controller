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

// DeleteAllocatorMetadataItemReader is a Reader for the DeleteAllocatorMetadataItem structure.
type DeleteAllocatorMetadataItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAllocatorMetadataItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAllocatorMetadataItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteAllocatorMetadataItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewDeleteAllocatorMetadataItemRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAllocatorMetadataItemOK creates a DeleteAllocatorMetadataItemOK with default headers values
func NewDeleteAllocatorMetadataItemOK() *DeleteAllocatorMetadataItemOK {
	return &DeleteAllocatorMetadataItemOK{}
}

/* DeleteAllocatorMetadataItemOK describes a response with status code 200, with default header values.

The allocator metadata was successfully changed (the updated JSON is returned)
*/
type DeleteAllocatorMetadataItemOK struct {
	Payload []*models.MetadataItem
}

func (o *DeleteAllocatorMetadataItemOK) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}/metadata/{key}][%d] deleteAllocatorMetadataItemOK  %+v", 200, o.Payload)
}
func (o *DeleteAllocatorMetadataItemOK) GetPayload() []*models.MetadataItem {
	return o.Payload
}

func (o *DeleteAllocatorMetadataItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllocatorMetadataItemNotFound creates a DeleteAllocatorMetadataItemNotFound with default headers values
func NewDeleteAllocatorMetadataItemNotFound() *DeleteAllocatorMetadataItemNotFound {
	return &DeleteAllocatorMetadataItemNotFound{}
}

/* DeleteAllocatorMetadataItemNotFound describes a response with status code 404, with default header values.

The allocator specified by {allocator_id} cannot be found. (code: `allocators.allocator_not_found`)
*/
type DeleteAllocatorMetadataItemNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteAllocatorMetadataItemNotFound) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}/metadata/{key}][%d] deleteAllocatorMetadataItemNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAllocatorMetadataItemNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteAllocatorMetadataItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAllocatorMetadataItemRetryWith creates a DeleteAllocatorMetadataItemRetryWith with default headers values
func NewDeleteAllocatorMetadataItemRetryWith() *DeleteAllocatorMetadataItemRetryWith {
	return &DeleteAllocatorMetadataItemRetryWith{}
}

/* DeleteAllocatorMetadataItemRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type DeleteAllocatorMetadataItemRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteAllocatorMetadataItemRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}/metadata/{key}][%d] deleteAllocatorMetadataItemRetryWith  %+v", 449, o.Payload)
}
func (o *DeleteAllocatorMetadataItemRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteAllocatorMetadataItemRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
