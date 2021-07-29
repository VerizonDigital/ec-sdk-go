package main

import (
	"flag"
	"fmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
)

func main() {

	//Setup
	apiToken := flag.String("api-token", "", "API Token provided to you")
	accountNumber := flag.String("account-number", "", "Account number you wish to retrieve all Access Rules for")

	flag.Parse()

	wafConfig := waf.NewConfig(*apiToken)
	wafConfig.Logger = edgecast.NewStandardLogger()
	wafService, err := waf.New(wafConfig)

	if err != nil {
		fmt.Printf("error creating WAF Service: %v\n", err)
		return
	}

	//Get All Access Rules Example
	accessRules, err := wafService.GetAccessRules(*accountNumber)

	if err != nil {
		fmt.Printf("Error retrieving all access rules: %v\n", err)
		return
	}

	for _, rule := range accessRules {
		fmt.Println(rule)
	}
}