# V1Initializers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | [**[]V1Initializer**](V1Initializer.md) | Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients. | 
**Result** | Pointer to [**V1Status**](V1Status.md) |  | [optional] 

## Methods

### NewV1Initializers

`func NewV1Initializers(pending []V1Initializer, ) *V1Initializers`

NewV1Initializers instantiates a new V1Initializers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InitializersWithDefaults

`func NewV1InitializersWithDefaults() *V1Initializers`

NewV1InitializersWithDefaults instantiates a new V1Initializers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *V1Initializers) GetPending() []V1Initializer`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *V1Initializers) GetPendingOk() (*[]V1Initializer, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *V1Initializers) SetPending(v []V1Initializer)`

SetPending sets Pending field to given value.


### GetResult

`func (o *V1Initializers) GetResult() V1Status`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *V1Initializers) GetResultOk() (*V1Status, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *V1Initializers) SetResult(v V1Status)`

SetResult sets Result field to given value.

### HasResult

`func (o *V1Initializers) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


