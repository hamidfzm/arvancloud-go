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

// checks if the PutBucketLoggingRequestBucketLoggingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketLoggingRequestBucketLoggingStatus{}

// PutBucketLoggingRequestBucketLoggingStatus struct for PutBucketLoggingRequestBucketLoggingStatus
type PutBucketLoggingRequestBucketLoggingStatus struct {
	LoggingEnabled *LoggingEnabled `json:"LoggingEnabled,omitempty"`
}

// NewPutBucketLoggingRequestBucketLoggingStatus instantiates a new PutBucketLoggingRequestBucketLoggingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketLoggingRequestBucketLoggingStatus() *PutBucketLoggingRequestBucketLoggingStatus {
	this := PutBucketLoggingRequestBucketLoggingStatus{}
	return &this
}

// NewPutBucketLoggingRequestBucketLoggingStatusWithDefaults instantiates a new PutBucketLoggingRequestBucketLoggingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketLoggingRequestBucketLoggingStatusWithDefaults() *PutBucketLoggingRequestBucketLoggingStatus {
	this := PutBucketLoggingRequestBucketLoggingStatus{}
	return &this
}

// GetLoggingEnabled returns the LoggingEnabled field value if set, zero value otherwise.
func (o *PutBucketLoggingRequestBucketLoggingStatus) GetLoggingEnabled() LoggingEnabled {
	if o == nil || IsNil(o.LoggingEnabled) {
		var ret LoggingEnabled
		return ret
	}
	return *o.LoggingEnabled
}

// GetLoggingEnabledOk returns a tuple with the LoggingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBucketLoggingRequestBucketLoggingStatus) GetLoggingEnabledOk() (*LoggingEnabled, bool) {
	if o == nil || IsNil(o.LoggingEnabled) {
		return nil, false
	}
	return o.LoggingEnabled, true
}

// HasLoggingEnabled returns a boolean if a field has been set.
func (o *PutBucketLoggingRequestBucketLoggingStatus) HasLoggingEnabled() bool {
	if o != nil && !IsNil(o.LoggingEnabled) {
		return true
	}

	return false
}

// SetLoggingEnabled gets a reference to the given LoggingEnabled and assigns it to the LoggingEnabled field.
func (o *PutBucketLoggingRequestBucketLoggingStatus) SetLoggingEnabled(v LoggingEnabled) {
	o.LoggingEnabled = &v
}

func (o PutBucketLoggingRequestBucketLoggingStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketLoggingRequestBucketLoggingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoggingEnabled) {
		toSerialize["LoggingEnabled"] = o.LoggingEnabled
	}
	return toSerialize, nil
}

type NullablePutBucketLoggingRequestBucketLoggingStatus struct {
	value *PutBucketLoggingRequestBucketLoggingStatus
	isSet bool
}

func (v NullablePutBucketLoggingRequestBucketLoggingStatus) Get() *PutBucketLoggingRequestBucketLoggingStatus {
	return v.value
}

func (v *NullablePutBucketLoggingRequestBucketLoggingStatus) Set(val *PutBucketLoggingRequestBucketLoggingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketLoggingRequestBucketLoggingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketLoggingRequestBucketLoggingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketLoggingRequestBucketLoggingStatus(val *PutBucketLoggingRequestBucketLoggingStatus) *NullablePutBucketLoggingRequestBucketLoggingStatus {
	return &NullablePutBucketLoggingRequestBucketLoggingStatus{value: val, isSet: true}
}

func (v NullablePutBucketLoggingRequestBucketLoggingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketLoggingRequestBucketLoggingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


