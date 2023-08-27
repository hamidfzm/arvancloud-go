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

// checks if the V1LocalObjectReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1LocalObjectReference{}

// V1LocalObjectReference LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
type V1LocalObjectReference struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `json:"name,omitempty"`
}

// NewV1LocalObjectReference instantiates a new V1LocalObjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1LocalObjectReference() *V1LocalObjectReference {
	this := V1LocalObjectReference{}
	return &this
}

// NewV1LocalObjectReferenceWithDefaults instantiates a new V1LocalObjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1LocalObjectReferenceWithDefaults() *V1LocalObjectReference {
	this := V1LocalObjectReference{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1LocalObjectReference) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LocalObjectReference) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1LocalObjectReference) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1LocalObjectReference) SetName(v string) {
	o.Name = &v
}

func (o V1LocalObjectReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1LocalObjectReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV1LocalObjectReference struct {
	value *V1LocalObjectReference
	isSet bool
}

func (v NullableV1LocalObjectReference) Get() *V1LocalObjectReference {
	return v.value
}

func (v *NullableV1LocalObjectReference) Set(val *V1LocalObjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableV1LocalObjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableV1LocalObjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1LocalObjectReference(val *V1LocalObjectReference) *NullableV1LocalObjectReference {
	return &NullableV1LocalObjectReference{value: val, isSet: true}
}

func (v NullableV1LocalObjectReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1LocalObjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

