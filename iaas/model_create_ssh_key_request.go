/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the CreateSSHKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSSHKeyRequest{}

// CreateSSHKeyRequest struct for CreateSSHKeyRequest
type CreateSSHKeyRequest struct {
	Name *string `json:"name,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
}

// NewCreateSSHKeyRequest instantiates a new CreateSSHKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSSHKeyRequest() *CreateSSHKeyRequest {
	this := CreateSSHKeyRequest{}
	return &this
}

// NewCreateSSHKeyRequestWithDefaults instantiates a new CreateSSHKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSSHKeyRequestWithDefaults() *CreateSSHKeyRequest {
	this := CreateSSHKeyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateSSHKeyRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHKeyRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateSSHKeyRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateSSHKeyRequest) SetName(v string) {
	o.Name = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *CreateSSHKeyRequest) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSSHKeyRequest) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *CreateSSHKeyRequest) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *CreateSSHKeyRequest) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o CreateSSHKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSSHKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	return toSerialize, nil
}

type NullableCreateSSHKeyRequest struct {
	value *CreateSSHKeyRequest
	isSet bool
}

func (v NullableCreateSSHKeyRequest) Get() *CreateSSHKeyRequest {
	return v.value
}

func (v *NullableCreateSSHKeyRequest) Set(val *CreateSSHKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSSHKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSSHKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSSHKeyRequest(val *CreateSSHKeyRequest) *NullableCreateSSHKeyRequest {
	return &NullableCreateSSHKeyRequest{value: val, isSet: true}
}

func (v NullableCreateSSHKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSSHKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


