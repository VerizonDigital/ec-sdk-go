// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license.
// See LICENSE file in project root for terms.

package customer

import (
	"fmt"
)

type Customer struct {
	AccountID                 string `json:"AccountId,omitempty"` // TODO: This might be completely unused. Verify
	Address1                  string
	Address2                  string
	BandwidthUsageLimit       int64
	BillingAccountTag         string
	BillingAddress1           string
	BillingAddress2           string
	BillingCity               string
	BillingContactEmail       string
	BillingContactFax         string
	BillingContactFirstName   string
	BillingContactLastName    string
	BillingContactMobile      string
	BillingContactPhone       string
	BillingContactTitle       string
	BillingCountry            string
	BillingRateInfo           string
	BillingState              string
	BillingZip                string
	City                      string
	CompanyName               string
	ContactEmail              string
	ContactFax                string
	ContactFirstName          string
	ContactLastName           string
	ContactMobile             string
	ContactPhone              string
	ContactTitle              string
	Country                   string
	DataTransferredUsageLimit int64
	Notes                     string
	PartnerUserID             int // Required when not providing a PCC token
	ServiceLevelCode          string
	State                     string
	Website                   string
	Zip                       string
	Status                    int
}

// AddCustomer -
func (svc *CustomerService) AddCustomer(customer *Customer) (string, error) {
	relURL := "v2/pcc/customers"
	if customer.PartnerUserID != 0 {
		relURL = relURL + fmt.Sprintf("?partneruserid=%d", customer.PartnerUserID)
	}

	request, err := svc.Client.BuildRequest("POST", relURL, customer)

	if err != nil {
		return "", fmt.Errorf("AddCustomer: %v", err)
	}

	parsedResponse := &struct {
		AccountNumber string
	}{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return "", fmt.Errorf("AddCustomer: %v", err)
	}

	return parsedResponse.AccountNumber, nil
}

// GetCustomerResponse -
type GetCustomerResponse struct {
	Customer
	ID                   int32  `json:"Id,omitempty"`
	CustomID             string `json:"CustomId,omitempty"`
	HexID                string
	UsageLimitUpdateDate string
	PartnerID            int `json:"PartnerId,omitempty"`
	PartnerName          string
	WholesaleID          int `json:"WholesaleId,omitempty"`
	WholesaleName        string
}

// GetCustomer retrieves a Customer's info using the Hex Account Number
func (svc *CustomerService) GetCustomer(accountNumber string) (*GetCustomerResponse, error) {
	relURL := fmt.Sprintf("v2/pcc/customers/%s", accountNumber)
	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetCustomer: %v", err)
	}

	parsedResponse := &GetCustomerResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("GetCustomer: %v", err)
	}

	return parsedResponse, nil
}

// AccessModule represents a module that a customer has access to
type AccessModule struct {
	ID       int
	Name     string
	ParentID *int
}

// GetCustomerAccessModules retrieves a Customer's Access Module info using the Hex Account Number
func (svc *CustomerService) GetCustomerAccessModules(accountNumber string) (*[]AccessModule, error) {
	relURL := fmt.Sprintf("v2/pcc/customers/%s/accessmodules", accountNumber)
	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerAccessModules: %v", err)
	}

	var accessModules []AccessModule

	_, err = svc.Client.SendRequest(request, &accessModules)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerAccessModules: %v", err)
	}

	return &accessModules, nil
}

// Service -
type Service struct {
	ID       int `json:"Id,omitempty"`
	Name     string
	ParentID int `json:"parentId,omitempty"`
	Status   int8
}

// GetAvailableCustomerServices gets all service information available for a partner to administor to thier customers
func (svc *CustomerService) GetAvailableCustomerServices() (*[]Service, error) {
	request, err := svc.Client.BuildRequest("GET", "v2/pcc/customers/services", nil)

	if err != nil {
		return nil, fmt.Errorf("GetAvailableCustomerServices: %v", err)
	}

	var services []Service

	_, err = svc.Client.SendRequest(request, &services)

	if err != nil {
		return nil, fmt.Errorf("GetAvailableCustomerServices: %v", err)
	}

	return &services, nil
}

// GetCustomerServices gets the list of services available to a customer and whether each is active for the customer
func (svc *CustomerService) GetCustomerServices(accountNumber string) ([]Service, error) {
	relURL := fmt.Sprintf("v2/pcc/customers/%s/services", accountNumber)
	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerServices: %v", err)
	}

	var services []Service

	_, err = svc.Client.SendRequest(request, &services)

	if err != nil {
		return nil, fmt.Errorf("GetCustomerServices: %v", err)
	}

	return services, nil
}

// UpdateCustomerServices -
func (svc *CustomerService) UpdateCustomerServices(accountNumber string, serviceIDs []int, status int8) error {
	relURL := fmt.Sprintf("v2/pcc/customers/%s/services", accountNumber)

	body := &struct {
		Status      int8
		ServiceCode []int
	}{
		Status:      status,
		ServiceCode: serviceIDs,
	}

	request, err := svc.Client.BuildRequest("PUT", relURL, body)

	if err != nil {
		return fmt.Errorf("UpdateCustomerServices: %v", err)
	}

	resp, err := svc.Client.SendRequest(request, nil)

	if err == nil && resp.StatusCode != 200 {
		return fmt.Errorf("failed to set customer services, please contact an administrator")
	}

	if err != nil {
		return fmt.Errorf("UpdateCustomerServices: %v", err)
	}

	return nil
}

// GetCustomerDeliveryRegion gets the current active delivery region set for the customer
func (svc *CustomerService) GetCustomerDeliveryRegion(accountNumber string) (int, error) {
	relURL := fmt.Sprintf("v2/pcc/customers/%s/deliveryregions", accountNumber)

	request, err := svc.Client.BuildRequest("GET", relURL, nil)

	if err != nil {
		return 0, fmt.Errorf("GetCustomerDeliveryRegion: %v", err)
	}

	parsedResponse := &struct {
		AccountNumber    string
		CustomID         string
		DeliveryRegionID int
	}{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return 0, fmt.Errorf("GetCustomerDeliveryRegion: %v", err)
	}

	return parsedResponse.DeliveryRegionID, nil
}

// UpdateCustomerDomainURL -
func (svc *CustomerService) UpdateCustomerDomainURL(accountNumber string, domainType int, url string, partnerID int) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf("v2/pcc/customers/domains/%d/url?idtype=an&id=%s", domainType, accountNumber)
	relURL := FormatURLAddPartnerID(baseURL, partnerID)

	body := &struct {
		URL string `json:"Url"`
	}{
		URL: url,
	}

	request, err := svc.Client.BuildRequest("PUT", relURL, body)

	if err != nil {
		return fmt.Errorf("UpdateCustomerDomainURL: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("UpdateCustomerDomainURL: %v", err)
	}

	return nil
}

type DeliveryRegion struct {
	AccountNumber    string
	DeliveryRegionID int
	PartnerID        int
}

// UpdateCustomerDeliveryRegion -
func (svc *CustomerService) UpdateCustomerDeliveryRegion(accountNumber string, deliveryRegion DeliveryRegion) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf("v2/pcc/customers/deliveryregions?idtype=an&id=%s", accountNumber)
	relURL := FormatURLAddPartnerID(baseURL, deliveryRegion.PartnerID)

	body := &struct {
		ID int `json:"Id"`
	}{
		ID: deliveryRegion.DeliveryRegionID,
	}

	request, err := svc.Client.BuildRequest("PUT", relURL, body)

	if err != nil {
		return fmt.Errorf("UpdateCustomerDeliveryRegion: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("UpdateCustomerDeliveryRegion: %v", err)
	}

	return nil
}

// DeleteCustomer -
func (svc *CustomerService) DeleteCustomer(accountNumber string, partnerID int) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf("v2/pcc/customers?idtype=an&id=%s", accountNumber)
	relURL := FormatURLAddPartnerID(baseURL, partnerID)

	request, err := svc.Client.BuildRequest("DELETE", relURL, nil)

	if err != nil {
		return fmt.Errorf("DeleteCustomer: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("DeleteCustomer: %v", err)
	}

	return nil
}

// UpdateCustomer -
func (svc *CustomerService) UpdateCustomer(accountNumber string, partnerId int, body *Customer) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf("v2/pcc/customers?idtype=an&id=%s", accountNumber)
	relURL := FormatURLAddPartnerID(baseURL, partnerId)

	request, err := svc.Client.BuildRequest("PUT", relURL, body)

	if err != nil {
		return fmt.Errorf("UpdateCustomer: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("UpdateCustomer: %v", err)
	}

	return nil
}

// UpdateCustomerAccessModule -
func (svc *CustomerService) UpdateCustomerAccessModule(accountNumber string, accessModuleID, partnerID int) error {
	// TODO: support custom ids for accounts
	baseURL := fmt.Sprintf("v2/pcc/customers/accessmodules/%d/status?idtype=an&id=%s", accessModuleID, accountNumber)
	relURL := FormatURLAddPartnerID(baseURL, partnerID)
	body := &struct{ Status int8 }{Status: 1}

	request, err := svc.Client.BuildRequest("PUT", relURL, body)

	if err != nil {
		return fmt.Errorf("UpdateCustomerAccessModule: %v", err)
	}

	_, err = svc.Client.SendRequest(request, nil)

	if err != nil {
		return fmt.Errorf("UpdateCustomerAccessModule: %v", err)
	}

	return nil
}

// FormatURLAddPartnerID is a utility function for adding the optional partner ID query string param
func FormatURLAddPartnerID(originalURL string, partnerID int) string {
	if partnerID != 0 {
		return originalURL + fmt.Sprintf("&partnerid=%d", partnerID)
	}

	return originalURL
}
