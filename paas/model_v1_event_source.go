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

// checks if the V1EventSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EventSource{}

// V1EventSource EventSource contains information for an event.
type V1EventSource struct {
	// Node name on which the event is generated.
	Host *string `json:"host,omitempty"`
	// Component from which the event is generated.
	Component *string `json:"component,omitempty"`
}

// NewV1EventSource instantiates a new V1EventSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EventSource() *V1EventSource {
	this := V1EventSource{}
	return &this
}

// NewV1EventSourceWithDefaults instantiates a new V1EventSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EventSourceWithDefaults() *V1EventSource {
	this := V1EventSource{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *V1EventSource) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EventSource) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *V1EventSource) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *V1EventSource) SetHost(v string) {
	o.Host = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *V1EventSource) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EventSource) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *V1EventSource) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *V1EventSource) SetComponent(v string) {
	o.Component = &v
}

func (o V1EventSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EventSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	return toSerialize, nil
}

type NullableV1EventSource struct {
	value *V1EventSource
	isSet bool
}

func (v NullableV1EventSource) Get() *V1EventSource {
	return v.value
}

func (v *NullableV1EventSource) Set(val *V1EventSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EventSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EventSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EventSource(val *V1EventSource) *NullableV1EventSource {
	return &NullableV1EventSource{value: val, isSet: true}
}

func (v NullableV1EventSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EventSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


