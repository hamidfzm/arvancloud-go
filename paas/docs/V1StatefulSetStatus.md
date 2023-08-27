# V1StatefulSetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollisionCount** | Pointer to **int32** | collisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision. | [optional] 
**Conditions** | Pointer to [**[]V1StatefulSetCondition**](V1StatefulSetCondition.md) | Represents the latest available observations of a statefulset&#39;s current state. | [optional] 
**CurrentReplicas** | Pointer to **int32** | currentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision. | [optional] 
**CurrentRevision** | Pointer to **string** | currentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas). | [optional] 
**ObservedGeneration** | Pointer to **int64** | observedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet&#39;s generation, which is updated on mutation by the API Server. | [optional] 
**ReadyReplicas** | Pointer to **int32** | readyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition. | [optional] 
**Replicas** | **int32** | replicas is the number of Pods created by the StatefulSet controller. | 
**UpdateRevision** | Pointer to **string** | updateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas) | [optional] 
**UpdatedReplicas** | Pointer to **int32** | updatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision. | [optional] 

## Methods

### NewV1StatefulSetStatus

`func NewV1StatefulSetStatus(replicas int32, ) *V1StatefulSetStatus`

NewV1StatefulSetStatus instantiates a new V1StatefulSetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1StatefulSetStatusWithDefaults

`func NewV1StatefulSetStatusWithDefaults() *V1StatefulSetStatus`

NewV1StatefulSetStatusWithDefaults instantiates a new V1StatefulSetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollisionCount

`func (o *V1StatefulSetStatus) GetCollisionCount() int32`

GetCollisionCount returns the CollisionCount field if non-nil, zero value otherwise.

### GetCollisionCountOk

`func (o *V1StatefulSetStatus) GetCollisionCountOk() (*int32, bool)`

GetCollisionCountOk returns a tuple with the CollisionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollisionCount

`func (o *V1StatefulSetStatus) SetCollisionCount(v int32)`

SetCollisionCount sets CollisionCount field to given value.

### HasCollisionCount

`func (o *V1StatefulSetStatus) HasCollisionCount() bool`

HasCollisionCount returns a boolean if a field has been set.

### GetConditions

`func (o *V1StatefulSetStatus) GetConditions() []V1StatefulSetCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1StatefulSetStatus) GetConditionsOk() (*[]V1StatefulSetCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1StatefulSetStatus) SetConditions(v []V1StatefulSetCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1StatefulSetStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCurrentReplicas

`func (o *V1StatefulSetStatus) GetCurrentReplicas() int32`

GetCurrentReplicas returns the CurrentReplicas field if non-nil, zero value otherwise.

### GetCurrentReplicasOk

`func (o *V1StatefulSetStatus) GetCurrentReplicasOk() (*int32, bool)`

GetCurrentReplicasOk returns a tuple with the CurrentReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReplicas

`func (o *V1StatefulSetStatus) SetCurrentReplicas(v int32)`

SetCurrentReplicas sets CurrentReplicas field to given value.

### HasCurrentReplicas

`func (o *V1StatefulSetStatus) HasCurrentReplicas() bool`

HasCurrentReplicas returns a boolean if a field has been set.

### GetCurrentRevision

`func (o *V1StatefulSetStatus) GetCurrentRevision() string`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *V1StatefulSetStatus) GetCurrentRevisionOk() (*string, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *V1StatefulSetStatus) SetCurrentRevision(v string)`

SetCurrentRevision sets CurrentRevision field to given value.

### HasCurrentRevision

`func (o *V1StatefulSetStatus) HasCurrentRevision() bool`

HasCurrentRevision returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *V1StatefulSetStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *V1StatefulSetStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *V1StatefulSetStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *V1StatefulSetStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *V1StatefulSetStatus) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *V1StatefulSetStatus) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *V1StatefulSetStatus) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *V1StatefulSetStatus) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *V1StatefulSetStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1StatefulSetStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1StatefulSetStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetUpdateRevision

`func (o *V1StatefulSetStatus) GetUpdateRevision() string`

GetUpdateRevision returns the UpdateRevision field if non-nil, zero value otherwise.

### GetUpdateRevisionOk

`func (o *V1StatefulSetStatus) GetUpdateRevisionOk() (*string, bool)`

GetUpdateRevisionOk returns a tuple with the UpdateRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRevision

`func (o *V1StatefulSetStatus) SetUpdateRevision(v string)`

SetUpdateRevision sets UpdateRevision field to given value.

### HasUpdateRevision

`func (o *V1StatefulSetStatus) HasUpdateRevision() bool`

HasUpdateRevision returns a boolean if a field has been set.

### GetUpdatedReplicas

`func (o *V1StatefulSetStatus) GetUpdatedReplicas() int32`

GetUpdatedReplicas returns the UpdatedReplicas field if non-nil, zero value otherwise.

### GetUpdatedReplicasOk

`func (o *V1StatefulSetStatus) GetUpdatedReplicasOk() (*int32, bool)`

GetUpdatedReplicasOk returns a tuple with the UpdatedReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedReplicas

`func (o *V1StatefulSetStatus) SetUpdatedReplicas(v int32)`

SetUpdatedReplicas sets UpdatedReplicas field to given value.

### HasUpdatedReplicas

`func (o *V1StatefulSetStatus) HasUpdatedReplicas() bool`

HasUpdatedReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


