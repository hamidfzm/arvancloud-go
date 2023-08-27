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

// checks if the SSEKMS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSEKMS{}

// SSEKMS Specifies the use of SSE-KMS to encrypt delivered inventory reports.
type SSEKMS struct {
	KeyId string `json:"KeyId"`
}

// NewSSEKMS instantiates a new SSEKMS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSEKMS(keyId string) *SSEKMS {
	this := SSEKMS{}
	this.KeyId = keyId
	return &this
}

// NewSSEKMSWithDefaults instantiates a new SSEKMS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSEKMSWithDefaults() *SSEKMS {
	this := SSEKMS{}
	return &this
}

// GetKeyId returns the KeyId field value
func (o *SSEKMS) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *SSEKMS) GetKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *SSEKMS) SetKeyId(v string) {
	o.KeyId = v
}

func (o SSEKMS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSEKMS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["KeyId"] = o.KeyId
	return toSerialize, nil
}

type NullableSSEKMS struct {
	value *SSEKMS
	isSet bool
}

func (v NullableSSEKMS) Get() *SSEKMS {
	return v.value
}

func (v *NullableSSEKMS) Set(val *SSEKMS) {
	v.value = val
	v.isSet = true
}

func (v NullableSSEKMS) IsSet() bool {
	return v.isSet
}

func (v *NullableSSEKMS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSEKMS(val *SSEKMS) *NullableSSEKMS {
	return &NullableSSEKMS{value: val, isSet: true}
}

func (v NullableSSEKMS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSEKMS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

