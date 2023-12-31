/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the DetachFloatIPRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetachFloatIPRequest{}

// DetachFloatIPRequest struct for DetachFloatIPRequest
type DetachFloatIPRequest struct {
	PortId *string `json:"port_id,omitempty"`
}

// NewDetachFloatIPRequest instantiates a new DetachFloatIPRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetachFloatIPRequest() *DetachFloatIPRequest {
	this := DetachFloatIPRequest{}
	return &this
}

// NewDetachFloatIPRequestWithDefaults instantiates a new DetachFloatIPRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetachFloatIPRequestWithDefaults() *DetachFloatIPRequest {
	this := DetachFloatIPRequest{}
	return &this
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *DetachFloatIPRequest) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetachFloatIPRequest) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *DetachFloatIPRequest) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *DetachFloatIPRequest) SetPortId(v string) {
	o.PortId = &v
}

func (o DetachFloatIPRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetachFloatIPRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortId) {
		toSerialize["port_id"] = o.PortId
	}
	return toSerialize, nil
}

type NullableDetachFloatIPRequest struct {
	value *DetachFloatIPRequest
	isSet bool
}

func (v NullableDetachFloatIPRequest) Get() *DetachFloatIPRequest {
	return v.value
}

func (v *NullableDetachFloatIPRequest) Set(val *DetachFloatIPRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDetachFloatIPRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDetachFloatIPRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetachFloatIPRequest(val *DetachFloatIPRequest) *NullableDetachFloatIPRequest {
	return &NullableDetachFloatIPRequest{value: val, isSet: true}
}

func (v NullableDetachFloatIPRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetachFloatIPRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


