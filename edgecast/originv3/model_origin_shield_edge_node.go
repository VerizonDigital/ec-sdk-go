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

// OriginShieldEdgeNode struct for OriginShieldEdgeNode
type OriginShieldEdgeNode struct {
	Id                   *float32          `json:"id,omitempty"`
	Code                 NullableString    `json:"code,omitempty"`
	City                 NullableString    `json:"city,omitempty"`
	Continent            NullableString    `json:"continent,omitempty"`
	Name                 *string           `json:"name,omitempty"`
	Pops                 []OriginShieldPop `json:"pops,omitempty"`
	OriginShieldRegionId *int32            `json:"origin_shield_region_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OriginShieldEdgeNode OriginShieldEdgeNode

// NewOriginShieldEdgeNode instantiates a new OriginShieldEdgeNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginShieldEdgeNode() *OriginShieldEdgeNode {
	this := OriginShieldEdgeNode{}
	return &this
}

// NewOriginShieldEdgeNodeWithDefaults instantiates a new OriginShieldEdgeNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginShieldEdgeNodeWithDefaults() *OriginShieldEdgeNode {
	this := OriginShieldEdgeNode{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OriginShieldEdgeNode) GetId() float32 {
	if o == nil || o.Id == nil {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginShieldEdgeNode) GetIdOk() (*float32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OriginShieldEdgeNode) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *OriginShieldEdgeNode) SetId(v float32) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginShieldEdgeNode) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginShieldEdgeNode) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *OriginShieldEdgeNode) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *OriginShieldEdgeNode) SetCode(v string) {
	o.Code.Set(&v)
}

// SetCodeNil sets the value for Code to be an explicit nil
func (o *OriginShieldEdgeNode) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *OriginShieldEdgeNode) UnsetCode() {
	o.Code.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginShieldEdgeNode) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginShieldEdgeNode) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *OriginShieldEdgeNode) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *OriginShieldEdgeNode) SetCity(v string) {
	o.City.Set(&v)
}

// SetCityNil sets the value for City to be an explicit nil
func (o *OriginShieldEdgeNode) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *OriginShieldEdgeNode) UnsetCity() {
	o.City.Unset()
}

// GetContinent returns the Continent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OriginShieldEdgeNode) GetContinent() string {
	if o == nil || o.Continent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Continent.Get()
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OriginShieldEdgeNode) GetContinentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Continent.Get(), o.Continent.IsSet()
}

// HasContinent returns a boolean if a field has been set.
func (o *OriginShieldEdgeNode) HasContinent() bool {
	if o != nil && o.Continent.IsSet() {
		return true
	}

	return false
}

// SetContinent gets a reference to the given NullableString and assigns it to the Continent field.
func (o *OriginShieldEdgeNode) SetContinent(v string) {
	o.Continent.Set(&v)
}

// SetContinentNil sets the value for Continent to be an explicit nil
func (o *OriginShieldEdgeNode) SetContinentNil() {
	o.Continent.Set(nil)
}

// UnsetContinent ensures that no value is present for Continent, not even an explicit nil
func (o *OriginShieldEdgeNode) UnsetContinent() {
	o.Continent.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OriginShieldEdgeNode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginShieldEdgeNode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OriginShieldEdgeNode) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OriginShieldEdgeNode) SetName(v string) {
	o.Name = &v
}

// GetPops returns the Pops field value if set, zero value otherwise.
func (o *OriginShieldEdgeNode) GetPops() []OriginShieldPop {
	if o == nil || o.Pops == nil {
		var ret []OriginShieldPop
		return ret
	}
	return o.Pops
}

// GetPopsOk returns a tuple with the Pops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginShieldEdgeNode) GetPopsOk() ([]OriginShieldPop, bool) {
	if o == nil || o.Pops == nil {
		return nil, false
	}
	return o.Pops, true
}

// HasPops returns a boolean if a field has been set.
func (o *OriginShieldEdgeNode) HasPops() bool {
	if o != nil && o.Pops != nil {
		return true
	}

	return false
}

// SetPops gets a reference to the given []OriginShieldPop and assigns it to the Pops field.
func (o *OriginShieldEdgeNode) SetPops(v []OriginShieldPop) {
	o.Pops = v
}

// GetOriginShieldRegionId returns the OriginShieldRegionId field value if set, zero value otherwise.
func (o *OriginShieldEdgeNode) GetOriginShieldRegionId() int32 {
	if o == nil || o.OriginShieldRegionId == nil {
		var ret int32
		return ret
	}
	return *o.OriginShieldRegionId
}

// GetOriginShieldRegionIdOk returns a tuple with the OriginShieldRegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginShieldEdgeNode) GetOriginShieldRegionIdOk() (*int32, bool) {
	if o == nil || o.OriginShieldRegionId == nil {
		return nil, false
	}
	return o.OriginShieldRegionId, true
}

// HasOriginShieldRegionId returns a boolean if a field has been set.
func (o *OriginShieldEdgeNode) HasOriginShieldRegionId() bool {
	if o != nil && o.OriginShieldRegionId != nil {
		return true
	}

	return false
}

// SetOriginShieldRegionId gets a reference to the given int32 and assigns it to the OriginShieldRegionId field.
func (o *OriginShieldEdgeNode) SetOriginShieldRegionId(v int32) {
	o.OriginShieldRegionId = &v
}

func (o OriginShieldEdgeNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Continent.IsSet() {
		toSerialize["continent"] = o.Continent.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Pops != nil {
		toSerialize["pops"] = o.Pops
	}
	if o.OriginShieldRegionId != nil {
		toSerialize["origin_shield_region_id"] = o.OriginShieldRegionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OriginShieldEdgeNode) UnmarshalJSON(bytes []byte) (err error) {
	varOriginShieldEdgeNode := _OriginShieldEdgeNode{}

	if err = json.Unmarshal(bytes, &varOriginShieldEdgeNode); err == nil {
		*o = OriginShieldEdgeNode(varOriginShieldEdgeNode)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "code")
		delete(additionalProperties, "city")
		delete(additionalProperties, "continent")
		delete(additionalProperties, "name")
		delete(additionalProperties, "pops")
		delete(additionalProperties, "origin_shield_region_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOriginShieldEdgeNode struct {
	value *OriginShieldEdgeNode
	isSet bool
}

func (v NullableOriginShieldEdgeNode) Get() *OriginShieldEdgeNode {
	return v.value
}

func (v *NullableOriginShieldEdgeNode) Set(val *OriginShieldEdgeNode) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginShieldEdgeNode) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginShieldEdgeNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginShieldEdgeNode(val *OriginShieldEdgeNode) *NullableOriginShieldEdgeNode {
	return &NullableOriginShieldEdgeNode{value: val, isSet: true}
}

func (v NullableOriginShieldEdgeNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginShieldEdgeNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
