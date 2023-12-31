/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
)

// checks if the AnalyticsAndOperator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsAndOperator{}

// AnalyticsAndOperator A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The operator must have at least two predicates in any combination, and an object must match all of the predicates for the filter to apply.
type AnalyticsAndOperator struct {
	Prefix *string `json:"Prefix,omitempty"`
	Tags *Array `json:"Tags,omitempty"`
}

// NewAnalyticsAndOperator instantiates a new AnalyticsAndOperator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsAndOperator() *AnalyticsAndOperator {
	this := AnalyticsAndOperator{}
	return &this
}

// NewAnalyticsAndOperatorWithDefaults instantiates a new AnalyticsAndOperator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsAndOperatorWithDefaults() *AnalyticsAndOperator {
	this := AnalyticsAndOperator{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *AnalyticsAndOperator) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsAndOperator) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *AnalyticsAndOperator) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *AnalyticsAndOperator) SetPrefix(v string) {
	o.Prefix = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AnalyticsAndOperator) GetTags() Array {
	if o == nil || IsNil(o.Tags) {
		var ret Array
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyticsAndOperator) GetTagsOk() (*Array, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AnalyticsAndOperator) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given Array and assigns it to the Tags field.
func (o *AnalyticsAndOperator) SetTags(v Array) {
	o.Tags = &v
}

func (o AnalyticsAndOperator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsAndOperator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prefix) {
		toSerialize["Prefix"] = o.Prefix
	}
	if !IsNil(o.Tags) {
		toSerialize["Tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableAnalyticsAndOperator struct {
	value *AnalyticsAndOperator
	isSet bool
}

func (v NullableAnalyticsAndOperator) Get() *AnalyticsAndOperator {
	return v.value
}

func (v *NullableAnalyticsAndOperator) Set(val *AnalyticsAndOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsAndOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsAndOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsAndOperator(val *AnalyticsAndOperator) *NullableAnalyticsAndOperator {
	return &NullableAnalyticsAndOperator{value: val, isSet: true}
}

func (v NullableAnalyticsAndOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsAndOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


