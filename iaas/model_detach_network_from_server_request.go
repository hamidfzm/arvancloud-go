/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the DetachNetworkFromServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetachNetworkFromServerRequest{}

// DetachNetworkFromServerRequest struct for DetachNetworkFromServerRequest
type DetachNetworkFromServerRequest struct {
	ServerId *string `json:"server_id,omitempty"`
}

// NewDetachNetworkFromServerRequest instantiates a new DetachNetworkFromServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetachNetworkFromServerRequest() *DetachNetworkFromServerRequest {
	this := DetachNetworkFromServerRequest{}
	return &this
}

// NewDetachNetworkFromServerRequestWithDefaults instantiates a new DetachNetworkFromServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetachNetworkFromServerRequestWithDefaults() *DetachNetworkFromServerRequest {
	this := DetachNetworkFromServerRequest{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *DetachNetworkFromServerRequest) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetachNetworkFromServerRequest) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *DetachNetworkFromServerRequest) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *DetachNetworkFromServerRequest) SetServerId(v string) {
	o.ServerId = &v
}

func (o DetachNetworkFromServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetachNetworkFromServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerId) {
		toSerialize["server_id"] = o.ServerId
	}
	return toSerialize, nil
}

type NullableDetachNetworkFromServerRequest struct {
	value *DetachNetworkFromServerRequest
	isSet bool
}

func (v NullableDetachNetworkFromServerRequest) Get() *DetachNetworkFromServerRequest {
	return v.value
}

func (v *NullableDetachNetworkFromServerRequest) Set(val *DetachNetworkFromServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDetachNetworkFromServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDetachNetworkFromServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetachNetworkFromServerRequest(val *DetachNetworkFromServerRequest) *NullableDetachNetworkFromServerRequest {
	return &NullableDetachNetworkFromServerRequest{value: val, isSet: true}
}

func (v NullableDetachNetworkFromServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetachNetworkFromServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


