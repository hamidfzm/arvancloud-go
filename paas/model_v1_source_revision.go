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

// checks if the V1SourceRevision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1SourceRevision{}

// V1SourceRevision SourceRevision is the revision or commit information from the source for the build
type V1SourceRevision struct {
	Git *V1GitSourceRevision `json:"git,omitempty"`
	// type of the build source, may be one of 'Source', 'Dockerfile', 'Binary', or 'Images'
	Type string `json:"type"`
}

// NewV1SourceRevision instantiates a new V1SourceRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SourceRevision(type_ string) *V1SourceRevision {
	this := V1SourceRevision{}
	this.Type = type_
	return &this
}

// NewV1SourceRevisionWithDefaults instantiates a new V1SourceRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SourceRevisionWithDefaults() *V1SourceRevision {
	this := V1SourceRevision{}
	return &this
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *V1SourceRevision) GetGit() V1GitSourceRevision {
	if o == nil || IsNil(o.Git) {
		var ret V1GitSourceRevision
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SourceRevision) GetGitOk() (*V1GitSourceRevision, bool) {
	if o == nil || IsNil(o.Git) {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *V1SourceRevision) HasGit() bool {
	if o != nil && !IsNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given V1GitSourceRevision and assigns it to the Git field.
func (o *V1SourceRevision) SetGit(v V1GitSourceRevision) {
	o.Git = &v
}

// GetType returns the Type field value
func (o *V1SourceRevision) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V1SourceRevision) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V1SourceRevision) SetType(v string) {
	o.Type = v
}

func (o V1SourceRevision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1SourceRevision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Git) {
		toSerialize["git"] = o.Git
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableV1SourceRevision struct {
	value *V1SourceRevision
	isSet bool
}

func (v NullableV1SourceRevision) Get() *V1SourceRevision {
	return v.value
}

func (v *NullableV1SourceRevision) Set(val *V1SourceRevision) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SourceRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SourceRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SourceRevision(val *V1SourceRevision) *NullableV1SourceRevision {
	return &NullableV1SourceRevision{value: val, isSet: true}
}

func (v NullableV1SourceRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SourceRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


