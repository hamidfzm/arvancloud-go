/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the UpdateVolumeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVolumeRequest{}

// UpdateVolumeRequest struct for UpdateVolumeRequest
type UpdateVolumeRequest struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewUpdateVolumeRequest instantiates a new UpdateVolumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVolumeRequest() *UpdateVolumeRequest {
	this := UpdateVolumeRequest{}
	return &this
}

// NewUpdateVolumeRequestWithDefaults instantiates a new UpdateVolumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVolumeRequestWithDefaults() *UpdateVolumeRequest {
	this := UpdateVolumeRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVolumeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVolumeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVolumeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVolumeRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVolumeRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVolumeRequest) SetName(v string) {
	o.Name = &v
}

func (o UpdateVolumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVolumeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdateVolumeRequest struct {
	value *UpdateVolumeRequest
	isSet bool
}

func (v NullableUpdateVolumeRequest) Get() *UpdateVolumeRequest {
	return v.value
}

func (v *NullableUpdateVolumeRequest) Set(val *UpdateVolumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVolumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVolumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVolumeRequest(val *UpdateVolumeRequest) *NullableUpdateVolumeRequest {
	return &NullableUpdateVolumeRequest{value: val, isSet: true}
}

func (v NullableUpdateVolumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVolumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

