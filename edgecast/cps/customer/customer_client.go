// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package customer

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/internal/ecclient"
)

// New creates a new customer API client.
func New(c ecclient.APIClient, baseAPIURL string) ClientService {
	return &Client{c, baseAPIURL}
}

// Client for customer API
type Client struct {
	client     ecclient.APIClient
	baseAPIURL string
}

// ClientService is the interface for Client methods
type ClientService interface {
	CustomerGetCustomerCommits(params CustomerGetCustomerCommitsParams) (*CustomerGetCustomerCommitsOK, error)

	CustomerGetCustomerNotifications(params CustomerGetCustomerNotificationsParams) (*CustomerGetCustomerNotificationsOK, error)

	CustomerUpdateCustomerNotifications(params CustomerUpdateCustomerNotificationsParams) (*CustomerUpdateCustomerNotificationsOK, error)
}

// CustomerGetCustomerCommits customer get customer commits API
func (a *Client) CustomerGetCustomerCommits(params CustomerGetCustomerCommitsParams) (*CustomerGetCustomerCommitsOK, error) {

	// Set parameters
	results, err := WriteToRequestCustomerGetCustomerCommitsParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("GET")
	if err != nil {
		return nil, fmt.Errorf("CustomerGetCustomerCommits: %v", err)
	}

	parsedResponse := &CustomerGetCustomerCommitsOK{}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.baseAPIURL + "/v2.0/customers/commits",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("CustomerGetCustomerCommits: %v", err)
	}

	return parsedResponse, nil
}

// CustomerGetCustomerNotifications customer get customer notifications API
func (a *Client) CustomerGetCustomerNotifications(params CustomerGetCustomerNotificationsParams) (*CustomerGetCustomerNotificationsOK, error) {

	// Set parameters
	results, err := WriteToRequestCustomerGetCustomerNotificationsParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("GET")
	if err != nil {
		return nil, fmt.Errorf("CustomerGetCustomerNotifications: %v", err)
	}

	parsedResponse := &CustomerGetCustomerNotificationsOK{}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.baseAPIURL + "/v2.0/customers/notifications",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("CustomerGetCustomerNotifications: %v", err)
	}

	return parsedResponse, nil
}

// CustomerUpdateCustomerNotifications customer update customer notifications API
func (a *Client) CustomerUpdateCustomerNotifications(params CustomerUpdateCustomerNotificationsParams) (*CustomerUpdateCustomerNotificationsOK, error) {

	// Set parameters
	results, err := WriteToRequestCustomerUpdateCustomerNotificationsParams(params)
	if err != nil {
		return nil, err
	}

	method, err := ecclient.ToHTTPMethod("PATCH")
	if err != nil {
		return nil, fmt.Errorf("CustomerUpdateCustomerNotifications: %v", err)
	}

	parsedResponse := &CustomerUpdateCustomerNotificationsOK{}

	_, err = a.client.SubmitRequest(ecclient.SubmitRequestParams{
		Method:         method,
		Path:           a.baseAPIURL + "/v2.0/customers/notifications",
		RawBody:        results.Body,
		PathParams:     results.PathParams,
		QueryParams:    results.QueryParams,
		ParsedResponse: parsedResponse,
	})

	if err != nil {
		return nil, fmt.Errorf("CustomerUpdateCustomerNotifications: %v", err)
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