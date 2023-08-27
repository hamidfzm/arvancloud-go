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

// checks if the V1beta1ScaleSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1ScaleSpec{}

// V1beta1ScaleSpec describes the attributes of a scale subresource
type V1beta1ScaleSpec struct {
	// desired number of instances for the scaled object.
	Replicas *int32 `json:"replicas,omitempty"`
}

// NewV1beta1ScaleSpec instantiates a new V1beta1ScaleSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1ScaleSpec() *V1beta1ScaleSpec {
	this := V1beta1ScaleSpec{}
	return &this
}

// NewV1beta1ScaleSpecWithDefaults instantiates a new V1beta1ScaleSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1ScaleSpecWithDefaults() *V1beta1ScaleSpec {
	this := V1beta1ScaleSpec{}
	return &this
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *V1beta1ScaleSpec) GetReplicas() int32 {
	if o == nil || IsNil(o.Replicas) {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1ScaleSpec) GetReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *V1beta1ScaleSpec) HasReplicas() bool {
	if o != nil && !IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *V1beta1ScaleSpec) SetReplicas(v int32) {
	o.Replicas = &v
}

func (o V1beta1ScaleSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1ScaleSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	return toSerialize, nil
}

type NullableV1beta1ScaleSpec struct {
	value *V1beta1ScaleSpec
	isSet bool
}

func (v NullableV1beta1ScaleSpec) Get() *V1beta1ScaleSpec {
	return v.value
}

func (v *NullableV1beta1ScaleSpec) Set(val *V1beta1ScaleSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1ScaleSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1ScaleSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1ScaleSpec(val *V1beta1ScaleSpec) *NullableV1beta1ScaleSpec {
	return &NullableV1beta1ScaleSpec{value: val, isSet: true}
}

func (v NullableV1beta1ScaleSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1ScaleSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

