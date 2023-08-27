# V1NodeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigSource** | Pointer to [**V1NodeConfigSource**](V1NodeConfigSource.md) |  | [optional] 
**ExternalID** | Pointer to **string** | Deprecated. Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966 | [optional] 
**PodCIDR** | Pointer to **string** | PodCIDR represents the pod IP range assigned to the node. | [optional] 
**ProviderID** | Pointer to **string** | ID of the node assigned by the cloud provider in the format: &lt;ProviderName&gt;://&lt;ProviderSpecificNodeID&gt; | [optional] 
**Taints** | Pointer to [**[]V1Taint**](V1Taint.md) | If specified, the node&#39;s taints. | [optional] 
**Unschedulable** | Pointer to **bool** | Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration | [optional] 

## Methods

### NewV1NodeSpec

`func NewV1NodeSpec() *V1NodeSpec`

NewV1NodeSpec instantiates a new V1NodeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NodeSpecWithDefaults

`func NewV1NodeSpecWithDefaults() *V1NodeSpec`

NewV1NodeSpecWithDefaults instantiates a new V1NodeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigSource

`func (o *V1NodeSpec) GetConfigSource() V1NodeConfigSource`

GetConfigSource returns the ConfigSource field if non-nil, zero value otherwise.

### GetConfigSourceOk

`func (o *V1NodeSpec) GetConfigSourceOk() (*V1NodeConfigSource, bool)`

GetConfigSourceOk returns a tuple with the ConfigSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSource

`func (o *V1NodeSpec) SetConfigSource(v V1NodeConfigSource)`

SetConfigSource sets ConfigSource field to given value.

### HasConfigSource

`func (o *V1NodeSpec) HasConfigSource() bool`

HasConfigSource returns a boolean if a field has been set.

### GetExternalID

`func (o *V1NodeSpec) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *V1NodeSpec) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *V1NodeSpec) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *V1NodeSpec) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetPodCIDR

`func (o *V1NodeSpec) GetPodCIDR() string`

GetPodCIDR returns the PodCIDR field if non-nil, zero value otherwise.

### GetPodCIDROk

`func (o *V1NodeSpec) GetPodCIDROk() (*string, bool)`

GetPodCIDROk returns a tuple with the PodCIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCIDR

`func (o *V1NodeSpec) SetPodCIDR(v string)`

SetPodCIDR sets PodCIDR field to given value.

### HasPodCIDR

`func (o *V1NodeSpec) HasPodCIDR() bool`

HasPodCIDR returns a boolean if a field has been set.

### GetProviderID

`func (o *V1NodeSpec) GetProviderID() string`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *V1NodeSpec) GetProviderIDOk() (*string, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *V1NodeSpec) SetProviderID(v string)`

SetProviderID sets ProviderID field to given value.

### HasProviderID

`func (o *V1NodeSpec) HasProviderID() bool`

HasProviderID returns a boolean if a field has been set.

### GetTaints

`func (o *V1NodeSpec) GetTaints() []V1Taint`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *V1NodeSpec) GetTaintsOk() (*[]V1Taint, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *V1NodeSpec) SetTaints(v []V1Taint)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *V1NodeSpec) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### GetUnschedulable

`func (o *V1NodeSpec) GetUnschedulable() bool`

GetUnschedulable returns the Unschedulable field if non-nil, zero value otherwise.

### GetUnschedulableOk

`func (o *V1NodeSpec) GetUnschedulableOk() (*bool, bool)`

GetUnschedulableOk returns a tuple with the Unschedulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnschedulable

`func (o *V1NodeSpec) SetUnschedulable(v bool)`

SetUnschedulable sets Unschedulable field to given value.

### HasUnschedulable

`func (o *V1NodeSpec) HasUnschedulable() bool`

HasUnschedulable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


