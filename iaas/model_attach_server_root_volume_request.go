/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the AttachServerRootVolumeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachServerRootVolumeRequest{}

// AttachServerRootVolumeRequest struct for AttachServerRootVolumeRequest
type AttachServerRootVolumeRequest struct {
	VolumeId *string `json:"volume_id,omitempty"`
}

// NewAttachServerRootVolumeRequest instantiates a new AttachServerRootVolumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachServerRootVolumeRequest() *AttachServerRootVolumeRequest {
	this := AttachServerRootVolumeRequest{}
	return &this
}

// NewAttachServerRootVolumeRequestWithDefaults instantiates a new AttachServerRootVolumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachServerRootVolumeRequestWithDefaults() *AttachServerRootVolumeRequest {
	this := AttachServerRootVolumeRequest{}
	return &this
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *AttachServerRootVolumeRequest) GetVolumeId() string {
	if o == nil || IsNil(o.VolumeId) {
		var ret string
		return ret
	}
	return *o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachServerRootVolumeRequest) GetVolumeIdOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeId) {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *AttachServerRootVolumeRequest) HasVolumeId() bool {
	if o != nil && !IsNil(o.VolumeId) {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *AttachServerRootVolumeRequest) SetVolumeId(v string) {
	o.VolumeId = &v
}

func (o AttachServerRootVolumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachServerRootVolumeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VolumeId) {
		toSerialize["volume_id"] = o.VolumeId
	}
	return toSerialize, nil
}

type NullableAttachServerRootVolumeRequest struct {
	value *AttachServerRootVolumeRequest
	isSet bool
}

func (v NullableAttachServerRootVolumeRequest) Get() *AttachServerRootVolumeRequest {
	return v.value
}

func (v *NullableAttachServerRootVolumeRequest) Set(val *AttachServerRootVolumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachServerRootVolumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachServerRootVolumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachServerRootVolumeRequest(val *AttachServerRootVolumeRequest) *NullableAttachServerRootVolumeRequest {
	return &NullableAttachServerRootVolumeRequest{value: val, isSet: true}
}

func (v NullableAttachServerRootVolumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachServerRootVolumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

