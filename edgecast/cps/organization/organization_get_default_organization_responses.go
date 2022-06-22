// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package organization

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"
)

// NewOrganizationGetDefaultOrganizationOK creates a OrganizationGetDefaultOrganizationOK with default headers values
func NewOrganizationGetDefaultOrganizationOK() *OrganizationGetDefaultOrganizationOK {
	return &OrganizationGetDefaultOrganizationOK{}
}

/* OrganizationGetDefaultOrganizationOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationGetDefaultOrganizationOK struct {
	models.OrganizationDetail
}
