/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the ServerAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerAddress{}

// ServerAddress struct for ServerAddress
type ServerAddress struct {
	Addr *string `json:"addr,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	MacAddr *string `json:"mac_addr,omitempty"`
	Type *string `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewServerAddress instantiates a new ServerAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerAddress() *ServerAddress {
	this := ServerAddress{}
	return &this
}

// NewServerAddressWithDefaults instantiates a new ServerAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerAddressWithDefaults() *ServerAddress {
	this := ServerAddress{}
	return &this
}

// GetAddr returns the Addr field value if set, zero value otherwise.
func (o *ServerAddress) GetAddr() string {
	if o == nil || IsNil(o.Addr) {
		var ret string
		return ret
	}
	return *o.Addr
}

// GetAddrOk returns a tuple with the Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAddress) GetAddrOk() (*string, bool) {
	if o == nil || IsNil(o.Addr) {
		return nil, false
	}
	return o.Addr, true
}

// HasAddr returns a boolean if a field has been set.
func (o *ServerAddress) HasAddr() bool {
	if o != nil && !IsNil(o.Addr) {
		return true
	}

	return false
}

// SetAddr gets a reference to the given string and assigns it to the Addr field.
func (o *ServerAddress) SetAddr(v string) {
	o.Addr = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *ServerAddress) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAddress) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *ServerAddress) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *ServerAddress) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetMacAddr returns the MacAddr field value if set, zero value otherwise.
func (o *ServerAddress) GetMacAddr() string {
	if o == nil || IsNil(o.MacAddr) {
		var ret string
		return ret
	}
	return *o.MacAddr
}

// GetMacAddrOk returns a tuple with the MacAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAddress) GetMacAddrOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddr) {
		return nil, false
	}
	return o.MacAddr, true
}

// HasMacAddr returns a boolean if a field has been set.
func (o *ServerAddress) HasMacAddr() bool {
	if o != nil && !IsNil(o.MacAddr) {
		return true
	}

	return false
}

// SetMacAddr gets a reference to the given string and assigns it to the MacAddr field.
func (o *ServerAddress) SetMacAddr(v string) {
	o.MacAddr = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServerAddress) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAddress) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServerAddress) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServerAddress) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServerAddress) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerAddress) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServerAddress) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ServerAddress) SetVersion(v string) {
	o.Version = &v
}

func (o ServerAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addr) {
		toSerialize["addr"] = o.Addr
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	if !IsNil(o.MacAddr) {
		toSerialize["mac_addr"] = o.MacAddr
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableServerAddress struct {
	value *ServerAddress
	isSet bool
}

func (v NullableServerAddress) Get() *ServerAddress {
	return v.value
}

func (v *NullableServerAddress) Set(val *ServerAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableServerAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableServerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerAddress(val *ServerAddress) *NullableServerAddress {
	return &NullableServerAddress{value: val, isSet: true}
}

func (v NullableServerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


