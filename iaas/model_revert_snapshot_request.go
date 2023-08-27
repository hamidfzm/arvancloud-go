/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the RevertSnapshotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevertSnapshotRequest{}

// RevertSnapshotRequest struct for RevertSnapshotRequest
type RevertSnapshotRequest struct {
	SnapshotId *string `json:"snapshot_id,omitempty"`
}

// NewRevertSnapshotRequest instantiates a new RevertSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevertSnapshotRequest() *RevertSnapshotRequest {
	this := RevertSnapshotRequest{}
	return &this
}

// NewRevertSnapshotRequestWithDefaults instantiates a new RevertSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevertSnapshotRequestWithDefaults() *RevertSnapshotRequest {
	this := RevertSnapshotRequest{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *RevertSnapshotRequest) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevertSnapshotRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *RevertSnapshotRequest) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *RevertSnapshotRequest) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

func (o RevertSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevertSnapshotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	return toSerialize, nil
}

type NullableRevertSnapshotRequest struct {
	value *RevertSnapshotRequest
	isSet bool
}

func (v NullableRevertSnapshotRequest) Get() *RevertSnapshotRequest {
	return v.value
}

func (v *NullableRevertSnapshotRequest) Set(val *RevertSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRevertSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRevertSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevertSnapshotRequest(val *RevertSnapshotRequest) *NullableRevertSnapshotRequest {
	return &NullableRevertSnapshotRequest{value: val, isSet: true}
}

func (v NullableRevertSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevertSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

