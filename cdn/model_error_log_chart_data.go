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

// checks if the ErrorLogChartData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorLogChartData{}

// ErrorLogChartData struct for ErrorLogChartData
type ErrorLogChartData struct {
	Data *ErrorLogChart `json:"data,omitempty"`
}

// NewErrorLogChartData instantiates a new ErrorLogChartData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorLogChartData() *ErrorLogChartData {
	this := ErrorLogChartData{}
	return &this
}

// NewErrorLogChartDataWithDefaults instantiates a new ErrorLogChartData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorLogChartDataWithDefaults() *ErrorLogChartData {
	this := ErrorLogChartData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ErrorLogChartData) GetData() ErrorLogChart {
	if o == nil || IsNil(o.Data) {
		var ret ErrorLogChart
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogChartData) GetDataOk() (*ErrorLogChart, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ErrorLogChartData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ErrorLogChart and assigns it to the Data field.
func (o *ErrorLogChartData) SetData(v ErrorLogChart) {
	o.Data = &v
}

func (o ErrorLogChartData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorLogChartData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableErrorLogChartData struct {
	value *ErrorLogChartData
	isSet bool
}

func (v NullableErrorLogChartData) Get() *ErrorLogChartData {
	return v.value
}

func (v *NullableErrorLogChartData) Set(val *ErrorLogChartData) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLogChartData) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLogChartData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLogChartData(val *ErrorLogChartData) *NullableErrorLogChartData {
	return &NullableErrorLogChartData{value: val, isSet: true}
}

func (v NullableErrorLogChartData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLogChartData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


