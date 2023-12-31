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

// checks if the V1TagImportPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1TagImportPolicy{}

// V1TagImportPolicy TagImportPolicy controls how images related to this tag will be imported.
type V1TagImportPolicy struct {
	// Insecure is true if the server may bypass certificate verification or connect directly over HTTP during image import.
	Insecure *bool `json:"insecure,omitempty"`
	// Scheduled indicates to the server that this tag should be periodically checked to ensure it is up to date, and imported
	Scheduled *bool `json:"scheduled,omitempty"`
}

// NewV1TagImportPolicy instantiates a new V1TagImportPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1TagImportPolicy() *V1TagImportPolicy {
	this := V1TagImportPolicy{}
	return &this
}

// NewV1TagImportPolicyWithDefaults instantiates a new V1TagImportPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1TagImportPolicyWithDefaults() *V1TagImportPolicy {
	this := V1TagImportPolicy{}
	return &this
}

// GetInsecure returns the Insecure field value if set, zero value otherwise.
func (o *V1TagImportPolicy) GetInsecure() bool {
	if o == nil || IsNil(o.Insecure) {
		var ret bool
		return ret
	}
	return *o.Insecure
}

// GetInsecureOk returns a tuple with the Insecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TagImportPolicy) GetInsecureOk() (*bool, bool) {
	if o == nil || IsNil(o.Insecure) {
		return nil, false
	}
	return o.Insecure, true
}

// HasInsecure returns a boolean if a field has been set.
func (o *V1TagImportPolicy) HasInsecure() bool {
	if o != nil && !IsNil(o.Insecure) {
		return true
	}

	return false
}

// SetInsecure gets a reference to the given bool and assigns it to the Insecure field.
func (o *V1TagImportPolicy) SetInsecure(v bool) {
	o.Insecure = &v
}

// GetScheduled returns the Scheduled field value if set, zero value otherwise.
func (o *V1TagImportPolicy) GetScheduled() bool {
	if o == nil || IsNil(o.Scheduled) {
		var ret bool
		return ret
	}
	return *o.Scheduled
}

// GetScheduledOk returns a tuple with the Scheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1TagImportPolicy) GetScheduledOk() (*bool, bool) {
	if o == nil || IsNil(o.Scheduled) {
		return nil, false
	}
	return o.Scheduled, true
}

// HasScheduled returns a boolean if a field has been set.
func (o *V1TagImportPolicy) HasScheduled() bool {
	if o != nil && !IsNil(o.Scheduled) {
		return true
	}

	return false
}

// SetScheduled gets a reference to the given bool and assigns it to the Scheduled field.
func (o *V1TagImportPolicy) SetScheduled(v bool) {
	o.Scheduled = &v
}

func (o V1TagImportPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1TagImportPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Insecure) {
		toSerialize["insecure"] = o.Insecure
	}
	if !IsNil(o.Scheduled) {
		toSerialize["scheduled"] = o.Scheduled
	}
	return toSerialize, nil
}

type NullableV1TagImportPolicy struct {
	value *V1TagImportPolicy
	isSet bool
}

func (v NullableV1TagImportPolicy) Get() *V1TagImportPolicy {
	return v.value
}

func (v *NullableV1TagImportPolicy) Set(val *V1TagImportPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableV1TagImportPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableV1TagImportPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1TagImportPolicy(val *V1TagImportPolicy) *NullableV1TagImportPolicy {
	return &NullableV1TagImportPolicy{value: val, isSet: true}
}

func (v NullableV1TagImportPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1TagImportPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


