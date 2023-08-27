# Transition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Days** | Pointer to **int32** |  | [optional] 
**StorageClass** | Pointer to [**TransitionStorageClass**](TransitionStorageClass.md) |  | [optional] 

## Methods

### NewTransition

`func NewTransition() *Transition`

NewTransition instantiates a new Transition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionWithDefaults

`func NewTransitionWithDefaults() *Transition`

NewTransitionWithDefaults instantiates a new Transition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *Transition) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Transition) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Transition) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *Transition) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDays

`func (o *Transition) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *Transition) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *Transition) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *Transition) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetStorageClass

`func (o *Transition) GetStorageClass() TransitionStorageClass`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *Transition) GetStorageClassOk() (*TransitionStorageClass, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *Transition) SetStorageClass(v TransitionStorageClass)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *Transition) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


