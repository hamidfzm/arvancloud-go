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

// checks if the CompletedMultipartUpload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompletedMultipartUpload{}

// CompletedMultipartUpload The container for the completed multipart upload details.
type CompletedMultipartUpload struct {
	Parts *Array `json:"Parts,omitempty"`
}

// NewCompletedMultipartUpload instantiates a new CompletedMultipartUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletedMultipartUpload() *CompletedMultipartUpload {
	this := CompletedMultipartUpload{}
	return &this
}

// NewCompletedMultipartUploadWithDefaults instantiates a new CompletedMultipartUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletedMultipartUploadWithDefaults() *CompletedMultipartUpload {
	this := CompletedMultipartUpload{}
	return &this
}

// GetParts returns the Parts field value if set, zero value otherwise.
func (o *CompletedMultipartUpload) GetParts() Array {
	if o == nil || IsNil(o.Parts) {
		var ret Array
		return ret
	}
	return *o.Parts
}

// GetPartsOk returns a tuple with the Parts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompletedMultipartUpload) GetPartsOk() (*Array, bool) {
	if o == nil || IsNil(o.Parts) {
		return nil, false
	}
	return o.Parts, true
}

// HasParts returns a boolean if a field has been set.
func (o *CompletedMultipartUpload) HasParts() bool {
	if o != nil && !IsNil(o.Parts) {
		return true
	}

	return false
}

// SetParts gets a reference to the given Array and assigns it to the Parts field.
func (o *CompletedMultipartUpload) SetParts(v Array) {
	o.Parts = &v
}

func (o CompletedMultipartUpload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompletedMultipartUpload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parts) {
		toSerialize["Parts"] = o.Parts
	}
	return toSerialize, nil
}

type NullableCompletedMultipartUpload struct {
	value *CompletedMultipartUpload
	isSet bool
}

func (v NullableCompletedMultipartUpload) Get() *CompletedMultipartUpload {
	return v.value
}

func (v *NullableCompletedMultipartUpload) Set(val *CompletedMultipartUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletedMultipartUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletedMultipartUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletedMultipartUpload(val *CompletedMultipartUpload) *NullableCompletedMultipartUpload {
	return &NullableCompletedMultipartUpload{value: val, isSet: true}
}

func (v NullableCompletedMultipartUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletedMultipartUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

