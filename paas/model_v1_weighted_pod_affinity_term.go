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

// checks if the V1WeightedPodAffinityTerm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1WeightedPodAffinityTerm{}

// V1WeightedPodAffinityTerm The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
type V1WeightedPodAffinityTerm struct {
	PodAffinityTerm V1PodAffinityTerm `json:"podAffinityTerm"`
	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	Weight int32 `json:"weight"`
}

// NewV1WeightedPodAffinityTerm instantiates a new V1WeightedPodAffinityTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1WeightedPodAffinityTerm(podAffinityTerm V1PodAffinityTerm, weight int32) *V1WeightedPodAffinityTerm {
	this := V1WeightedPodAffinityTerm{}
	this.PodAffinityTerm = podAffinityTerm
	this.Weight = weight
	return &this
}

// NewV1WeightedPodAffinityTermWithDefaults instantiates a new V1WeightedPodAffinityTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1WeightedPodAffinityTermWithDefaults() *V1WeightedPodAffinityTerm {
	this := V1WeightedPodAffinityTerm{}
	return &this
}

// GetPodAffinityTerm returns the PodAffinityTerm field value
func (o *V1WeightedPodAffinityTerm) GetPodAffinityTerm() V1PodAffinityTerm {
	if o == nil {
		var ret V1PodAffinityTerm
		return ret
	}

	return o.PodAffinityTerm
}

// GetPodAffinityTermOk returns a tuple with the PodAffinityTerm field value
// and a boolean to check if the value has been set.
func (o *V1WeightedPodAffinityTerm) GetPodAffinityTermOk() (*V1PodAffinityTerm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodAffinityTerm, true
}

// SetPodAffinityTerm sets field value
func (o *V1WeightedPodAffinityTerm) SetPodAffinityTerm(v V1PodAffinityTerm) {
	o.PodAffinityTerm = v
}

// GetWeight returns the Weight field value
func (o *V1WeightedPodAffinityTerm) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *V1WeightedPodAffinityTerm) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *V1WeightedPodAffinityTerm) SetWeight(v int32) {
	o.Weight = v
}

func (o V1WeightedPodAffinityTerm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1WeightedPodAffinityTerm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["podAffinityTerm"] = o.PodAffinityTerm
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

type NullableV1WeightedPodAffinityTerm struct {
	value *V1WeightedPodAffinityTerm
	isSet bool
}

func (v NullableV1WeightedPodAffinityTerm) Get() *V1WeightedPodAffinityTerm {
	return v.value
}

func (v *NullableV1WeightedPodAffinityTerm) Set(val *V1WeightedPodAffinityTerm) {
	v.value = val
	v.isSet = true
}

func (v NullableV1WeightedPodAffinityTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableV1WeightedPodAffinityTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1WeightedPodAffinityTerm(val *V1WeightedPodAffinityTerm) *NullableV1WeightedPodAffinityTerm {
	return &NullableV1WeightedPodAffinityTerm{value: val, isSet: true}
}

func (v NullableV1WeightedPodAffinityTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1WeightedPodAffinityTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


