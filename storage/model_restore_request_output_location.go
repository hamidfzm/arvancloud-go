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

// checks if the RestoreRequestOutputLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreRequestOutputLocation{}

// RestoreRequestOutputLocation struct for RestoreRequestOutputLocation
type RestoreRequestOutputLocation struct {
	S3 *OutputLocationS3 `json:"S3,omitempty"`
}

// NewRestoreRequestOutputLocation instantiates a new RestoreRequestOutputLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreRequestOutputLocation() *RestoreRequestOutputLocation {
	this := RestoreRequestOutputLocation{}
	return &this
}

// NewRestoreRequestOutputLocationWithDefaults instantiates a new RestoreRequestOutputLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreRequestOutputLocationWithDefaults() *RestoreRequestOutputLocation {
	this := RestoreRequestOutputLocation{}
	return &this
}

// GetS3 returns the S3 field value if set, zero value otherwise.
func (o *RestoreRequestOutputLocation) GetS3() OutputLocationS3 {
	if o == nil || IsNil(o.S3) {
		var ret OutputLocationS3
		return ret
	}
	return *o.S3
}

// GetS3Ok returns a tuple with the S3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreRequestOutputLocation) GetS3Ok() (*OutputLocationS3, bool) {
	if o == nil || IsNil(o.S3) {
		return nil, false
	}
	return o.S3, true
}

// HasS3 returns a boolean if a field has been set.
func (o *RestoreRequestOutputLocation) HasS3() bool {
	if o != nil && !IsNil(o.S3) {
		return true
	}

	return false
}

// SetS3 gets a reference to the given OutputLocationS3 and assigns it to the S3 field.
func (o *RestoreRequestOutputLocation) SetS3(v OutputLocationS3) {
	o.S3 = &v
}

func (o RestoreRequestOutputLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreRequestOutputLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.S3) {
		toSerialize["S3"] = o.S3
	}
	return toSerialize, nil
}

type NullableRestoreRequestOutputLocation struct {
	value *RestoreRequestOutputLocation
	isSet bool
}

func (v NullableRestoreRequestOutputLocation) Get() *RestoreRequestOutputLocation {
	return v.value
}

func (v *NullableRestoreRequestOutputLocation) Set(val *RestoreRequestOutputLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreRequestOutputLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreRequestOutputLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreRequestOutputLocation(val *RestoreRequestOutputLocation) *NullableRestoreRequestOutputLocation {
	return &NullableRestoreRequestOutputLocation{value: val, isSet: true}
}

func (v NullableRestoreRequestOutputLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreRequestOutputLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


