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

// checks if the V1KeyToPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1KeyToPath{}

// V1KeyToPath Maps a string key to a path within a volume.
type V1KeyToPath struct {
	// The key to project.
	Key string `json:"key"`
	// Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Mode *int32 `json:"mode,omitempty"`
	// The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
	Path string `json:"path"`
}

// NewV1KeyToPath instantiates a new V1KeyToPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1KeyToPath(key string, path string) *V1KeyToPath {
	this := V1KeyToPath{}
	this.Key = key
	this.Path = path
	return &this
}

// NewV1KeyToPathWithDefaults instantiates a new V1KeyToPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1KeyToPathWithDefaults() *V1KeyToPath {
	this := V1KeyToPath{}
	return &this
}

// GetKey returns the Key field value
func (o *V1KeyToPath) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *V1KeyToPath) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *V1KeyToPath) SetKey(v string) {
	o.Key = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *V1KeyToPath) GetMode() int32 {
	if o == nil || IsNil(o.Mode) {
		var ret int32
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1KeyToPath) GetModeOk() (*int32, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *V1KeyToPath) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given int32 and assigns it to the Mode field.
func (o *V1KeyToPath) SetMode(v int32) {
	o.Mode = &v
}

// GetPath returns the Path field value
func (o *V1KeyToPath) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *V1KeyToPath) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *V1KeyToPath) SetPath(v string) {
	o.Path = v
}

func (o V1KeyToPath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1KeyToPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

type NullableV1KeyToPath struct {
	value *V1KeyToPath
	isSet bool
}

func (v NullableV1KeyToPath) Get() *V1KeyToPath {
	return v.value
}

func (v *NullableV1KeyToPath) Set(val *V1KeyToPath) {
	v.value = val
	v.isSet = true
}

func (v NullableV1KeyToPath) IsSet() bool {
	return v.isSet
}

func (v *NullableV1KeyToPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1KeyToPath(val *V1KeyToPath) *NullableV1KeyToPath {
	return &NullableV1KeyToPath{value: val, isSet: true}
}

func (v NullableV1KeyToPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1KeyToPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


