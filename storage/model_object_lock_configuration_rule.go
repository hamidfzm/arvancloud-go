/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
)

// checks if the ObjectLockConfigurationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectLockConfigurationRule{}

// ObjectLockConfigurationRule struct for ObjectLockConfigurationRule
type ObjectLockConfigurationRule struct {
	DefaultRetention *ObjectLockRuleDefaultRetention `json:"DefaultRetention,omitempty"`
}

// NewObjectLockConfigurationRule instantiates a new ObjectLockConfigurationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectLockConfigurationRule() *ObjectLockConfigurationRule {
	this := ObjectLockConfigurationRule{}
	return &this
}

// NewObjectLockConfigurationRuleWithDefaults instantiates a new ObjectLockConfigurationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectLockConfigurationRuleWithDefaults() *ObjectLockConfigurationRule {
	this := ObjectLockConfigurationRule{}
	return &this
}

// GetDefaultRetention returns the DefaultRetention field value if set, zero value otherwise.
func (o *ObjectLockConfigurationRule) GetDefaultRetention() ObjectLockRuleDefaultRetention {
	if o == nil || IsNil(o.DefaultRetention) {
		var ret ObjectLockRuleDefaultRetention
		return ret
	}
	return *o.DefaultRetention
}

// GetDefaultRetentionOk returns a tuple with the DefaultRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectLockConfigurationRule) GetDefaultRetentionOk() (*ObjectLockRuleDefaultRetention, bool) {
	if o == nil || IsNil(o.DefaultRetention) {
		return nil, false
	}
	return o.DefaultRetention, true
}

// HasDefaultRetention returns a boolean if a field has been set.
func (o *ObjectLockConfigurationRule) HasDefaultRetention() bool {
	if o != nil && !IsNil(o.DefaultRetention) {
		return true
	}

	return false
}

// SetDefaultRetention gets a reference to the given ObjectLockRuleDefaultRetention and assigns it to the DefaultRetention field.
func (o *ObjectLockConfigurationRule) SetDefaultRetention(v ObjectLockRuleDefaultRetention) {
	o.DefaultRetention = &v
}

func (o ObjectLockConfigurationRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectLockConfigurationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultRetention) {
		toSerialize["DefaultRetention"] = o.DefaultRetention
	}
	return toSerialize, nil
}

type NullableObjectLockConfigurationRule struct {
	value *ObjectLockConfigurationRule
	isSet bool
}

func (v NullableObjectLockConfigurationRule) Get() *ObjectLockConfigurationRule {
	return v.value
}

func (v *NullableObjectLockConfigurationRule) Set(val *ObjectLockConfigurationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectLockConfigurationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectLockConfigurationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectLockConfigurationRule(val *ObjectLockConfigurationRule) *NullableObjectLockConfigurationRule {
	return &NullableObjectLockConfigurationRule{value: val, isSet: true}
}

func (v NullableObjectLockConfigurationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectLockConfigurationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


