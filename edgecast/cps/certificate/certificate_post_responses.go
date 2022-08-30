// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package certificate

import "github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

// NewCertificatePostCreated creates a CertificatePostCreated with default headers values
func NewCertificatePostCreated() *CertificatePostCreated {
	return &CertificatePostCreated{}
}

/*
	CertificatePostCreated describes a response with status code 201, with default header values.

Success
*/
type CertificatePostCreated struct {
	models.CertificateRequestBase
}

// NewCertificatePostBadRequest creates a CertificatePostBadRequest with default headers values
func NewCertificatePostBadRequest() *CertificatePostBadRequest {
	return &CertificatePostBadRequest{}
}

/*
	CertificatePostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CertificatePostBadRequest struct {
	models.HyperionErrorReponse
}
