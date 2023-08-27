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

// checks if the V1NodeSystemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1NodeSystemInfo{}

// V1NodeSystemInfo NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
type V1NodeSystemInfo struct {
	// The Architecture reported by the node
	Architecture string `json:"architecture"`
	// Boot ID reported by the node.
	BootID string `json:"bootID"`
	// ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	// Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	KernelVersion string `json:"kernelVersion"`
	// KubeProxy Version reported by the node.
	KubeProxyVersion string `json:"kubeProxyVersion"`
	// Kubelet Version reported by the node.
	KubeletVersion string `json:"kubeletVersion"`
	// MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
	MachineID string `json:"machineID"`
	// The Operating System reported by the node
	OperatingSystem string `json:"operatingSystem"`
	// OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	OsImage string `json:"osImage"`
	// SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html
	SystemUUID string `json:"systemUUID"`
}

// NewV1NodeSystemInfo instantiates a new V1NodeSystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1NodeSystemInfo(architecture string, bootID string, containerRuntimeVersion string, kernelVersion string, kubeProxyVersion string, kubeletVersion string, machineID string, operatingSystem string, osImage string, systemUUID string) *V1NodeSystemInfo {
	this := V1NodeSystemInfo{}
	this.Architecture = architecture
	this.BootID = bootID
	this.ContainerRuntimeVersion = containerRuntimeVersion
	this.KernelVersion = kernelVersion
	this.KubeProxyVersion = kubeProxyVersion
	this.KubeletVersion = kubeletVersion
	this.MachineID = machineID
	this.OperatingSystem = operatingSystem
	this.OsImage = osImage
	this.SystemUUID = systemUUID
	return &this
}

// NewV1NodeSystemInfoWithDefaults instantiates a new V1NodeSystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1NodeSystemInfoWithDefaults() *V1NodeSystemInfo {
	this := V1NodeSystemInfo{}
	return &this
}

// GetArchitecture returns the Architecture field value
func (o *V1NodeSystemInfo) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *V1NodeSystemInfo) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *V1NodeSystemInfo) SetArchitecture(v string) {
	o.Architecture = v
}

// GetBootID returns the BootID field value
func (o *V1NodeSystemInfo) GetBootID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BootID
}

// GetBootIDOk returns a tuple with the BootID field value
// and a boolean to check if the value has been set.
func (o *V1NodeSystemInfo) GetBootIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BootID, true
}

// SetBootID sets field value
func (o *V1NodeSystemInfo) SetBootID(v string) {
	o.BootID = v
}

// GetContainerRuntimeVersion returns the ContainerRuntimeVersion field value
func (o *V1NodeSystemInfo) GetContainerRuntimeVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerRuntimeVersion
}

// GetContainerRuntimeVersionOk returns a tuple with the ContainerRuntimeVersion field value
// and a boolean to check if the value has been set.
func (o *V1NodeSystemInfo) GetContainerRuntimeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerRuntimeVersion, true
}

// SetContainerRuntimeVersion sets field value
func (o *V1NodeSystemInfo) SetContainerRuntimeVersion(v string) {
	o.ContainerRuntimeVersion = v
}

// GetKernelVersion returns the KernelVersion field value
func (o *V1NodeSystemInfo) GetKernelVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KernelVersion
}

// GetKernelVersionOk returns a tuple with the KernelVersion field value
// and a boolean to check if the value has been set.
func (o *V1NodeSystemInfo) GetKernelVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KernelVersion, true
}

// SetKernelVersion sets field value
func (o *V1NodeSystemInfo) SetKernelVersion(v string) {
	o.KernelVersion = v
}

// GetKubeProxyVersion returns the KubeProxyVersion field value
func (o *V1NodeSystemInfo) GetKubeProxyVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubeProxyVersion
}

// GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field value
// and a boolean to check if the value has been set.
func (o *V1NodeSystemInfo) GetKubeProxyVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubeProxyVersion, true
}

// SetKubeProxyVersion sets field value
func (o *V1NodeSystemInfo) SetKubeProxyVersion(v string) {
	o.KubeProxyVersion = v
}

// GetKubeletVersion returns the KubeletVersion field value
func (o *V1NodeSystemInfo) GetKubeletVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubeletVersion
}

// GetKubeletVersionOk returns a tuple with the KubeletVersion field value
// and a boolean to check if the value has been set.
func (o *V1NodeSystemInfo) GetKubeletVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubeletVersion, true
}

// SetKubeletVersion sets field value
func (o *V1NodeSystemInfo) SetKubeletVersion(v string) {
	o.KubeletVersion = v
}

// GetMachineID returns the MachineID field value
func (o *V1NodeSystemInfo) GetMachineID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MachineID
}

// GetMachineIDOk returns a tuple with the MachineID field value
// and a boolean to check if the value has been set.
func (o *V1NodeSystemInfo) GetMachineIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MachineID, true
}

// SetMachineID sets field value
func (o *V1NodeSystemInfo) SetMachineID(v string) {
	o.MachineID = v
}

// GetOperatingSystem returns the OperatingSystem field value
func (o *V1NodeSystemInfo) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *V1NodeSystemInfo) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value
func (o *V1NodeSystemInfo) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

// GetOsImage returns the OsImage field value
func (o *V1NodeSystemInfo) GetOsImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OsImage
}

// GetOsImageOk returns a tuple with the OsImage field value
// and a boolean to check if the value has been set.
func (o *V1NodeSystemInfo) GetOsImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OsImage, true
}

// SetOsImage sets field value
func (o *V1NodeSystemInfo) SetOsImage(v string) {
	o.OsImage = v
}

// GetSystemUUID returns the SystemUUID field value
func (o *V1NodeSystemInfo) GetSystemUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemUUID
}

// GetSystemUUIDOk returns a tuple with the SystemUUID field value
// and a boolean to check if the value has been set.
func (o *V1NodeSystemInfo) GetSystemUUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemUUID, true
}

// SetSystemUUID sets field value
func (o *V1NodeSystemInfo) SetSystemUUID(v string) {
	o.SystemUUID = v
}

func (o V1NodeSystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1NodeSystemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["architecture"] = o.Architecture
	toSerialize["bootID"] = o.BootID
	toSerialize["containerRuntimeVersion"] = o.ContainerRuntimeVersion
	toSerialize["kernelVersion"] = o.KernelVersion
	toSerialize["kubeProxyVersion"] = o.KubeProxyVersion
	toSerialize["kubeletVersion"] = o.KubeletVersion
	toSerialize["machineID"] = o.MachineID
	toSerialize["operatingSystem"] = o.OperatingSystem
	toSerialize["osImage"] = o.OsImage
	toSerialize["systemUUID"] = o.SystemUUID
	return toSerialize, nil
}

type NullableV1NodeSystemInfo struct {
	value *V1NodeSystemInfo
	isSet bool
}

func (v NullableV1NodeSystemInfo) Get() *V1NodeSystemInfo {
	return v.value
}

func (v *NullableV1NodeSystemInfo) Set(val *V1NodeSystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV1NodeSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV1NodeSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1NodeSystemInfo(val *V1NodeSystemInfo) *NullableV1NodeSystemInfo {
	return &NullableV1NodeSystemInfo{value: val, isSet: true}
}

func (v NullableV1NodeSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1NodeSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


