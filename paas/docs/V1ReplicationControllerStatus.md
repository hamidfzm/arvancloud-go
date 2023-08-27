# V1ReplicationControllerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | Pointer to **int32** | The number of available replicas (ready for at least minReadySeconds) for this replication controller. | [optional] 
**Conditions** | Pointer to [**[]V1ReplicationControllerCondition**](V1ReplicationControllerCondition.md) | Represents the latest available observations of a replication controller&#39;s current state. | [optional] 
**FullyLabeledReplicas** | Pointer to **int32** | The number of pods that have labels matching the labels of the pod template of the replication controller. | [optional] 
**ObservedGeneration** | Pointer to **int64** | ObservedGeneration reflects the generation of the most recently observed replication controller. | [optional] 
**ReadyReplicas** | Pointer to **int32** | The number of ready replicas for this replication controller. | [optional] 
**Replicas** | **int32** | Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller | 

## Methods

### NewV1ReplicationControllerStatus

`func NewV1ReplicationControllerStatus(replicas int32, ) *V1ReplicationControllerStatus`

NewV1ReplicationControllerStatus instantiates a new V1ReplicationControllerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReplicationControllerStatusWithDefaults

`func NewV1ReplicationControllerStatusWithDefaults() *V1ReplicationControllerStatus`

NewV1ReplicationControllerStatusWithDefaults instantiates a new V1ReplicationControllerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableReplicas

`func (o *V1ReplicationControllerStatus) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *V1ReplicationControllerStatus) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *V1ReplicationControllerStatus) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *V1ReplicationControllerStatus) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetConditions

`func (o *V1ReplicationControllerStatus) GetConditions() []V1ReplicationControllerCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1ReplicationControllerStatus) GetConditionsOk() (*[]V1ReplicationControllerCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1ReplicationControllerStatus) SetConditions(v []V1ReplicationControllerCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1ReplicationControllerStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetFullyLabeledReplicas

`func (o *V1ReplicationControllerStatus) GetFullyLabeledReplicas() int32`

GetFullyLabeledReplicas returns the FullyLabeledReplicas field if non-nil, zero value otherwise.

### GetFullyLabeledReplicasOk

`func (o *V1ReplicationControllerStatus) GetFullyLabeledReplicasOk() (*int32, bool)`

GetFullyLabeledReplicasOk returns a tuple with the FullyLabeledReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyLabeledReplicas

`func (o *V1ReplicationControllerStatus) SetFullyLabeledReplicas(v int32)`

SetFullyLabeledReplicas sets FullyLabeledReplicas field to given value.

### HasFullyLabeledReplicas

`func (o *V1ReplicationControllerStatus) HasFullyLabeledReplicas() bool`

HasFullyLabeledReplicas returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *V1ReplicationControllerStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *V1ReplicationControllerStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *V1ReplicationControllerStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *V1ReplicationControllerStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *V1ReplicationControllerStatus) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *V1ReplicationControllerStatus) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *V1ReplicationControllerStatus) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *V1ReplicationControllerStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *V1ReplicationControllerStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1ReplicationControllerStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1ReplicationControllerStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


