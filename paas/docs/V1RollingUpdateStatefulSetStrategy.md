# V1RollingUpdateStatefulSetStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Partition** | Pointer to **int32** | Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0. | [optional] 

## Methods

### NewV1RollingUpdateStatefulSetStrategy

`func NewV1RollingUpdateStatefulSetStrategy() *V1RollingUpdateStatefulSetStrategy`

NewV1RollingUpdateStatefulSetStrategy instantiates a new V1RollingUpdateStatefulSetStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RollingUpdateStatefulSetStrategyWithDefaults

`func NewV1RollingUpdateStatefulSetStrategyWithDefaults() *V1RollingUpdateStatefulSetStrategy`

NewV1RollingUpdateStatefulSetStrategyWithDefaults instantiates a new V1RollingUpdateStatefulSetStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartition

`func (o *V1RollingUpdateStatefulSetStrategy) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *V1RollingUpdateStatefulSetStrategy) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *V1RollingUpdateStatefulSetStrategy) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *V1RollingUpdateStatefulSetStrategy) HasPartition() bool`

HasPartition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


