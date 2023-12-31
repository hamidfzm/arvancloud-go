/*
Arvancloud PaaS REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paas

import (
	"encoding/json"
)

// checks if the V1EndpointAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EndpointAddress{}

// V1EndpointAddress EndpointAddress is a tuple that describes single IP address.
type V1EndpointAddress struct {
	// The Hostname of this endpoint
	Hostname *string `json:"hostname,omitempty"`
	// The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready.
	Ip string `json:"ip"`
	// Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.
	NodeName *string `json:"nodeName,omitempty"`
	TargetRef *V1ObjectReference `json:"targetRef,omitempty"`
}

// NewV1EndpointAddress instantiates a new V1EndpointAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EndpointAddress(ip string) *V1EndpointAddress {
	this := V1EndpointAddress{}
	this.Ip = ip
	return &this
}

// NewV1EndpointAddressWithDefaults instantiates a new V1EndpointAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EndpointAddressWithDefaults() *V1EndpointAddress {
	this := V1EndpointAddress{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *V1EndpointAddress) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EndpointAddress) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *V1EndpointAddress) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *V1EndpointAddress) SetHostname(v string) {
	o.Hostname = &v
}

// GetIp returns the Ip field value
func (o *V1EndpointAddress) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *V1EndpointAddress) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *V1EndpointAddress) SetIp(v string) {
	o.Ip = v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *V1EndpointAddress) GetNodeName() string {
	if o == nil || IsNil(o.NodeName) {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EndpointAddress) GetNodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.NodeName) {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *V1EndpointAddress) HasNodeName() bool {
	if o != nil && !IsNil(o.NodeName) {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *V1EndpointAddress) SetNodeName(v string) {
	o.NodeName = &v
}

// GetTargetRef returns the TargetRef field value if set, zero value otherwise.
func (o *V1EndpointAddress) GetTargetRef() V1ObjectReference {
	if o == nil || IsNil(o.TargetRef) {
		var ret V1ObjectReference
		return ret
	}
	return *o.TargetRef
}

// GetTargetRefOk returns a tuple with the TargetRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EndpointAddress) GetTargetRefOk() (*V1ObjectReference, bool) {
	if o == nil || IsNil(o.TargetRef) {
		return nil, false
	}
	return o.TargetRef, true
}

// HasTargetRef returns a boolean if a field has been set.
func (o *V1EndpointAddress) HasTargetRef() bool {
	if o != nil && !IsNil(o.TargetRef) {
		return true
	}

	return false
}

// SetTargetRef gets a reference to the given V1ObjectReference and assigns it to the TargetRef field.
func (o *V1EndpointAddress) SetTargetRef(v V1ObjectReference) {
	o.TargetRef = &v
}

func (o V1EndpointAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EndpointAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	toSerialize["ip"] = o.Ip
	if !IsNil(o.NodeName) {
		toSerialize["nodeName"] = o.NodeName
	}
	if !IsNil(o.TargetRef) {
		toSerialize["targetRef"] = o.TargetRef
	}
	return toSerialize, nil
}

type NullableV1EndpointAddress struct {
	value *V1EndpointAddress
	isSet bool
}

func (v NullableV1EndpointAddress) Get() *V1EndpointAddress {
	return v.value
}

func (v *NullableV1EndpointAddress) Set(val *V1EndpointAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EndpointAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EndpointAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EndpointAddress(val *V1EndpointAddress) *NullableV1EndpointAddress {
	return &NullableV1EndpointAddress{value: val, isSet: true}
}

func (v NullableV1EndpointAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EndpointAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


