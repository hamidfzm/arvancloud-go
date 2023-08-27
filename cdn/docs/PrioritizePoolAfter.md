# PrioritizePoolAfter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolId** | **string** | ID of the pool you want to move | 
**AfterPoolId** | **string** | ID of the pool you want to be prior to the selected pool | 

## Methods

### NewPrioritizePoolAfter

`func NewPrioritizePoolAfter(poolId string, afterPoolId string, ) *PrioritizePoolAfter`

NewPrioritizePoolAfter instantiates a new PrioritizePoolAfter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrioritizePoolAfterWithDefaults

`func NewPrioritizePoolAfterWithDefaults() *PrioritizePoolAfter`

NewPrioritizePoolAfterWithDefaults instantiates a new PrioritizePoolAfter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolId

`func (o *PrioritizePoolAfter) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *PrioritizePoolAfter) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *PrioritizePoolAfter) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetAfterPoolId

`func (o *PrioritizePoolAfter) GetAfterPoolId() string`

GetAfterPoolId returns the AfterPoolId field if non-nil, zero value otherwise.

### GetAfterPoolIdOk

`func (o *PrioritizePoolAfter) GetAfterPoolIdOk() (*string, bool)`

GetAfterPoolIdOk returns a tuple with the AfterPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterPoolId

`func (o *PrioritizePoolAfter) SetAfterPoolId(v string)`

SetAfterPoolId sets AfterPoolId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


