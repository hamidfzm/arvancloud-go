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

// checks if the PutBucketLifecycleRequestLifecycleConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketLifecycleRequestLifecycleConfiguration{}

// PutBucketLifecycleRequestLifecycleConfiguration struct for PutBucketLifecycleRequestLifecycleConfiguration
type PutBucketLifecycleRequestLifecycleConfiguration struct {
	Rules Rules `json:"Rules"`
}

// NewPutBucketLifecycleRequestLifecycleConfiguration instantiates a new PutBucketLifecycleRequestLifecycleConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketLifecycleRequestLifecycleConfiguration(rules Rules) *PutBucketLifecycleRequestLifecycleConfiguration {
	this := PutBucketLifecycleRequestLifecycleConfiguration{}
	this.Rules = rules
	return &this
}

// NewPutBucketLifecycleRequestLifecycleConfigurationWithDefaults instantiates a new PutBucketLifecycleRequestLifecycleConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketLifecycleRequestLifecycleConfigurationWithDefaults() *PutBucketLifecycleRequestLifecycleConfiguration {
	this := PutBucketLifecycleRequestLifecycleConfiguration{}
	return &this
}

// GetRules returns the Rules field value
func (o *PutBucketLifecycleRequestLifecycleConfiguration) GetRules() Rules {
	if o == nil {
		var ret Rules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *PutBucketLifecycleRequestLifecycleConfiguration) GetRulesOk() (*Rules, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value
func (o *PutBucketLifecycleRequestLifecycleConfiguration) SetRules(v Rules) {
	o.Rules = v
}

func (o PutBucketLifecycleRequestLifecycleConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketLifecycleRequestLifecycleConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Rules"] = o.Rules
	return toSerialize, nil
}

type NullablePutBucketLifecycleRequestLifecycleConfiguration struct {
	value *PutBucketLifecycleRequestLifecycleConfiguration
	isSet bool
}

func (v NullablePutBucketLifecycleRequestLifecycleConfiguration) Get() *PutBucketLifecycleRequestLifecycleConfiguration {
	return v.value
}

func (v *NullablePutBucketLifecycleRequestLifecycleConfiguration) Set(val *PutBucketLifecycleRequestLifecycleConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketLifecycleRequestLifecycleConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketLifecycleRequestLifecycleConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketLifecycleRequestLifecycleConfiguration(val *PutBucketLifecycleRequestLifecycleConfiguration) *NullablePutBucketLifecycleRequestLifecycleConfiguration {
	return &NullablePutBucketLifecycleRequestLifecycleConfiguration{value: val, isSet: true}
}

func (v NullablePutBucketLifecycleRequestLifecycleConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketLifecycleRequestLifecycleConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


