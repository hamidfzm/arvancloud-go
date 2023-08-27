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

// checks if the V1ConfigMapBuildSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ConfigMapBuildSource{}

// V1ConfigMapBuildSource ConfigMapBuildSource describes a configmap and its destination directory that will be used only at the build time. The content of the configmap referenced here will be copied into the destination directory instead of mounting.
type V1ConfigMapBuildSource struct {
	ConfigMap V1LocalObjectReference `json:"configMap"`
	// destinationDir is the directory where the files from the configmap should be available for the build time. For the Source build strategy, these will be injected into a container where the assemble script runs. For the Docker build strategy, these will be copied into the build directory, where the Dockerfile is located, so users can ADD or COPY them during docker build.
	DestinationDir *string `json:"destinationDir,omitempty"`
}

// NewV1ConfigMapBuildSource instantiates a new V1ConfigMapBuildSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ConfigMapBuildSource(configMap V1LocalObjectReference) *V1ConfigMapBuildSource {
	this := V1ConfigMapBuildSource{}
	this.ConfigMap = configMap
	return &this
}

// NewV1ConfigMapBuildSourceWithDefaults instantiates a new V1ConfigMapBuildSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ConfigMapBuildSourceWithDefaults() *V1ConfigMapBuildSource {
	this := V1ConfigMapBuildSource{}
	return &this
}

// GetConfigMap returns the ConfigMap field value
func (o *V1ConfigMapBuildSource) GetConfigMap() V1LocalObjectReference {
	if o == nil {
		var ret V1LocalObjectReference
		return ret
	}

	return o.ConfigMap
}

// GetConfigMapOk returns a tuple with the ConfigMap field value
// and a boolean to check if the value has been set.
func (o *V1ConfigMapBuildSource) GetConfigMapOk() (*V1LocalObjectReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigMap, true
}

// SetConfigMap sets field value
func (o *V1ConfigMapBuildSource) SetConfigMap(v V1LocalObjectReference) {
	o.ConfigMap = v
}

// GetDestinationDir returns the DestinationDir field value if set, zero value otherwise.
func (o *V1ConfigMapBuildSource) GetDestinationDir() string {
	if o == nil || IsNil(o.DestinationDir) {
		var ret string
		return ret
	}
	return *o.DestinationDir
}

// GetDestinationDirOk returns a tuple with the DestinationDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigMapBuildSource) GetDestinationDirOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationDir) {
		return nil, false
	}
	return o.DestinationDir, true
}

// HasDestinationDir returns a boolean if a field has been set.
func (o *V1ConfigMapBuildSource) HasDestinationDir() bool {
	if o != nil && !IsNil(o.DestinationDir) {
		return true
	}

	return false
}

// SetDestinationDir gets a reference to the given string and assigns it to the DestinationDir field.
func (o *V1ConfigMapBuildSource) SetDestinationDir(v string) {
	o.DestinationDir = &v
}

func (o V1ConfigMapBuildSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ConfigMapBuildSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configMap"] = o.ConfigMap
	if !IsNil(o.DestinationDir) {
		toSerialize["destinationDir"] = o.DestinationDir
	}
	return toSerialize, nil
}

type NullableV1ConfigMapBuildSource struct {
	value *V1ConfigMapBuildSource
	isSet bool
}

func (v NullableV1ConfigMapBuildSource) Get() *V1ConfigMapBuildSource {
	return v.value
}

func (v *NullableV1ConfigMapBuildSource) Set(val *V1ConfigMapBuildSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ConfigMapBuildSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ConfigMapBuildSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ConfigMapBuildSource(val *V1ConfigMapBuildSource) *NullableV1ConfigMapBuildSource {
	return &NullableV1ConfigMapBuildSource{value: val, isSet: true}
}

func (v NullableV1ConfigMapBuildSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ConfigMapBuildSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


