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

// checks if the RuntimeRawExtension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuntimeRawExtension{}

// RuntimeRawExtension struct for RuntimeRawExtension
type RuntimeRawExtension struct {
	Raw string `json:"Raw"`
}

// NewRuntimeRawExtension instantiates a new RuntimeRawExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuntimeRawExtension(raw string) *RuntimeRawExtension {
	this := RuntimeRawExtension{}
	this.Raw = raw
	return &this
}

// NewRuntimeRawExtensionWithDefaults instantiates a new RuntimeRawExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuntimeRawExtensionWithDefaults() *RuntimeRawExtension {
	this := RuntimeRawExtension{}
	return &this
}

// GetRaw returns the Raw field value
func (o *RuntimeRawExtension) GetRaw() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Raw
}

// GetRawOk returns a tuple with the Raw field value
// and a boolean to check if the value has been set.
func (o *RuntimeRawExtension) GetRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Raw, true
}

// SetRaw sets field value
func (o *RuntimeRawExtension) SetRaw(v string) {
	o.Raw = v
}

func (o RuntimeRawExtension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuntimeRawExtension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Raw"] = o.Raw
	return toSerialize, nil
}

type NullableRuntimeRawExtension struct {
	value *RuntimeRawExtension
	isSet bool
}

func (v NullableRuntimeRawExtension) Get() *RuntimeRawExtension {
	return v.value
}

func (v *NullableRuntimeRawExtension) Set(val *RuntimeRawExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableRuntimeRawExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableRuntimeRawExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuntimeRawExtension(val *RuntimeRawExtension) *NullableRuntimeRawExtension {
	return &NullableRuntimeRawExtension{value: val, isSet: true}
}

func (v NullableRuntimeRawExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuntimeRawExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


