/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the SecurityGroupName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupName{}

// SecurityGroupName struct for SecurityGroupName
type SecurityGroupName struct {
	Name *string `json:"name,omitempty"`
}

// NewSecurityGroupName instantiates a new SecurityGroupName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupName() *SecurityGroupName {
	this := SecurityGroupName{}
	return &this
}

// NewSecurityGroupNameWithDefaults instantiates a new SecurityGroupName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupNameWithDefaults() *SecurityGroupName {
	this := SecurityGroupName{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityGroupName) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupName) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityGroupName) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityGroupName) SetName(v string) {
	o.Name = &v
}

func (o SecurityGroupName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityGroupName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSecurityGroupName struct {
	value *SecurityGroupName
	isSet bool
}

func (v NullableSecurityGroupName) Get() *SecurityGroupName {
	return v.value
}

func (v *NullableSecurityGroupName) Set(val *SecurityGroupName) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupName) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupName(val *SecurityGroupName) *NullableSecurityGroupName {
	return &NullableSecurityGroupName{value: val, isSet: true}
}

func (v NullableSecurityGroupName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


