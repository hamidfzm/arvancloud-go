# V1ScaleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicas** | **int32** | actual number of observed instances of the scaled object. | 
**Selector** | Pointer to **string** | label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors | [optional] 

## Methods

### NewV1ScaleStatus

`func NewV1ScaleStatus(replicas int32, ) *V1ScaleStatus`

NewV1ScaleStatus instantiates a new V1ScaleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ScaleStatusWithDefaults

`func NewV1ScaleStatusWithDefaults() *V1ScaleStatus`

NewV1ScaleStatusWithDefaults instantiates a new V1ScaleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicas

`func (o *V1ScaleStatus) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1ScaleStatus) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1ScaleStatus) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetSelector

`func (o *V1ScaleStatus) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1ScaleStatus) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1ScaleStatus) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1ScaleStatus) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


