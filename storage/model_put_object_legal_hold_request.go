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

// checks if the PutObjectLegalHoldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutObjectLegalHoldRequest{}

// PutObjectLegalHoldRequest struct for PutObjectLegalHoldRequest
type PutObjectLegalHoldRequest struct {
	LegalHold *PutObjectLegalHoldRequestLegalHold `json:"LegalHold,omitempty"`
}

// NewPutObjectLegalHoldRequest instantiates a new PutObjectLegalHoldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutObjectLegalHoldRequest() *PutObjectLegalHoldRequest {
	this := PutObjectLegalHoldRequest{}
	return &this
}

// NewPutObjectLegalHoldRequestWithDefaults instantiates a new PutObjectLegalHoldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutObjectLegalHoldRequestWithDefaults() *PutObjectLegalHoldRequest {
	this := PutObjectLegalHoldRequest{}
	return &this
}

// GetLegalHold returns the LegalHold field value if set, zero value otherwise.
func (o *PutObjectLegalHoldRequest) GetLegalHold() PutObjectLegalHoldRequestLegalHold {
	if o == nil || IsNil(o.LegalHold) {
		var ret PutObjectLegalHoldRequestLegalHold
		return ret
	}
	return *o.LegalHold
}

// GetLegalHoldOk returns a tuple with the LegalHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutObjectLegalHoldRequest) GetLegalHoldOk() (*PutObjectLegalHoldRequestLegalHold, bool) {
	if o == nil || IsNil(o.LegalHold) {
		return nil, false
	}
	return o.LegalHold, true
}

// HasLegalHold returns a boolean if a field has been set.
func (o *PutObjectLegalHoldRequest) HasLegalHold() bool {
	if o != nil && !IsNil(o.LegalHold) {
		return true
	}

	return false
}

// SetLegalHold gets a reference to the given PutObjectLegalHoldRequestLegalHold and assigns it to the LegalHold field.
func (o *PutObjectLegalHoldRequest) SetLegalHold(v PutObjectLegalHoldRequestLegalHold) {
	o.LegalHold = &v
}

func (o PutObjectLegalHoldRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutObjectLegalHoldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LegalHold) {
		toSerialize["LegalHold"] = o.LegalHold
	}
	return toSerialize, nil
}

type NullablePutObjectLegalHoldRequest struct {
	value *PutObjectLegalHoldRequest
	isSet bool
}

func (v NullablePutObjectLegalHoldRequest) Get() *PutObjectLegalHoldRequest {
	return v.value
}

func (v *NullablePutObjectLegalHoldRequest) Set(val *PutObjectLegalHoldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutObjectLegalHoldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutObjectLegalHoldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutObjectLegalHoldRequest(val *PutObjectLegalHoldRequest) *NullablePutObjectLegalHoldRequest {
	return &NullablePutObjectLegalHoldRequest{value: val, isSet: true}
}

func (v NullablePutObjectLegalHoldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutObjectLegalHoldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

