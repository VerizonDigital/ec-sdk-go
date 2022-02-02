// Copyright 2021 Edgecast Inc., Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package rtld

// This file was generated by codegen-sdk-go.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/url"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"

	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/lookups"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_cdn"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_rl"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/settings_internal"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/rtld"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

//New creates a new real time log delivery API client
func New(config edgecast.SDKConfig) (*RealTimeLogDeliveryAPI, error) {

	apiURL, err := url.Parse(config.BaseAPIURL.String() + DefaultBasePath)
	if err != nil {
		return nil, fmt.Errorf("RealTimeLogDeliveryAPI.New(): %v", err)
	}

	c := client.NewClient(client.ClientConfig{
		BaseAPIURL: *apiURL,
		UserAgent:  config.UserAgent,
		Logger:     config.Logger,
	})

	// OAuth2 authentication
	authProvider, err := auth.NewIDSAuthorizationProvider(config.BaseIDSURL, config.IDSCredentials)
	if err != nil {

		//Token authentication
		authTokenProvider, err := auth.NewTokenAuthorizationProvider(config.APIToken)
		if err != nil {
			return nil, fmt.Errorf("RealTimeLogDeliveryAPI.New(): %v", err)
		}
		c.Config.AuthProvider = authTokenProvider
	}
	c.Config.AuthProvider = authProvider

	return &RealTimeLogDeliveryAPI{
		Client:           c,
		Logger:           config.Logger,
		Lookups:          lookups.New(c),
		ProfilesCdn:      profiles_cdn.New(c),
		ProfilesRl:       profiles_rl.New(c),
		ProfilesWaf:      profiles_waf.New(c),
		SettingsInternal: settings_internal.New(c),
	}, nil

}

// RealTimeLogDeliveryAPI is a client for real time log delivery API
type RealTimeLogDeliveryAPI struct {
	Lookups lookups.ClientService

	ProfilesCdn profiles_cdn.ClientService

	ProfilesRl profiles_rl.ClientService

	ProfilesWaf profiles_waf.ClientService

	SettingsInternal settings_internal.ClientService

	client.Client

	Logger logging.Logger
}
