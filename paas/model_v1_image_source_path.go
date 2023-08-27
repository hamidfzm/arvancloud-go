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

// checks if the V1ImageSourcePath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ImageSourcePath{}

// V1ImageSourcePath ImageSourcePath describes a path to be copied from a source image and its destination within the build directory.
type V1ImageSourcePath struct {
	// destinationDir is the relative directory within the build directory where files copied from the image are placed.
	DestinationDir string `json:"destinationDir"`
	// sourcePath is the absolute path of the file or directory inside the image to copy to the build directory.  If the source path ends in /. then the content of the directory will be copied, but the directory itself will not be created at the destination.
	SourcePath string `json:"sourcePath"`
}

// NewV1ImageSourcePath instantiates a new V1ImageSourcePath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ImageSourcePath(destinationDir string, sourcePath string) *V1ImageSourcePath {
	this := V1ImageSourcePath{}
	this.DestinationDir = destinationDir
	this.SourcePath = sourcePath
	return &this
}

// NewV1ImageSourcePathWithDefaults instantiates a new V1ImageSourcePath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ImageSourcePathWithDefaults() *V1ImageSourcePath {
	this := V1ImageSourcePath{}
	return &this
}

// GetDestinationDir returns the DestinationDir field value
func (o *V1ImageSourcePath) GetDestinationDir() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationDir
}

// GetDestinationDirOk returns a tuple with the DestinationDir field value
// and a boolean to check if the value has been set.
func (o *V1ImageSourcePath) GetDestinationDirOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationDir, true
}

// SetDestinationDir sets field value
func (o *V1ImageSourcePath) SetDestinationDir(v string) {
	o.DestinationDir = v
}

// GetSourcePath returns the SourcePath field value
func (o *V1ImageSourcePath) GetSourcePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourcePath
}

// GetSourcePathOk returns a tuple with the SourcePath field value
// and a boolean to check if the value has been set.
func (o *V1ImageSourcePath) GetSourcePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcePath, true
}

// SetSourcePath sets field value
func (o *V1ImageSourcePath) SetSourcePath(v string) {
	o.SourcePath = v
}

func (o V1ImageSourcePath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ImageSourcePath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destinationDir"] = o.DestinationDir
	toSerialize["sourcePath"] = o.SourcePath
	return toSerialize, nil
}

type NullableV1ImageSourcePath struct {
	value *V1ImageSourcePath
	isSet bool
}

func (v NullableV1ImageSourcePath) Get() *V1ImageSourcePath {
	return v.value
}

func (v *NullableV1ImageSourcePath) Set(val *V1ImageSourcePath) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ImageSourcePath) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ImageSourcePath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ImageSourcePath(val *V1ImageSourcePath) *NullableV1ImageSourcePath {
	return &NullableV1ImageSourcePath{value: val, isSet: true}
}

func (v NullableV1ImageSourcePath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ImageSourcePath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


