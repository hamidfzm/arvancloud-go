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

// checks if the V1BuildStatusOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BuildStatusOutput{}

// V1BuildStatusOutput BuildStatusOutput contains the status of the built image.
type V1BuildStatusOutput struct {
	To *V1BuildStatusOutputTo `json:"to,omitempty"`
}

// NewV1BuildStatusOutput instantiates a new V1BuildStatusOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BuildStatusOutput() *V1BuildStatusOutput {
	this := V1BuildStatusOutput{}
	return &this
}

// NewV1BuildStatusOutputWithDefaults instantiates a new V1BuildStatusOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BuildStatusOutputWithDefaults() *V1BuildStatusOutput {
	this := V1BuildStatusOutput{}
	return &this
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *V1BuildStatusOutput) GetTo() V1BuildStatusOutputTo {
	if o == nil || IsNil(o.To) {
		var ret V1BuildStatusOutputTo
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildStatusOutput) GetToOk() (*V1BuildStatusOutputTo, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *V1BuildStatusOutput) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given V1BuildStatusOutputTo and assigns it to the To field.
func (o *V1BuildStatusOutput) SetTo(v V1BuildStatusOutputTo) {
	o.To = &v
}

func (o V1BuildStatusOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BuildStatusOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableV1BuildStatusOutput struct {
	value *V1BuildStatusOutput
	isSet bool
}

func (v NullableV1BuildStatusOutput) Get() *V1BuildStatusOutput {
	return v.value
}

func (v *NullableV1BuildStatusOutput) Set(val *V1BuildStatusOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BuildStatusOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BuildStatusOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BuildStatusOutput(val *V1BuildStatusOutput) *NullableV1BuildStatusOutput {
	return &NullableV1BuildStatusOutput{value: val, isSet: true}
}

func (v NullableV1BuildStatusOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BuildStatusOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

