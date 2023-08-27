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

// checks if the V1RollingUpdateDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1RollingUpdateDeployment{}

// V1RollingUpdateDeployment Spec to control the desired behavior of rolling update.
type V1RollingUpdateDeployment struct {
	// The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.
	MaxSurge *string `json:"maxSurge,omitempty"`
	// The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.
	MaxUnavailable *string `json:"maxUnavailable,omitempty"`
}

// NewV1RollingUpdateDeployment instantiates a new V1RollingUpdateDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RollingUpdateDeployment() *V1RollingUpdateDeployment {
	this := V1RollingUpdateDeployment{}
	return &this
}

// NewV1RollingUpdateDeploymentWithDefaults instantiates a new V1RollingUpdateDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RollingUpdateDeploymentWithDefaults() *V1RollingUpdateDeployment {
	this := V1RollingUpdateDeployment{}
	return &this
}

// GetMaxSurge returns the MaxSurge field value if set, zero value otherwise.
func (o *V1RollingUpdateDeployment) GetMaxSurge() string {
	if o == nil || IsNil(o.MaxSurge) {
		var ret string
		return ret
	}
	return *o.MaxSurge
}

// GetMaxSurgeOk returns a tuple with the MaxSurge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RollingUpdateDeployment) GetMaxSurgeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxSurge) {
		return nil, false
	}
	return o.MaxSurge, true
}

// HasMaxSurge returns a boolean if a field has been set.
func (o *V1RollingUpdateDeployment) HasMaxSurge() bool {
	if o != nil && !IsNil(o.MaxSurge) {
		return true
	}

	return false
}

// SetMaxSurge gets a reference to the given string and assigns it to the MaxSurge field.
func (o *V1RollingUpdateDeployment) SetMaxSurge(v string) {
	o.MaxSurge = &v
}

// GetMaxUnavailable returns the MaxUnavailable field value if set, zero value otherwise.
func (o *V1RollingUpdateDeployment) GetMaxUnavailable() string {
	if o == nil || IsNil(o.MaxUnavailable) {
		var ret string
		return ret
	}
	return *o.MaxUnavailable
}

// GetMaxUnavailableOk returns a tuple with the MaxUnavailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RollingUpdateDeployment) GetMaxUnavailableOk() (*string, bool) {
	if o == nil || IsNil(o.MaxUnavailable) {
		return nil, false
	}
	return o.MaxUnavailable, true
}

// HasMaxUnavailable returns a boolean if a field has been set.
func (o *V1RollingUpdateDeployment) HasMaxUnavailable() bool {
	if o != nil && !IsNil(o.MaxUnavailable) {
		return true
	}

	return false
}

// SetMaxUnavailable gets a reference to the given string and assigns it to the MaxUnavailable field.
func (o *V1RollingUpdateDeployment) SetMaxUnavailable(v string) {
	o.MaxUnavailable = &v
}

func (o V1RollingUpdateDeployment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1RollingUpdateDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxSurge) {
		toSerialize["maxSurge"] = o.MaxSurge
	}
	if !IsNil(o.MaxUnavailable) {
		toSerialize["maxUnavailable"] = o.MaxUnavailable
	}
	return toSerialize, nil
}

type NullableV1RollingUpdateDeployment struct {
	value *V1RollingUpdateDeployment
	isSet bool
}

func (v NullableV1RollingUpdateDeployment) Get() *V1RollingUpdateDeployment {
	return v.value
}

func (v *NullableV1RollingUpdateDeployment) Set(val *V1RollingUpdateDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RollingUpdateDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RollingUpdateDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RollingUpdateDeployment(val *V1RollingUpdateDeployment) *NullableV1RollingUpdateDeployment {
	return &NullableV1RollingUpdateDeployment{value: val, isSet: true}
}

func (v NullableV1RollingUpdateDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RollingUpdateDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


