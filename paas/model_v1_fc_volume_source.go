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

// checks if the V1FCVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1FCVolumeSource{}

// V1FCVolumeSource Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
type V1FCVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified.
	FsType *string `json:"fsType,omitempty"`
	// Optional: FC target lun number
	Lun *int32 `json:"lun,omitempty"`
	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// Optional: FC target worldwide names (WWNs)
	TargetWWNs []string `json:"targetWWNs,omitempty"`
	// Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
	Wwids []string `json:"wwids,omitempty"`
}

// NewV1FCVolumeSource instantiates a new V1FCVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1FCVolumeSource() *V1FCVolumeSource {
	this := V1FCVolumeSource{}
	return &this
}

// NewV1FCVolumeSourceWithDefaults instantiates a new V1FCVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1FCVolumeSourceWithDefaults() *V1FCVolumeSource {
	this := V1FCVolumeSource{}
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *V1FCVolumeSource) GetFsType() string {
	if o == nil || IsNil(o.FsType) {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1FCVolumeSource) GetFsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FsType) {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *V1FCVolumeSource) HasFsType() bool {
	if o != nil && !IsNil(o.FsType) {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *V1FCVolumeSource) SetFsType(v string) {
	o.FsType = &v
}

// GetLun returns the Lun field value if set, zero value otherwise.
func (o *V1FCVolumeSource) GetLun() int32 {
	if o == nil || IsNil(o.Lun) {
		var ret int32
		return ret
	}
	return *o.Lun
}

// GetLunOk returns a tuple with the Lun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1FCVolumeSource) GetLunOk() (*int32, bool) {
	if o == nil || IsNil(o.Lun) {
		return nil, false
	}
	return o.Lun, true
}

// HasLun returns a boolean if a field has been set.
func (o *V1FCVolumeSource) HasLun() bool {
	if o != nil && !IsNil(o.Lun) {
		return true
	}

	return false
}

// SetLun gets a reference to the given int32 and assigns it to the Lun field.
func (o *V1FCVolumeSource) SetLun(v int32) {
	o.Lun = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *V1FCVolumeSource) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1FCVolumeSource) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *V1FCVolumeSource) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *V1FCVolumeSource) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetTargetWWNs returns the TargetWWNs field value if set, zero value otherwise.
func (o *V1FCVolumeSource) GetTargetWWNs() []string {
	if o == nil || IsNil(o.TargetWWNs) {
		var ret []string
		return ret
	}
	return o.TargetWWNs
}

// GetTargetWWNsOk returns a tuple with the TargetWWNs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1FCVolumeSource) GetTargetWWNsOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetWWNs) {
		return nil, false
	}
	return o.TargetWWNs, true
}

// HasTargetWWNs returns a boolean if a field has been set.
func (o *V1FCVolumeSource) HasTargetWWNs() bool {
	if o != nil && !IsNil(o.TargetWWNs) {
		return true
	}

	return false
}

// SetTargetWWNs gets a reference to the given []string and assigns it to the TargetWWNs field.
func (o *V1FCVolumeSource) SetTargetWWNs(v []string) {
	o.TargetWWNs = v
}

// GetWwids returns the Wwids field value if set, zero value otherwise.
func (o *V1FCVolumeSource) GetWwids() []string {
	if o == nil || IsNil(o.Wwids) {
		var ret []string
		return ret
	}
	return o.Wwids
}

// GetWwidsOk returns a tuple with the Wwids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1FCVolumeSource) GetWwidsOk() ([]string, bool) {
	if o == nil || IsNil(o.Wwids) {
		return nil, false
	}
	return o.Wwids, true
}

// HasWwids returns a boolean if a field has been set.
func (o *V1FCVolumeSource) HasWwids() bool {
	if o != nil && !IsNil(o.Wwids) {
		return true
	}

	return false
}

// SetWwids gets a reference to the given []string and assigns it to the Wwids field.
func (o *V1FCVolumeSource) SetWwids(v []string) {
	o.Wwids = v
}

func (o V1FCVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1FCVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsType) {
		toSerialize["fsType"] = o.FsType
	}
	if !IsNil(o.Lun) {
		toSerialize["lun"] = o.Lun
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if !IsNil(o.TargetWWNs) {
		toSerialize["targetWWNs"] = o.TargetWWNs
	}
	if !IsNil(o.Wwids) {
		toSerialize["wwids"] = o.Wwids
	}
	return toSerialize, nil
}

type NullableV1FCVolumeSource struct {
	value *V1FCVolumeSource
	isSet bool
}

func (v NullableV1FCVolumeSource) Get() *V1FCVolumeSource {
	return v.value
}

func (v *NullableV1FCVolumeSource) Set(val *V1FCVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableV1FCVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableV1FCVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1FCVolumeSource(val *V1FCVolumeSource) *NullableV1FCVolumeSource {
	return &NullableV1FCVolumeSource{value: val, isSet: true}
}

func (v NullableV1FCVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1FCVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


