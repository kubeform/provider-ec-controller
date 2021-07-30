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

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewUpdateCommentParams creates a new UpdateCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateCommentParams() *UpdateCommentParams {
	return &UpdateCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCommentParamsWithTimeout creates a new UpdateCommentParams object
// with the ability to set a timeout on a request.
func NewUpdateCommentParamsWithTimeout(timeout time.Duration) *UpdateCommentParams {
	return &UpdateCommentParams{
		timeout: timeout,
	}
}

// NewUpdateCommentParamsWithContext creates a new UpdateCommentParams object
// with the ability to set a context for a request.
func NewUpdateCommentParamsWithContext(ctx context.Context) *UpdateCommentParams {
	return &UpdateCommentParams{
		Context: ctx,
	}
}

// NewUpdateCommentParamsWithHTTPClient creates a new UpdateCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateCommentParamsWithHTTPClient(client *http.Client) *UpdateCommentParams {
	return &UpdateCommentParams{
		HTTPClient: client,
	}
}

/* UpdateCommentParams contains all the parameters to send to the API endpoint
   for the update comment operation.

   Typically these are written to a http.Request.
*/
type UpdateCommentParams struct {

	/* Body.

	   The Comment update data.
	*/
	Body *models.CommentUpdateRequest

	/* CommentID.

	   Id of a Comment
	*/
	CommentID string

	/* ResourceID.

	   Id of the Resource that a Comment belongs to.
	*/
	ResourceID string

	/* ResourceType.

	   The kind of Resource that a Comment belongs to. Should be one of [elasticsearch, kibana, apm, appsearch, enterprise_search, allocator, constructor, runner, proxy].
	*/
	ResourceType string

	/* Version.

	   If specified then checks for conflicts against the version stored in the persistent store (returned in 'x-cloud-resource-version' of the GET request)
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCommentParams) WithDefaults() *UpdateCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update comment params
func (o *UpdateCommentParams) WithTimeout(timeout time.Duration) *UpdateCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update comment params
func (o *UpdateCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update comment params
func (o *UpdateCommentParams) WithContext(ctx context.Context) *UpdateCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update comment params
func (o *UpdateCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update comment params
func (o *UpdateCommentParams) WithHTTPClient(client *http.Client) *UpdateCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update comment params
func (o *UpdateCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update comment params
func (o *UpdateCommentParams) WithBody(body *models.CommentUpdateRequest) *UpdateCommentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update comment params
func (o *UpdateCommentParams) SetBody(body *models.CommentUpdateRequest) {
	o.Body = body
}

// WithCommentID adds the commentID to the update comment params
func (o *UpdateCommentParams) WithCommentID(commentID string) *UpdateCommentParams {
	o.SetCommentID(commentID)
	return o
}

// SetCommentID adds the commentId to the update comment params
func (o *UpdateCommentParams) SetCommentID(commentID string) {
	o.CommentID = commentID
}

// WithResourceID adds the resourceID to the update comment params
func (o *UpdateCommentParams) WithResourceID(resourceID string) *UpdateCommentParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the update comment params
func (o *UpdateCommentParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the update comment params
func (o *UpdateCommentParams) WithResourceType(resourceType string) *UpdateCommentParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the update comment params
func (o *UpdateCommentParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WithVersion adds the version to the update comment params
func (o *UpdateCommentParams) WithVersion(version *string) *UpdateCommentParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update comment params
func (o *UpdateCommentParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param comment_id
	if err := r.SetPathParam("comment_id", o.CommentID); err != nil {
		return err
	}

	// path param resource_id
	if err := r.SetPathParam("resource_id", o.ResourceID); err != nil {
		return err
	}

	// path param resource_type
	if err := r.SetPathParam("resource_type", o.ResourceType); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}