# PrioritizePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolId** | **string** | ID of the pool you want to move | 
**AfterPoolId** | **string** | ID of the pool you want to be prior to the selected pool | 
**BeforePoolId** | **string** | ID of the pool you want to follow the selected pool | 

## Methods

### NewPrioritizePool

`func NewPrioritizePool(poolId string, afterPoolId string, beforePoolId string, ) *PrioritizePool`

NewPrioritizePool instantiates a new PrioritizePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrioritizePoolWithDefaults

`func NewPrioritizePoolWithDefaults() *PrioritizePool`

NewPrioritizePoolWithDefaults instantiates a new PrioritizePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolId

`func (o *PrioritizePool) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PrioritizePool) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PrioritizePool) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetAfterPoolId

`func (o *PrioritizePool) GetAfterPoolId() string`

GetAfterPoolId returns the AfterPoolId field if non-nil, zero value otherwise.

### GetAfterPoolIdOk

`func (o *PrioritizePool) GetAfterPoolIdOk() (*string, bool)`

GetAfterPoolIdOk returns a tuple with the AfterPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterPoolId

`func (o *PrioritizePool) SetAfterPoolId(v string)`

SetAfterPoolId sets AfterPoolId field to given value.


### GetBeforePoolId

`func (o *PrioritizePool) GetBeforePoolId() string`

GetBeforePoolId returns the BeforePoolId field if non-nil, zero value otherwise.

### GetBeforePoolIdOk

`func (o *PrioritizePool) GetBeforePoolIdOk() (*string, bool)`

GetBeforePoolIdOk returns a tuple with the BeforePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforePoolId

`func (o *PrioritizePool) SetBeforePoolId(v string)`

SetBeforePoolId sets BeforePoolId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


