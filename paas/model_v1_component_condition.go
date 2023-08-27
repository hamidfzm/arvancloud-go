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

// checks if the V1ComponentCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ComponentCondition{}

// V1ComponentCondition Information about the condition of a component.
type V1ComponentCondition struct {
	// Condition error code for a component. For example, a health check error code.
	Error *string `json:"error,omitempty"`
	// Message about the condition for a component. For example, information about a health check.
	Message *string `json:"message,omitempty"`
	// Status of the condition for a component. Valid values for \"Healthy\": \"True\", \"False\", or \"Unknown\".
	Status string `json:"status"`
	// Type of condition for a component. Valid value: \"Healthy\"
	Type string `json:"type"`
}

// NewV1ComponentCondition instantiates a new V1ComponentCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ComponentCondition(status string, type_ string) *V1ComponentCondition {
	this := V1ComponentCondition{}
	this.Status = status
	this.Type = type_
	return &this
}

// NewV1ComponentConditionWithDefaults instantiates a new V1ComponentCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ComponentConditionWithDefaults() *V1ComponentCondition {
	this := V1ComponentCondition{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V1ComponentCondition) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ComponentCondition) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V1ComponentCondition) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *V1ComponentCondition) SetError(v string) {
	o.Error = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1ComponentCondition) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ComponentCondition) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1ComponentCondition) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1ComponentCondition) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value
func (o *V1ComponentCondition) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *V1ComponentCondition) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *V1ComponentCondition) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *V1ComponentCondition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1ComponentCondition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1ComponentCondition) SetType(v string) {
	o.Type = v
}

func (o V1ComponentCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ComponentCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableV1ComponentCondition struct {
	value *V1ComponentCondition
	isSet bool
}

func (v NullableV1ComponentCondition) Get() *V1ComponentCondition {
	return v.value
}

func (v *NullableV1ComponentCondition) Set(val *V1ComponentCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ComponentCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ComponentCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ComponentCondition(val *V1ComponentCondition) *NullableV1ComponentCondition {
	return &NullableV1ComponentCondition{value: val, isSet: true}
}

func (v NullableV1ComponentCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ComponentCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


