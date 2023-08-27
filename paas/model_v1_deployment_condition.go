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

// checks if the V1DeploymentCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DeploymentCondition{}

// V1DeploymentCondition DeploymentCondition describes the state of a deployment at a certain point.
type V1DeploymentCondition struct {
	// Last time the condition transitioned from one status to another.
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`
	// The last time this condition was updated.
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`
	// A human readable message indicating details about the transition.
	Message *string `json:"message,omitempty"`
	// The reason for the condition's last transition.
	Reason *string `json:"reason,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status"`
	// Type of deployment condition.
	Type string `json:"type"`
}

// NewV1DeploymentCondition instantiates a new V1DeploymentCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DeploymentCondition(status string, type_ string) *V1DeploymentCondition {
	this := V1DeploymentCondition{}
	this.Status = status
	this.Type = type_
	return &this
}

// NewV1DeploymentConditionWithDefaults instantiates a new V1DeploymentCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DeploymentConditionWithDefaults() *V1DeploymentCondition {
	this := V1DeploymentCondition{}
	return &this
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *V1DeploymentCondition) GetLastTransitionTime() string {
	if o == nil || IsNil(o.LastTransitionTime) {
		var ret string
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentCondition) GetLastTransitionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastTransitionTime) {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *V1DeploymentCondition) HasLastTransitionTime() bool {
	if o != nil && !IsNil(o.LastTransitionTime) {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given string and assigns it to the LastTransitionTime field.
func (o *V1DeploymentCondition) SetLastTransitionTime(v string) {
	o.LastTransitionTime = &v
}

// GetLastUpdateTime returns the LastUpdateTime field value if set, zero value otherwise.
func (o *V1DeploymentCondition) GetLastUpdateTime() string {
	if o == nil || IsNil(o.LastUpdateTime) {
		var ret string
		return ret
	}
	return *o.LastUpdateTime
}

// GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentCondition) GetLastUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdateTime) {
		return nil, false
	}
	return o.LastUpdateTime, true
}

// HasLastUpdateTime returns a boolean if a field has been set.
func (o *V1DeploymentCondition) HasLastUpdateTime() bool {
	if o != nil && !IsNil(o.LastUpdateTime) {
		return true
	}

	return false
}

// SetLastUpdateTime gets a reference to the given string and assigns it to the LastUpdateTime field.
func (o *V1DeploymentCondition) SetLastUpdateTime(v string) {
	o.LastUpdateTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1DeploymentCondition) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentCondition) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1DeploymentCondition) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1DeploymentCondition) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V1DeploymentCondition) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentCondition) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V1DeploymentCondition) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V1DeploymentCondition) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value
func (o *V1DeploymentCondition) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *V1DeploymentCondition) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *V1DeploymentCondition) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *V1DeploymentCondition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1DeploymentCondition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1DeploymentCondition) SetType(v string) {
	o.Type = v
}

func (o V1DeploymentCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DeploymentCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastTransitionTime) {
		toSerialize["lastTransitionTime"] = o.LastTransitionTime
	}
	if !IsNil(o.LastUpdateTime) {
		toSerialize["lastUpdateTime"] = o.LastUpdateTime
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

type NullableV1DeploymentCondition struct {
	value *V1DeploymentCondition
	isSet bool
}

func (v NullableV1DeploymentCondition) Get() *V1DeploymentCondition {
	return v.value
}

func (v *NullableV1DeploymentCondition) Set(val *V1DeploymentCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DeploymentCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DeploymentCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DeploymentCondition(val *V1DeploymentCondition) *NullableV1DeploymentCondition {
	return &NullableV1DeploymentCondition{value: val, isSet: true}
}

func (v NullableV1DeploymentCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DeploymentCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


