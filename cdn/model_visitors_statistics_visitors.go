/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
	"time"
)

// checks if the VisitorsStatisticsVisitors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisitorsStatisticsVisitors{}

// VisitorsStatisticsVisitors struct for VisitorsStatisticsVisitors
type VisitorsStatisticsVisitors struct {
	TopVisitors time.Time `json:"top_visitors"`
	TotalVisitors int32 `json:"total_visitors"`
}

// NewVisitorsStatisticsVisitors instantiates a new VisitorsStatisticsVisitors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisitorsStatisticsVisitors(topVisitors time.Time, totalVisitors int32) *VisitorsStatisticsVisitors {
	this := VisitorsStatisticsVisitors{}
	this.TopVisitors = topVisitors
	this.TotalVisitors = totalVisitors
	return &this
}

// NewVisitorsStatisticsVisitorsWithDefaults instantiates a new VisitorsStatisticsVisitors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisitorsStatisticsVisitorsWithDefaults() *VisitorsStatisticsVisitors {
	this := VisitorsStatisticsVisitors{}
	return &this
}

// GetTopVisitors returns the TopVisitors field value
func (o *VisitorsStatisticsVisitors) GetTopVisitors() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TopVisitors
}

// GetTopVisitorsOk returns a tuple with the TopVisitors field value
// and a boolean to check if the value has been set.
func (o *VisitorsStatisticsVisitors) GetTopVisitorsOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopVisitors, true
}

// SetTopVisitors sets field value
func (o *VisitorsStatisticsVisitors) SetTopVisitors(v time.Time) {
	o.TopVisitors = v
}

// GetTotalVisitors returns the TotalVisitors field value
func (o *VisitorsStatisticsVisitors) GetTotalVisitors() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalVisitors
}

// GetTotalVisitorsOk returns a tuple with the TotalVisitors field value
// and a boolean to check if the value has been set.
func (o *VisitorsStatisticsVisitors) GetTotalVisitorsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalVisitors, true
}

// SetTotalVisitors sets field value
func (o *VisitorsStatisticsVisitors) SetTotalVisitors(v int32) {
	o.TotalVisitors = v
}

func (o VisitorsStatisticsVisitors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisitorsStatisticsVisitors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["top_visitors"] = o.TopVisitors
	toSerialize["total_visitors"] = o.TotalVisitors
	return toSerialize, nil
}

type NullableVisitorsStatisticsVisitors struct {
	value *VisitorsStatisticsVisitors
	isSet bool
}

func (v NullableVisitorsStatisticsVisitors) Get() *VisitorsStatisticsVisitors {
	return v.value
}

func (v *NullableVisitorsStatisticsVisitors) Set(val *VisitorsStatisticsVisitors) {
	v.value = val
	v.isSet = true
}

func (v NullableVisitorsStatisticsVisitors) IsSet() bool {
	return v.isSet
}

func (v *NullableVisitorsStatisticsVisitors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisitorsStatisticsVisitors(val *VisitorsStatisticsVisitors) *NullableVisitorsStatisticsVisitors {
	return &NullableVisitorsStatisticsVisitors{value: val, isSet: true}
}

func (v NullableVisitorsStatisticsVisitors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisitorsStatisticsVisitors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


