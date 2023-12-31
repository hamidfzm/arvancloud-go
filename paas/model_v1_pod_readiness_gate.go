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

// checks if the V1PodReadinessGate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PodReadinessGate{}

// V1PodReadinessGate PodReadinessGate contains the reference to a pod condition
type V1PodReadinessGate struct {
	// ConditionType refers to a condition in the pod's condition list with matching type.
	ConditionType string `json:"conditionType"`
}

// NewV1PodReadinessGate instantiates a new V1PodReadinessGate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PodReadinessGate(conditionType string) *V1PodReadinessGate {
	this := V1PodReadinessGate{}
	this.ConditionType = conditionType
	return &this
}

// NewV1PodReadinessGateWithDefaults instantiates a new V1PodReadinessGate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PodReadinessGateWithDefaults() *V1PodReadinessGate {
	this := V1PodReadinessGate{}
	return &this
}

// GetConditionType returns the ConditionType field value
func (o *V1PodReadinessGate) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *V1PodReadinessGate) GetConditionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *V1PodReadinessGate) SetConditionType(v string) {
	o.ConditionType = v
}

func (o V1PodReadinessGate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PodReadinessGate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conditionType"] = o.ConditionType
	return toSerialize, nil
}

type NullableV1PodReadinessGate struct {
	value *V1PodReadinessGate
	isSet bool
}

func (v NullableV1PodReadinessGate) Get() *V1PodReadinessGate {
	return v.value
}

func (v *NullableV1PodReadinessGate) Set(val *V1PodReadinessGate) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PodReadinessGate) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PodReadinessGate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PodReadinessGate(val *V1PodReadinessGate) *NullableV1PodReadinessGate {
	return &NullableV1PodReadinessGate{value: val, isSet: true}
}

func (v NullableV1PodReadinessGate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PodReadinessGate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


