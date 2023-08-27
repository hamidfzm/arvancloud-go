# PrioritizePoolBefore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolId** | **string** | ID of the pool you want to move | 
**BeforePoolId** | **string** | ID of the pool you want to follow the selected pool | 

## Methods

### NewPrioritizePoolBefore

`func NewPrioritizePoolBefore(poolId string, beforePoolId string, ) *PrioritizePoolBefore`

NewPrioritizePoolBefore instantiates a new PrioritizePoolBefore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrioritizePoolBeforeWithDefaults

`func NewPrioritizePoolBeforeWithDefaults() *PrioritizePoolBefore`

NewPrioritizePoolBeforeWithDefaults instantiates a new PrioritizePoolBefore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolId

`func (o *PrioritizePoolBefore) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PrioritizePoolBefore) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PrioritizePoolBefore) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetBeforePoolId

`func (o *PrioritizePoolBefore) GetBeforePoolId() string`

GetBeforePoolId returns the BeforePoolId field if non-nil, zero value otherwise.

### GetBeforePoolIdOk

`func (o *PrioritizePoolBefore) GetBeforePoolIdOk() (*string, bool)`

GetBeforePoolIdOk returns a tuple with the BeforePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforePoolId

`func (o *PrioritizePoolBefore) SetBeforePoolId(v string)`

SetBeforePoolId sets BeforePoolId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


