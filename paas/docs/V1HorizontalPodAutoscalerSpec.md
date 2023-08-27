# V1HorizontalPodAutoscalerSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxReplicas** | **int32** | upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas. | 
**MinReplicas** | Pointer to **int32** | lower limit for the number of pods that can be set by the autoscaler, default 1. | [optional] 
**ScaleTargetRef** | [**V1CrossVersionObjectReference**](V1CrossVersionObjectReference.md) |  | 
**TargetCPUUtilizationPercentage** | Pointer to **int32** | target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used. | [optional] 

## Methods

### NewV1HorizontalPodAutoscalerSpec

`func NewV1HorizontalPodAutoscalerSpec(maxReplicas int32, scaleTargetRef V1CrossVersionObjectReference, ) *V1HorizontalPodAutoscalerSpec`

NewV1HorizontalPodAutoscalerSpec instantiates a new V1HorizontalPodAutoscalerSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HorizontalPodAutoscalerSpecWithDefaults

`func NewV1HorizontalPodAutoscalerSpecWithDefaults() *V1HorizontalPodAutoscalerSpec`

NewV1HorizontalPodAutoscalerSpecWithDefaults instantiates a new V1HorizontalPodAutoscalerSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxReplicas

`func (o *V1HorizontalPodAutoscalerSpec) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *V1HorizontalPodAutoscalerSpec) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *V1HorizontalPodAutoscalerSpec) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.


### GetMinReplicas

`func (o *V1HorizontalPodAutoscalerSpec) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *V1HorizontalPodAutoscalerSpec) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *V1HorizontalPodAutoscalerSpec) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *V1HorizontalPodAutoscalerSpec) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetScaleTargetRef

`func (o *V1HorizontalPodAutoscalerSpec) GetScaleTargetRef() V1CrossVersionObjectReference`

GetScaleTargetRef returns the ScaleTargetRef field if non-nil, zero value otherwise.

### GetScaleTargetRefOk

`func (o *V1HorizontalPodAutoscalerSpec) GetScaleTargetRefOk() (*V1CrossVersionObjectReference, bool)`

GetScaleTargetRefOk returns a tuple with the ScaleTargetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleTargetRef

`func (o *V1HorizontalPodAutoscalerSpec) SetScaleTargetRef(v V1CrossVersionObjectReference)`

SetScaleTargetRef sets ScaleTargetRef field to given value.


### GetTargetCPUUtilizationPercentage

`func (o *V1HorizontalPodAutoscalerSpec) GetTargetCPUUtilizationPercentage() int32`

GetTargetCPUUtilizationPercentage returns the TargetCPUUtilizationPercentage field if non-nil, zero value otherwise.

### GetTargetCPUUtilizationPercentageOk

`func (o *V1HorizontalPodAutoscalerSpec) GetTargetCPUUtilizationPercentageOk() (*int32, bool)`

GetTargetCPUUtilizationPercentageOk returns a tuple with the TargetCPUUtilizationPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCPUUtilizationPercentage

`func (o *V1HorizontalPodAutoscalerSpec) SetTargetCPUUtilizationPercentage(v int32)`

SetTargetCPUUtilizationPercentage sets TargetCPUUtilizationPercentage field to given value.

### HasTargetCPUUtilizationPercentage

`func (o *V1HorizontalPodAutoscalerSpec) HasTargetCPUUtilizationPercentage() bool`

HasTargetCPUUtilizationPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


