// Code generated by the Code Generation Kit (CGK) using OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
//
// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

/*
Customer Origins API v3

List of API of config Customer Origin.

API version: 0.5.0
Contact: portals-streaming@edgecast.com
*/

package originv3

import (
	"encoding/json"
)

// AdnGateway struct for AdnGateway
type AdnGateway struct {
	Code                 *string `json:"code,omitempty"`
	Ordinal              *int32  `json:"ordinal,omitempty"`
	City                 *string `json:"city,omitempty"`
	Region               *string `json:"region,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdnGateway AdnGateway

// NewAdnGateway instantiates a new AdnGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdnGateway() *AdnGateway {
	this := AdnGateway{}
	return &this
}

// NewAdnGatewayWithDefaults instantiates a new AdnGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdnGatewayWithDefaults() *AdnGateway {
	this := AdnGateway{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AdnGateway) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdnGateway) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AdnGateway) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AdnGateway) SetCode(v string) {
	o.Code = &v
}

// GetOrdinal returns the Ordinal field value if set, zero value otherwise.
func (o *AdnGateway) GetOrdinal() int32 {
	if o == nil || o.Ordinal == nil {
		var ret int32
		return ret
	}
	return *o.Ordinal
}

// GetOrdinalOk returns a tuple with the Ordinal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdnGateway) GetOrdinalOk() (*int32, bool) {
	if o == nil || o.Ordinal == nil {
		return nil, false
	}
	return o.Ordinal, true
}

// HasOrdinal returns a boolean if a field has been set.
func (o *AdnGateway) HasOrdinal() bool {
	if o != nil && o.Ordinal != nil {
		return true
	}

	return false
}

// SetOrdinal gets a reference to the given int32 and assigns it to the Ordinal field.
func (o *AdnGateway) SetOrdinal(v int32) {
	o.Ordinal = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AdnGateway) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdnGateway) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AdnGateway) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AdnGateway) SetCity(v string) {
	o.City = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AdnGateway) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdnGateway) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AdnGateway) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AdnGateway) SetRegion(v string) {
	o.Region = &v
}

func (o AdnGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Ordinal != nil {
		toSerialize["ordinal"] = o.Ordinal
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdnGateway) UnmarshalJSON(bytes []byte) (err error) {
	varAdnGateway := _AdnGateway{}

	if err = json.Unmarshal(bytes, &varAdnGateway); err == nil {
		*o = AdnGateway(varAdnGateway)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "ordinal")
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdnGateway struct {
	value *AdnGateway
	isSet bool
}

func (v NullableAdnGateway) Get() *AdnGateway {
	return v.value
}

func (v *NullableAdnGateway) Set(val *AdnGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableAdnGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableAdnGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdnGateway(val *AdnGateway) *NullableAdnGateway {
	return &NullableAdnGateway{value: val, isSet: true}
}

func (v NullableAdnGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdnGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
