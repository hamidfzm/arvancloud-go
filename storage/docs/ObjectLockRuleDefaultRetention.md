# ObjectLockRuleDefaultRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**ObjectLockRetentionMode**](ObjectLockRetentionMode.md) |  | [optional] 
**Days** | Pointer to **int32** |  | [optional] 
**Years** | Pointer to **int32** |  | [optional] 

## Methods

### NewObjectLockRuleDefaultRetention

`func NewObjectLockRuleDefaultRetention() *ObjectLockRuleDefaultRetention`

NewObjectLockRuleDefaultRetention instantiates a new ObjectLockRuleDefaultRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectLockRuleDefaultRetentionWithDefaults

`func NewObjectLockRuleDefaultRetentionWithDefaults() *ObjectLockRuleDefaultRetention`

NewObjectLockRuleDefaultRetentionWithDefaults instantiates a new ObjectLockRuleDefaultRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ObjectLockRuleDefaultRetention) GetMode() ObjectLockRetentionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ObjectLockRuleDefaultRetention) GetModeOk() (*ObjectLockRetentionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ObjectLockRuleDefaultRetention) SetMode(v ObjectLockRetentionMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ObjectLockRuleDefaultRetention) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetDays

`func (o *ObjectLockRuleDefaultRetention) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ObjectLockRuleDefaultRetention) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ObjectLockRuleDefaultRetention) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *ObjectLockRuleDefaultRetention) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetYears

`func (o *ObjectLockRuleDefaultRetention) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *ObjectLockRuleDefaultRetention) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *ObjectLockRuleDefaultRetention) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *ObjectLockRuleDefaultRetention) HasYears() bool`

HasYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


