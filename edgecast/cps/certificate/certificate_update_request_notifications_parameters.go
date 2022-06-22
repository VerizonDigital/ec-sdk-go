// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package certificate

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"

	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"
)

// NewCertificateUpdateRequestNotificationsParams creates a new CertificateUpdateRequestNotificationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCertificateUpdateRequestNotificationsParams() CertificateUpdateRequestNotificationsParams {
	return CertificateUpdateRequestNotificationsParams{}
}

// CertificateUpdateRequestNotificationsParams contains all the parameters to send to the API
// endpoint for the certificate update request notifications operation. Typically these are written
// to a http.Request.
type CertificateUpdateRequestNotificationsParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	// Notifications.
	Notifications []*models.EmailNotification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the certificate update request notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CertificateUpdateRequestNotificationsParams) WithDefaults() *CertificateUpdateRequestNotificationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the certificate update request notifications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CertificateUpdateRequestNotificationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WriteToRequest extracts parameters and sets for the request to be consumed
func WriteToRequestCertificateUpdateRequestNotificationsParams(o CertificateUpdateRequestNotificationsParams) (RequestParameters, error) {

	var res []error

	params := NewRequestParameters()

	// path param id
	params.PathParams["id"] = swag.FormatInt64(o.ID)
	if o.Notifications != nil {
		params.Body = o.Notifications
	}

	if len(res) > 0 {
		return *params, errors.CompositeValidationError(res...)
	}
	return *params, nil
}
