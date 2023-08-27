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

// checks if the ResponseTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTime{}

// ResponseTime struct for ResponseTime
type ResponseTime struct {
	Statistics *ResponseTimeStatistics `json:"statistics,omitempty"`
	Charts *ResponseTimeCharts `json:"charts,omitempty"`
}

// NewResponseTime instantiates a new ResponseTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTime() *ResponseTime {
	this := ResponseTime{}
	return &this
}

// NewResponseTimeWithDefaults instantiates a new ResponseTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTimeWithDefaults() *ResponseTime {
	this := ResponseTime{}
	return &this
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *ResponseTime) GetStatistics() ResponseTimeStatistics {
	if o == nil || IsNil(o.Statistics) {
		var ret ResponseTimeStatistics
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTime) GetStatisticsOk() (*ResponseTimeStatistics, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *ResponseTime) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given ResponseTimeStatistics and assigns it to the Statistics field.
func (o *ResponseTime) SetStatistics(v ResponseTimeStatistics) {
	o.Statistics = &v
}

// GetCharts returns the Charts field value if set, zero value otherwise.
func (o *ResponseTime) GetCharts() ResponseTimeCharts {
	if o == nil || IsNil(o.Charts) {
		var ret ResponseTimeCharts
		return ret
	}
	return *o.Charts
}

// GetChartsOk returns a tuple with the Charts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTime) GetChartsOk() (*ResponseTimeCharts, bool) {
	if o == nil || IsNil(o.Charts) {
		return nil, false
	}
	return o.Charts, true
}

// HasCharts returns a boolean if a field has been set.
func (o *ResponseTime) HasCharts() bool {
	if o != nil && !IsNil(o.Charts) {
		return true
	}

	return false
}

// SetCharts gets a reference to the given ResponseTimeCharts and assigns it to the Charts field.
func (o *ResponseTime) SetCharts(v ResponseTimeCharts) {
	o.Charts = &v
}

func (o ResponseTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !IsNil(o.Charts) {
		toSerialize["charts"] = o.Charts
	}
	return toSerialize, nil
}

type NullableResponseTime struct {
	value *ResponseTime
	isSet bool
}

func (v NullableResponseTime) Get() *ResponseTime {
	return v.value
}

func (v *NullableResponseTime) Set(val *ResponseTime) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTime) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTime(val *ResponseTime) *NullableResponseTime {
	return &NullableResponseTime{value: val, isSet: true}
}

func (v NullableResponseTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


