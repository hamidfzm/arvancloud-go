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

// checks if the PutObjectLegalHoldRequestLegalHold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutObjectLegalHoldRequestLegalHold{}

// PutObjectLegalHoldRequestLegalHold struct for PutObjectLegalHoldRequestLegalHold
type PutObjectLegalHoldRequestLegalHold struct {
	Status *ObjectLockLegalHoldStatus `json:"Status,omitempty"`
}

// NewPutObjectLegalHoldRequestLegalHold instantiates a new PutObjectLegalHoldRequestLegalHold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutObjectLegalHoldRequestLegalHold() *PutObjectLegalHoldRequestLegalHold {
	this := PutObjectLegalHoldRequestLegalHold{}
	return &this
}

// NewPutObjectLegalHoldRequestLegalHoldWithDefaults instantiates a new PutObjectLegalHoldRequestLegalHold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutObjectLegalHoldRequestLegalHoldWithDefaults() *PutObjectLegalHoldRequestLegalHold {
	this := PutObjectLegalHoldRequestLegalHold{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PutObjectLegalHoldRequestLegalHold) GetStatus() ObjectLockLegalHoldStatus {
	if o == nil || IsNil(o.Status) {
		var ret ObjectLockLegalHoldStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutObjectLegalHoldRequestLegalHold) GetStatusOk() (*ObjectLockLegalHoldStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PutObjectLegalHoldRequestLegalHold) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ObjectLockLegalHoldStatus and assigns it to the Status field.
func (o *PutObjectLegalHoldRequestLegalHold) SetStatus(v ObjectLockLegalHoldStatus) {
	o.Status = &v
}

func (o PutObjectLegalHoldRequestLegalHold) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutObjectLegalHoldRequestLegalHold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePutObjectLegalHoldRequestLegalHold struct {
	value *PutObjectLegalHoldRequestLegalHold
	isSet bool
}

func (v NullablePutObjectLegalHoldRequestLegalHold) Get() *PutObjectLegalHoldRequestLegalHold {
	return v.value
}

func (v *NullablePutObjectLegalHoldRequestLegalHold) Set(val *PutObjectLegalHoldRequestLegalHold) {
	v.value = val
	v.isSet = true
}

func (v NullablePutObjectLegalHoldRequestLegalHold) IsSet() bool {
	return v.isSet
}

func (v *NullablePutObjectLegalHoldRequestLegalHold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutObjectLegalHoldRequestLegalHold(val *PutObjectLegalHoldRequestLegalHold) *NullablePutObjectLegalHoldRequestLegalHold {
	return &NullablePutObjectLegalHoldRequestLegalHold{value: val, isSet: true}
}

func (v NullablePutObjectLegalHoldRequestLegalHold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutObjectLegalHoldRequestLegalHold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


