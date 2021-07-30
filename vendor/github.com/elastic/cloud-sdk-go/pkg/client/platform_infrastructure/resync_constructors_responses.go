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

// ResyncConstructorsReader is a Reader for the ResyncConstructors structure.
type ResyncConstructorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResyncConstructorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewResyncConstructorsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 449:
		result := NewResyncConstructorsRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResyncConstructorsAccepted creates a ResyncConstructorsAccepted with default headers values
func NewResyncConstructorsAccepted() *ResyncConstructorsAccepted {
	return &ResyncConstructorsAccepted{}
}

/* ResyncConstructorsAccepted describes a response with status code 202, with default header values.

The ids of documents, organized by model version, that will be synchronized.
*/
type ResyncConstructorsAccepted struct {
	Payload *models.ModelVersionIndexSynchronizationResults
}

func (o *ResyncConstructorsAccepted) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/constructors/_resync][%d] resyncConstructorsAccepted  %+v", 202, o.Payload)
}
func (o *ResyncConstructorsAccepted) GetPayload() *models.ModelVersionIndexSynchronizationResults {
	return o.Payload
}

func (o *ResyncConstructorsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelVersionIndexSynchronizationResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncConstructorsRetryWith creates a ResyncConstructorsRetryWith with default headers values
func NewResyncConstructorsRetryWith() *ResyncConstructorsRetryWith {
	return &ResyncConstructorsRetryWith{}
}

/* ResyncConstructorsRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ResyncConstructorsRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncConstructorsRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/constructors/_resync][%d] resyncConstructorsRetryWith  %+v", 449, o.Payload)
}
func (o *ResyncConstructorsRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResyncConstructorsRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
