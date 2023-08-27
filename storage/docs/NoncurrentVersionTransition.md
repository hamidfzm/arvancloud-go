# NoncurrentVersionTransition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoncurrentDays** | Pointer to **int32** |  | [optional] 
**StorageClass** | Pointer to [**TransitionStorageClass**](TransitionStorageClass.md) |  | [optional] 

## Methods

### NewNoncurrentVersionTransition

`func NewNoncurrentVersionTransition() *NoncurrentVersionTransition`

NewNoncurrentVersionTransition instantiates a new NoncurrentVersionTransition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoncurrentVersionTransitionWithDefaults

`func NewNoncurrentVersionTransitionWithDefaults() *NoncurrentVersionTransition`

NewNoncurrentVersionTransitionWithDefaults instantiates a new NoncurrentVersionTransition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoncurrentDays

`func (o *NoncurrentVersionTransition) GetNoncurrentDays() int32`

GetNoncurrentDays returns the NoncurrentDays field if non-nil, zero value otherwise.

### GetNoncurrentDaysOk

`func (o *NoncurrentVersionTransition) GetNoncurrentDaysOk() (*int32, bool)`

GetNoncurrentDaysOk returns a tuple with the NoncurrentDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoncurrentDays

`func (o *NoncurrentVersionTransition) SetNoncurrentDays(v int32)`

SetNoncurrentDays sets NoncurrentDays field to given value.

### HasNoncurrentDays

`func (o *NoncurrentVersionTransition) HasNoncurrentDays() bool`

HasNoncurrentDays returns a boolean if a field has been set.

### GetStorageClass

`func (o *NoncurrentVersionTransition) GetStorageClass() TransitionStorageClass`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *NoncurrentVersionTransition) GetStorageClassOk() (*TransitionStorageClass, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *NoncurrentVersionTransition) SetStorageClass(v TransitionStorageClass)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *NoncurrentVersionTransition) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


