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

// checks if the V1CSIPersistentVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1CSIPersistentVolumeSource{}

// V1CSIPersistentVolumeSource Represents storage that is managed by an external CSI volume driver (Beta feature)
type V1CSIPersistentVolumeSource struct {
	ControllerPublishSecretRef *V1SecretReference `json:"controllerPublishSecretRef,omitempty"`
	// Driver is the name of the driver to use for this volume. Required.
	Driver string `json:"driver"`
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\".
	FsType *string `json:"fsType,omitempty"`
	NodePublishSecretRef *V1SecretReference `json:"nodePublishSecretRef,omitempty"`
	NodeStageSecretRef *V1SecretReference `json:"nodeStageSecretRef,omitempty"`
	// Optional: The value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).
	ReadOnly *bool `json:"readOnly,omitempty"`
	// Attributes of the volume to publish.
	VolumeAttributes map[string]interface{} `json:"volumeAttributes,omitempty"`
	// VolumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.
	VolumeHandle string `json:"volumeHandle"`
}

// NewV1CSIPersistentVolumeSource instantiates a new V1CSIPersistentVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CSIPersistentVolumeSource(driver string, volumeHandle string) *V1CSIPersistentVolumeSource {
	this := V1CSIPersistentVolumeSource{}
	this.Driver = driver
	this.VolumeHandle = volumeHandle
	return &this
}

// NewV1CSIPersistentVolumeSourceWithDefaults instantiates a new V1CSIPersistentVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CSIPersistentVolumeSourceWithDefaults() *V1CSIPersistentVolumeSource {
	this := V1CSIPersistentVolumeSource{}
	return &this
}

// GetControllerPublishSecretRef returns the ControllerPublishSecretRef field value if set, zero value otherwise.
func (o *V1CSIPersistentVolumeSource) GetControllerPublishSecretRef() V1SecretReference {
	if o == nil || IsNil(o.ControllerPublishSecretRef) {
		var ret V1SecretReference
		return ret
	}
	return *o.ControllerPublishSecretRef
}

// GetControllerPublishSecretRefOk returns a tuple with the ControllerPublishSecretRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CSIPersistentVolumeSource) GetControllerPublishSecretRefOk() (*V1SecretReference, bool) {
	if o == nil || IsNil(o.ControllerPublishSecretRef) {
		return nil, false
	}
	return o.ControllerPublishSecretRef, true
}

// HasControllerPublishSecretRef returns a boolean if a field has been set.
func (o *V1CSIPersistentVolumeSource) HasControllerPublishSecretRef() bool {
	if o != nil && !IsNil(o.ControllerPublishSecretRef) {
		return true
	}

	return false
}

// SetControllerPublishSecretRef gets a reference to the given V1SecretReference and assigns it to the ControllerPublishSecretRef field.
func (o *V1CSIPersistentVolumeSource) SetControllerPublishSecretRef(v V1SecretReference) {
	o.ControllerPublishSecretRef = &v
}

// GetDriver returns the Driver field value
func (o *V1CSIPersistentVolumeSource) GetDriver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Driver
}

// GetDriverOk returns a tuple with the Driver field value
// and a boolean to check if the value has been set.
func (o *V1CSIPersistentVolumeSource) GetDriverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Driver, true
}

// SetDriver sets field value
func (o *V1CSIPersistentVolumeSource) SetDriver(v string) {
	o.Driver = v
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *V1CSIPersistentVolumeSource) GetFsType() string {
	if o == nil || IsNil(o.FsType) {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CSIPersistentVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FsType) {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *V1CSIPersistentVolumeSource) HasFsType() bool {
	if o != nil && !IsNil(o.FsType) {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *V1CSIPersistentVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetNodePublishSecretRef returns the NodePublishSecretRef field value if set, zero value otherwise.
func (o *V1CSIPersistentVolumeSource) GetNodePublishSecretRef() V1SecretReference {
	if o == nil || IsNil(o.NodePublishSecretRef) {
		var ret V1SecretReference
		return ret
	}
	return *o.NodePublishSecretRef
}

// GetNodePublishSecretRefOk returns a tuple with the NodePublishSecretRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CSIPersistentVolumeSource) GetNodePublishSecretRefOk() (*V1SecretReference, bool) {
	if o == nil || IsNil(o.NodePublishSecretRef) {
		return nil, false
	}
	return o.NodePublishSecretRef, true
}

// HasNodePublishSecretRef returns a boolean if a field has been set.
func (o *V1CSIPersistentVolumeSource) HasNodePublishSecretRef() bool {
	if o != nil && !IsNil(o.NodePublishSecretRef) {
		return true
	}

	return false
}

// SetNodePublishSecretRef gets a reference to the given V1SecretReference and assigns it to the NodePublishSecretRef field.
func (o *V1CSIPersistentVolumeSource) SetNodePublishSecretRef(v V1SecretReference) {
	o.NodePublishSecretRef = &v
}

// GetNodeStageSecretRef returns the NodeStageSecretRef field value if set, zero value otherwise.
func (o *V1CSIPersistentVolumeSource) GetNodeStageSecretRef() V1SecretReference {
	if o == nil || IsNil(o.NodeStageSecretRef) {
		var ret V1SecretReference
		return ret
	}
	return *o.NodeStageSecretRef
}

// GetNodeStageSecretRefOk returns a tuple with the NodeStageSecretRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CSIPersistentVolumeSource) GetNodeStageSecretRefOk() (*V1SecretReference, bool) {
	if o == nil || IsNil(o.NodeStageSecretRef) {
		return nil, false
	}
	return o.NodeStageSecretRef, true
}

// HasNodeStageSecretRef returns a boolean if a field has been set.
func (o *V1CSIPersistentVolumeSource) HasNodeStageSecretRef() bool {
	if o != nil && !IsNil(o.NodeStageSecretRef) {
		return true
	}

	return false
}

// SetNodeStageSecretRef gets a reference to the given V1SecretReference and assigns it to the NodeStageSecretRef field.
func (o *V1CSIPersistentVolumeSource) SetNodeStageSecretRef(v V1SecretReference) {
	o.NodeStageSecretRef = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *V1CSIPersistentVolumeSource) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CSIPersistentVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *V1CSIPersistentVolumeSource) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *V1CSIPersistentVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetVolumeAttributes returns the VolumeAttributes field value if set, zero value otherwise.
func (o *V1CSIPersistentVolumeSource) GetVolumeAttributes() map[string]interface{} {
	if o == nil || IsNil(o.VolumeAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.VolumeAttributes
}

// GetVolumeAttributesOk returns a tuple with the VolumeAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CSIPersistentVolumeSource) GetVolumeAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.VolumeAttributes) {
		return map[string]interface{}{}, false
	}
	return o.VolumeAttributes, true
}

// HasVolumeAttributes returns a boolean if a field has been set.
func (o *V1CSIPersistentVolumeSource) HasVolumeAttributes() bool {
	if o != nil && !IsNil(o.VolumeAttributes) {
		return true
	}

	return false
}

// SetVolumeAttributes gets a reference to the given map[string]interface{} and assigns it to the VolumeAttributes field.
func (o *V1CSIPersistentVolumeSource) SetVolumeAttributes(v map[string]interface{}) {
	o.VolumeAttributes = v
}

// GetVolumeHandle returns the VolumeHandle field value
func (o *V1CSIPersistentVolumeSource) GetVolumeHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeHandle
}

// GetVolumeHandleOk returns a tuple with the VolumeHandle field value
// and a boolean to check if the value has been set.
func (o *V1CSIPersistentVolumeSource) GetVolumeHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeHandle, true
}

// SetVolumeHandle sets field value
func (o *V1CSIPersistentVolumeSource) SetVolumeHandle(v string) {
	o.VolumeHandle = v
}

func (o V1CSIPersistentVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1CSIPersistentVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ControllerPublishSecretRef) {
		toSerialize["controllerPublishSecretRef"] = o.ControllerPublishSecretRef
	}
	toSerialize["driver"] = o.Driver
	if !IsNil(o.FsType) {
		toSerialize["fsType"] = o.FsType
	}
	if !IsNil(o.NodePublishSecretRef) {
		toSerialize["nodePublishSecretRef"] = o.NodePublishSecretRef
	}
	if !IsNil(o.NodeStageSecretRef) {
		toSerialize["nodeStageSecretRef"] = o.NodeStageSecretRef
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.VolumeAttributes) {
		toSerialize["volumeAttributes"] = o.VolumeAttributes
	}
	toSerialize["volumeHandle"] = o.VolumeHandle
	return toSerialize, nil
}

type NullableV1CSIPersistentVolumeSource struct {
	value *V1CSIPersistentVolumeSource
	isSet bool
}

func (v NullableV1CSIPersistentVolumeSource) Get() *V1CSIPersistentVolumeSource {
	return v.value
}

func (v *NullableV1CSIPersistentVolumeSource) Set(val *V1CSIPersistentVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CSIPersistentVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CSIPersistentVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CSIPersistentVolumeSource(val *V1CSIPersistentVolumeSource) *NullableV1CSIPersistentVolumeSource {
	return &NullableV1CSIPersistentVolumeSource{value: val, isSet: true}
}

func (v NullableV1CSIPersistentVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CSIPersistentVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


