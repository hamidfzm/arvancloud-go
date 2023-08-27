/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the NetworkServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkServer{}

// NetworkServer struct for NetworkServer
type NetworkServer struct {
	Addresses *map[string][]ServerAddress `json:"addresses,omitempty"`
	Id *string `json:"id,omitempty"`
	Ips []FullIP `json:"ips,omitempty"`
	Name *string `json:"name,omitempty"`
	PublicIp []PublicIP `json:"public_ip,omitempty"`
	SecurityGroups []string `json:"security_groups,omitempty"`
}

// NewNetworkServer instantiates a new NetworkServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkServer() *NetworkServer {
	this := NetworkServer{}
	return &this
}

// NewNetworkServerWithDefaults instantiates a new NetworkServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkServerWithDefaults() *NetworkServer {
	this := NetworkServer{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *NetworkServer) GetAddresses() map[string][]ServerAddress {
	if o == nil || IsNil(o.Addresses) {
		var ret map[string][]ServerAddress
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkServer) GetAddressesOk() (*map[string][]ServerAddress, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *NetworkServer) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given map[string][]ServerAddress and assigns it to the Addresses field.
func (o *NetworkServer) SetAddresses(v map[string][]ServerAddress) {
	o.Addresses = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkServer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkServer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkServer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkServer) SetId(v string) {
	o.Id = &v
}

// GetIps returns the Ips field value if set, zero value otherwise.
func (o *NetworkServer) GetIps() []FullIP {
	if o == nil || IsNil(o.Ips) {
		var ret []FullIP
		return ret
	}
	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkServer) GetIpsOk() ([]FullIP, bool) {
	if o == nil || IsNil(o.Ips) {
		return nil, false
	}
	return o.Ips, true
}

// HasIps returns a boolean if a field has been set.
func (o *NetworkServer) HasIps() bool {
	if o != nil && !IsNil(o.Ips) {
		return true
	}

	return false
}

// SetIps gets a reference to the given []FullIP and assigns it to the Ips field.
func (o *NetworkServer) SetIps(v []FullIP) {
	o.Ips = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkServer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkServer) SetName(v string) {
	o.Name = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *NetworkServer) GetPublicIp() []PublicIP {
	if o == nil || IsNil(o.PublicIp) {
		var ret []PublicIP
		return ret
	}
	return o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkServer) GetPublicIpOk() ([]PublicIP, bool) {
	if o == nil || IsNil(o.PublicIp) {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *NetworkServer) HasPublicIp() bool {
	if o != nil && !IsNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given []PublicIP and assigns it to the PublicIp field.
func (o *NetworkServer) SetPublicIp(v []PublicIP) {
	o.PublicIp = v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *NetworkServer) GetSecurityGroups() []string {
	if o == nil || IsNil(o.SecurityGroups) {
		var ret []string
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkServer) GetSecurityGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.SecurityGroups) {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *NetworkServer) HasSecurityGroups() bool {
	if o != nil && !IsNil(o.SecurityGroups) {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *NetworkServer) SetSecurityGroups(v []string) {
	o.SecurityGroups = v
}

func (o NetworkServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ips) {
		toSerialize["ips"] = o.Ips
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PublicIp) {
		toSerialize["public_ip"] = o.PublicIp
	}
	if !IsNil(o.SecurityGroups) {
		toSerialize["security_groups"] = o.SecurityGroups
	}
	return toSerialize, nil
}

type NullableNetworkServer struct {
	value *NetworkServer
	isSet bool
}

func (v NullableNetworkServer) Get() *NetworkServer {
	return v.value
}

func (v *NullableNetworkServer) Set(val *NetworkServer) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkServer) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkServer(val *NetworkServer) *NullableNetworkServer {
	return &NullableNetworkServer{value: val, isSet: true}
}

func (v NullableNetworkServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


