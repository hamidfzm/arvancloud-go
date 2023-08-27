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

// checks if the SelectObjectContentEventStreamProgress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectObjectContentEventStreamProgress{}

// SelectObjectContentEventStreamProgress struct for SelectObjectContentEventStreamProgress
type SelectObjectContentEventStreamProgress struct {
	Details *ProgressEventDetails `json:"Details,omitempty"`
}

// NewSelectObjectContentEventStreamProgress instantiates a new SelectObjectContentEventStreamProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectObjectContentEventStreamProgress() *SelectObjectContentEventStreamProgress {
	this := SelectObjectContentEventStreamProgress{}
	return &this
}

// NewSelectObjectContentEventStreamProgressWithDefaults instantiates a new SelectObjectContentEventStreamProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectObjectContentEventStreamProgressWithDefaults() *SelectObjectContentEventStreamProgress {
	this := SelectObjectContentEventStreamProgress{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SelectObjectContentEventStreamProgress) GetDetails() ProgressEventDetails {
	if o == nil || IsNil(o.Details) {
		var ret ProgressEventDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectObjectContentEventStreamProgress) GetDetailsOk() (*ProgressEventDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SelectObjectContentEventStreamProgress) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given ProgressEventDetails and assigns it to the Details field.
func (o *SelectObjectContentEventStreamProgress) SetDetails(v ProgressEventDetails) {
	o.Details = &v
}

func (o SelectObjectContentEventStreamProgress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectObjectContentEventStreamProgress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["Details"] = o.Details
	}
	return toSerialize, nil
}

type NullableSelectObjectContentEventStreamProgress struct {
	value *SelectObjectContentEventStreamProgress
	isSet bool
}

func (v NullableSelectObjectContentEventStreamProgress) Get() *SelectObjectContentEventStreamProgress {
	return v.value
}

func (v *NullableSelectObjectContentEventStreamProgress) Set(val *SelectObjectContentEventStreamProgress) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectObjectContentEventStreamProgress) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectObjectContentEventStreamProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectObjectContentEventStreamProgress(val *SelectObjectContentEventStreamProgress) *NullableSelectObjectContentEventStreamProgress {
	return &NullableSelectObjectContentEventStreamProgress{value: val, isSet: true}
}

func (v NullableSelectObjectContentEventStreamProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectObjectContentEventStreamProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

