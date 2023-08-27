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

// checks if the AttackReportMap type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportMap{}

// AttackReportMap struct for AttackReportMap
type AttackReportMap struct {
	// Deprecated
	Statistics []AttackReportMapStatisticsInner `json:"statistics,omitempty"`
	Charts *AttackReportMapCharts `json:"charts,omitempty"`
	Lists []AttackReportMapStatisticsInner `json:"lists,omitempty"`
}

// NewAttackReportMap instantiates a new AttackReportMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportMap() *AttackReportMap {
	this := AttackReportMap{}
	return &this
}

// NewAttackReportMapWithDefaults instantiates a new AttackReportMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportMapWithDefaults() *AttackReportMap {
	this := AttackReportMap{}
	return &this
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
// Deprecated
func (o *AttackReportMap) GetStatistics() []AttackReportMapStatisticsInner {
	if o == nil || IsNil(o.Statistics) {
		var ret []AttackReportMapStatisticsInner
		return ret
	}
	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AttackReportMap) GetStatisticsOk() ([]AttackReportMapStatisticsInner, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *AttackReportMap) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given []AttackReportMapStatisticsInner and assigns it to the Statistics field.
// Deprecated
func (o *AttackReportMap) SetStatistics(v []AttackReportMapStatisticsInner) {
	o.Statistics = v
}

// GetCharts returns the Charts field value if set, zero value otherwise.
func (o *AttackReportMap) GetCharts() AttackReportMapCharts {
	if o == nil || IsNil(o.Charts) {
		var ret AttackReportMapCharts
		return ret
	}
	return *o.Charts
}

// GetChartsOk returns a tuple with the Charts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMap) GetChartsOk() (*AttackReportMapCharts, bool) {
	if o == nil || IsNil(o.Charts) {
		return nil, false
	}
	return o.Charts, true
}

// HasCharts returns a boolean if a field has been set.
func (o *AttackReportMap) HasCharts() bool {
	if o != nil && !IsNil(o.Charts) {
		return true
	}

	return false
}

// SetCharts gets a reference to the given AttackReportMapCharts and assigns it to the Charts field.
func (o *AttackReportMap) SetCharts(v AttackReportMapCharts) {
	o.Charts = &v
}

// GetLists returns the Lists field value if set, zero value otherwise.
func (o *AttackReportMap) GetLists() []AttackReportMapStatisticsInner {
	if o == nil || IsNil(o.Lists) {
		var ret []AttackReportMapStatisticsInner
		return ret
	}
	return o.Lists
}

// GetListsOk returns a tuple with the Lists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMap) GetListsOk() ([]AttackReportMapStatisticsInner, bool) {
	if o == nil || IsNil(o.Lists) {
		return nil, false
	}
	return o.Lists, true
}

// HasLists returns a boolean if a field has been set.
func (o *AttackReportMap) HasLists() bool {
	if o != nil && !IsNil(o.Lists) {
		return true
	}

	return false
}

// SetLists gets a reference to the given []AttackReportMapStatisticsInner and assigns it to the Lists field.
func (o *AttackReportMap) SetLists(v []AttackReportMapStatisticsInner) {
	o.Lists = v
}

func (o AttackReportMap) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportMap) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !IsNil(o.Charts) {
		toSerialize["charts"] = o.Charts
	}
	if !IsNil(o.Lists) {
		toSerialize["lists"] = o.Lists
	}
	return toSerialize, nil
}

type NullableAttackReportMap struct {
	value *AttackReportMap
	isSet bool
}

func (v NullableAttackReportMap) Get() *AttackReportMap {
	return v.value
}

func (v *NullableAttackReportMap) Set(val *AttackReportMap) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportMap) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportMap(val *AttackReportMap) *NullableAttackReportMap {
	return &NullableAttackReportMap{value: val, isSet: true}
}

func (v NullableAttackReportMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


