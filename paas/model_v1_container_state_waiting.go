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

// checks if the V1ContainerStateWaiting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ContainerStateWaiting{}

// V1ContainerStateWaiting ContainerStateWaiting is a waiting state of a container.
type V1ContainerStateWaiting struct {
	// Message regarding why the container is not yet running.
	Message *string `json:"message,omitempty"`
	// (brief) reason the container is not yet running.
	Reason *string `json:"reason,omitempty"`
}

// NewV1ContainerStateWaiting instantiates a new V1ContainerStateWaiting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ContainerStateWaiting() *V1ContainerStateWaiting {
	this := V1ContainerStateWaiting{}
	return &this
}

// NewV1ContainerStateWaitingWithDefaults instantiates a new V1ContainerStateWaiting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ContainerStateWaitingWithDefaults() *V1ContainerStateWaiting {
	this := V1ContainerStateWaiting{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1ContainerStateWaiting) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ContainerStateWaiting) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1ContainerStateWaiting) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1ContainerStateWaiting) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V1ContainerStateWaiting) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ContainerStateWaiting) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V1ContainerStateWaiting) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V1ContainerStateWaiting) SetReason(v string) {
	o.Reason = &v
}

func (o V1ContainerStateWaiting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ContainerStateWaiting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableV1ContainerStateWaiting struct {
	value *V1ContainerStateWaiting
	isSet bool
}

func (v NullableV1ContainerStateWaiting) Get() *V1ContainerStateWaiting {
	return v.value
}

func (v *NullableV1ContainerStateWaiting) Set(val *V1ContainerStateWaiting) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ContainerStateWaiting) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ContainerStateWaiting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ContainerStateWaiting(val *V1ContainerStateWaiting) *NullableV1ContainerStateWaiting {
	return &NullableV1ContainerStateWaiting{value: val, isSet: true}
}

func (v NullableV1ContainerStateWaiting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ContainerStateWaiting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


