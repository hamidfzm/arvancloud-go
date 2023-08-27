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

// checks if the VisitorsStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisitorsStatistics{}

// VisitorsStatistics struct for VisitorsStatistics
type VisitorsStatistics struct {
	Visitors *VisitorsStatisticsVisitors `json:"visitors,omitempty"`
}

// NewVisitorsStatistics instantiates a new VisitorsStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisitorsStatistics() *VisitorsStatistics {
	this := VisitorsStatistics{}
	return &this
}

// NewVisitorsStatisticsWithDefaults instantiates a new VisitorsStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisitorsStatisticsWithDefaults() *VisitorsStatistics {
	this := VisitorsStatistics{}
	return &this
}

// GetVisitors returns the Visitors field value if set, zero value otherwise.
func (o *VisitorsStatistics) GetVisitors() VisitorsStatisticsVisitors {
	if o == nil || IsNil(o.Visitors) {
		var ret VisitorsStatisticsVisitors
		return ret
	}
	return *o.Visitors
}

// GetVisitorsOk returns a tuple with the Visitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisitorsStatistics) GetVisitorsOk() (*VisitorsStatisticsVisitors, bool) {
	if o == nil || IsNil(o.Visitors) {
		return nil, false
	}
	return o.Visitors, true
}

// HasVisitors returns a boolean if a field has been set.
func (o *VisitorsStatistics) HasVisitors() bool {
	if o != nil && !IsNil(o.Visitors) {
		return true
	}

	return false
}

// SetVisitors gets a reference to the given VisitorsStatisticsVisitors and assigns it to the Visitors field.
func (o *VisitorsStatistics) SetVisitors(v VisitorsStatisticsVisitors) {
	o.Visitors = &v
}

func (o VisitorsStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisitorsStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Visitors) {
		toSerialize["visitors"] = o.Visitors
	}
	return toSerialize, nil
}

type NullableVisitorsStatistics struct {
	value *VisitorsStatistics
	isSet bool
}

func (v NullableVisitorsStatistics) Get() *VisitorsStatistics {
	return v.value
}

func (v *NullableVisitorsStatistics) Set(val *VisitorsStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableVisitorsStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableVisitorsStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisitorsStatistics(val *VisitorsStatistics) *NullableVisitorsStatistics {
	return &NullableVisitorsStatistics{value: val, isSet: true}
}

func (v NullableVisitorsStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisitorsStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


