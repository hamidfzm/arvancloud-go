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

// checks if the VisitorsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisitorsData{}

// VisitorsData struct for VisitorsData
type VisitorsData struct {
	Data *Visitors `json:"data,omitempty"`
}

// NewVisitorsData instantiates a new VisitorsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisitorsData() *VisitorsData {
	this := VisitorsData{}
	return &this
}

// NewVisitorsDataWithDefaults instantiates a new VisitorsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisitorsDataWithDefaults() *VisitorsData {
	this := VisitorsData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VisitorsData) GetData() Visitors {
	if o == nil || IsNil(o.Data) {
		var ret Visitors
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisitorsData) GetDataOk() (*Visitors, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VisitorsData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Visitors and assigns it to the Data field.
func (o *VisitorsData) SetData(v Visitors) {
	o.Data = &v
}

func (o VisitorsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisitorsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableVisitorsData struct {
	value *VisitorsData
	isSet bool
}

func (v NullableVisitorsData) Get() *VisitorsData {
	return v.value
}

func (v *NullableVisitorsData) Set(val *VisitorsData) {
	v.value = val
	v.isSet = true
}

func (v NullableVisitorsData) IsSet() bool {
	return v.isSet
}

func (v *NullableVisitorsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisitorsData(val *VisitorsData) *NullableVisitorsData {
	return &NullableVisitorsData{value: val, isSet: true}
}

func (v NullableVisitorsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisitorsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


