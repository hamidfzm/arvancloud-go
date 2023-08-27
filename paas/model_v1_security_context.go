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

// checks if the V1SecurityContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1SecurityContext{}

// V1SecurityContext SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext.  When both are set, the values in SecurityContext take precedence.
type V1SecurityContext struct {
	// AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
	AllowPrivilegeEscalation *bool `json:"allowPrivilegeEscalation,omitempty"`
	Capabilities *V1Capabilities `json:"capabilities,omitempty"`
	// Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
	Privileged *bool `json:"privileged,omitempty"`
	// Whether this container has a read-only root filesystem. Default is false.
	ReadOnlyRootFilesystem *bool `json:"readOnlyRootFilesystem,omitempty"`
	// The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsGroup *int64 `json:"runAsGroup,omitempty"`
	// Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty"`
	// The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsUser *int64 `json:"runAsUser,omitempty"`
	SeLinuxOptions *V1SELinuxOptions `json:"seLinuxOptions,omitempty"`
}

// NewV1SecurityContext instantiates a new V1SecurityContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1SecurityContext() *V1SecurityContext {
	this := V1SecurityContext{}
	return &this
}

// NewV1SecurityContextWithDefaults instantiates a new V1SecurityContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1SecurityContextWithDefaults() *V1SecurityContext {
	this := V1SecurityContext{}
	return &this
}

// GetAllowPrivilegeEscalation returns the AllowPrivilegeEscalation field value if set, zero value otherwise.
func (o *V1SecurityContext) GetAllowPrivilegeEscalation() bool {
	if o == nil || IsNil(o.AllowPrivilegeEscalation) {
		var ret bool
		return ret
	}
	return *o.AllowPrivilegeEscalation
}

// GetAllowPrivilegeEscalationOk returns a tuple with the AllowPrivilegeEscalation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContext) GetAllowPrivilegeEscalationOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPrivilegeEscalation) {
		return nil, false
	}
	return o.AllowPrivilegeEscalation, true
}

// HasAllowPrivilegeEscalation returns a boolean if a field has been set.
func (o *V1SecurityContext) HasAllowPrivilegeEscalation() bool {
	if o != nil && !IsNil(o.AllowPrivilegeEscalation) {
		return true
	}

	return false
}

// SetAllowPrivilegeEscalation gets a reference to the given bool and assigns it to the AllowPrivilegeEscalation field.
func (o *V1SecurityContext) SetAllowPrivilegeEscalation(v bool) {
	o.AllowPrivilegeEscalation = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *V1SecurityContext) GetCapabilities() V1Capabilities {
	if o == nil || IsNil(o.Capabilities) {
		var ret V1Capabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContext) GetCapabilitiesOk() (*V1Capabilities, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *V1SecurityContext) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given V1Capabilities and assigns it to the Capabilities field.
func (o *V1SecurityContext) SetCapabilities(v V1Capabilities) {
	o.Capabilities = &v
}

// GetPrivileged returns the Privileged field value if set, zero value otherwise.
func (o *V1SecurityContext) GetPrivileged() bool {
	if o == nil || IsNil(o.Privileged) {
		var ret bool
		return ret
	}
	return *o.Privileged
}

// GetPrivilegedOk returns a tuple with the Privileged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContext) GetPrivilegedOk() (*bool, bool) {
	if o == nil || IsNil(o.Privileged) {
		return nil, false
	}
	return o.Privileged, true
}

// HasPrivileged returns a boolean if a field has been set.
func (o *V1SecurityContext) HasPrivileged() bool {
	if o != nil && !IsNil(o.Privileged) {
		return true
	}

	return false
}

// SetPrivileged gets a reference to the given bool and assigns it to the Privileged field.
func (o *V1SecurityContext) SetPrivileged(v bool) {
	o.Privileged = &v
}

// GetReadOnlyRootFilesystem returns the ReadOnlyRootFilesystem field value if set, zero value otherwise.
func (o *V1SecurityContext) GetReadOnlyRootFilesystem() bool {
	if o == nil || IsNil(o.ReadOnlyRootFilesystem) {
		var ret bool
		return ret
	}
	return *o.ReadOnlyRootFilesystem
}

// GetReadOnlyRootFilesystemOk returns a tuple with the ReadOnlyRootFilesystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContext) GetReadOnlyRootFilesystemOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnlyRootFilesystem) {
		return nil, false
	}
	return o.ReadOnlyRootFilesystem, true
}

// HasReadOnlyRootFilesystem returns a boolean if a field has been set.
func (o *V1SecurityContext) HasReadOnlyRootFilesystem() bool {
	if o != nil && !IsNil(o.ReadOnlyRootFilesystem) {
		return true
	}

	return false
}

// SetReadOnlyRootFilesystem gets a reference to the given bool and assigns it to the ReadOnlyRootFilesystem field.
func (o *V1SecurityContext) SetReadOnlyRootFilesystem(v bool) {
	o.ReadOnlyRootFilesystem = &v
}

// GetRunAsGroup returns the RunAsGroup field value if set, zero value otherwise.
func (o *V1SecurityContext) GetRunAsGroup() int64 {
	if o == nil || IsNil(o.RunAsGroup) {
		var ret int64
		return ret
	}
	return *o.RunAsGroup
}

// GetRunAsGroupOk returns a tuple with the RunAsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContext) GetRunAsGroupOk() (*int64, bool) {
	if o == nil || IsNil(o.RunAsGroup) {
		return nil, false
	}
	return o.RunAsGroup, true
}

// HasRunAsGroup returns a boolean if a field has been set.
func (o *V1SecurityContext) HasRunAsGroup() bool {
	if o != nil && !IsNil(o.RunAsGroup) {
		return true
	}

	return false
}

// SetRunAsGroup gets a reference to the given int64 and assigns it to the RunAsGroup field.
func (o *V1SecurityContext) SetRunAsGroup(v int64) {
	o.RunAsGroup = &v
}

// GetRunAsNonRoot returns the RunAsNonRoot field value if set, zero value otherwise.
func (o *V1SecurityContext) GetRunAsNonRoot() bool {
	if o == nil || IsNil(o.RunAsNonRoot) {
		var ret bool
		return ret
	}
	return *o.RunAsNonRoot
}

// GetRunAsNonRootOk returns a tuple with the RunAsNonRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContext) GetRunAsNonRootOk() (*bool, bool) {
	if o == nil || IsNil(o.RunAsNonRoot) {
		return nil, false
	}
	return o.RunAsNonRoot, true
}

// HasRunAsNonRoot returns a boolean if a field has been set.
func (o *V1SecurityContext) HasRunAsNonRoot() bool {
	if o != nil && !IsNil(o.RunAsNonRoot) {
		return true
	}

	return false
}

// SetRunAsNonRoot gets a reference to the given bool and assigns it to the RunAsNonRoot field.
func (o *V1SecurityContext) SetRunAsNonRoot(v bool) {
	o.RunAsNonRoot = &v
}

// GetRunAsUser returns the RunAsUser field value if set, zero value otherwise.
func (o *V1SecurityContext) GetRunAsUser() int64 {
	if o == nil || IsNil(o.RunAsUser) {
		var ret int64
		return ret
	}
	return *o.RunAsUser
}

// GetRunAsUserOk returns a tuple with the RunAsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContext) GetRunAsUserOk() (*int64, bool) {
	if o == nil || IsNil(o.RunAsUser) {
		return nil, false
	}
	return o.RunAsUser, true
}

// HasRunAsUser returns a boolean if a field has been set.
func (o *V1SecurityContext) HasRunAsUser() bool {
	if o != nil && !IsNil(o.RunAsUser) {
		return true
	}

	return false
}

// SetRunAsUser gets a reference to the given int64 and assigns it to the RunAsUser field.
func (o *V1SecurityContext) SetRunAsUser(v int64) {
	o.RunAsUser = &v
}

// GetSeLinuxOptions returns the SeLinuxOptions field value if set, zero value otherwise.
func (o *V1SecurityContext) GetSeLinuxOptions() V1SELinuxOptions {
	if o == nil || IsNil(o.SeLinuxOptions) {
		var ret V1SELinuxOptions
		return ret
	}
	return *o.SeLinuxOptions
}

// GetSeLinuxOptionsOk returns a tuple with the SeLinuxOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1SecurityContext) GetSeLinuxOptionsOk() (*V1SELinuxOptions, bool) {
	if o == nil || IsNil(o.SeLinuxOptions) {
		return nil, false
	}
	return o.SeLinuxOptions, true
}

// HasSeLinuxOptions returns a boolean if a field has been set.
func (o *V1SecurityContext) HasSeLinuxOptions() bool {
	if o != nil && !IsNil(o.SeLinuxOptions) {
		return true
	}

	return false
}

// SetSeLinuxOptions gets a reference to the given V1SELinuxOptions and assigns it to the SeLinuxOptions field.
func (o *V1SecurityContext) SetSeLinuxOptions(v V1SELinuxOptions) {
	o.SeLinuxOptions = &v
}

func (o V1SecurityContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1SecurityContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowPrivilegeEscalation) {
		toSerialize["allowPrivilegeEscalation"] = o.AllowPrivilegeEscalation
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.Privileged) {
		toSerialize["privileged"] = o.Privileged
	}
	if !IsNil(o.ReadOnlyRootFilesystem) {
		toSerialize["readOnlyRootFilesystem"] = o.ReadOnlyRootFilesystem
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
	return toSerialize, nil
}

type NullableV1SecurityContext struct {
	value *V1SecurityContext
	isSet bool
}

func (v NullableV1SecurityContext) Get() *V1SecurityContext {
	return v.value
}

func (v *NullableV1SecurityContext) Set(val *V1SecurityContext) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SecurityContext) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SecurityContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SecurityContext(val *V1SecurityContext) *NullableV1SecurityContext {
	return &NullableV1SecurityContext{value: val, isSet: true}
}

func (v NullableV1SecurityContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SecurityContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


