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

// checks if the V1CinderVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1CinderVolumeSource{}

// V1CinderVolumeSource Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
type V1CinderVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	FsType *string `json:"fsType,omitempty"`
	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	ReadOnly *bool `json:"readOnly,omitempty"`
	SecretRef *V1LocalObjectReference `json:"secretRef,omitempty"`
	// volume id used to identify the volume in cinder More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
	VolumeID string `json:"volumeID"`
}

// NewV1CinderVolumeSource instantiates a new V1CinderVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CinderVolumeSource(volumeID string) *V1CinderVolumeSource {
	this := V1CinderVolumeSource{}
	this.VolumeID = volumeID
	return &this
}

// NewV1CinderVolumeSourceWithDefaults instantiates a new V1CinderVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CinderVolumeSourceWithDefaults() *V1CinderVolumeSource {
	this := V1CinderVolumeSource{}
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *V1CinderVolumeSource) GetFsType() string {
	if o == nil || IsNil(o.FsType) {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CinderVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FsType) {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *V1CinderVolumeSource) HasFsType() bool {
	if o != nil && !IsNil(o.FsType) {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *V1CinderVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *V1CinderVolumeSource) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CinderVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *V1CinderVolumeSource) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *V1CinderVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetSecretRef returns the SecretRef field value if set, zero value otherwise.
func (o *V1CinderVolumeSource) GetSecretRef() V1LocalObjectReference {
	if o == nil || IsNil(o.SecretRef) {
		var ret V1LocalObjectReference
		return ret
	}
	return *o.SecretRef
}

// GetSecretRefOk returns a tuple with the SecretRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CinderVolumeSource) GetSecretRefOk() (*V1LocalObjectReference, bool) {
	if o == nil || IsNil(o.SecretRef) {
		return nil, false
	}
	return o.SecretRef, true
}

// HasSecretRef returns a boolean if a field has been set.
func (o *V1CinderVolumeSource) HasSecretRef() bool {
	if o != nil && !IsNil(o.SecretRef) {
		return true
	}

	return false
}

// SetSecretRef gets a reference to the given V1LocalObjectReference and assigns it to the SecretRef field.
func (o *V1CinderVolumeSource) SetSecretRef(v V1LocalObjectReference) {
	o.SecretRef = &v
}

// GetVolumeID returns the VolumeID field value
func (o *V1CinderVolumeSource) GetVolumeID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeID
}

// GetVolumeIDOk returns a tuple with the VolumeID field value
// and a boolean to check if the value has been set.
func (o *V1CinderVolumeSource) GetVolumeIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeID, true
}

// SetVolumeID sets field value
func (o *V1CinderVolumeSource) SetVolumeID(v string) {
	o.VolumeID = v
}

func (o V1CinderVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1CinderVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsType) {
		toSerialize["fsType"] = o.FsType
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.SecretRef) {
		toSerialize["secretRef"] = o.SecretRef
	}
	toSerialize["volumeID"] = o.VolumeID
	return toSerialize, nil
}

type NullableV1CinderVolumeSource struct {
	value *V1CinderVolumeSource
	isSet bool
}

func (v NullableV1CinderVolumeSource) Get() *V1CinderVolumeSource {
	return v.value
}

func (v *NullableV1CinderVolumeSource) Set(val *V1CinderVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CinderVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CinderVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CinderVolumeSource(val *V1CinderVolumeSource) *NullableV1CinderVolumeSource {
	return &NullableV1CinderVolumeSource{value: val, isSet: true}
}

func (v NullableV1CinderVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CinderVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


