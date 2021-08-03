// Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license. See LICENSE file in project root for terms.

package waf

import (
	"fmt"
)

/*
	This file contains operations and types specific to WAF Custom Rule Sets.
*/

// DeleteCustomRuleResponse contains the response from the WAF API when deleting a new custom rule
type DeleteCustomRuleResponse struct {

	// ID indicates the generated ID for the newly deleted Rule
	ID string

	WAFResponse
}

// AddCustomRuleResponse contains the response from the WAF API when adding a new custom rule
type AddCustomRuleResponse struct {
	AddRuleResponse
}

// A custom rule set defines custom threat assessment criteria.
type CustomRule struct {

	// Contains custom rules. Each directive object defines a custom rule via the sec_rule object.
	// You may create up to 10 custom rules.
	Directives []Directive `json:"directive"`

	// Indicates the name of the custom rule.
	Name string `json:"name,omitempty"`
}

//Contains custom rules. Each directive object defines a custom rule via the sec_rule object.
type Directive struct {

	//Defines a custom rule
	SecRule SecRule `json:"sec_rule"`
}

// Defines a custom rule
type SecRule struct {

	// Determines whether the string identified in a variable object will be transformed and the metadata that will be assigned to malicious traffic.
	Action Action `json:"action"`

	// Contains additional criteria that must be satisfied to identify a malicious request.
	ChainedRules []ChainedRule `json:"chained_rule,omitempty"`

	// Indicates the name assigned to this custom rule.
	Name string `json:"name,omitempty"`

	// Indicates the comparison that will be performed against the request element(s) identified within a variable object.
	Operator Operator `json:"operator"`

	// Contains criteria that identifies a request element.
	Variables []Variable `json:"variable"`
}

// The action object determines whether the value derived from the request element identified in a variable object will be transformed and the metadata that will be used to identify malicious traffic.
type Action struct {

	/*
		Determines the custom ID that will be assigned to this custom rule. This custom ID is exposed via the Threats Dashboard.

		Valid values fall within this range: 66000000 - 66999999
			Note: This field is only applicable for the action object that resides in the root of the sec_rule object.

		Default Value: Random number
	*/
	ID string `json:"id,omitempty"`

	/*
		Determines the rule message that will be assigned to this custom rule. This message is exposed via the Threats Dashboard.
			Note: This field is only applicable for the action object that resides in the root of the sec_rule object.

		Default Value: Blank
	*/
	MSG string `json:"msg,omitempty"`

	/*
		Determines the set of transformations that will be applied to the value derived from the request element identified in a variable object (i.e., source value).
		Transformations are always applied to the source value, regardless of the number of transformations that have been defined.

		Valid values are:

			NONE: Indicates that the source value should not be modified.
			LOWERCASE: Indicates that the source value should be converted to lowercase characters.
			URLDECODE: Indicates that the source value should be URL decoded. This transformation is useful when the source value has been URL encoded twice.
			REMOVENULLS: Indicates that null values should be removed from the source value.

		Note: A criterion is satisfied if the source value or any of the modified string values meet the conditions defined by the operator object.
	*/
	T []string `json:"t,omitempty"`
}

// Each object within the chained_rule array describes an additional set of criteria that must be satisfied in order to identify a malicious request.
type ChainedRule struct {

	// Determines whether the string value derived from the request element identified in a variable object will be transformed and the metadata that will be used to identify malicious traffic.
	Action Action `json:"action"`

	// Indicates the comparison that will be performed on the string value(s) derived from the request element(s) defined within the variable array.
	Operator Operator `json:"operator"`

	// Identifies each request element for which a comparison will be made.
	Variables []Variable `json:"variable"`
}

// The variable array identifies each request element for which a comparison will be made using its properties
type Variable struct {

	/*
		Determines the request element that will be assessed.
		Valid values are:
		ARGS_POST | GEO | QUERY_STRING | REMOTE_ADDR | REQUEST_BODY | REQUEST_COOKIES | REQUEST_HEADERS | REQUEST_METHOD | REQUEST_URI

		Note: If a request element consists of one or more key-value pairs, then you may identify a key via a match object. If is_count has been disabled, then you may identify a specific value via the operator object.
	*/
	Type string `json:"type"`

	//Contains comparison settings for the request element identified by the type property.
	Matches []Match `json:"match,omitempty"`

	/*
		Determines whether a comparison will be performed between the operator object and a string value or the number of matches found.
		Valid values are:

		true: A counter will increment whenever the request element defined by this variable object is found. The operator object will perform a comparison against this number.
			Note: If you enable is_count, then you must also set the type property to EQ.

		false: The operator object will perform a comparison against the string value derived from the request element defined by this variable object.

	*/
	IsCount bool `json:"is_count,omitempty"`
}

//The operator object describes the comparison that will be performed on the request element(s) defined within a variable object using its properties:
type Operator struct {

	/*
		Indicates whether a condition will be satisfied when the value derived from the request element defined within a variable object matches or does not match the value property.

		Valid values are:

			True: Does not match
			False: Matches
	*/
	IsNegated bool `json:"is_negated,omitempty"`

	/*
		Indicates how the system will interpret the comparison between the value property and the value derived from the request element defined within a variable object.
		Valid values are:

			RX: Indicates that the string value derived from the request element must satisfy the regular expression defined in the value property.
			STREQ: Indicates that the string value derived from the request element must be an exact match to the value property.
			CONTAINS: Indicates that the value property must contain the string value derived from the request element.
			BEGINSWITH: Indicates that the value property must start with the string value derived from the request element.
			ENDSWITH: Indicates that the value property must end with the string value derived from the request element.
			EQ: Indicates that the number derived from the variable object must be an exact match to the value property.
				Note: You should only use EQ when the is_count property has been enabled.
			IPMATCH: Requires that the request's IP address either be contained by an IP block or be an exact match to an IP address defined in the values property. Only use IPMATCH with the REMOTE_ADDR variable.
	*/
	Type string `json:"type"`

	/*
		Indicates a value that will be compared against the string or number value derived from the request element defined within a variable object.

		Note: If you are identifying traffic via a URL path (REQUEST_URI), then you should specify a URL path pattern that starts directly after the hostname. Exclude a protocol or a hostname when defining this property.
				Sample values:
				/marketing
				/800001/mycustomerorigin
	*/
	Value string `json:"value,omitempty"`
}

// The match array determines the comparison conditions for the request element identified by the type property.
type Match struct {

	/*
		Determines whether this condition is satisfied when the request element identified by the variable object is found or not found.

			True: Not found
			False: Found
	*/
	IsNegated bool `json:"is_negated,omitempty"`

	/*
		Determines whether the value property will be interpreted as a regular expression. Valid values are:

			True: Regular expression
			False: Default value. Literal value.
	*/
	IsRegex bool `json:"is_regex,omitempty"`

	/*
		Restricts the match condition defined by the type property to the specified value.

		Example:

			If the type property is set to REQUEST_HEADERS and this property is set to User-Agent, then this match condition is restricted to the User-Agent request header. If the value property is omitted, then this match condition applies to all request headers.
	*/
	Value string `json:"value,omitempty"`
}

type CustomRuleSet struct {
	//Indicates the system-defined ID for the custom rule set.
	Id string `json:"id"`

	//Indicates the date and time at which the custom rule was last modified.
	//Syntax:
	//MM/DD/YYYYhh:mm:ss [AM|PM]
	LastModifiedDate string `json:"last_modified_date"`

	//Indicates the name of the custom rule set.
	Name string `json:"name"`
}

// Creates a custom rule set that defines custom threat assessment criteria.
func (svc *WAFService) AddCustomRuleSet(customRuleSet CustomRule, accountNumber string) (*AddCustomRuleResponse, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/rules", accountNumber)

	request, err := svc.Client.BuildRequest("POST", url, customRuleSet)

	if err != nil {
		return nil, fmt.Errorf("AddCustomRuleSet: %v", err)
	}

	parsedResponse := &AddCustomRuleResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("AddCustomRuleSet: %v", err)
	}

	return parsedResponse, nil
}

// Retrieves a list of custom rule sets. A custom rule set allows you to define custom threat assessment criterion.
func (svc *WAFService) GetAllCustomRuleSets(accountNumber string) ([]CustomRuleSet, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/rules", accountNumber)

	request, err := svc.Client.BuildRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("GetAllCustomRuleSets: %v", err)
	}

	var customRuleSets = &[]CustomRuleSet{}

	_, err = svc.Client.SendRequest(request, &customRuleSets)

	if err != nil {
		return nil, fmt.Errorf("GetAllCustomRuleSets: %v", err)
	}

	return *customRuleSets, nil
}

// Deletes a custom rule.
func (svc *WAFService) DeleteCustomRuleSet(accountNumber string, customRuleID string) (*DeleteCustomRuleResponse, error) {
	url := fmt.Sprintf("/v2/mcc/customers/%s/waf/v1.0/rules/%v", accountNumber, customRuleID)

	request, err := svc.Client.BuildRequest("DELETE", url, nil)

	if err != nil {
		return nil, fmt.Errorf("DeleteCustomRuleSet: %v", err)
	}

	parsedResponse := &DeleteCustomRuleResponse{}

	_, err = svc.Client.SendRequest(request, &parsedResponse)

	if err != nil {
		return nil, fmt.Errorf("DeleteCustomRuleSet: %v", err)
	}

	return parsedResponse, nil
}
