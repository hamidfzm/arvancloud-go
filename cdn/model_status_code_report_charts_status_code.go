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

// checks if the StatusCodeReportChartsStatusCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCodeReportChartsStatusCode{}

// StatusCodeReportChartsStatusCode struct for StatusCodeReportChartsStatusCode
type StatusCodeReportChartsStatusCode struct {
	Name *string `json:"name,omitempty"`
	Categories []time.Time `json:"categories,omitempty"`
	Series []StatusCodeReportChartsStatusCodeSeriesInner `json:"series,omitempty"`
}

// NewStatusCodeReportChartsStatusCode instantiates a new StatusCodeReportChartsStatusCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCodeReportChartsStatusCode() *StatusCodeReportChartsStatusCode {
	this := StatusCodeReportChartsStatusCode{}
	return &this
}

// NewStatusCodeReportChartsStatusCodeWithDefaults instantiates a new StatusCodeReportChartsStatusCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCodeReportChartsStatusCodeWithDefaults() *StatusCodeReportChartsStatusCode {
	this := StatusCodeReportChartsStatusCode{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StatusCodeReportChartsStatusCode) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportChartsStatusCode) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StatusCodeReportChartsStatusCode) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StatusCodeReportChartsStatusCode) SetName(v string) {
	o.Name = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *StatusCodeReportChartsStatusCode) GetCategories() []time.Time {
	if o == nil || IsNil(o.Categories) {
		var ret []time.Time
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportChartsStatusCode) GetCategoriesOk() ([]time.Time, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *StatusCodeReportChartsStatusCode) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []time.Time and assigns it to the Categories field.
func (o *StatusCodeReportChartsStatusCode) SetCategories(v []time.Time) {
	o.Categories = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *StatusCodeReportChartsStatusCode) GetSeries() []StatusCodeReportChartsStatusCodeSeriesInner {
	if o == nil || IsNil(o.Series) {
		var ret []StatusCodeReportChartsStatusCodeSeriesInner
		return ret
	}
	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeReportChartsStatusCode) GetSeriesOk() ([]StatusCodeReportChartsStatusCodeSeriesInner, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *StatusCodeReportChartsStatusCode) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []StatusCodeReportChartsStatusCodeSeriesInner and assigns it to the Series field.
func (o *StatusCodeReportChartsStatusCode) SetSeries(v []StatusCodeReportChartsStatusCodeSeriesInner) {
	o.Series = v
}

func (o StatusCodeReportChartsStatusCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCodeReportChartsStatusCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	return toSerialize, nil
}

type NullableStatusCodeReportChartsStatusCode struct {
	value *StatusCodeReportChartsStatusCode
	isSet bool
}

func (v NullableStatusCodeReportChartsStatusCode) Get() *StatusCodeReportChartsStatusCode {
	return v.value
}

func (v *NullableStatusCodeReportChartsStatusCode) Set(val *StatusCodeReportChartsStatusCode) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCodeReportChartsStatusCode) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCodeReportChartsStatusCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCodeReportChartsStatusCode(val *StatusCodeReportChartsStatusCode) *NullableStatusCodeReportChartsStatusCode {
	return &NullableStatusCodeReportChartsStatusCode{value: val, isSet: true}
}

func (v NullableStatusCodeReportChartsStatusCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCodeReportChartsStatusCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

