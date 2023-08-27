# V1beta1ScaleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | **int32** | actual number of observed instances of the scaled object. | 
**Selector** | Pointer to **map[string]interface{}** | label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors | [optional] 
**TargetSelector** | Pointer to **string** | label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors | [optional] 

## Methods

### NewV1beta1ScaleStatus

`func NewV1beta1ScaleStatus(replicas int32, ) *V1beta1ScaleStatus`

NewV1beta1ScaleStatus instantiates a new V1beta1ScaleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1ScaleStatusWithDefaults

`func NewV1beta1ScaleStatusWithDefaults() *V1beta1ScaleStatus`

NewV1beta1ScaleStatusWithDefaults instantiates a new V1beta1ScaleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicas

`func (o *V1beta1ScaleStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1beta1ScaleStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1beta1ScaleStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetSelector

`func (o *V1beta1ScaleStatus) GetSelector() map[string]interface{}`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1beta1ScaleStatus) GetSelectorOk() (*map[string]interface{}, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1beta1ScaleStatus) SetSelector(v map[string]interface{})`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1beta1ScaleStatus) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetTargetSelector

`func (o *V1beta1ScaleStatus) GetTargetSelector() string`

GetTargetSelector returns the TargetSelector field if non-nil, zero value otherwise.

### GetTargetSelectorOk

`func (o *V1beta1ScaleStatus) GetTargetSelectorOk() (*string, bool)`

GetTargetSelectorOk returns a tuple with the TargetSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSelector

`func (o *V1beta1ScaleStatus) SetTargetSelector(v string)`

SetTargetSelector sets TargetSelector field to given value.

### HasTargetSelector

`func (o *V1beta1ScaleStatus) HasTargetSelector() bool`

HasTargetSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


