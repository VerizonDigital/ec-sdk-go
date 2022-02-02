// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package profiles_rl

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NewProfilesRlDeleteCustomerSettingsByIDParams creates a new ProfilesRlDeleteCustomerSettingsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProfilesRlDeleteCustomerSettingsByIDParams() *ProfilesRlDeleteCustomerSettingsByIDParams {
	return &ProfilesRlDeleteCustomerSettingsByIDParams{}
}

/* ProfilesRlDeleteCustomerSettingsByIDParams contains all the parameters to send to the API endpoint
   for the profiles rl delete customer settings by Id operation.

   Typically these are written to a http.Request.
*/
type ProfilesRlDeleteCustomerSettingsByIDParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the profiles rl delete customer settings by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProfilesRlDeleteCustomerSettingsByIDParams) WithDefaults() *ProfilesRlDeleteCustomerSettingsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the profiles rl delete customer settings by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProfilesRlDeleteCustomerSettingsByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WriteToRequest extracts parameters and sets for the request to be consumed
func WriteToRequestProfilesRlDeleteCustomerSettingsByIDParams(o *ProfilesRlDeleteCustomerSettingsByIDParams) (RequestParameters, error) {

	var res []error

	params := NewRequestParameters()

	// path param id
	params.PathParams["id"] = swag.FormatInt32(o.ID)

	if len(res) > 0 {
		return *params, errors.CompositeValidationError(res...)
	}
	return *params, nil
}
