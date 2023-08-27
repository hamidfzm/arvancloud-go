# V1SecurityContextConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowHostDirVolumePlugin** | **bool** | AllowHostDirVolumePlugin determines if the policy allow containers to use the HostDir volume plugin | 
**AllowHostIPC** | **bool** | AllowHostIPC determines if the policy allows host ipc in the containers. | 
**AllowHostNetwork** | **bool** | AllowHostNetwork determines if the policy allows the use of HostNetwork in the pod spec. | 
**AllowHostPID** | **bool** | AllowHostPID determines if the policy allows host pid in the containers. | 
**AllowHostPorts** | **bool** | AllowHostPorts determines if the policy allows host ports in the containers. | 
**AllowPrivilegeEscalation** | Pointer to **bool** | AllowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true. | [optional] 
**AllowPrivilegedContainer** | **bool** | AllowPrivilegedContainer determines if a container can request to be run as privileged. | 
**AllowedCapabilities** | [**[]V1Capability**](V1Capability.md) | AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field maybe added at the pod author&#39;s discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities. To allow all capabilities you may use &#39;*&#39;. | 
**AllowedFlexVolumes** | Pointer to [**[]V1AllowedFlexVolume**](V1AllowedFlexVolume.md) | AllowedFlexVolumes is a whitelist of allowed Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the \&quot;Volumes\&quot; field. | [optional] 
**AllowedUnsafeSysctls** | Pointer to **[]string** | AllowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in \&quot;*\&quot; in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to whitelist all allowed unsafe sysctls explicitly to avoid rejection.  Examples: e.g. \&quot;foo/_*\&quot; allows \&quot;foo/bar\&quot;, \&quot;foo/baz\&quot;, etc. e.g. \&quot;foo.*\&quot; allows \&quot;foo.bar\&quot;, \&quot;foo.baz\&quot;, etc. | [optional] 
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**DefaultAddCapabilities** | [**[]V1Capability**](V1Capability.md) | DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capabiility in both DefaultAddCapabilities and RequiredDropCapabilities. | 
**DefaultAllowPrivilegeEscalation** | Pointer to **bool** | DefaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process. | [optional] 
**ForbiddenSysctls** | Pointer to **[]string** | ForbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in \&quot;*\&quot; in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.  Examples: e.g. \&quot;foo/_*\&quot; forbids \&quot;foo/bar\&quot;, \&quot;foo/baz\&quot;, etc. e.g. \&quot;foo.*\&quot; forbids \&quot;foo.bar\&quot;, \&quot;foo.baz\&quot;, etc. | [optional] 
**FsGroup** | Pointer to [**V1FSGroupStrategyOptions**](V1FSGroupStrategyOptions.md) |  | [optional] 
**Groups** | **[]string** | The groups that have permission to use this security context constraints | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Priority** | **int32** | Priority influences the sort order of SCCs when evaluating which SCCs to try first for a given pod request based on access in the Users and Groups fields.  The higher the int, the higher priority. An unset value is considered a 0 priority. If scores for multiple SCCs are equal they will be sorted from most restrictive to least restrictive. If both priorities and restrictions are equal the SCCs will be sorted by name. | 
**ReadOnlyRootFilesystem** | **bool** | ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the SCC should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to. | 
**RequiredDropCapabilities** | [**[]V1Capability**](V1Capability.md) | RequiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added. | 
**RunAsUser** | Pointer to [**V1RunAsUserStrategyOptions**](V1RunAsUserStrategyOptions.md) |  | [optional] 
**SeLinuxContext** | Pointer to [**V1SELinuxContextStrategyOptions**](V1SELinuxContextStrategyOptions.md) |  | [optional] 
**SeccompProfiles** | Pointer to **[]string** | SeccompProfiles lists the allowed profiles that may be set for the pod or container&#39;s seccomp annotations.  An unset (nil) or empty value means that no profiles may be specifid by the pod or container. The wildcard &#39;*&#39; may be used to allow all profiles.  When used to generate a value for a pod the first non-wildcard profile will be used as the default. | [optional] 
**SupplementalGroups** | Pointer to [**V1SupplementalGroupsStrategyOptions**](V1SupplementalGroupsStrategyOptions.md) |  | [optional] 
**Users** | **[]string** | The users who have permissions to use this security context constraints | 
**Volumes** | [**[]V1FSType**](V1FSType.md) | Volumes is a white list of allowed volume plugins.  FSType corresponds directly with the field names of a VolumeSource (azureFile, configMap, emptyDir).  To allow all volumes you may use \&quot;*\&quot;. To allow no volumes, set to [\&quot;none\&quot;]. | 

## Methods

### NewV1SecurityContextConstraints

`func NewV1SecurityContextConstraints(allowHostDirVolumePlugin bool, allowHostIPC bool, allowHostNetwork bool, allowHostPID bool, allowHostPorts bool, allowPrivilegedContainer bool, allowedCapabilities []V1Capability, defaultAddCapabilities []V1Capability, groups []string, priority int32, readOnlyRootFilesystem bool, requiredDropCapabilities []V1Capability, users []string, volumes []V1FSType, ) *V1SecurityContextConstraints`

NewV1SecurityContextConstraints instantiates a new V1SecurityContextConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SecurityContextConstraintsWithDefaults

`func NewV1SecurityContextConstraintsWithDefaults() *V1SecurityContextConstraints`

NewV1SecurityContextConstraintsWithDefaults instantiates a new V1SecurityContextConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowHostDirVolumePlugin

`func (o *V1SecurityContextConstraints) GetAllowHostDirVolumePlugin() bool`

GetAllowHostDirVolumePlugin returns the AllowHostDirVolumePlugin field if non-nil, zero value otherwise.

### GetAllowHostDirVolumePluginOk

`func (o *V1SecurityContextConstraints) GetAllowHostDirVolumePluginOk() (*bool, bool)`

GetAllowHostDirVolumePluginOk returns a tuple with the AllowHostDirVolumePlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHostDirVolumePlugin

`func (o *V1SecurityContextConstraints) SetAllowHostDirVolumePlugin(v bool)`

SetAllowHostDirVolumePlugin sets AllowHostDirVolumePlugin field to given value.


### GetAllowHostIPC

`func (o *V1SecurityContextConstraints) GetAllowHostIPC() bool`

GetAllowHostIPC returns the AllowHostIPC field if non-nil, zero value otherwise.

### GetAllowHostIPCOk

`func (o *V1SecurityContextConstraints) GetAllowHostIPCOk() (*bool, bool)`

GetAllowHostIPCOk returns a tuple with the AllowHostIPC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHostIPC

`func (o *V1SecurityContextConstraints) SetAllowHostIPC(v bool)`

SetAllowHostIPC sets AllowHostIPC field to given value.


### GetAllowHostNetwork

`func (o *V1SecurityContextConstraints) GetAllowHostNetwork() bool`

GetAllowHostNetwork returns the AllowHostNetwork field if non-nil, zero value otherwise.

### GetAllowHostNetworkOk

`func (o *V1SecurityContextConstraints) GetAllowHostNetworkOk() (*bool, bool)`

GetAllowHostNetworkOk returns a tuple with the AllowHostNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHostNetwork

`func (o *V1SecurityContextConstraints) SetAllowHostNetwork(v bool)`

SetAllowHostNetwork sets AllowHostNetwork field to given value.


### GetAllowHostPID

`func (o *V1SecurityContextConstraints) GetAllowHostPID() bool`

GetAllowHostPID returns the AllowHostPID field if non-nil, zero value otherwise.

### GetAllowHostPIDOk

`func (o *V1SecurityContextConstraints) GetAllowHostPIDOk() (*bool, bool)`

GetAllowHostPIDOk returns a tuple with the AllowHostPID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHostPID

`func (o *V1SecurityContextConstraints) SetAllowHostPID(v bool)`

SetAllowHostPID sets AllowHostPID field to given value.


### GetAllowHostPorts

`func (o *V1SecurityContextConstraints) GetAllowHostPorts() bool`

GetAllowHostPorts returns the AllowHostPorts field if non-nil, zero value otherwise.

### GetAllowHostPortsOk

`func (o *V1SecurityContextConstraints) GetAllowHostPortsOk() (*bool, bool)`

GetAllowHostPortsOk returns a tuple with the AllowHostPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHostPorts

`func (o *V1SecurityContextConstraints) SetAllowHostPorts(v bool)`

SetAllowHostPorts sets AllowHostPorts field to given value.


### GetAllowPrivilegeEscalation

`func (o *V1SecurityContextConstraints) GetAllowPrivilegeEscalation() bool`

GetAllowPrivilegeEscalation returns the AllowPrivilegeEscalation field if non-nil, zero value otherwise.

### GetAllowPrivilegeEscalationOk

`func (o *V1SecurityContextConstraints) GetAllowPrivilegeEscalationOk() (*bool, bool)`

GetAllowPrivilegeEscalationOk returns a tuple with the AllowPrivilegeEscalation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrivilegeEscalation

`func (o *V1SecurityContextConstraints) SetAllowPrivilegeEscalation(v bool)`

SetAllowPrivilegeEscalation sets AllowPrivilegeEscalation field to given value.

### HasAllowPrivilegeEscalation

`func (o *V1SecurityContextConstraints) HasAllowPrivilegeEscalation() bool`

HasAllowPrivilegeEscalation returns a boolean if a field has been set.

### GetAllowPrivilegedContainer

`func (o *V1SecurityContextConstraints) GetAllowPrivilegedContainer() bool`

GetAllowPrivilegedContainer returns the AllowPrivilegedContainer field if non-nil, zero value otherwise.

### GetAllowPrivilegedContainerOk

`func (o *V1SecurityContextConstraints) GetAllowPrivilegedContainerOk() (*bool, bool)`

GetAllowPrivilegedContainerOk returns a tuple with the AllowPrivilegedContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrivilegedContainer

`func (o *V1SecurityContextConstraints) SetAllowPrivilegedContainer(v bool)`

SetAllowPrivilegedContainer sets AllowPrivilegedContainer field to given value.


### GetAllowedCapabilities

`func (o *V1SecurityContextConstraints) GetAllowedCapabilities() []V1Capability`

GetAllowedCapabilities returns the AllowedCapabilities field if non-nil, zero value otherwise.

### GetAllowedCapabilitiesOk

`func (o *V1SecurityContextConstraints) GetAllowedCapabilitiesOk() (*[]V1Capability, bool)`

GetAllowedCapabilitiesOk returns a tuple with the AllowedCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCapabilities

`func (o *V1SecurityContextConstraints) SetAllowedCapabilities(v []V1Capability)`

SetAllowedCapabilities sets AllowedCapabilities field to given value.


### GetAllowedFlexVolumes

`func (o *V1SecurityContextConstraints) GetAllowedFlexVolumes() []V1AllowedFlexVolume`

GetAllowedFlexVolumes returns the AllowedFlexVolumes field if non-nil, zero value otherwise.

### GetAllowedFlexVolumesOk

`func (o *V1SecurityContextConstraints) GetAllowedFlexVolumesOk() (*[]V1AllowedFlexVolume, bool)`

GetAllowedFlexVolumesOk returns a tuple with the AllowedFlexVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFlexVolumes

`func (o *V1SecurityContextConstraints) SetAllowedFlexVolumes(v []V1AllowedFlexVolume)`

SetAllowedFlexVolumes sets AllowedFlexVolumes field to given value.

### HasAllowedFlexVolumes

`func (o *V1SecurityContextConstraints) HasAllowedFlexVolumes() bool`

HasAllowedFlexVolumes returns a boolean if a field has been set.

### GetAllowedUnsafeSysctls

`func (o *V1SecurityContextConstraints) GetAllowedUnsafeSysctls() []string`

GetAllowedUnsafeSysctls returns the AllowedUnsafeSysctls field if non-nil, zero value otherwise.

### GetAllowedUnsafeSysctlsOk

`func (o *V1SecurityContextConstraints) GetAllowedUnsafeSysctlsOk() (*[]string, bool)`

GetAllowedUnsafeSysctlsOk returns a tuple with the AllowedUnsafeSysctls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUnsafeSysctls

`func (o *V1SecurityContextConstraints) SetAllowedUnsafeSysctls(v []string)`

SetAllowedUnsafeSysctls sets AllowedUnsafeSysctls field to given value.

### HasAllowedUnsafeSysctls

`func (o *V1SecurityContextConstraints) HasAllowedUnsafeSysctls() bool`

HasAllowedUnsafeSysctls returns a boolean if a field has been set.

### GetApiVersion

`func (o *V1SecurityContextConstraints) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1SecurityContextConstraints) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1SecurityContextConstraints) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1SecurityContextConstraints) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDefaultAddCapabilities

`func (o *V1SecurityContextConstraints) GetDefaultAddCapabilities() []V1Capability`

GetDefaultAddCapabilities returns the DefaultAddCapabilities field if non-nil, zero value otherwise.

### GetDefaultAddCapabilitiesOk

`func (o *V1SecurityContextConstraints) GetDefaultAddCapabilitiesOk() (*[]V1Capability, bool)`

GetDefaultAddCapabilitiesOk returns a tuple with the DefaultAddCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAddCapabilities

`func (o *V1SecurityContextConstraints) SetDefaultAddCapabilities(v []V1Capability)`

SetDefaultAddCapabilities sets DefaultAddCapabilities field to given value.


### GetDefaultAllowPrivilegeEscalation

`func (o *V1SecurityContextConstraints) GetDefaultAllowPrivilegeEscalation() bool`

GetDefaultAllowPrivilegeEscalation returns the DefaultAllowPrivilegeEscalation field if non-nil, zero value otherwise.

### GetDefaultAllowPrivilegeEscalationOk

`func (o *V1SecurityContextConstraints) GetDefaultAllowPrivilegeEscalationOk() (*bool, bool)`

GetDefaultAllowPrivilegeEscalationOk returns a tuple with the DefaultAllowPrivilegeEscalation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAllowPrivilegeEscalation

`func (o *V1SecurityContextConstraints) SetDefaultAllowPrivilegeEscalation(v bool)`

SetDefaultAllowPrivilegeEscalation sets DefaultAllowPrivilegeEscalation field to given value.

### HasDefaultAllowPrivilegeEscalation

`func (o *V1SecurityContextConstraints) HasDefaultAllowPrivilegeEscalation() bool`

HasDefaultAllowPrivilegeEscalation returns a boolean if a field has been set.

### GetForbiddenSysctls

`func (o *V1SecurityContextConstraints) GetForbiddenSysctls() []string`

GetForbiddenSysctls returns the ForbiddenSysctls field if non-nil, zero value otherwise.

### GetForbiddenSysctlsOk

`func (o *V1SecurityContextConstraints) GetForbiddenSysctlsOk() (*[]string, bool)`

GetForbiddenSysctlsOk returns a tuple with the ForbiddenSysctls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenSysctls

`func (o *V1SecurityContextConstraints) SetForbiddenSysctls(v []string)`

SetForbiddenSysctls sets ForbiddenSysctls field to given value.

### HasForbiddenSysctls

`func (o *V1SecurityContextConstraints) HasForbiddenSysctls() bool`

HasForbiddenSysctls returns a boolean if a field has been set.

### GetFsGroup

`func (o *V1SecurityContextConstraints) GetFsGroup() V1FSGroupStrategyOptions`

GetFsGroup returns the FsGroup field if non-nil, zero value otherwise.

### GetFsGroupOk

`func (o *V1SecurityContextConstraints) GetFsGroupOk() (*V1FSGroupStrategyOptions, bool)`

GetFsGroupOk returns a tuple with the FsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsGroup

`func (o *V1SecurityContextConstraints) SetFsGroup(v V1FSGroupStrategyOptions)`

SetFsGroup sets FsGroup field to given value.

### HasFsGroup

`func (o *V1SecurityContextConstraints) HasFsGroup() bool`

HasFsGroup returns a boolean if a field has been set.

### GetGroups

`func (o *V1SecurityContextConstraints) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V1SecurityContextConstraints) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V1SecurityContextConstraints) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetKind

`func (o *V1SecurityContextConstraints) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1SecurityContextConstraints) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1SecurityContextConstraints) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1SecurityContextConstraints) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1SecurityContextConstraints) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1SecurityContextConstraints) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1SecurityContextConstraints) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1SecurityContextConstraints) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPriority

`func (o *V1SecurityContextConstraints) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V1SecurityContextConstraints) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V1SecurityContextConstraints) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetReadOnlyRootFilesystem

`func (o *V1SecurityContextConstraints) GetReadOnlyRootFilesystem() bool`

GetReadOnlyRootFilesystem returns the ReadOnlyRootFilesystem field if non-nil, zero value otherwise.

### GetReadOnlyRootFilesystemOk

`func (o *V1SecurityContextConstraints) GetReadOnlyRootFilesystemOk() (*bool, bool)`

GetReadOnlyRootFilesystemOk returns a tuple with the ReadOnlyRootFilesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyRootFilesystem

`func (o *V1SecurityContextConstraints) SetReadOnlyRootFilesystem(v bool)`

SetReadOnlyRootFilesystem sets ReadOnlyRootFilesystem field to given value.


### GetRequiredDropCapabilities

`func (o *V1SecurityContextConstraints) GetRequiredDropCapabilities() []V1Capability`

GetRequiredDropCapabilities returns the RequiredDropCapabilities field if non-nil, zero value otherwise.

### GetRequiredDropCapabilitiesOk

`func (o *V1SecurityContextConstraints) GetRequiredDropCapabilitiesOk() (*[]V1Capability, bool)`

GetRequiredDropCapabilitiesOk returns a tuple with the RequiredDropCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDropCapabilities

`func (o *V1SecurityContextConstraints) SetRequiredDropCapabilities(v []V1Capability)`

SetRequiredDropCapabilities sets RequiredDropCapabilities field to given value.


### GetRunAsUser

`func (o *V1SecurityContextConstraints) GetRunAsUser() V1RunAsUserStrategyOptions`

GetRunAsUser returns the RunAsUser field if non-nil, zero value otherwise.

### GetRunAsUserOk

`func (o *V1SecurityContextConstraints) GetRunAsUserOk() (*V1RunAsUserStrategyOptions, bool)`

GetRunAsUserOk returns a tuple with the RunAsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAsUser

`func (o *V1SecurityContextConstraints) SetRunAsUser(v V1RunAsUserStrategyOptions)`

SetRunAsUser sets RunAsUser field to given value.

### HasRunAsUser

`func (o *V1SecurityContextConstraints) HasRunAsUser() bool`

HasRunAsUser returns a boolean if a field has been set.

### GetSeLinuxContext

`func (o *V1SecurityContextConstraints) GetSeLinuxContext() V1SELinuxContextStrategyOptions`

GetSeLinuxContext returns the SeLinuxContext field if non-nil, zero value otherwise.

### GetSeLinuxContextOk

`func (o *V1SecurityContextConstraints) GetSeLinuxContextOk() (*V1SELinuxContextStrategyOptions, bool)`

GetSeLinuxContextOk returns a tuple with the SeLinuxContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeLinuxContext

`func (o *V1SecurityContextConstraints) SetSeLinuxContext(v V1SELinuxContextStrategyOptions)`

SetSeLinuxContext sets SeLinuxContext field to given value.

### HasSeLinuxContext

`func (o *V1SecurityContextConstraints) HasSeLinuxContext() bool`

HasSeLinuxContext returns a boolean if a field has been set.

### GetSeccompProfiles

`func (o *V1SecurityContextConstraints) GetSeccompProfiles() []string`

GetSeccompProfiles returns the SeccompProfiles field if non-nil, zero value otherwise.

### GetSeccompProfilesOk

`func (o *V1SecurityContextConstraints) GetSeccompProfilesOk() (*[]string, bool)`

GetSeccompProfilesOk returns a tuple with the SeccompProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeccompProfiles

`func (o *V1SecurityContextConstraints) SetSeccompProfiles(v []string)`

SetSeccompProfiles sets SeccompProfiles field to given value.

### HasSeccompProfiles

`func (o *V1SecurityContextConstraints) HasSeccompProfiles() bool`

HasSeccompProfiles returns a boolean if a field has been set.

### GetSupplementalGroups

`func (o *V1SecurityContextConstraints) GetSupplementalGroups() V1SupplementalGroupsStrategyOptions`

GetSupplementalGroups returns the SupplementalGroups field if non-nil, zero value otherwise.

### GetSupplementalGroupsOk

`func (o *V1SecurityContextConstraints) GetSupplementalGroupsOk() (*V1SupplementalGroupsStrategyOptions, bool)`

GetSupplementalGroupsOk returns a tuple with the SupplementalGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementalGroups

`func (o *V1SecurityContextConstraints) SetSupplementalGroups(v V1SupplementalGroupsStrategyOptions)`

SetSupplementalGroups sets SupplementalGroups field to given value.

### HasSupplementalGroups

`func (o *V1SecurityContextConstraints) HasSupplementalGroups() bool`

HasSupplementalGroups returns a boolean if a field has been set.

### GetUsers

`func (o *V1SecurityContextConstraints) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V1SecurityContextConstraints) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V1SecurityContextConstraints) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetVolumes

`func (o *V1SecurityContextConstraints) GetVolumes() []V1FSType`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *V1SecurityContextConstraints) GetVolumesOk() (*[]V1FSType, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *V1SecurityContextConstraints) SetVolumes(v []V1FSType)`

SetVolumes sets Volumes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


