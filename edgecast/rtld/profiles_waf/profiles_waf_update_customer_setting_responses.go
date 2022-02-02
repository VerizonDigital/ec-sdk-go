// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package profiles_waf

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtldmodels"
)

// NewProfilesWafUpdateCustomerSettingOK creates a ProfilesWafUpdateCustomerSettingOK with default headers values
func NewProfilesWafUpdateCustomerSettingOK() *ProfilesWafUpdateCustomerSettingOK {
	return &ProfilesWafUpdateCustomerSettingOK{}
}

/* ProfilesWafUpdateCustomerSettingOK describes a response with status code 200, with default header values.

Success
*/
type ProfilesWafUpdateCustomerSettingOK struct {
	rtldmodels.WafProfileDto
}
