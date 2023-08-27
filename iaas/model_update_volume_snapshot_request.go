/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the UpdateVolumeSnapshotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVolumeSnapshotRequest{}

// UpdateVolumeSnapshotRequest struct for UpdateVolumeSnapshotRequest
type UpdateVolumeSnapshotRequest struct {
	Description *string `json:"description,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	Name *string `json:"name,omitempty"`
	VolumeName *string `json:"volume_name,omitempty"`
}

// NewUpdateVolumeSnapshotRequest instantiates a new UpdateVolumeSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVolumeSnapshotRequest() *UpdateVolumeSnapshotRequest {
	this := UpdateVolumeSnapshotRequest{}
	return &this
}

// NewUpdateVolumeSnapshotRequestWithDefaults instantiates a new UpdateVolumeSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVolumeSnapshotRequestWithDefaults() *UpdateVolumeSnapshotRequest {
	this := UpdateVolumeSnapshotRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVolumeSnapshotRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeSnapshotRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVolumeSnapshotRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVolumeSnapshotRequest) SetDescription(v string) {
	o.Description = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *UpdateVolumeSnapshotRequest) GetInstanceName() string {
	if o == nil || IsNil(o.InstanceName) {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeSnapshotRequest) GetInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceName) {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *UpdateVolumeSnapshotRequest) HasInstanceName() bool {
	if o != nil && !IsNil(o.InstanceName) {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *UpdateVolumeSnapshotRequest) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVolumeSnapshotRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeSnapshotRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVolumeSnapshotRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVolumeSnapshotRequest) SetName(v string) {
	o.Name = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *UpdateVolumeSnapshotRequest) GetVolumeName() string {
	if o == nil || IsNil(o.VolumeName) {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeSnapshotRequest) GetVolumeNameOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeName) {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *UpdateVolumeSnapshotRequest) HasVolumeName() bool {
	if o != nil && !IsNil(o.VolumeName) {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *UpdateVolumeSnapshotRequest) SetVolumeName(v string) {
	o.VolumeName = &v
}

func (o UpdateVolumeSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVolumeSnapshotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InstanceName) {
		toSerialize["instance_name"] = o.InstanceName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.VolumeName) {
		toSerialize["volume_name"] = o.VolumeName
	}
	return toSerialize, nil
}

type NullableUpdateVolumeSnapshotRequest struct {
	value *UpdateVolumeSnapshotRequest
	isSet bool
}

func (v NullableUpdateVolumeSnapshotRequest) Get() *UpdateVolumeSnapshotRequest {
	return v.value
}

func (v *NullableUpdateVolumeSnapshotRequest) Set(val *UpdateVolumeSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVolumeSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVolumeSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVolumeSnapshotRequest(val *UpdateVolumeSnapshotRequest) *NullableUpdateVolumeSnapshotRequest {
	return &NullableUpdateVolumeSnapshotRequest{value: val, isSet: true}
}

func (v NullableUpdateVolumeSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVolumeSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


