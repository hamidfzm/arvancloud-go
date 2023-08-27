/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the FloatIPAttachRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FloatIPAttachRequest{}

// FloatIPAttachRequest struct for FloatIPAttachRequest
type FloatIPAttachRequest struct {
	PortId *string `json:"port_id,omitempty"`
	ServerId *string `json:"server_id,omitempty"`
	SubnetId *string `json:"subnet_id,omitempty"`
}

// NewFloatIPAttachRequest instantiates a new FloatIPAttachRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFloatIPAttachRequest() *FloatIPAttachRequest {
	this := FloatIPAttachRequest{}
	return &this
}

// NewFloatIPAttachRequestWithDefaults instantiates a new FloatIPAttachRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFloatIPAttachRequestWithDefaults() *FloatIPAttachRequest {
	this := FloatIPAttachRequest{}
	return &this
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *FloatIPAttachRequest) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FloatIPAttachRequest) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *FloatIPAttachRequest) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *FloatIPAttachRequest) SetPortId(v string) {
	o.PortId = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *FloatIPAttachRequest) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FloatIPAttachRequest) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *FloatIPAttachRequest) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *FloatIPAttachRequest) SetServerId(v string) {
	o.ServerId = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *FloatIPAttachRequest) GetSubnetId() string {
	if o == nil || IsNil(o.SubnetId) {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FloatIPAttachRequest) GetSubnetIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetId) {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *FloatIPAttachRequest) HasSubnetId() bool {
	if o != nil && !IsNil(o.SubnetId) {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *FloatIPAttachRequest) SetSubnetId(v string) {
	o.SubnetId = &v
}

func (o FloatIPAttachRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FloatIPAttachRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortId) {
		toSerialize["port_id"] = o.PortId
	}
	if !IsNil(o.ServerId) {
		toSerialize["server_id"] = o.ServerId
	}
	if !IsNil(o.SubnetId) {
		toSerialize["subnet_id"] = o.SubnetId
	}
	return toSerialize, nil
}

type NullableFloatIPAttachRequest struct {
	value *FloatIPAttachRequest
	isSet bool
}

func (v NullableFloatIPAttachRequest) Get() *FloatIPAttachRequest {
	return v.value
}

func (v *NullableFloatIPAttachRequest) Set(val *FloatIPAttachRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFloatIPAttachRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFloatIPAttachRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFloatIPAttachRequest(val *FloatIPAttachRequest) *NullableFloatIPAttachRequest {
	return &NullableFloatIPAttachRequest{value: val, isSet: true}
}

func (v NullableFloatIPAttachRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFloatIPAttachRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


