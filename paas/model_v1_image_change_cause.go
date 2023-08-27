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

// checks if the V1ImageChangeCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ImageChangeCause{}

// V1ImageChangeCause ImageChangeCause contains information about the image that triggered a build
type V1ImageChangeCause struct {
	FromRef *V1ObjectReference `json:"fromRef,omitempty"`
	// imageID is the ID of the image that triggered a a new build.
	ImageID *string `json:"imageID,omitempty"`
}

// NewV1ImageChangeCause instantiates a new V1ImageChangeCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImageChangeCause() *V1ImageChangeCause {
	this := V1ImageChangeCause{}
	return &this
}

// NewV1ImageChangeCauseWithDefaults instantiates a new V1ImageChangeCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImageChangeCauseWithDefaults() *V1ImageChangeCause {
	this := V1ImageChangeCause{}
	return &this
}

// GetFromRef returns the FromRef field value if set, zero value otherwise.
func (o *V1ImageChangeCause) GetFromRef() V1ObjectReference {
	if o == nil || IsNil(o.FromRef) {
		var ret V1ObjectReference
		return ret
	}
	return *o.FromRef
}

// GetFromRefOk returns a tuple with the FromRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageChangeCause) GetFromRefOk() (*V1ObjectReference, bool) {
	if o == nil || IsNil(o.FromRef) {
		return nil, false
	}
	return o.FromRef, true
}

// HasFromRef returns a boolean if a field has been set.
func (o *V1ImageChangeCause) HasFromRef() bool {
	if o != nil && !IsNil(o.FromRef) {
		return true
	}

	return false
}

// SetFromRef gets a reference to the given V1ObjectReference and assigns it to the FromRef field.
func (o *V1ImageChangeCause) SetFromRef(v V1ObjectReference) {
	o.FromRef = &v
}

// GetImageID returns the ImageID field value if set, zero value otherwise.
func (o *V1ImageChangeCause) GetImageID() string {
	if o == nil || IsNil(o.ImageID) {
		var ret string
		return ret
	}
	return *o.ImageID
}

// GetImageIDOk returns a tuple with the ImageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ImageChangeCause) GetImageIDOk() (*string, bool) {
	if o == nil || IsNil(o.ImageID) {
		return nil, false
	}
	return o.ImageID, true
}

// HasImageID returns a boolean if a field has been set.
func (o *V1ImageChangeCause) HasImageID() bool {
	if o != nil && !IsNil(o.ImageID) {
		return true
	}

	return false
}

// SetImageID gets a reference to the given string and assigns it to the ImageID field.
func (o *V1ImageChangeCause) SetImageID(v string) {
	o.ImageID = &v
}

func (o V1ImageChangeCause) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ImageChangeCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromRef) {
		toSerialize["fromRef"] = o.FromRef
	}
	if !IsNil(o.ImageID) {
		toSerialize["imageID"] = o.ImageID
	}
	return toSerialize, nil
}

type NullableV1ImageChangeCause struct {
	value *V1ImageChangeCause
	isSet bool
}

func (v NullableV1ImageChangeCause) Get() *V1ImageChangeCause {
	return v.value
}

func (v *NullableV1ImageChangeCause) Set(val *V1ImageChangeCause) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImageChangeCause) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImageChangeCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImageChangeCause(val *V1ImageChangeCause) *NullableV1ImageChangeCause {
	return &NullableV1ImageChangeCause{value: val, isSet: true}
}

func (v NullableV1ImageChangeCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImageChangeCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

