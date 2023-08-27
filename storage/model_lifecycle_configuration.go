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

// checks if the LifecycleConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifecycleConfiguration{}

// LifecycleConfiguration Container for lifecycle rules. You can add as many as 1000 rules.
type LifecycleConfiguration struct {
	Rules Rules `json:"Rules"`
}

// NewLifecycleConfiguration instantiates a new LifecycleConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleConfiguration(rules Rules) *LifecycleConfiguration {
	this := LifecycleConfiguration{}
	this.Rules = rules
	return &this
}

// NewLifecycleConfigurationWithDefaults instantiates a new LifecycleConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleConfigurationWithDefaults() *LifecycleConfiguration {
	this := LifecycleConfiguration{}
	return &this
}

// GetRules returns the Rules field value
func (o *LifecycleConfiguration) GetRules() Rules {
	if o == nil {
		var ret Rules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *LifecycleConfiguration) GetRulesOk() (*Rules, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value
func (o *LifecycleConfiguration) SetRules(v Rules) {
	o.Rules = v
}

func (o LifecycleConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Rules"] = o.Rules
	return toSerialize, nil
}

type NullableLifecycleConfiguration struct {
	value *LifecycleConfiguration
	isSet bool
}

func (v NullableLifecycleConfiguration) Get() *LifecycleConfiguration {
	return v.value
}

func (v *NullableLifecycleConfiguration) Set(val *LifecycleConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleConfiguration(val *LifecycleConfiguration) *NullableLifecycleConfiguration {
	return &NullableLifecycleConfiguration{value: val, isSet: true}
}

func (v NullableLifecycleConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

