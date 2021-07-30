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

package deployments_traffic_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetTrafficFilterRulesetsReader is a Reader for the GetTrafficFilterRulesets structure.
type GetTrafficFilterRulesetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTrafficFilterRulesetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTrafficFilterRulesetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetTrafficFilterRulesetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTrafficFilterRulesetsOK creates a GetTrafficFilterRulesetsOK with default headers values
func NewGetTrafficFilterRulesetsOK() *GetTrafficFilterRulesetsOK {
	return &GetTrafficFilterRulesetsOK{}
}

/* GetTrafficFilterRulesetsOK describes a response with status code 200, with default header values.

The collection of traffic filter routes
*/
type GetTrafficFilterRulesetsOK struct {
	Payload *models.TrafficFilterRulesets
}

func (o *GetTrafficFilterRulesetsOK) Error() string {
	return fmt.Sprintf("[GET /deployments/traffic-filter/rulesets][%d] getTrafficFilterRulesetsOK  %+v", 200, o.Payload)
}
func (o *GetTrafficFilterRulesetsOK) GetPayload() *models.TrafficFilterRulesets {
	return o.Payload
}

func (o *GetTrafficFilterRulesetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrafficFilterRulesets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTrafficFilterRulesetsInternalServerError creates a GetTrafficFilterRulesetsInternalServerError with default headers values
func NewGetTrafficFilterRulesetsInternalServerError() *GetTrafficFilterRulesetsInternalServerError {
	return &GetTrafficFilterRulesetsInternalServerError{}
}

/* GetTrafficFilterRulesetsInternalServerError describes a response with status code 500, with default header values.

Error fetching traffic filter rulesets. (code: `traffic_filter.request_execution_failed`)
*/
type GetTrafficFilterRulesetsInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetTrafficFilterRulesetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/traffic-filter/rulesets][%d] getTrafficFilterRulesetsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetTrafficFilterRulesetsInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetTrafficFilterRulesetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
