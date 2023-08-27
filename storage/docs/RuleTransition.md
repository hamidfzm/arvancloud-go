# RuleTransition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Days** | Pointer to **int32** |  | [optional] 
**StorageClass** | Pointer to [**TransitionStorageClass**](TransitionStorageClass.md) |  | [optional] 

## Methods

### NewRuleTransition

`func NewRuleTransition() *RuleTransition`

NewRuleTransition instantiates a new RuleTransition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTransitionWithDefaults

`func NewRuleTransitionWithDefaults() *RuleTransition`

NewRuleTransitionWithDefaults instantiates a new RuleTransition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *RuleTransition) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RuleTransition) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RuleTransition) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *RuleTransition) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDays

`func (o *RuleTransition) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *RuleTransition) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *RuleTransition) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *RuleTransition) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetStorageClass

`func (o *RuleTransition) GetStorageClass() TransitionStorageClass`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *RuleTransition) GetStorageClassOk() (*TransitionStorageClass, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *RuleTransition) SetStorageClass(v TransitionStorageClass)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *RuleTransition) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


