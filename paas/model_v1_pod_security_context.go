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

// checks if the V1PodSecurityContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1PodSecurityContext{}

// V1PodSecurityContext PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext.  Field values of container.securityContext take precedence over field values of PodSecurityContext.
type V1PodSecurityContext struct {
	// A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw 
	FsGroup *int64 `json:"fsGroup,omitempty"`
	// The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	RunAsGroup *int64 `json:"runAsGroup,omitempty"`
	// Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty"`
	// The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	RunAsUser *int64 `json:"runAsUser,omitempty"`
	SeLinuxOptions *V1SELinuxOptions `json:"seLinuxOptions,omitempty"`
	// A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container.
	SupplementalGroups []int32 `json:"supplementalGroups,omitempty"`
	// Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch.
	Sysctls []V1Sysctl `json:"sysctls,omitempty"`
}

// NewV1PodSecurityContext instantiates a new V1PodSecurityContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1PodSecurityContext() *V1PodSecurityContext {
	this := V1PodSecurityContext{}
	return &this
}

// NewV1PodSecurityContextWithDefaults instantiates a new V1PodSecurityContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1PodSecurityContextWithDefaults() *V1PodSecurityContext {
	this := V1PodSecurityContext{}
	return &this
}

// GetFsGroup returns the FsGroup field value if set, zero value otherwise.
func (o *V1PodSecurityContext) GetFsGroup() int64 {
	if o == nil || IsNil(o.FsGroup) {
		var ret int64
		return ret
	}
	return *o.FsGroup
}

// GetFsGroupOk returns a tuple with the FsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodSecurityContext) GetFsGroupOk() (*int64, bool) {
	if o == nil || IsNil(o.FsGroup) {
		return nil, false
	}
	return o.FsGroup, true
}

// HasFsGroup returns a boolean if a field has been set.
func (o *V1PodSecurityContext) HasFsGroup() bool {
	if o != nil && !IsNil(o.FsGroup) {
		return true
	}

	return false
}

// SetFsGroup gets a reference to the given int64 and assigns it to the FsGroup field.
func (o *V1PodSecurityContext) SetFsGroup(v int64) {
	o.FsGroup = &v
}

// GetRunAsGroup returns the RunAsGroup field value if set, zero value otherwise.
func (o *V1PodSecurityContext) GetRunAsGroup() int64 {
	if o == nil || IsNil(o.RunAsGroup) {
		var ret int64
		return ret
	}
	return *o.RunAsGroup
}

// GetRunAsGroupOk returns a tuple with the RunAsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodSecurityContext) GetRunAsGroupOk() (*int64, bool) {
	if o == nil || IsNil(o.RunAsGroup) {
		return nil, false
	}
	return o.RunAsGroup, true
}

// HasRunAsGroup returns a boolean if a field has been set.
func (o *V1PodSecurityContext) HasRunAsGroup() bool {
	if o != nil && !IsNil(o.RunAsGroup) {
		return true
	}

	return false
}

// SetRunAsGroup gets a reference to the given int64 and assigns it to the RunAsGroup field.
func (o *V1PodSecurityContext) SetRunAsGroup(v int64) {
	o.RunAsGroup = &v
}

// GetRunAsNonRoot returns the RunAsNonRoot field value if set, zero value otherwise.
func (o *V1PodSecurityContext) GetRunAsNonRoot() bool {
	if o == nil || IsNil(o.RunAsNonRoot) {
		var ret bool
		return ret
	}
	return *o.RunAsNonRoot
}

// GetRunAsNonRootOk returns a tuple with the RunAsNonRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodSecurityContext) GetRunAsNonRootOk() (*bool, bool) {
	if o == nil || IsNil(o.RunAsNonRoot) {
		return nil, false
	}
	return o.RunAsNonRoot, true
}

// HasRunAsNonRoot returns a boolean if a field has been set.
func (o *V1PodSecurityContext) HasRunAsNonRoot() bool {
	if o != nil && !IsNil(o.RunAsNonRoot) {
		return true
	}

	return false
}

// SetRunAsNonRoot gets a reference to the given bool and assigns it to the RunAsNonRoot field.
func (o *V1PodSecurityContext) SetRunAsNonRoot(v bool) {
	o.RunAsNonRoot = &v
}

// GetRunAsUser returns the RunAsUser field value if set, zero value otherwise.
func (o *V1PodSecurityContext) GetRunAsUser() int64 {
	if o == nil || IsNil(o.RunAsUser) {
		var ret int64
		return ret
	}
	return *o.RunAsUser
}

// GetRunAsUserOk returns a tuple with the RunAsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodSecurityContext) GetRunAsUserOk() (*int64, bool) {
	if o == nil || IsNil(o.RunAsUser) {
		return nil, false
	}
	return o.RunAsUser, true
}

// HasRunAsUser returns a boolean if a field has been set.
func (o *V1PodSecurityContext) HasRunAsUser() bool {
	if o != nil && !IsNil(o.RunAsUser) {
		return true
	}

	return false
}

// SetRunAsUser gets a reference to the given int64 and assigns it to the RunAsUser field.
func (o *V1PodSecurityContext) SetRunAsUser(v int64) {
	o.RunAsUser = &v
}

// GetSeLinuxOptions returns the SeLinuxOptions field value if set, zero value otherwise.
func (o *V1PodSecurityContext) GetSeLinuxOptions() V1SELinuxOptions {
	if o == nil || IsNil(o.SeLinuxOptions) {
		var ret V1SELinuxOptions
		return ret
	}
	return *o.SeLinuxOptions
}

// GetSeLinuxOptionsOk returns a tuple with the SeLinuxOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodSecurityContext) GetSeLinuxOptionsOk() (*V1SELinuxOptions, bool) {
	if o == nil || IsNil(o.SeLinuxOptions) {
		return nil, false
	}
	return o.SeLinuxOptions, true
}

// HasSeLinuxOptions returns a boolean if a field has been set.
func (o *V1PodSecurityContext) HasSeLinuxOptions() bool {
	if o != nil && !IsNil(o.SeLinuxOptions) {
		return true
	}

	return false
}

// SetSeLinuxOptions gets a reference to the given V1SELinuxOptions and assigns it to the SeLinuxOptions field.
func (o *V1PodSecurityContext) SetSeLinuxOptions(v V1SELinuxOptions) {
	o.SeLinuxOptions = &v
}

// GetSupplementalGroups returns the SupplementalGroups field value if set, zero value otherwise.
func (o *V1PodSecurityContext) GetSupplementalGroups() []int32 {
	if o == nil || IsNil(o.SupplementalGroups) {
		var ret []int32
		return ret
	}
	return o.SupplementalGroups
}

// GetSupplementalGroupsOk returns a tuple with the SupplementalGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodSecurityContext) GetSupplementalGroupsOk() ([]int32, bool) {
	if o == nil || IsNil(o.SupplementalGroups) {
		return nil, false
	}
	return o.SupplementalGroups, true
}

// HasSupplementalGroups returns a boolean if a field has been set.
func (o *V1PodSecurityContext) HasSupplementalGroups() bool {
	if o != nil && !IsNil(o.SupplementalGroups) {
		return true
	}

	return false
}

// SetSupplementalGroups gets a reference to the given []int32 and assigns it to the SupplementalGroups field.
func (o *V1PodSecurityContext) SetSupplementalGroups(v []int32) {
	o.SupplementalGroups = v
}

// GetSysctls returns the Sysctls field value if set, zero value otherwise.
func (o *V1PodSecurityContext) GetSysctls() []V1Sysctl {
	if o == nil || IsNil(o.Sysctls) {
		var ret []V1Sysctl
		return ret
	}
	return o.Sysctls
}

// GetSysctlsOk returns a tuple with the Sysctls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1PodSecurityContext) GetSysctlsOk() ([]V1Sysctl, bool) {
	if o == nil || IsNil(o.Sysctls) {
		return nil, false
	}
	return o.Sysctls, true
}

// HasSysctls returns a boolean if a field has been set.
func (o *V1PodSecurityContext) HasSysctls() bool {
	if o != nil && !IsNil(o.Sysctls) {
		return true
	}

	return false
}

// SetSysctls gets a reference to the given []V1Sysctl and assigns it to the Sysctls field.
func (o *V1PodSecurityContext) SetSysctls(v []V1Sysctl) {
	o.Sysctls = v
}

func (o V1PodSecurityContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1PodSecurityContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsGroup) {
		toSerialize["fsGroup"] = o.FsGroup
	}
	if !IsNil(o.RunAsGroup) {
		toSerialize["runAsGroup"] = o.RunAsGroup
	}
	if !IsNil(o.RunAsNonRoot) {
		toSerialize["runAsNonRoot"] = o.RunAsNonRoot
	}
	if !IsNil(o.RunAsUser) {
		toSerialize["runAsUser"] = o.RunAsUser
	}
	if !IsNil(o.SeLinuxOptions) {
		toSerialize["seLinuxOptions"] = o.SeLinuxOptions
	}
	if !IsNil(o.SupplementalGroups) {
		toSerialize["supplementalGroups"] = o.SupplementalGroups
	}
	if !IsNil(o.Sysctls) {
		toSerialize["sysctls"] = o.Sysctls
	}
	return toSerialize, nil
}

type NullableV1PodSecurityContext struct {
	value *V1PodSecurityContext
	isSet bool
}

func (v NullableV1PodSecurityContext) Get() *V1PodSecurityContext {
	return v.value
}

func (v *NullableV1PodSecurityContext) Set(val *V1PodSecurityContext) {
	v.value = val
	v.isSet = true
}

func (v NullableV1PodSecurityContext) IsSet() bool {
	return v.isSet
}

func (v *NullableV1PodSecurityContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1PodSecurityContext(val *V1PodSecurityContext) *NullableV1PodSecurityContext {
	return &NullableV1PodSecurityContext{value: val, isSet: true}
}

func (v NullableV1PodSecurityContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1PodSecurityContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

