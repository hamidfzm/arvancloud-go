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

// checks if the OwnershipControlsRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnershipControlsRule{}

// OwnershipControlsRule The container element for an ownership control rule.
type OwnershipControlsRule struct {
	ObjectOwnership ObjectOwnership `json:"ObjectOwnership"`
}

// NewOwnershipControlsRule instantiates a new OwnershipControlsRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnershipControlsRule(objectOwnership ObjectOwnership) *OwnershipControlsRule {
	this := OwnershipControlsRule{}
	this.ObjectOwnership = objectOwnership
	return &this
}

// NewOwnershipControlsRuleWithDefaults instantiates a new OwnershipControlsRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnershipControlsRuleWithDefaults() *OwnershipControlsRule {
	this := OwnershipControlsRule{}
	return &this
}

// GetObjectOwnership returns the ObjectOwnership field value
func (o *OwnershipControlsRule) GetObjectOwnership() ObjectOwnership {
	if o == nil {
		var ret ObjectOwnership
		return ret
	}

	return o.ObjectOwnership
}

// GetObjectOwnershipOk returns a tuple with the ObjectOwnership field value
// and a boolean to check if the value has been set.
func (o *OwnershipControlsRule) GetObjectOwnershipOk() (*ObjectOwnership, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectOwnership, true
}

// SetObjectOwnership sets field value
func (o *OwnershipControlsRule) SetObjectOwnership(v ObjectOwnership) {
	o.ObjectOwnership = v
}

func (o OwnershipControlsRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnershipControlsRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ObjectOwnership"] = o.ObjectOwnership
	return toSerialize, nil
}

type NullableOwnershipControlsRule struct {
	value *OwnershipControlsRule
	isSet bool
}

func (v NullableOwnershipControlsRule) Get() *OwnershipControlsRule {
	return v.value
}

func (v *NullableOwnershipControlsRule) Set(val *OwnershipControlsRule) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnershipControlsRule) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnershipControlsRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnershipControlsRule(val *OwnershipControlsRule) *NullableOwnershipControlsRule {
	return &NullableOwnershipControlsRule{value: val, isSet: true}
}

func (v NullableOwnershipControlsRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnershipControlsRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


