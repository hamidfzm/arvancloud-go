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

// checks if the V1DeploymentCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1DeploymentCause{}

// V1DeploymentCause DeploymentCause captures information about a particular cause of a deployment.
type V1DeploymentCause struct {
	ImageTrigger *V1DeploymentCauseImageTrigger `json:"imageTrigger,omitempty"`
	// Type of the trigger that resulted in the creation of a new deployment
	Type string `json:"type"`
}

// NewV1DeploymentCause instantiates a new V1DeploymentCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1DeploymentCause(type_ string) *V1DeploymentCause {
	this := V1DeploymentCause{}
	this.Type = type_
	return &this
}

// NewV1DeploymentCauseWithDefaults instantiates a new V1DeploymentCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1DeploymentCauseWithDefaults() *V1DeploymentCause {
	this := V1DeploymentCause{}
	return &this
}

// GetImageTrigger returns the ImageTrigger field value if set, zero value otherwise.
func (o *V1DeploymentCause) GetImageTrigger() V1DeploymentCauseImageTrigger {
	if o == nil || IsNil(o.ImageTrigger) {
		var ret V1DeploymentCauseImageTrigger
		return ret
	}
	return *o.ImageTrigger
}

// GetImageTriggerOk returns a tuple with the ImageTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1DeploymentCause) GetImageTriggerOk() (*V1DeploymentCauseImageTrigger, bool) {
	if o == nil || IsNil(o.ImageTrigger) {
		return nil, false
	}
	return o.ImageTrigger, true
}

// HasImageTrigger returns a boolean if a field has been set.
func (o *V1DeploymentCause) HasImageTrigger() bool {
	if o != nil && !IsNil(o.ImageTrigger) {
		return true
	}

	return false
}

// SetImageTrigger gets a reference to the given V1DeploymentCauseImageTrigger and assigns it to the ImageTrigger field.
func (o *V1DeploymentCause) SetImageTrigger(v V1DeploymentCauseImageTrigger) {
	o.ImageTrigger = &v
}

// GetType returns the Type field value
func (o *V1DeploymentCause) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1DeploymentCause) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1DeploymentCause) SetType(v string) {
	o.Type = v
}

func (o V1DeploymentCause) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1DeploymentCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageTrigger) {
		toSerialize["imageTrigger"] = o.ImageTrigger
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableV1DeploymentCause struct {
	value *V1DeploymentCause
	isSet bool
}

func (v NullableV1DeploymentCause) Get() *V1DeploymentCause {
	return v.value
}

func (v *NullableV1DeploymentCause) Set(val *V1DeploymentCause) {
	v.value = val
	v.isSet = true
}

func (v NullableV1DeploymentCause) IsSet() bool {
	return v.isSet
}

func (v *NullableV1DeploymentCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1DeploymentCause(val *V1DeploymentCause) *NullableV1DeploymentCause {
	return &NullableV1DeploymentCause{value: val, isSet: true}
}

func (v NullableV1DeploymentCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1DeploymentCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


