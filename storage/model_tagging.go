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

// checks if the Tagging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tagging{}

// Tagging Container for <code>TagSet</code> elements.
type Tagging struct {
	TagSet Array `json:"TagSet"`
}

// NewTagging instantiates a new Tagging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagging(tagSet Array) *Tagging {
	this := Tagging{}
	this.TagSet = tagSet
	return &this
}

// NewTaggingWithDefaults instantiates a new Tagging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaggingWithDefaults() *Tagging {
	this := Tagging{}
	return &this
}

// GetTagSet returns the TagSet field value
func (o *Tagging) GetTagSet() Array {
	if o == nil {
		var ret Array
		return ret
	}

	return o.TagSet
}

// GetTagSetOk returns a tuple with the TagSet field value
// and a boolean to check if the value has been set.
func (o *Tagging) GetTagSetOk() (*Array, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagSet, true
}

// SetTagSet sets field value
func (o *Tagging) SetTagSet(v Array) {
	o.TagSet = v
}

func (o Tagging) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tagging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["TagSet"] = o.TagSet
	return toSerialize, nil
}

type NullableTagging struct {
	value *Tagging
	isSet bool
}

func (v NullableTagging) Get() *Tagging {
	return v.value
}

func (v *NullableTagging) Set(val *Tagging) {
	v.value = val
	v.isSet = true
}

func (v NullableTagging) IsSet() bool {
	return v.isSet
}

func (v *NullableTagging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagging(val *Tagging) *NullableTagging {
	return &NullableTagging{value: val, isSet: true}
}

func (v NullableTagging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

