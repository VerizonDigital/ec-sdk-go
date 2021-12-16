// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package customer

import (
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/auth"
	"github.com/EdgeCast/ec-sdk-go/edgecast/client"
	"github.com/EdgeCast/ec-sdk-go/edgecast/logging"
)

// Customer service interacts with the EdgeCast API for Customer
type CustomerService struct {
	client.Client
	Logger logging.Logger
}

// New creates a new Customer service
func New(config edgecast.SDKConfig) (*CustomerService, error) {

	authProvider, err := auth.NewTokenAuthorizationProvider(config.APIToken)

	if err != nil {
		return nil, fmt.Errorf("customer.New(): %v", err)
	}

	c := client.NewClient(client.ClientConfig{
		AuthProvider: authProvider,
		BaseAPIURL:   config.BaseAPIURLLegacy,
		UserAgent:    config.UserAgent,
		Logger:       config.Logger,
	})

	return &CustomerService{
		Client: c,
		Logger: config.Logger,
	}, nil
}