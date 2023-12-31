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

// checks if the V1PodCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PodCondition{}

// V1PodCondition PodCondition contains details for the current condition of this pod.
type V1PodCondition struct {
	// Last time we probed the condition.
	LastProbeTime *string `json:"lastProbeTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`
	// Human-readable message indicating details about last transition.
	Message *string `json:"message,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason *string `json:"reason,omitempty"`
	// Status is the status of the condition. Can be True, False, Unknown. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Status string `json:"status"`
	// Type is the type of the condition. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Type string `json:"type"`
}

// NewV1PodCondition instantiates a new V1PodCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PodCondition(status string, type_ string) *V1PodCondition {
	this := V1PodCondition{}
	this.Status = status
	this.Type = type_
	return &this
}

// NewV1PodConditionWithDefaults instantiates a new V1PodCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PodConditionWithDefaults() *V1PodCondition {
	this := V1PodCondition{}
	return &this
}

// GetLastProbeTime returns the LastProbeTime field value if set, zero value otherwise.
func (o *V1PodCondition) GetLastProbeTime() string {
	if o == nil || IsNil(o.LastProbeTime) {
		var ret string
		return ret
	}
	return *o.LastProbeTime
}

// GetLastProbeTimeOk returns a tuple with the LastProbeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodCondition) GetLastProbeTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastProbeTime) {
		return nil, false
	}
	return o.LastProbeTime, true
}

// HasLastProbeTime returns a boolean if a field has been set.
func (o *V1PodCondition) HasLastProbeTime() bool {
	if o != nil && !IsNil(o.LastProbeTime) {
		return true
	}

	return false
}

// SetLastProbeTime gets a reference to the given string and assigns it to the LastProbeTime field.
func (o *V1PodCondition) SetLastProbeTime(v string) {
	o.LastProbeTime = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *V1PodCondition) GetLastTransitionTime() string {
	if o == nil || IsNil(o.LastTransitionTime) {
		var ret string
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodCondition) GetLastTransitionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastTransitionTime) {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *V1PodCondition) HasLastTransitionTime() bool {
	if o != nil && !IsNil(o.LastTransitionTime) {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given string and assigns it to the LastTransitionTime field.
func (o *V1PodCondition) SetLastTransitionTime(v string) {
	o.LastTransitionTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1PodCondition) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodCondition) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1PodCondition) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1PodCondition) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V1PodCondition) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodCondition) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V1PodCondition) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V1PodCondition) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value
func (o *V1PodCondition) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *V1PodCondition) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *V1PodCondition) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *V1PodCondition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1PodCondition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1PodCondition) SetType(v string) {
	o.Type = v
}

func (o V1PodCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PodCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastProbeTime) {
		toSerialize["lastProbeTime"] = o.LastProbeTime
	}
	if !IsNil(o.LastTransitionTime) {
		toSerialize["lastTransitionTime"] = o.LastTransitionTime
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableV1PodCondition struct {
	value *V1PodCondition
	isSet bool
}

func (v NullableV1PodCondition) Get() *V1PodCondition {
	return v.value
}

func (v *NullableV1PodCondition) Set(val *V1PodCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PodCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PodCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PodCondition(val *V1PodCondition) *NullableV1PodCondition {
	return &NullableV1PodCondition{value: val, isSet: true}
}

func (v NullableV1PodCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PodCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


