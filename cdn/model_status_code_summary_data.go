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

// checks if the StatusCodeSummaryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCodeSummaryData{}

// StatusCodeSummaryData struct for StatusCodeSummaryData
type StatusCodeSummaryData struct {
	Data *StatusCodeSummary `json:"data,omitempty"`
}

// NewStatusCodeSummaryData instantiates a new StatusCodeSummaryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCodeSummaryData() *StatusCodeSummaryData {
	this := StatusCodeSummaryData{}
	return &this
}

// NewStatusCodeSummaryDataWithDefaults instantiates a new StatusCodeSummaryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCodeSummaryDataWithDefaults() *StatusCodeSummaryData {
	this := StatusCodeSummaryData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StatusCodeSummaryData) GetData() StatusCodeSummary {
	if o == nil || IsNil(o.Data) {
		var ret StatusCodeSummary
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeSummaryData) GetDataOk() (*StatusCodeSummary, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StatusCodeSummaryData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given StatusCodeSummary and assigns it to the Data field.
func (o *StatusCodeSummaryData) SetData(v StatusCodeSummary) {
	o.Data = &v
}

func (o StatusCodeSummaryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCodeSummaryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableStatusCodeSummaryData struct {
	value *StatusCodeSummaryData
	isSet bool
}

func (v NullableStatusCodeSummaryData) Get() *StatusCodeSummaryData {
	return v.value
}

func (v *NullableStatusCodeSummaryData) Set(val *StatusCodeSummaryData) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCodeSummaryData) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCodeSummaryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCodeSummaryData(val *StatusCodeSummaryData) *NullableStatusCodeSummaryData {
	return &NullableStatusCodeSummaryData{value: val, isSet: true}
}

func (v NullableStatusCodeSummaryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCodeSummaryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

