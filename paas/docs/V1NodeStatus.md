# V1NodeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]V1NodeAddress**](V1NodeAddress.md) | List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses | [optional] 
**Allocatable** | Pointer to **map[string]interface{}** | Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity. | [optional] 
**Capacity** | Pointer to **map[string]interface{}** | Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity | [optional] 
**Conditions** | Pointer to [**[]V1NodeCondition**](V1NodeCondition.md) | Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition | [optional] 
**Config** | Pointer to [**V1NodeConfigStatus**](V1NodeConfigStatus.md) |  | [optional] 
**DaemonEndpoints** | Pointer to [**V1NodeDaemonEndpoints**](V1NodeDaemonEndpoints.md) |  | [optional] 
**Images** | Pointer to [**[]V1ContainerImage**](V1ContainerImage.md) | List of container images on this node | [optional] 
**NodeInfo** | Pointer to [**V1NodeSystemInfo**](V1NodeSystemInfo.md) |  | [optional] 
**Phase** | Pointer to **string** | NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated. | [optional] 
**VolumesAttached** | Pointer to [**[]V1AttachedVolume**](V1AttachedVolume.md) | List of volumes that are attached to the node. | [optional] 
**VolumesInUse** | Pointer to [**[]V1UniqueVolumeName**](V1UniqueVolumeName.md) | List of attachable volumes in use (mounted) by the node. | [optional] 

## Methods

### NewV1NodeStatus

`func NewV1NodeStatus() *V1NodeStatus`

NewV1NodeStatus instantiates a new V1NodeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NodeStatusWithDefaults

`func NewV1NodeStatusWithDefaults() *V1NodeStatus`

NewV1NodeStatusWithDefaults instantiates a new V1NodeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *V1NodeStatus) GetAddresses() []V1NodeAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *V1NodeStatus) GetAddressesOk() (*[]V1NodeAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *V1NodeStatus) SetAddresses(v []V1NodeAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *V1NodeStatus) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAllocatable

`func (o *V1NodeStatus) GetAllocatable() map[string]interface{}`

GetAllocatable returns the Allocatable field if non-nil, zero value otherwise.

### GetAllocatableOk

`func (o *V1NodeStatus) GetAllocatableOk() (*map[string]interface{}, bool)`

GetAllocatableOk returns a tuple with the Allocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatable

`func (o *V1NodeStatus) SetAllocatable(v map[string]interface{})`

SetAllocatable sets Allocatable field to given value.

### HasAllocatable

`func (o *V1NodeStatus) HasAllocatable() bool`

HasAllocatable returns a boolean if a field has been set.

### GetCapacity

`func (o *V1NodeStatus) GetCapacity() map[string]interface{}`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *V1NodeStatus) GetCapacityOk() (*map[string]interface{}, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *V1NodeStatus) SetCapacity(v map[string]interface{})`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *V1NodeStatus) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetConditions

`func (o *V1NodeStatus) GetConditions() []V1NodeCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1NodeStatus) GetConditionsOk() (*[]V1NodeCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1NodeStatus) SetConditions(v []V1NodeCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1NodeStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetConfig

`func (o *V1NodeStatus) GetConfig() V1NodeConfigStatus`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1NodeStatus) GetConfigOk() (*V1NodeConfigStatus, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1NodeStatus) SetConfig(v V1NodeConfigStatus)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1NodeStatus) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDaemonEndpoints

`func (o *V1NodeStatus) GetDaemonEndpoints() V1NodeDaemonEndpoints`

GetDaemonEndpoints returns the DaemonEndpoints field if non-nil, zero value otherwise.

### GetDaemonEndpointsOk

`func (o *V1NodeStatus) GetDaemonEndpointsOk() (*V1NodeDaemonEndpoints, bool)`

GetDaemonEndpointsOk returns a tuple with the DaemonEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonEndpoints

`func (o *V1NodeStatus) SetDaemonEndpoints(v V1NodeDaemonEndpoints)`

SetDaemonEndpoints sets DaemonEndpoints field to given value.

### HasDaemonEndpoints

`func (o *V1NodeStatus) HasDaemonEndpoints() bool`

HasDaemonEndpoints returns a boolean if a field has been set.

### GetImages

`func (o *V1NodeStatus) GetImages() []V1ContainerImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *V1NodeStatus) GetImagesOk() (*[]V1ContainerImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *V1NodeStatus) SetImages(v []V1ContainerImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *V1NodeStatus) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetNodeInfo

`func (o *V1NodeStatus) GetNodeInfo() V1NodeSystemInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *V1NodeStatus) GetNodeInfoOk() (*V1NodeSystemInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *V1NodeStatus) SetNodeInfo(v V1NodeSystemInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *V1NodeStatus) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetPhase

`func (o *V1NodeStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1NodeStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1NodeStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1NodeStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetVolumesAttached

`func (o *V1NodeStatus) GetVolumesAttached() []V1AttachedVolume`

GetVolumesAttached returns the VolumesAttached field if non-nil, zero value otherwise.

### GetVolumesAttachedOk

`func (o *V1NodeStatus) GetVolumesAttachedOk() (*[]V1AttachedVolume, bool)`

GetVolumesAttachedOk returns a tuple with the VolumesAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesAttached

`func (o *V1NodeStatus) SetVolumesAttached(v []V1AttachedVolume)`

SetVolumesAttached sets VolumesAttached field to given value.

### HasVolumesAttached

`func (o *V1NodeStatus) HasVolumesAttached() bool`

HasVolumesAttached returns a boolean if a field has been set.

### GetVolumesInUse

`func (o *V1NodeStatus) GetVolumesInUse() []V1UniqueVolumeName`

GetVolumesInUse returns the VolumesInUse field if non-nil, zero value otherwise.

### GetVolumesInUseOk

`func (o *V1NodeStatus) GetVolumesInUseOk() (*[]V1UniqueVolumeName, bool)`

GetVolumesInUseOk returns a tuple with the VolumesInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesInUse

`func (o *V1NodeStatus) SetVolumesInUse(v []V1UniqueVolumeName)`

SetVolumesInUse sets VolumesInUse field to given value.

### HasVolumesInUse

`func (o *V1NodeStatus) HasVolumesInUse() bool`

HasVolumesInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


