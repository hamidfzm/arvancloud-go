# V1StatefulSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodManagementPolicy** | Pointer to **string** | podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down. The default policy is &#x60;OrderedReady&#x60;, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is &#x60;Parallel&#x60; which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once. | [optional] 
**Replicas** | Pointer to **int32** | replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1. | [optional] 
**RevisionHistoryLimit** | Pointer to **int32** | revisionHistoryLimit is the maximum number of revisions that will be maintained in the StatefulSet&#39;s revision history. The revision history consists of all revisions not represented by a currently applied StatefulSetSpec version. The default value is 10. | [optional] 
**Selector** | [**V1LabelSelector**](V1LabelSelector.md) |  | 
**ServiceName** | **string** | serviceName is the name of the service that governs this StatefulSet. This service must exist before the StatefulSet, and is responsible for the network identity of the set. Pods get DNS/hostnames that follow the pattern: pod-specific-string.serviceName.default.svc.cluster.local where \&quot;pod-specific-string\&quot; is managed by the StatefulSet controller. | 
**Template** | [**V1PodTemplateSpec**](V1PodTemplateSpec.md) |  | 
**UpdateStrategy** | Pointer to [**V1StatefulSetUpdateStrategy**](V1StatefulSetUpdateStrategy.md) |  | [optional] 
**VolumeClaimTemplates** | Pointer to [**[]V1PersistentVolumeClaim**](V1PersistentVolumeClaim.md) | volumeClaimTemplates is a list of claims that pods are allowed to reference. The StatefulSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pod. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name. | [optional] 

## Methods

### NewV1StatefulSetSpec

`func NewV1StatefulSetSpec(selector V1LabelSelector, serviceName string, template V1PodTemplateSpec, ) *V1StatefulSetSpec`

NewV1StatefulSetSpec instantiates a new V1StatefulSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1StatefulSetSpecWithDefaults

`func NewV1StatefulSetSpecWithDefaults() *V1StatefulSetSpec`

NewV1StatefulSetSpecWithDefaults instantiates a new V1StatefulSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodManagementPolicy

`func (o *V1StatefulSetSpec) GetPodManagementPolicy() string`

GetPodManagementPolicy returns the PodManagementPolicy field if non-nil, zero value otherwise.

### GetPodManagementPolicyOk

`func (o *V1StatefulSetSpec) GetPodManagementPolicyOk() (*string, bool)`

GetPodManagementPolicyOk returns a tuple with the PodManagementPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodManagementPolicy

`func (o *V1StatefulSetSpec) SetPodManagementPolicy(v string)`

SetPodManagementPolicy sets PodManagementPolicy field to given value.

### HasPodManagementPolicy

`func (o *V1StatefulSetSpec) HasPodManagementPolicy() bool`

HasPodManagementPolicy returns a boolean if a field has been set.

### GetReplicas

`func (o *V1StatefulSetSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1StatefulSetSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1StatefulSetSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *V1StatefulSetSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRevisionHistoryLimit

`func (o *V1StatefulSetSpec) GetRevisionHistoryLimit() int32`

GetRevisionHistoryLimit returns the RevisionHistoryLimit field if non-nil, zero value otherwise.

### GetRevisionHistoryLimitOk

`func (o *V1StatefulSetSpec) GetRevisionHistoryLimitOk() (*int32, bool)`

GetRevisionHistoryLimitOk returns a tuple with the RevisionHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistoryLimit

`func (o *V1StatefulSetSpec) SetRevisionHistoryLimit(v int32)`

SetRevisionHistoryLimit sets RevisionHistoryLimit field to given value.

### HasRevisionHistoryLimit

`func (o *V1StatefulSetSpec) HasRevisionHistoryLimit() bool`

HasRevisionHistoryLimit returns a boolean if a field has been set.

### GetSelector

`func (o *V1StatefulSetSpec) GetSelector() V1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1StatefulSetSpec) GetSelectorOk() (*V1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1StatefulSetSpec) SetSelector(v V1LabelSelector)`

SetSelector sets Selector field to given value.


### GetServiceName

`func (o *V1StatefulSetSpec) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1StatefulSetSpec) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1StatefulSetSpec) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetTemplate

`func (o *V1StatefulSetSpec) GetTemplate() V1PodTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1StatefulSetSpec) GetTemplateOk() (*V1PodTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1StatefulSetSpec) SetTemplate(v V1PodTemplateSpec)`

SetTemplate sets Template field to given value.


### GetUpdateStrategy

`func (o *V1StatefulSetSpec) GetUpdateStrategy() V1StatefulSetUpdateStrategy`

GetUpdateStrategy returns the UpdateStrategy field if non-nil, zero value otherwise.

### GetUpdateStrategyOk

`func (o *V1StatefulSetSpec) GetUpdateStrategyOk() (*V1StatefulSetUpdateStrategy, bool)`

GetUpdateStrategyOk returns a tuple with the UpdateStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStrategy

`func (o *V1StatefulSetSpec) SetUpdateStrategy(v V1StatefulSetUpdateStrategy)`

SetUpdateStrategy sets UpdateStrategy field to given value.

### HasUpdateStrategy

`func (o *V1StatefulSetSpec) HasUpdateStrategy() bool`

HasUpdateStrategy returns a boolean if a field has been set.

### GetVolumeClaimTemplates

`func (o *V1StatefulSetSpec) GetVolumeClaimTemplates() []V1PersistentVolumeClaim`

GetVolumeClaimTemplates returns the VolumeClaimTemplates field if non-nil, zero value otherwise.

### GetVolumeClaimTemplatesOk

`func (o *V1StatefulSetSpec) GetVolumeClaimTemplatesOk() (*[]V1PersistentVolumeClaim, bool)`

GetVolumeClaimTemplatesOk returns a tuple with the VolumeClaimTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClaimTemplates

`func (o *V1StatefulSetSpec) SetVolumeClaimTemplates(v []V1PersistentVolumeClaim)`

SetVolumeClaimTemplates sets VolumeClaimTemplates field to given value.

### HasVolumeClaimTemplates

`func (o *V1StatefulSetSpec) HasVolumeClaimTemplates() bool`

HasVolumeClaimTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


