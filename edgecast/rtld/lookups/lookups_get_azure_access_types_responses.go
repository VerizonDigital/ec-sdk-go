// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package lookups

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtldmodels"
)

// NewLookupsGetAzureAccessTypesOK creates a LookupsGetAzureAccessTypesOK with default headers values
func NewLookupsGetAzureAccessTypesOK() *LookupsGetAzureAccessTypesOK {
	return &LookupsGetAzureAccessTypesOK{}
}

/* LookupsGetAzureAccessTypesOK describes a response with status code 200, with default header values.

Success
*/
type LookupsGetAzureAccessTypesOK struct {
	*rtldmodels.HyperionCollectionCodeName
}
