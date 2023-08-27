/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the ServerReports type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerReports{}

// ServerReports struct for ServerReports
type ServerReports struct {
	Charts *ChartCollection `json:"charts,omitempty"`
}

// NewServerReports instantiates a new ServerReports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerReports() *ServerReports {
	this := ServerReports{}
	return &this
}

// NewServerReportsWithDefaults instantiates a new ServerReports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerReportsWithDefaults() *ServerReports {
	this := ServerReports{}
	return &this
}

// GetCharts returns the Charts field value if set, zero value otherwise.
func (o *ServerReports) GetCharts() ChartCollection {
	if o == nil || IsNil(o.Charts) {
		var ret ChartCollection
		return ret
	}
	return *o.Charts
}

// GetChartsOk returns a tuple with the Charts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerReports) GetChartsOk() (*ChartCollection, bool) {
	if o == nil || IsNil(o.Charts) {
		return nil, false
	}
	return o.Charts, true
}

// HasCharts returns a boolean if a field has been set.
func (o *ServerReports) HasCharts() bool {
	if o != nil && !IsNil(o.Charts) {
		return true
	}

	return false
}

// SetCharts gets a reference to the given ChartCollection and assigns it to the Charts field.
func (o *ServerReports) SetCharts(v ChartCollection) {
	o.Charts = &v
}

func (o ServerReports) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerReports) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Charts) {
		toSerialize["charts"] = o.Charts
	}
	return toSerialize, nil
}

type NullableServerReports struct {
	value *ServerReports
	isSet bool
}

func (v NullableServerReports) Get() *ServerReports {
	return v.value
}

func (v *NullableServerReports) Set(val *ServerReports) {
	v.value = val
	v.isSet = true
}

func (v NullableServerReports) IsSet() bool {
	return v.isSet
}

func (v *NullableServerReports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerReports(val *ServerReports) *NullableServerReports {
	return &NullableServerReports{value: val, isSet: true}
}

func (v NullableServerReports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerReports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

