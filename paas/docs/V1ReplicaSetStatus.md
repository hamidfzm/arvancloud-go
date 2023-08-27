# V1ReplicaSetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableReplicas** | Pointer to **int32** | The number of available replicas (ready for at least minReadySeconds) for this replica set. | [optional] 
**Conditions** | Pointer to [**[]V1ReplicaSetCondition**](V1ReplicaSetCondition.md) | Represents the latest available observations of a replica set&#39;s current state. | [optional] 
**FullyLabeledReplicas** | Pointer to **int32** | The number of pods that have labels matching the labels of the pod template of the replicaset. | [optional] 
**ObservedGeneration** | Pointer to **int64** | ObservedGeneration reflects the generation of the most recently observed ReplicaSet. | [optional] 
**ReadyReplicas** | Pointer to **int32** | The number of ready replicas for this replica set. | [optional] 
**Replicas** | **int32** | Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller | 

## Methods

### NewV1ReplicaSetStatus

`func NewV1ReplicaSetStatus(replicas int32, ) *V1ReplicaSetStatus`

NewV1ReplicaSetStatus instantiates a new V1ReplicaSetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReplicaSetStatusWithDefaults

`func NewV1ReplicaSetStatusWithDefaults() *V1ReplicaSetStatus`

NewV1ReplicaSetStatusWithDefaults instantiates a new V1ReplicaSetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableReplicas

`func (o *V1ReplicaSetStatus) GetAvailableReplicas() int32`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *V1ReplicaSetStatus) GetAvailableReplicasOk() (*int32, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *V1ReplicaSetStatus) SetAvailableReplicas(v int32)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *V1ReplicaSetStatus) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetConditions

`func (o *V1ReplicaSetStatus) GetConditions() []V1ReplicaSetCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1ReplicaSetStatus) GetConditionsOk() (*[]V1ReplicaSetCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1ReplicaSetStatus) SetConditions(v []V1ReplicaSetCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1ReplicaSetStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetFullyLabeledReplicas

`func (o *V1ReplicaSetStatus) GetFullyLabeledReplicas() int32`

GetFullyLabeledReplicas returns the FullyLabeledReplicas field if non-nil, zero value otherwise.

### GetFullyLabeledReplicasOk

`func (o *V1ReplicaSetStatus) GetFullyLabeledReplicasOk() (*int32, bool)`

GetFullyLabeledReplicasOk returns a tuple with the FullyLabeledReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyLabeledReplicas

`func (o *V1ReplicaSetStatus) SetFullyLabeledReplicas(v int32)`

SetFullyLabeledReplicas sets FullyLabeledReplicas field to given value.

### HasFullyLabeledReplicas

`func (o *V1ReplicaSetStatus) HasFullyLabeledReplicas() bool`

HasFullyLabeledReplicas returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *V1ReplicaSetStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *V1ReplicaSetStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *V1ReplicaSetStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *V1ReplicaSetStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *V1ReplicaSetStatus) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *V1ReplicaSetStatus) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *V1ReplicaSetStatus) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *V1ReplicaSetStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *V1ReplicaSetStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1ReplicaSetStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1ReplicaSetStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


