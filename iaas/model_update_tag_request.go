/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the UpdateTagRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTagRequest{}

// UpdateTagRequest struct for UpdateTagRequest
type UpdateTagRequest struct {
	TagName *string `json:"tag_name,omitempty"`
}

// NewUpdateTagRequest instantiates a new UpdateTagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTagRequest() *UpdateTagRequest {
	this := UpdateTagRequest{}
	return &this
}

// NewUpdateTagRequestWithDefaults instantiates a new UpdateTagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTagRequestWithDefaults() *UpdateTagRequest {
	this := UpdateTagRequest{}
	return &this
}

// GetTagName returns the TagName field value if set, zero value otherwise.
func (o *UpdateTagRequest) GetTagName() string {
	if o == nil || IsNil(o.TagName) {
		var ret string
		return ret
	}
	return *o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTagRequest) GetTagNameOk() (*string, bool) {
	if o == nil || IsNil(o.TagName) {
		return nil, false
	}
	return o.TagName, true
}

// HasTagName returns a boolean if a field has been set.
func (o *UpdateTagRequest) HasTagName() bool {
	if o != nil && !IsNil(o.TagName) {
		return true
	}

	return false
}

// SetTagName gets a reference to the given string and assigns it to the TagName field.
func (o *UpdateTagRequest) SetTagName(v string) {
	o.TagName = &v
}

func (o UpdateTagRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTagRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TagName) {
		toSerialize["tag_name"] = o.TagName
	}
	return toSerialize, nil
}

type NullableUpdateTagRequest struct {
	value *UpdateTagRequest
	isSet bool
}

func (v NullableUpdateTagRequest) Get() *UpdateTagRequest {
	return v.value
}

func (v *NullableUpdateTagRequest) Set(val *UpdateTagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTagRequest(val *UpdateTagRequest) *NullableUpdateTagRequest {
	return &NullableUpdateTagRequest{value: val, isSet: true}
}

func (v NullableUpdateTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

