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

// checks if the ObjectLockLegalHold type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectLockLegalHold{}

// ObjectLockLegalHold A Legal Hold configuration for an object.
type ObjectLockLegalHold struct {
	Status *ObjectLockLegalHoldStatus `json:"Status,omitempty"`
}

// NewObjectLockLegalHold instantiates a new ObjectLockLegalHold object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectLockLegalHold() *ObjectLockLegalHold {
	this := ObjectLockLegalHold{}
	return &this
}

// NewObjectLockLegalHoldWithDefaults instantiates a new ObjectLockLegalHold object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectLockLegalHoldWithDefaults() *ObjectLockLegalHold {
	this := ObjectLockLegalHold{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ObjectLockLegalHold) GetStatus() ObjectLockLegalHoldStatus {
	if o == nil || IsNil(o.Status) {
		var ret ObjectLockLegalHoldStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectLockLegalHold) GetStatusOk() (*ObjectLockLegalHoldStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ObjectLockLegalHold) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ObjectLockLegalHoldStatus and assigns it to the Status field.
func (o *ObjectLockLegalHold) SetStatus(v ObjectLockLegalHoldStatus) {
	o.Status = &v
}

func (o ObjectLockLegalHold) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectLockLegalHold) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	return toSerialize, nil
}

type NullableObjectLockLegalHold struct {
	value *ObjectLockLegalHold
	isSet bool
}

func (v NullableObjectLockLegalHold) Get() *ObjectLockLegalHold {
	return v.value
}

func (v *NullableObjectLockLegalHold) Set(val *ObjectLockLegalHold) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectLockLegalHold) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectLockLegalHold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectLockLegalHold(val *ObjectLockLegalHold) *NullableObjectLockLegalHold {
	return &NullableObjectLockLegalHold{value: val, isSet: true}
}

func (v NullableObjectLockLegalHold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectLockLegalHold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


