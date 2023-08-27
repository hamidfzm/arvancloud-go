/*
Arvancloud PaaS REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paas

import (
	"encoding/json"
)

// checks if the V1PodDNSConfigOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PodDNSConfigOption{}

// V1PodDNSConfigOption PodDNSConfigOption defines DNS resolver options of a pod.
type V1PodDNSConfigOption struct {
	// Required.
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewV1PodDNSConfigOption instantiates a new V1PodDNSConfigOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PodDNSConfigOption() *V1PodDNSConfigOption {
	this := V1PodDNSConfigOption{}
	return &this
}

// NewV1PodDNSConfigOptionWithDefaults instantiates a new V1PodDNSConfigOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PodDNSConfigOptionWithDefaults() *V1PodDNSConfigOption {
	this := V1PodDNSConfigOption{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1PodDNSConfigOption) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodDNSConfigOption) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1PodDNSConfigOption) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1PodDNSConfigOption) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V1PodDNSConfigOption) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodDNSConfigOption) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V1PodDNSConfigOption) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *V1PodDNSConfigOption) SetValue(v string) {
	o.Value = &v
}

func (o V1PodDNSConfigOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PodDNSConfigOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableV1PodDNSConfigOption struct {
	value *V1PodDNSConfigOption
	isSet bool
}

func (v NullableV1PodDNSConfigOption) Get() *V1PodDNSConfigOption {
	return v.value
}

func (v *NullableV1PodDNSConfigOption) Set(val *V1PodDNSConfigOption) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PodDNSConfigOption) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PodDNSConfigOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PodDNSConfigOption(val *V1PodDNSConfigOption) *NullableV1PodDNSConfigOption {
	return &NullableV1PodDNSConfigOption{value: val, isSet: true}
}

func (v NullableV1PodDNSConfigOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PodDNSConfigOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

