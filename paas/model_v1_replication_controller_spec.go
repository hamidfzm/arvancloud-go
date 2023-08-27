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

// checks if the V1ReplicationControllerSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ReplicationControllerSpec{}

// V1ReplicationControllerSpec ReplicationControllerSpec is the specification of a replication controller.
type V1ReplicationControllerSpec struct {
	// Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
	MinReadySeconds *int32 `json:"minReadySeconds,omitempty"`
	// Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
	Replicas *int32 `json:"replicas,omitempty"`
	// Selector is a label query over pods that should match the Replicas count. If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector map[string]interface{} `json:"selector,omitempty"`
	Template *V1PodTemplateSpec `json:"template,omitempty"`
}

// NewV1ReplicationControllerSpec instantiates a new V1ReplicationControllerSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ReplicationControllerSpec() *V1ReplicationControllerSpec {
	this := V1ReplicationControllerSpec{}
	return &this
}

// NewV1ReplicationControllerSpecWithDefaults instantiates a new V1ReplicationControllerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ReplicationControllerSpecWithDefaults() *V1ReplicationControllerSpec {
	this := V1ReplicationControllerSpec{}
	return &this
}

// GetMinReadySeconds returns the MinReadySeconds field value if set, zero value otherwise.
func (o *V1ReplicationControllerSpec) GetMinReadySeconds() int32 {
	if o == nil || IsNil(o.MinReadySeconds) {
		var ret int32
		return ret
	}
	return *o.MinReadySeconds
}

// GetMinReadySecondsOk returns a tuple with the MinReadySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ReplicationControllerSpec) GetMinReadySecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.MinReadySeconds) {
		return nil, false
	}
	return o.MinReadySeconds, true
}

// HasMinReadySeconds returns a boolean if a field has been set.
func (o *V1ReplicationControllerSpec) HasMinReadySeconds() bool {
	if o != nil && !IsNil(o.MinReadySeconds) {
		return true
	}

	return false
}

// SetMinReadySeconds gets a reference to the given int32 and assigns it to the MinReadySeconds field.
func (o *V1ReplicationControllerSpec) SetMinReadySeconds(v int32) {
	o.MinReadySeconds = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *V1ReplicationControllerSpec) GetReplicas() int32 {
	if o == nil || IsNil(o.Replicas) {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ReplicationControllerSpec) GetReplicasOk() (*int32, bool) {
	if o == nil || IsNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *V1ReplicationControllerSpec) HasReplicas() bool {
	if o != nil && !IsNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *V1ReplicationControllerSpec) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *V1ReplicationControllerSpec) GetSelector() map[string]interface{} {
	if o == nil || IsNil(o.Selector) {
		var ret map[string]interface{}
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ReplicationControllerSpec) GetSelectorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Selector) {
		return map[string]interface{}{}, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *V1ReplicationControllerSpec) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given map[string]interface{} and assigns it to the Selector field.
func (o *V1ReplicationControllerSpec) SetSelector(v map[string]interface{}) {
	o.Selector = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *V1ReplicationControllerSpec) GetTemplate() V1PodTemplateSpec {
	if o == nil || IsNil(o.Template) {
		var ret V1PodTemplateSpec
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ReplicationControllerSpec) GetTemplateOk() (*V1PodTemplateSpec, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *V1ReplicationControllerSpec) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given V1PodTemplateSpec and assigns it to the Template field.
func (o *V1ReplicationControllerSpec) SetTemplate(v V1PodTemplateSpec) {
	o.Template = &v
}

func (o V1ReplicationControllerSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ReplicationControllerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinReadySeconds) {
		toSerialize["minReadySeconds"] = o.MinReadySeconds
	}
	if !IsNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableV1ReplicationControllerSpec struct {
	value *V1ReplicationControllerSpec
	isSet bool
}

func (v NullableV1ReplicationControllerSpec) Get() *V1ReplicationControllerSpec {
	return v.value
}

func (v *NullableV1ReplicationControllerSpec) Set(val *V1ReplicationControllerSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ReplicationControllerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ReplicationControllerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ReplicationControllerSpec(val *V1ReplicationControllerSpec) *NullableV1ReplicationControllerSpec {
	return &NullableV1ReplicationControllerSpec{value: val, isSet: true}
}

func (v NullableV1ReplicationControllerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ReplicationControllerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


