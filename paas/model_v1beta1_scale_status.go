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

// checks if the V1beta1ScaleStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1ScaleStatus{}

// V1beta1ScaleStatus represents the current status of a scale subresource.
type V1beta1ScaleStatus struct {
	// actual number of observed instances of the scaled object.
	Replicas int32 `json:"replicas"`
	// label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
	Selector map[string]interface{} `json:"selector,omitempty"`
	// label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	TargetSelector *string `json:"targetSelector,omitempty"`
}

// NewV1beta1ScaleStatus instantiates a new V1beta1ScaleStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1ScaleStatus(replicas int32) *V1beta1ScaleStatus {
	this := V1beta1ScaleStatus{}
	this.Replicas = replicas
	return &this
}

// NewV1beta1ScaleStatusWithDefaults instantiates a new V1beta1ScaleStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1ScaleStatusWithDefaults() *V1beta1ScaleStatus {
	this := V1beta1ScaleStatus{}
	return &this
}

// GetReplicas returns the Replicas field value
func (o *V1beta1ScaleStatus) GetReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *V1beta1ScaleStatus) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value
func (o *V1beta1ScaleStatus) SetReplicas(v int32) {
	o.Replicas = v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *V1beta1ScaleStatus) GetSelector() map[string]interface{} {
	if o == nil || IsNil(o.Selector) {
		var ret map[string]interface{}
		return ret
	}
	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1ScaleStatus) GetSelectorOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Selector) {
		return map[string]interface{}{}, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *V1beta1ScaleStatus) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given map[string]interface{} and assigns it to the Selector field.
func (o *V1beta1ScaleStatus) SetSelector(v map[string]interface{}) {
	o.Selector = v
}

// GetTargetSelector returns the TargetSelector field value if set, zero value otherwise.
func (o *V1beta1ScaleStatus) GetTargetSelector() string {
	if o == nil || IsNil(o.TargetSelector) {
		var ret string
		return ret
	}
	return *o.TargetSelector
}

// GetTargetSelectorOk returns a tuple with the TargetSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1ScaleStatus) GetTargetSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.TargetSelector) {
		return nil, false
	}
	return o.TargetSelector, true
}

// HasTargetSelector returns a boolean if a field has been set.
func (o *V1beta1ScaleStatus) HasTargetSelector() bool {
	if o != nil && !IsNil(o.TargetSelector) {
		return true
	}

	return false
}

// SetTargetSelector gets a reference to the given string and assigns it to the TargetSelector field.
func (o *V1beta1ScaleStatus) SetTargetSelector(v string) {
	o.TargetSelector = &v
}

func (o V1beta1ScaleStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1ScaleStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["replicas"] = o.Replicas
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !IsNil(o.TargetSelector) {
		toSerialize["targetSelector"] = o.TargetSelector
	}
	return toSerialize, nil
}

type NullableV1beta1ScaleStatus struct {
	value *V1beta1ScaleStatus
	isSet bool
}

func (v NullableV1beta1ScaleStatus) Get() *V1beta1ScaleStatus {
	return v.value
}

func (v *NullableV1beta1ScaleStatus) Set(val *V1beta1ScaleStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1ScaleStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1ScaleStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1ScaleStatus(val *V1beta1ScaleStatus) *NullableV1beta1ScaleStatus {
	return &NullableV1beta1ScaleStatus{value: val, isSet: true}
}

func (v NullableV1beta1ScaleStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1ScaleStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


