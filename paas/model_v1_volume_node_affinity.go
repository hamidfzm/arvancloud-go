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

// checks if the V1VolumeNodeAffinity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1VolumeNodeAffinity{}

// V1VolumeNodeAffinity VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
type V1VolumeNodeAffinity struct {
	Required *V1NodeSelector `json:"required,omitempty"`
}

// NewV1VolumeNodeAffinity instantiates a new V1VolumeNodeAffinity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1VolumeNodeAffinity() *V1VolumeNodeAffinity {
	this := V1VolumeNodeAffinity{}
	return &this
}

// NewV1VolumeNodeAffinityWithDefaults instantiates a new V1VolumeNodeAffinity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1VolumeNodeAffinityWithDefaults() *V1VolumeNodeAffinity {
	this := V1VolumeNodeAffinity{}
	return &this
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *V1VolumeNodeAffinity) GetRequired() V1NodeSelector {
	if o == nil || IsNil(o.Required) {
		var ret V1NodeSelector
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1VolumeNodeAffinity) GetRequiredOk() (*V1NodeSelector, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *V1VolumeNodeAffinity) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given V1NodeSelector and assigns it to the Required field.
func (o *V1VolumeNodeAffinity) SetRequired(v V1NodeSelector) {
	o.Required = &v
}

func (o V1VolumeNodeAffinity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1VolumeNodeAffinity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	return toSerialize, nil
}

type NullableV1VolumeNodeAffinity struct {
	value *V1VolumeNodeAffinity
	isSet bool
}

func (v NullableV1VolumeNodeAffinity) Get() *V1VolumeNodeAffinity {
	return v.value
}

func (v *NullableV1VolumeNodeAffinity) Set(val *V1VolumeNodeAffinity) {
	v.value = val
	v.isSet = true
}

func (v NullableV1VolumeNodeAffinity) IsSet() bool {
	return v.isSet
}

func (v *NullableV1VolumeNodeAffinity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1VolumeNodeAffinity(val *V1VolumeNodeAffinity) *NullableV1VolumeNodeAffinity {
	return &NullableV1VolumeNodeAffinity{value: val, isSet: true}
}

func (v NullableV1VolumeNodeAffinity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1VolumeNodeAffinity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


