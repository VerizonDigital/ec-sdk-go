// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package dcv

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// New creates a new dcv API client.
func New(c ecclient.APIClient, baseAPIURL string) ClientService {
	return &Client{c, baseAPIURL}
}

// Client for dcv API
type Client struct {
	client     ecclient.APIClient
	baseAPIURL string
}

// ClientService is the interface for Client methods
type ClientService interface {
	DcvCheckDcvTokens(params DcvCheckDcvTokensParams) error

	DcvGetCertificateDomainDetails(params DcvGetCertificateDomainDetailsParams) (*DcvGetCertificateDomainDetailsOK, error)

	DcvPostEmailResend(params DcvPostEmailResendParams) (*DcvPostEmailResendNoContent, error)

	DcvRegenerateDcvTokens(params DcvRegenerateDcvTokensParams) (*DcvRegenerateDcvTokensOK, error)

	DcvSetCertificateDcvMethod(params DcvSetCertificateDcvMethodParams) (*DcvSetCertificateDcvMethodNoContent, error)
}

// DcvCheckDcvTokens dcv check dcv tokens API
func (a *Client) DcvCheckDcvTokens(params DcvCheckDcvTokensParams) error {

	// Set parameters
	results, err := WriteToRequestDcvCheckDcvTokensParams(params)
	if err != nil {
		return err
	}

	method, err := ecclient.ToHTTPMethod("PUT")
	if err != nil {
		return fmt.Errorf("DcvCheckDcvTokens: %v", err)
	}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:      method,
		Path:        a.baseAPIURL + "/v2.0/dcv/certificates/{id}/check",
		RawBody:     results.Body,
		PathParams:  results.PathParams,
		QueryParams: results.QueryParams,
	})

	if err != nil {
		return fmt.Errorf("DcvCheckDcvTokens: %v", err)
	}

	return nil
}

// DcvGetCertificateDomainDetails dcv get certificate domain details API
func (a *Client) DcvGetCertificateDomainDetails(params DcvGetCertificateDomainDetailsParams) (*DcvGetCertificateDomainDetailsOK, error) {

	// Set parameters
	results, err := WriteToRequestDcvGetCertificateDomainDetailsParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("GET")
	if err != nil {
		return nil, fmt.Errorf("DcvGetCertificateDomainDetails: %v", err)
	}

	parsedResponse := &DcvGetCertificateDomainDetailsOK{}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.baseAPIURL + "/v2.0/dcv/certificates/{id}",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("DcvGetCertificateDomainDetails: %v", err)
	}

	return parsedResponse, nil
}

// DcvPostEmailResend dcv post email resend API
func (a *Client) DcvPostEmailResend(params DcvPostEmailResendParams) (*DcvPostEmailResendNoContent, error) {

	// Set parameters
	results, err := WriteToRequestDcvPostEmailResendParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("POST")
	if err != nil {
		return nil, fmt.Errorf("DcvPostEmailResend: %v", err)
	}

	parsedResponse := &DcvPostEmailResendNoContent{}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.baseAPIURL + "/v2.0/dcv/certificates/{id}/emails/resend",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("DcvPostEmailResend: %v", err)
	}

	return parsedResponse, nil
}

// DcvRegenerateDcvTokens dcv regenerate dcv tokens API
func (a *Client) DcvRegenerateDcvTokens(params DcvRegenerateDcvTokensParams) (*DcvRegenerateDcvTokensOK, error) {

	// Set parameters
	results, err := WriteToRequestDcvRegenerateDcvTokensParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("PUT")
	if err != nil {
		return nil, fmt.Errorf("DcvRegenerateDcvTokens: %v", err)
	}

	parsedResponse := &DcvRegenerateDcvTokensOK{}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.baseAPIURL + "/v2.0/dcv/certificates/{id}/token",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("DcvRegenerateDcvTokens: %v", err)
	}

	return parsedResponse, nil
}

// DcvSetCertificateDcvMethod dcv set certificate dcv method API
func (a *Client) DcvSetCertificateDcvMethod(params DcvSetCertificateDcvMethodParams) (*DcvSetCertificateDcvMethodNoContent, error) {

	// Set parameters
	results, err := WriteToRequestDcvSetCertificateDcvMethodParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("PUT")
	if err != nil {
		return nil, fmt.Errorf("DcvSetCertificateDcvMethod: %v", err)
	}

	parsedResponse := &DcvSetCertificateDcvMethodNoContent{}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.baseAPIURL + "/v2.0/dcv/certificates/{id}/method",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("DcvSetCertificateDcvMethod: %v", err)
	}

	return parsedResponse, nil
}

type RequestParameters struct {
	QueryParams map[string]string
	PathParams  map[string]string
	Body        interface{}
}

func NewRequestParameters() *RequestParameters {
	return &RequestParameters{
		QueryParams: make(map[string]string),
		PathParams:  make(map[string]string),
	}
}