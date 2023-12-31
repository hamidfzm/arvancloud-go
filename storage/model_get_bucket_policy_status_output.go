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

// checks if the GetBucketPolicyStatusOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBucketPolicyStatusOutput{}

// GetBucketPolicyStatusOutput struct for GetBucketPolicyStatusOutput
type GetBucketPolicyStatusOutput struct {
	PolicyStatus *GetBucketPolicyStatusOutputPolicyStatus `json:"PolicyStatus,omitempty"`
}

// NewGetBucketPolicyStatusOutput instantiates a new GetBucketPolicyStatusOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBucketPolicyStatusOutput() *GetBucketPolicyStatusOutput {
	this := GetBucketPolicyStatusOutput{}
	return &this
}

// NewGetBucketPolicyStatusOutputWithDefaults instantiates a new GetBucketPolicyStatusOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBucketPolicyStatusOutputWithDefaults() *GetBucketPolicyStatusOutput {
	this := GetBucketPolicyStatusOutput{}
	return &this
}

// GetPolicyStatus returns the PolicyStatus field value if set, zero value otherwise.
func (o *GetBucketPolicyStatusOutput) GetPolicyStatus() GetBucketPolicyStatusOutputPolicyStatus {
	if o == nil || IsNil(o.PolicyStatus) {
		var ret GetBucketPolicyStatusOutputPolicyStatus
		return ret
	}
	return *o.PolicyStatus
}

// GetPolicyStatusOk returns a tuple with the PolicyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBucketPolicyStatusOutput) GetPolicyStatusOk() (*GetBucketPolicyStatusOutputPolicyStatus, bool) {
	if o == nil || IsNil(o.PolicyStatus) {
		return nil, false
	}
	return o.PolicyStatus, true
}

// HasPolicyStatus returns a boolean if a field has been set.
func (o *GetBucketPolicyStatusOutput) HasPolicyStatus() bool {
	if o != nil && !IsNil(o.PolicyStatus) {
		return true
	}

	return false
}

// SetPolicyStatus gets a reference to the given GetBucketPolicyStatusOutputPolicyStatus and assigns it to the PolicyStatus field.
func (o *GetBucketPolicyStatusOutput) SetPolicyStatus(v GetBucketPolicyStatusOutputPolicyStatus) {
	o.PolicyStatus = &v
}

func (o GetBucketPolicyStatusOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBucketPolicyStatusOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PolicyStatus) {
		toSerialize["PolicyStatus"] = o.PolicyStatus
	}
	return toSerialize, nil
}

type NullableGetBucketPolicyStatusOutput struct {
	value *GetBucketPolicyStatusOutput
	isSet bool
}

func (v NullableGetBucketPolicyStatusOutput) Get() *GetBucketPolicyStatusOutput {
	return v.value
}

func (v *NullableGetBucketPolicyStatusOutput) Set(val *GetBucketPolicyStatusOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBucketPolicyStatusOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBucketPolicyStatusOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBucketPolicyStatusOutput(val *GetBucketPolicyStatusOutput) *NullableGetBucketPolicyStatusOutput {
	return &NullableGetBucketPolicyStatusOutput{value: val, isSet: true}
}

func (v NullableGetBucketPolicyStatusOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBucketPolicyStatusOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


