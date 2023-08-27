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

// checks if the V1SecretLocalReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1SecretLocalReference{}

// V1SecretLocalReference SecretLocalReference contains information that points to the local secret being used
type V1SecretLocalReference struct {
	// Name is the name of the resource in the same namespace being referenced
	Name string `json:"name"`
}

// NewV1SecretLocalReference instantiates a new V1SecretLocalReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SecretLocalReference(name string) *V1SecretLocalReference {
	this := V1SecretLocalReference{}
	this.Name = name
	return &this
}

// NewV1SecretLocalReferenceWithDefaults instantiates a new V1SecretLocalReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SecretLocalReferenceWithDefaults() *V1SecretLocalReference {
	this := V1SecretLocalReference{}
	return &this
}

// GetName returns the Name field value
func (o *V1SecretLocalReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1SecretLocalReference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1SecretLocalReference) SetName(v string) {
	o.Name = v
}

func (o V1SecretLocalReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1SecretLocalReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableV1SecretLocalReference struct {
	value *V1SecretLocalReference
	isSet bool
}

func (v NullableV1SecretLocalReference) Get() *V1SecretLocalReference {
	return v.value
}

func (v *NullableV1SecretLocalReference) Set(val *V1SecretLocalReference) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SecretLocalReference) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SecretLocalReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SecretLocalReference(val *V1SecretLocalReference) *NullableV1SecretLocalReference {
	return &NullableV1SecretLocalReference{value: val, isSet: true}
}

func (v NullableV1SecretLocalReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SecretLocalReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


