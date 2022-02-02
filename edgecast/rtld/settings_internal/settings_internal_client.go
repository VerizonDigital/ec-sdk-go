// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package settings_internal

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
)

// New creates a new settings internal API client.
func New(c client.Client) ClientService {
	return &Client{Client: c}
}

/*
Client for settings internal API
*/
type Client struct {
	client.Client
}

// ClientService is the interface for Client methods
type ClientService interface {
	SettingsGetRlSettings(params *SettingsGetRlSettingsParams) (*SettingsGetRlSettingsOK, error)

	SettingsGetSettingsByPlatform(params *SettingsGetSettingsByPlatformParams) (*SettingsGetSettingsByPlatformOK, error)

	SettingsGetWafSettings(params *SettingsGetWafSettingsParams) (*SettingsGetWafSettingsOK, error)
}

/*
  SettingsGetRlSettings settings get rl settings API
*/
func (a *Client) SettingsGetRlSettings(params *SettingsGetRlSettingsParams) (*SettingsGetRlSettingsOK, error) {
	// Validate the params before sending
	if params == nil {
		params = NewSettingsGetRlSettingsParams()
	}

	//Set parameters
	results, err := WriteToRequestSettingsGetRlSettingsParams(params)
	if err != nil {
		return nil, err
	}

	url := a.Config.BaseAPIURL.String() + "/v1.0/rl/settings"

	request, err := a.Client.PrepareRequest("GET", url, results.Body, results.QueryParams, results.PathParams)
	if err != nil {
		return nil, fmt.Errorf("SettingsGetRlSettings: %v", err)
	}

	parsedResponse := &SettingsGetRlSettingsOK{}

	_, err = a.Client.SendRequest(request, parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("SettingsGetRlSettings: %v", err)
	}

	return parsedResponse, nil
}

/*
  SettingsGetSettingsByPlatform settings get settings by platform API
*/
func (a *Client) SettingsGetSettingsByPlatform(params *SettingsGetSettingsByPlatformParams) (*SettingsGetSettingsByPlatformOK, error) {
	// Validate the params before sending
	if params == nil {
		params = NewSettingsGetSettingsByPlatformParams()
	}

	//Set parameters
	results, err := WriteToRequestSettingsGetSettingsByPlatformParams(params)
	if err != nil {
		return nil, err
	}

	url := a.Config.BaseAPIURL.String() + "/v1.0/platforms/{platformId}/settings"

	request, err := a.Client.PrepareRequest("GET", url, results.Body, results.QueryParams, results.PathParams)
	if err != nil {
		return nil, fmt.Errorf("SettingsGetSettingsByPlatform: %v", err)
	}

	parsedResponse := &SettingsGetSettingsByPlatformOK{}

	_, err = a.Client.SendRequest(request, parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("SettingsGetSettingsByPlatform: %v", err)
	}

	return parsedResponse, nil
}

/*
  SettingsGetWafSettings settings get waf settings API
*/
func (a *Client) SettingsGetWafSettings(params *SettingsGetWafSettingsParams) (*SettingsGetWafSettingsOK, error) {
	// Validate the params before sending
	if params == nil {
		params = NewSettingsGetWafSettingsParams()
	}

	//Set parameters
	results, err := WriteToRequestSettingsGetWafSettingsParams(params)
	if err != nil {
		return nil, err
	}

	url := a.Config.BaseAPIURL.String() + "/v1.0/waf/settings"

	request, err := a.Client.PrepareRequest("GET", url, results.Body, results.QueryParams, results.PathParams)
	if err != nil {
		return nil, fmt.Errorf("SettingsGetWafSettings: %v", err)
	}

	parsedResponse := &SettingsGetWafSettingsOK{}

	_, err = a.Client.SendRequest(request, parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("SettingsGetWafSettings: %v", err)
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
