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

// checks if the ObjectLockRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectLockRule{}

// ObjectLockRule The container element for an Object Lock rule.
type ObjectLockRule struct {
	DefaultRetention *ObjectLockRuleDefaultRetention `json:"DefaultRetention,omitempty"`
}

// NewObjectLockRule instantiates a new ObjectLockRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectLockRule() *ObjectLockRule {
	this := ObjectLockRule{}
	return &this
}

// NewObjectLockRuleWithDefaults instantiates a new ObjectLockRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectLockRuleWithDefaults() *ObjectLockRule {
	this := ObjectLockRule{}
	return &this
}

// GetDefaultRetention returns the DefaultRetention field value if set, zero value otherwise.
func (o *ObjectLockRule) GetDefaultRetention() ObjectLockRuleDefaultRetention {
	if o == nil || IsNil(o.DefaultRetention) {
		var ret ObjectLockRuleDefaultRetention
		return ret
	}
	return *o.DefaultRetention
}

// GetDefaultRetentionOk returns a tuple with the DefaultRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectLockRule) GetDefaultRetentionOk() (*ObjectLockRuleDefaultRetention, bool) {
	if o == nil || IsNil(o.DefaultRetention) {
		return nil, false
	}
	return o.DefaultRetention, true
}

// HasDefaultRetention returns a boolean if a field has been set.
func (o *ObjectLockRule) HasDefaultRetention() bool {
	if o != nil && !IsNil(o.DefaultRetention) {
		return true
	}

	return false
}

// SetDefaultRetention gets a reference to the given ObjectLockRuleDefaultRetention and assigns it to the DefaultRetention field.
func (o *ObjectLockRule) SetDefaultRetention(v ObjectLockRuleDefaultRetention) {
	o.DefaultRetention = &v
}

func (o ObjectLockRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectLockRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultRetention) {
		toSerialize["DefaultRetention"] = o.DefaultRetention
	}
	return toSerialize, nil
}

type NullableObjectLockRule struct {
	value *ObjectLockRule
	isSet bool
}

func (v NullableObjectLockRule) Get() *ObjectLockRule {
	return v.value
}

func (v *NullableObjectLockRule) Set(val *ObjectLockRule) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectLockRule) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectLockRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectLockRule(val *ObjectLockRule) *NullableObjectLockRule {
	return &NullableObjectLockRule{value: val, isSet: true}
}

func (v NullableObjectLockRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectLockRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


