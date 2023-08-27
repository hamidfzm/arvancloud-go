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

// checks if the V1ReplicaSetSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ReplicaSetSpec{}

// V1ReplicaSetSpec ReplicaSetSpec is the specification of a ReplicaSet.
type V1ReplicaSetSpec struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	MinReadySeconds *int32 `json:"minReadySeconds,omitempty"`
	// Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Replicas *int32 `json:"replicas,omitempty"`
	Selector V1LabelSelector `json:"selector"`
	Template *V1PodTemplateSpec `json:"template,omitempty"`
}

// NewV1ReplicaSetSpec instantiates a new V1ReplicaSetSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ReplicaSetSpec(selector V1LabelSelector) *V1ReplicaSetSpec {
	this := V1ReplicaSetSpec{}
	this.Selector = selector
	return &this
}

// NewV1ReplicaSetSpecWithDefaults instantiates a new V1ReplicaSetSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ReplicaSetSpecWithDefaults() *V1ReplicaSetSpec {
	this := V1ReplicaSetSpec{}
	return &this
}

// GetMinReadySeconds returns the MinReadySeconds field value if set, zero value otherwise.
func (o *V1ReplicaSetSpec) GetMinReadySeconds() int32 {
	if o == nil || IsNil(o.MinReadySeconds) {
		var ret int32
		return ret
	}
	return *o.MinReadySeconds
}

// GetMinReadySecondsOk returns a tuple with the MinReadySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ReplicaSetSpec) GetMinReadySecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.MinReadySeconds) {
		return nil, false
	}
	return o.MinReadySeconds, true
}

// HasMinReadySeconds returns a boolean if a field has been set.
func (o *V1ReplicaSetSpec) HasMinReadySeconds() bool {
	if o != nil && !IsNil(o.MinReadySeconds) {
		return true
	}

	return false
}

// SetMinReadySeconds gets a reference to the given int32 and assigns it to the MinReadySeconds field.
func (o *V1ReplicaSetSpec) SetMinReadySeconds(v int32) {
	o.MinReadySeconds = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *V1ReplicaSetSpec) GetReplicas() int32 {
	if o == nil || IsNil(o.Replicas) {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ReplicaSetSpec) GetReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *V1ReplicaSetSpec) HasReplicas() bool {
	if o != nil && !IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *V1ReplicaSetSpec) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetSelector returns the Selector field value
func (o *V1ReplicaSetSpec) GetSelector() V1LabelSelector {
	if o == nil {
		var ret V1LabelSelector
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *V1ReplicaSetSpec) GetSelectorOk() (*V1LabelSelector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *V1ReplicaSetSpec) SetSelector(v V1LabelSelector) {
	o.Selector = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *V1ReplicaSetSpec) GetTemplate() V1PodTemplateSpec {
	if o == nil || IsNil(o.Template) {
		var ret V1PodTemplateSpec
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ReplicaSetSpec) GetTemplateOk() (*V1PodTemplateSpec, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *V1ReplicaSetSpec) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given V1PodTemplateSpec and assigns it to the Template field.
func (o *V1ReplicaSetSpec) SetTemplate(v V1PodTemplateSpec) {
	o.Template = &v
}

func (o V1ReplicaSetSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ReplicaSetSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinReadySeconds) {
		toSerialize["minReadySeconds"] = o.MinReadySeconds
	}
	if !IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	toSerialize["selector"] = o.Selector
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableV1ReplicaSetSpec struct {
	value *V1ReplicaSetSpec
	isSet bool
}

func (v NullableV1ReplicaSetSpec) Get() *V1ReplicaSetSpec {
	return v.value
}

func (v *NullableV1ReplicaSetSpec) Set(val *V1ReplicaSetSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ReplicaSetSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ReplicaSetSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ReplicaSetSpec(val *V1ReplicaSetSpec) *NullableV1ReplicaSetSpec {
	return &NullableV1ReplicaSetSpec{value: val, isSet: true}
}

func (v NullableV1ReplicaSetSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ReplicaSetSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


