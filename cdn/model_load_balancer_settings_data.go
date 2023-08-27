/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the LoadBalancerSettingsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerSettingsData{}

// LoadBalancerSettingsData struct for LoadBalancerSettingsData
type LoadBalancerSettingsData struct {
	Data *LoadBalancerSetting `json:"data,omitempty"`
}

// NewLoadBalancerSettingsData instantiates a new LoadBalancerSettingsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerSettingsData() *LoadBalancerSettingsData {
	this := LoadBalancerSettingsData{}
	return &this
}

// NewLoadBalancerSettingsDataWithDefaults instantiates a new LoadBalancerSettingsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerSettingsDataWithDefaults() *LoadBalancerSettingsData {
	this := LoadBalancerSettingsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoadBalancerSettingsData) GetData() LoadBalancerSetting {
	if o == nil || IsNil(o.Data) {
		var ret LoadBalancerSetting
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerSettingsData) GetDataOk() (*LoadBalancerSetting, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoadBalancerSettingsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LoadBalancerSetting and assigns it to the Data field.
func (o *LoadBalancerSettingsData) SetData(v LoadBalancerSetting) {
	o.Data = &v
}

func (o LoadBalancerSettingsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerSettingsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLoadBalancerSettingsData struct {
	value *LoadBalancerSettingsData
	isSet bool
}

func (v NullableLoadBalancerSettingsData) Get() *LoadBalancerSettingsData {
	return v.value
}

func (v *NullableLoadBalancerSettingsData) Set(val *LoadBalancerSettingsData) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerSettingsData) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerSettingsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerSettingsData(val *LoadBalancerSettingsData) *NullableLoadBalancerSettingsData {
	return &NullableLoadBalancerSettingsData{value: val, isSet: true}
}

func (v NullableLoadBalancerSettingsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerSettingsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


