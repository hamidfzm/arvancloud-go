# V1HorizontalPodAutoscalerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentCPUUtilizationPercentage** | Pointer to **int32** | current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU. | [optional] 
**CurrentReplicas** | **int32** | current number of replicas of pods managed by this autoscaler. | 
**DesiredReplicas** | **int32** | desired number of replicas of pods managed by this autoscaler. | 
**LastScaleTime** | Pointer to **string** | last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed. | [optional] 
**ObservedGeneration** | Pointer to **int64** | most recent generation observed by this autoscaler. | [optional] 

## Methods

### NewV1HorizontalPodAutoscalerStatus

`func NewV1HorizontalPodAutoscalerStatus(currentReplicas int32, desiredReplicas int32, ) *V1HorizontalPodAutoscalerStatus`

NewV1HorizontalPodAutoscalerStatus instantiates a new V1HorizontalPodAutoscalerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HorizontalPodAutoscalerStatusWithDefaults

`func NewV1HorizontalPodAutoscalerStatusWithDefaults() *V1HorizontalPodAutoscalerStatus`

NewV1HorizontalPodAutoscalerStatusWithDefaults instantiates a new V1HorizontalPodAutoscalerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentCPUUtilizationPercentage

`func (o *V1HorizontalPodAutoscalerStatus) GetCurrentCPUUtilizationPercentage() int32`

GetCurrentCPUUtilizationPercentage returns the CurrentCPUUtilizationPercentage field if non-nil, zero value otherwise.

### GetCurrentCPUUtilizationPercentageOk

`func (o *V1HorizontalPodAutoscalerStatus) GetCurrentCPUUtilizationPercentageOk() (*int32, bool)`

GetCurrentCPUUtilizationPercentageOk returns a tuple with the CurrentCPUUtilizationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCPUUtilizationPercentage

`func (o *V1HorizontalPodAutoscalerStatus) SetCurrentCPUUtilizationPercentage(v int32)`

SetCurrentCPUUtilizationPercentage sets CurrentCPUUtilizationPercentage field to given value.

### HasCurrentCPUUtilizationPercentage

`func (o *V1HorizontalPodAutoscalerStatus) HasCurrentCPUUtilizationPercentage() bool`

HasCurrentCPUUtilizationPercentage returns a boolean if a field has been set.

### GetCurrentReplicas

`func (o *V1HorizontalPodAutoscalerStatus) GetCurrentReplicas() int32`

GetCurrentReplicas returns the CurrentReplicas field if non-nil, zero value otherwise.

### GetCurrentReplicasOk

`func (o *V1HorizontalPodAutoscalerStatus) GetCurrentReplicasOk() (*int32, bool)`

GetCurrentReplicasOk returns a tuple with the CurrentReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentReplicas

`func (o *V1HorizontalPodAutoscalerStatus) SetCurrentReplicas(v int32)`

SetCurrentReplicas sets CurrentReplicas field to given value.


### GetDesiredReplicas

`func (o *V1HorizontalPodAutoscalerStatus) GetDesiredReplicas() int32`

GetDesiredReplicas returns the DesiredReplicas field if non-nil, zero value otherwise.

### GetDesiredReplicasOk

`func (o *V1HorizontalPodAutoscalerStatus) GetDesiredReplicasOk() (*int32, bool)`

GetDesiredReplicasOk returns a tuple with the DesiredReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredReplicas

`func (o *V1HorizontalPodAutoscalerStatus) SetDesiredReplicas(v int32)`

SetDesiredReplicas sets DesiredReplicas field to given value.


### GetLastScaleTime

`func (o *V1HorizontalPodAutoscalerStatus) GetLastScaleTime() string`

GetLastScaleTime returns the LastScaleTime field if non-nil, zero value otherwise.

### GetLastScaleTimeOk

`func (o *V1HorizontalPodAutoscalerStatus) GetLastScaleTimeOk() (*string, bool)`

GetLastScaleTimeOk returns a tuple with the LastScaleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScaleTime

`func (o *V1HorizontalPodAutoscalerStatus) SetLastScaleTime(v string)`

SetLastScaleTime sets LastScaleTime field to given value.

### HasLastScaleTime

`func (o *V1HorizontalPodAutoscalerStatus) HasLastScaleTime() bool`

HasLastScaleTime returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *V1HorizontalPodAutoscalerStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *V1HorizontalPodAutoscalerStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *V1HorizontalPodAutoscalerStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *V1HorizontalPodAutoscalerStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


