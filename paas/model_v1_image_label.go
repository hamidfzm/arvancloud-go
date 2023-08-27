/*
Arvancloud PaaS REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paas

import (
	"encoding/json"
)

// checks if the V1ImageLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ImageLabel{}

// V1ImageLabel ImageLabel represents a label applied to the resulting image.
type V1ImageLabel struct {
	// name defines the name of the label. It must have non-zero length.
	Name string `json:"name"`
	// value defines the literal value of the label.
	Value *string `json:"value,omitempty"`
}

// NewV1ImageLabel instantiates a new V1ImageLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImageLabel(name string) *V1ImageLabel {
	this := V1ImageLabel{}
	this.Name = name
	return &this
}

// NewV1ImageLabelWithDefaults instantiates a new V1ImageLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImageLabelWithDefaults() *V1ImageLabel {
	this := V1ImageLabel{}
	return &this
}

// GetName returns the Name field value
func (o *V1ImageLabel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1ImageLabel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1ImageLabel) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V1ImageLabel) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageLabel) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V1ImageLabel) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *V1ImageLabel) SetValue(v string) {
	o.Value = &v
}

func (o V1ImageLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ImageLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableV1ImageLabel struct {
	value *V1ImageLabel
	isSet bool
}

func (v NullableV1ImageLabel) Get() *V1ImageLabel {
	return v.value
}

func (v *NullableV1ImageLabel) Set(val *V1ImageLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImageLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImageLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImageLabel(val *V1ImageLabel) *NullableV1ImageLabel {
	return &NullableV1ImageLabel{value: val, isSet: true}
}

func (v NullableV1ImageLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImageLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

