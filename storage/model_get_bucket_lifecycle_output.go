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

// checks if the GetBucketLifecycleOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBucketLifecycleOutput{}

// GetBucketLifecycleOutput struct for GetBucketLifecycleOutput
type GetBucketLifecycleOutput struct {
	Rules *Rules `json:"Rules,omitempty"`
}

// NewGetBucketLifecycleOutput instantiates a new GetBucketLifecycleOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBucketLifecycleOutput() *GetBucketLifecycleOutput {
	this := GetBucketLifecycleOutput{}
	return &this
}

// NewGetBucketLifecycleOutputWithDefaults instantiates a new GetBucketLifecycleOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBucketLifecycleOutputWithDefaults() *GetBucketLifecycleOutput {
	this := GetBucketLifecycleOutput{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GetBucketLifecycleOutput) GetRules() Rules {
	if o == nil || IsNil(o.Rules) {
		var ret Rules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBucketLifecycleOutput) GetRulesOk() (*Rules, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GetBucketLifecycleOutput) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given Rules and assigns it to the Rules field.
func (o *GetBucketLifecycleOutput) SetRules(v Rules) {
	o.Rules = &v
}

func (o GetBucketLifecycleOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBucketLifecycleOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["Rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableGetBucketLifecycleOutput struct {
	value *GetBucketLifecycleOutput
	isSet bool
}

func (v NullableGetBucketLifecycleOutput) Get() *GetBucketLifecycleOutput {
	return v.value
}

func (v *NullableGetBucketLifecycleOutput) Set(val *GetBucketLifecycleOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBucketLifecycleOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBucketLifecycleOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBucketLifecycleOutput(val *GetBucketLifecycleOutput) *NullableGetBucketLifecycleOutput {
	return &NullableGetBucketLifecycleOutput{value: val, isSet: true}
}

func (v NullableGetBucketLifecycleOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBucketLifecycleOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


