/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the AttachPortSecurityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachPortSecurityRequest{}

// AttachPortSecurityRequest struct for AttachPortSecurityRequest
type AttachPortSecurityRequest struct {
	SecurityGroupId *string `json:"security_group_id,omitempty"`
}

// NewAttachPortSecurityRequest instantiates a new AttachPortSecurityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachPortSecurityRequest() *AttachPortSecurityRequest {
	this := AttachPortSecurityRequest{}
	return &this
}

// NewAttachPortSecurityRequestWithDefaults instantiates a new AttachPortSecurityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachPortSecurityRequestWithDefaults() *AttachPortSecurityRequest {
	this := AttachPortSecurityRequest{}
	return &this
}

// GetSecurityGroupId returns the SecurityGroupId field value if set, zero value otherwise.
func (o *AttachPortSecurityRequest) GetSecurityGroupId() string {
	if o == nil || IsNil(o.SecurityGroupId) {
		var ret string
		return ret
	}
	return *o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachPortSecurityRequest) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityGroupId) {
		return nil, false
	}
	return o.SecurityGroupId, true
}

// HasSecurityGroupId returns a boolean if a field has been set.
func (o *AttachPortSecurityRequest) HasSecurityGroupId() bool {
	if o != nil && !IsNil(o.SecurityGroupId) {
		return true
	}

	return false
}

// SetSecurityGroupId gets a reference to the given string and assigns it to the SecurityGroupId field.
func (o *AttachPortSecurityRequest) SetSecurityGroupId(v string) {
	o.SecurityGroupId = &v
}

func (o AttachPortSecurityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachPortSecurityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecurityGroupId) {
		toSerialize["security_group_id"] = o.SecurityGroupId
	}
	return toSerialize, nil
}

type NullableAttachPortSecurityRequest struct {
	value *AttachPortSecurityRequest
	isSet bool
}

func (v NullableAttachPortSecurityRequest) Get() *AttachPortSecurityRequest {
	return v.value
}

func (v *NullableAttachPortSecurityRequest) Set(val *AttachPortSecurityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachPortSecurityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachPortSecurityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachPortSecurityRequest(val *AttachPortSecurityRequest) *NullableAttachPortSecurityRequest {
	return &NullableAttachPortSecurityRequest{value: val, isSet: true}
}

func (v NullableAttachPortSecurityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachPortSecurityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


