/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the ChartDataset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChartDataset{}

// ChartDataset struct for ChartDataset
type ChartDataset struct {
	Data []float64 `json:"data,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewChartDataset instantiates a new ChartDataset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartDataset() *ChartDataset {
	this := ChartDataset{}
	return &this
}

// NewChartDatasetWithDefaults instantiates a new ChartDataset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartDatasetWithDefaults() *ChartDataset {
	this := ChartDataset{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ChartDataset) GetData() []float64 {
	if o == nil || IsNil(o.Data) {
		var ret []float64
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartDataset) GetDataOk() ([]float64, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ChartDataset) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []float64 and assigns it to the Data field.
func (o *ChartDataset) SetData(v []float64) {
	o.Data = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChartDataset) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartDataset) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChartDataset) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChartDataset) SetName(v string) {
	o.Name = &v
}

func (o ChartDataset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChartDataset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableChartDataset struct {
	value *ChartDataset
	isSet bool
}

func (v NullableChartDataset) Get() *ChartDataset {
	return v.value
}

func (v *NullableChartDataset) Set(val *ChartDataset) {
	v.value = val
	v.isSet = true
}

func (v NullableChartDataset) IsSet() bool {
	return v.isSet
}

func (v *NullableChartDataset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartDataset(val *ChartDataset) *NullableChartDataset {
	return &NullableChartDataset{value: val, isSet: true}
}

func (v NullableChartDataset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartDataset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


