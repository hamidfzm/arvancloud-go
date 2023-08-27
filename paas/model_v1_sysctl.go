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

// checks if the V1Sysctl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1Sysctl{}

// V1Sysctl Sysctl defines a kernel parameter to be set
type V1Sysctl struct {
	// Name of a property to set
	Name string `json:"name"`
	// Value of a property to set
	Value string `json:"value"`
}

// NewV1Sysctl instantiates a new V1Sysctl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Sysctl(name string, value string) *V1Sysctl {
	this := V1Sysctl{}
	this.Name = name
	this.Value = value
	return &this
}

// NewV1SysctlWithDefaults instantiates a new V1Sysctl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SysctlWithDefaults() *V1Sysctl {
	this := V1Sysctl{}
	return &this
}

// GetName returns the Name field value
func (o *V1Sysctl) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1Sysctl) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1Sysctl) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *V1Sysctl) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *V1Sysctl) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *V1Sysctl) SetValue(v string) {
	o.Value = v
}

func (o V1Sysctl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1Sysctl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableV1Sysctl struct {
	value *V1Sysctl
	isSet bool
}

func (v NullableV1Sysctl) Get() *V1Sysctl {
	return v.value
}

func (v *NullableV1Sysctl) Set(val *V1Sysctl) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Sysctl) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Sysctl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Sysctl(val *V1Sysctl) *NullableV1Sysctl {
	return &NullableV1Sysctl{value: val, isSet: true}
}

func (v NullableV1Sysctl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Sysctl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


