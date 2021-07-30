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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetEsClusterCcsReader is a Reader for the GetEsClusterCcs structure.
type GetEsClusterCcsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEsClusterCcsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEsClusterCcsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEsClusterCcsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEsClusterCcsOK creates a GetEsClusterCcsOK with default headers values
func NewGetEsClusterCcsOK() *GetEsClusterCcsOK {
	return &GetEsClusterCcsOK{}
}

/* GetEsClusterCcsOK describes a response with status code 200, with default header values.

List of cross-cluster search clusters' IDs for the remote cluster
*/
type GetEsClusterCcsOK struct {
	Payload *models.CrossClusterSearchClusters
}

func (o *GetEsClusterCcsOK) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/{cluster_id}/ccs][%d] getEsClusterCcsOK  %+v", 200, o.Payload)
}
func (o *GetEsClusterCcsOK) GetPayload() *models.CrossClusterSearchClusters {
	return o.Payload
}

func (o *GetEsClusterCcsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CrossClusterSearchClusters)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEsClusterCcsNotFound creates a GetEsClusterCcsNotFound with default headers values
func NewGetEsClusterCcsNotFound() *GetEsClusterCcsNotFound {
	return &GetEsClusterCcsNotFound{}
}

/* GetEsClusterCcsNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type GetEsClusterCcsNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *GetEsClusterCcsNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/{cluster_id}/ccs][%d] getEsClusterCcsNotFound  %+v", 404, o.Payload)
}
func (o *GetEsClusterCcsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetEsClusterCcsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
