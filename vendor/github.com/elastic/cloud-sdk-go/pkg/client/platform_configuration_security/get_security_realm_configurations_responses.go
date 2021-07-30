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

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetSecurityRealmConfigurationsReader is a Reader for the GetSecurityRealmConfigurations structure.
type GetSecurityRealmConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityRealmConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecurityRealmConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSecurityRealmConfigurationsOK creates a GetSecurityRealmConfigurationsOK with default headers values
func NewGetSecurityRealmConfigurationsOK() *GetSecurityRealmConfigurationsOK {
	return &GetSecurityRealmConfigurationsOK{}
}

/* GetSecurityRealmConfigurationsOK describes a response with status code 200, with default header values.

The security realm configurations were successfully returned
*/
type GetSecurityRealmConfigurationsOK struct {
	Payload *models.SecurityRealmInfoList
}

func (o *GetSecurityRealmConfigurationsOK) Error() string {
	return fmt.Sprintf("[GET /platform/configuration/security/realms][%d] getSecurityRealmConfigurationsOK  %+v", 200, o.Payload)
}
func (o *GetSecurityRealmConfigurationsOK) GetPayload() *models.SecurityRealmInfoList {
	return o.Payload
}

func (o *GetSecurityRealmConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityRealmInfoList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
