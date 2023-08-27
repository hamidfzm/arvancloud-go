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

// checks if the PutBucketOwnershipControlsRequestOwnershipControls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketOwnershipControlsRequestOwnershipControls{}

// PutBucketOwnershipControlsRequestOwnershipControls struct for PutBucketOwnershipControlsRequestOwnershipControls
type PutBucketOwnershipControlsRequestOwnershipControls struct {
	Rules OwnershipControlsRules `json:"Rules"`
}

// NewPutBucketOwnershipControlsRequestOwnershipControls instantiates a new PutBucketOwnershipControlsRequestOwnershipControls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketOwnershipControlsRequestOwnershipControls(rules OwnershipControlsRules) *PutBucketOwnershipControlsRequestOwnershipControls {
	this := PutBucketOwnershipControlsRequestOwnershipControls{}
	this.Rules = rules
	return &this
}

// NewPutBucketOwnershipControlsRequestOwnershipControlsWithDefaults instantiates a new PutBucketOwnershipControlsRequestOwnershipControls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketOwnershipControlsRequestOwnershipControlsWithDefaults() *PutBucketOwnershipControlsRequestOwnershipControls {
	this := PutBucketOwnershipControlsRequestOwnershipControls{}
	return &this
}

// GetRules returns the Rules field value
func (o *PutBucketOwnershipControlsRequestOwnershipControls) GetRules() OwnershipControlsRules {
	if o == nil {
		var ret OwnershipControlsRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *PutBucketOwnershipControlsRequestOwnershipControls) GetRulesOk() (*OwnershipControlsRules, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value
func (o *PutBucketOwnershipControlsRequestOwnershipControls) SetRules(v OwnershipControlsRules) {
	o.Rules = v
}

func (o PutBucketOwnershipControlsRequestOwnershipControls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketOwnershipControlsRequestOwnershipControls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Rules"] = o.Rules
	return toSerialize, nil
}

type NullablePutBucketOwnershipControlsRequestOwnershipControls struct {
	value *PutBucketOwnershipControlsRequestOwnershipControls
	isSet bool
}

func (v NullablePutBucketOwnershipControlsRequestOwnershipControls) Get() *PutBucketOwnershipControlsRequestOwnershipControls {
	return v.value
}

func (v *NullablePutBucketOwnershipControlsRequestOwnershipControls) Set(val *PutBucketOwnershipControlsRequestOwnershipControls) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketOwnershipControlsRequestOwnershipControls) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketOwnershipControlsRequestOwnershipControls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketOwnershipControlsRequestOwnershipControls(val *PutBucketOwnershipControlsRequestOwnershipControls) *NullablePutBucketOwnershipControlsRequestOwnershipControls {
	return &NullablePutBucketOwnershipControlsRequestOwnershipControls{value: val, isSet: true}
}

func (v NullablePutBucketOwnershipControlsRequestOwnershipControls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketOwnershipControlsRequestOwnershipControls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

