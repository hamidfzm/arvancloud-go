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

// checks if the PutBucketVersioningRequestVersioningConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutBucketVersioningRequestVersioningConfiguration{}

// PutBucketVersioningRequestVersioningConfiguration struct for PutBucketVersioningRequestVersioningConfiguration
type PutBucketVersioningRequestVersioningConfiguration struct {
	MFADelete *MFADelete `json:"MFADelete,omitempty"`
	Status *BucketVersioningStatus `json:"Status,omitempty"`
}

// NewPutBucketVersioningRequestVersioningConfiguration instantiates a new PutBucketVersioningRequestVersioningConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutBucketVersioningRequestVersioningConfiguration() *PutBucketVersioningRequestVersioningConfiguration {
	this := PutBucketVersioningRequestVersioningConfiguration{}
	return &this
}

// NewPutBucketVersioningRequestVersioningConfigurationWithDefaults instantiates a new PutBucketVersioningRequestVersioningConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutBucketVersioningRequestVersioningConfigurationWithDefaults() *PutBucketVersioningRequestVersioningConfiguration {
	this := PutBucketVersioningRequestVersioningConfiguration{}
	return &this
}

// GetMFADelete returns the MFADelete field value if set, zero value otherwise.
func (o *PutBucketVersioningRequestVersioningConfiguration) GetMFADelete() MFADelete {
	if o == nil || IsNil(o.MFADelete) {
		var ret MFADelete
		return ret
	}
	return *o.MFADelete
}

// GetMFADeleteOk returns a tuple with the MFADelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBucketVersioningRequestVersioningConfiguration) GetMFADeleteOk() (*MFADelete, bool) {
	if o == nil || IsNil(o.MFADelete) {
		return nil, false
	}
	return o.MFADelete, true
}

// HasMFADelete returns a boolean if a field has been set.
func (o *PutBucketVersioningRequestVersioningConfiguration) HasMFADelete() bool {
	if o != nil && !IsNil(o.MFADelete) {
		return true
	}

	return false
}

// SetMFADelete gets a reference to the given MFADelete and assigns it to the MFADelete field.
func (o *PutBucketVersioningRequestVersioningConfiguration) SetMFADelete(v MFADelete) {
	o.MFADelete = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PutBucketVersioningRequestVersioningConfiguration) GetStatus() BucketVersioningStatus {
	if o == nil || IsNil(o.Status) {
		var ret BucketVersioningStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutBucketVersioningRequestVersioningConfiguration) GetStatusOk() (*BucketVersioningStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PutBucketVersioningRequestVersioningConfiguration) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BucketVersioningStatus and assigns it to the Status field.
func (o *PutBucketVersioningRequestVersioningConfiguration) SetStatus(v BucketVersioningStatus) {
	o.Status = &v
}

func (o PutBucketVersioningRequestVersioningConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutBucketVersioningRequestVersioningConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MFADelete) {
		toSerialize["MFADelete"] = o.MFADelete
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePutBucketVersioningRequestVersioningConfiguration struct {
	value *PutBucketVersioningRequestVersioningConfiguration
	isSet bool
}

func (v NullablePutBucketVersioningRequestVersioningConfiguration) Get() *PutBucketVersioningRequestVersioningConfiguration {
	return v.value
}

func (v *NullablePutBucketVersioningRequestVersioningConfiguration) Set(val *PutBucketVersioningRequestVersioningConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePutBucketVersioningRequestVersioningConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePutBucketVersioningRequestVersioningConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutBucketVersioningRequestVersioningConfiguration(val *PutBucketVersioningRequestVersioningConfiguration) *NullablePutBucketVersioningRequestVersioningConfiguration {
	return &NullablePutBucketVersioningRequestVersioningConfiguration{value: val, isSet: true}
}

func (v NullablePutBucketVersioningRequestVersioningConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutBucketVersioningRequestVersioningConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


