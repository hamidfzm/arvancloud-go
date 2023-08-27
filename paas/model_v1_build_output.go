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

// checks if the V1BuildOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1BuildOutput{}

// V1BuildOutput BuildOutput is input to a build strategy and describes the Docker image that the strategy should produce.
type V1BuildOutput struct {
	// imageLabels define a list of labels that are applied to the resulting image. If there are multiple labels with the same name then the last one in the list is used.
	ImageLabels []V1ImageLabel `json:"imageLabels,omitempty"`
	PushSecret *V1LocalObjectReference `json:"pushSecret,omitempty"`
	To *V1ObjectReference `json:"to,omitempty"`
}

// NewV1BuildOutput instantiates a new V1BuildOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1BuildOutput() *V1BuildOutput {
	this := V1BuildOutput{}
	return &this
}

// NewV1BuildOutputWithDefaults instantiates a new V1BuildOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1BuildOutputWithDefaults() *V1BuildOutput {
	this := V1BuildOutput{}
	return &this
}

// GetImageLabels returns the ImageLabels field value if set, zero value otherwise.
func (o *V1BuildOutput) GetImageLabels() []V1ImageLabel {
	if o == nil || IsNil(o.ImageLabels) {
		var ret []V1ImageLabel
		return ret
	}
	return o.ImageLabels
}

// GetImageLabelsOk returns a tuple with the ImageLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildOutput) GetImageLabelsOk() ([]V1ImageLabel, bool) {
	if o == nil || IsNil(o.ImageLabels) {
		return nil, false
	}
	return o.ImageLabels, true
}

// HasImageLabels returns a boolean if a field has been set.
func (o *V1BuildOutput) HasImageLabels() bool {
	if o != nil && !IsNil(o.ImageLabels) {
		return true
	}

	return false
}

// SetImageLabels gets a reference to the given []V1ImageLabel and assigns it to the ImageLabels field.
func (o *V1BuildOutput) SetImageLabels(v []V1ImageLabel) {
	o.ImageLabels = v
}

// GetPushSecret returns the PushSecret field value if set, zero value otherwise.
func (o *V1BuildOutput) GetPushSecret() V1LocalObjectReference {
	if o == nil || IsNil(o.PushSecret) {
		var ret V1LocalObjectReference
		return ret
	}
	return *o.PushSecret
}

// GetPushSecretOk returns a tuple with the PushSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildOutput) GetPushSecretOk() (*V1LocalObjectReference, bool) {
	if o == nil || IsNil(o.PushSecret) {
		return nil, false
	}
	return o.PushSecret, true
}

// HasPushSecret returns a boolean if a field has been set.
func (o *V1BuildOutput) HasPushSecret() bool {
	if o != nil && !IsNil(o.PushSecret) {
		return true
	}

	return false
}

// SetPushSecret gets a reference to the given V1LocalObjectReference and assigns it to the PushSecret field.
func (o *V1BuildOutput) SetPushSecret(v V1LocalObjectReference) {
	o.PushSecret = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *V1BuildOutput) GetTo() V1ObjectReference {
	if o == nil || IsNil(o.To) {
		var ret V1ObjectReference
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1BuildOutput) GetToOk() (*V1ObjectReference, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *V1BuildOutput) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given V1ObjectReference and assigns it to the To field.
func (o *V1BuildOutput) SetTo(v V1ObjectReference) {
	o.To = &v
}

func (o V1BuildOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1BuildOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageLabels) {
		toSerialize["imageLabels"] = o.ImageLabels
	}
	if !IsNil(o.PushSecret) {
		toSerialize["pushSecret"] = o.PushSecret
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableV1BuildOutput struct {
	value *V1BuildOutput
	isSet bool
}

func (v NullableV1BuildOutput) Get() *V1BuildOutput {
	return v.value
}

func (v *NullableV1BuildOutput) Set(val *V1BuildOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableV1BuildOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableV1BuildOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1BuildOutput(val *V1BuildOutput) *NullableV1BuildOutput {
	return &NullableV1BuildOutput{value: val, isSet: true}
}

func (v NullableV1BuildOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1BuildOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


