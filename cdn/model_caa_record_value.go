/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the CAARecordValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CAARecordValue{}

// CAARecordValue struct for CAARecordValue
type CAARecordValue struct {
	// Should be a valid domain
	Value string `json:"value"`
	Tag string `json:"tag"`
}

// NewCAARecordValue instantiates a new CAARecordValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCAARecordValue(value string, tag string) *CAARecordValue {
	this := CAARecordValue{}
	this.Value = value
	this.Tag = tag
	return &this
}

// NewCAARecordValueWithDefaults instantiates a new CAARecordValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCAARecordValueWithDefaults() *CAARecordValue {
	this := CAARecordValue{}
	return &this
}

// GetValue returns the Value field value
func (o *CAARecordValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CAARecordValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CAARecordValue) SetValue(v string) {
	o.Value = v
}

// GetTag returns the Tag field value
func (o *CAARecordValue) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *CAARecordValue) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *CAARecordValue) SetTag(v string) {
	o.Tag = v
}

func (o CAARecordValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CAARecordValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["tag"] = o.Tag
	return toSerialize, nil
}

type NullableCAARecordValue struct {
	value *CAARecordValue
	isSet bool
}

func (v NullableCAARecordValue) Get() *CAARecordValue {
	return v.value
}

func (v *NullableCAARecordValue) Set(val *CAARecordValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCAARecordValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCAARecordValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCAARecordValue(val *CAARecordValue) *NullableCAARecordValue {
	return &NullableCAARecordValue{value: val, isSet: true}
}

func (v NullableCAARecordValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCAARecordValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


