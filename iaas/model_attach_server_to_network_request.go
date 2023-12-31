/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the AttachServerToNetworkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachServerToNetworkRequest{}

// AttachServerToNetworkRequest struct for AttachServerToNetworkRequest
type AttachServerToNetworkRequest struct {
	EnablePortSecurity *bool `json:"enablePortSecurity,omitempty"`
	Ip *string `json:"ip,omitempty"`
	ServerId *string `json:"server_id,omitempty"`
	SubnetId *string `json:"subnet_id,omitempty"`
}

// NewAttachServerToNetworkRequest instantiates a new AttachServerToNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachServerToNetworkRequest() *AttachServerToNetworkRequest {
	this := AttachServerToNetworkRequest{}
	return &this
}

// NewAttachServerToNetworkRequestWithDefaults instantiates a new AttachServerToNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachServerToNetworkRequestWithDefaults() *AttachServerToNetworkRequest {
	this := AttachServerToNetworkRequest{}
	return &this
}

// GetEnablePortSecurity returns the EnablePortSecurity field value if set, zero value otherwise.
func (o *AttachServerToNetworkRequest) GetEnablePortSecurity() bool {
	if o == nil || IsNil(o.EnablePortSecurity) {
		var ret bool
		return ret
	}
	return *o.EnablePortSecurity
}

// GetEnablePortSecurityOk returns a tuple with the EnablePortSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachServerToNetworkRequest) GetEnablePortSecurityOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePortSecurity) {
		return nil, false
	}
	return o.EnablePortSecurity, true
}

// HasEnablePortSecurity returns a boolean if a field has been set.
func (o *AttachServerToNetworkRequest) HasEnablePortSecurity() bool {
	if o != nil && !IsNil(o.EnablePortSecurity) {
		return true
	}

	return false
}

// SetEnablePortSecurity gets a reference to the given bool and assigns it to the EnablePortSecurity field.
func (o *AttachServerToNetworkRequest) SetEnablePortSecurity(v bool) {
	o.EnablePortSecurity = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *AttachServerToNetworkRequest) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachServerToNetworkRequest) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *AttachServerToNetworkRequest) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *AttachServerToNetworkRequest) SetIp(v string) {
	o.Ip = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *AttachServerToNetworkRequest) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachServerToNetworkRequest) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *AttachServerToNetworkRequest) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *AttachServerToNetworkRequest) SetServerId(v string) {
	o.ServerId = &v
}

// GetSubnetId returns the SubnetId field value if set, zero value otherwise.
func (o *AttachServerToNetworkRequest) GetSubnetId() string {
	if o == nil || IsNil(o.SubnetId) {
		var ret string
		return ret
	}
	return *o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachServerToNetworkRequest) GetSubnetIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubnetId) {
		return nil, false
	}
	return o.SubnetId, true
}

// HasSubnetId returns a boolean if a field has been set.
func (o *AttachServerToNetworkRequest) HasSubnetId() bool {
	if o != nil && !IsNil(o.SubnetId) {
		return true
	}

	return false
}

// SetSubnetId gets a reference to the given string and assigns it to the SubnetId field.
func (o *AttachServerToNetworkRequest) SetSubnetId(v string) {
	o.SubnetId = &v
}

func (o AttachServerToNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachServerToNetworkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnablePortSecurity) {
		toSerialize["enablePortSecurity"] = o.EnablePortSecurity
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.ServerId) {
		toSerialize["server_id"] = o.ServerId
	}
	if !IsNil(o.SubnetId) {
		toSerialize["subnet_id"] = o.SubnetId
	}
	return toSerialize, nil
}

type NullableAttachServerToNetworkRequest struct {
	value *AttachServerToNetworkRequest
	isSet bool
}

func (v NullableAttachServerToNetworkRequest) Get() *AttachServerToNetworkRequest {
	return v.value
}

func (v *NullableAttachServerToNetworkRequest) Set(val *AttachServerToNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachServerToNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachServerToNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachServerToNetworkRequest(val *AttachServerToNetworkRequest) *NullableAttachServerToNetworkRequest {
	return &NullableAttachServerToNetworkRequest{value: val, isSet: true}
}

func (v NullableAttachServerToNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachServerToNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


