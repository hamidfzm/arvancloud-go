# LifecycleRuleFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**LifecycleRuleFilterTag**](LifecycleRuleFilterTag.md) |  | [optional] 
**And** | Pointer to [**LifecycleRuleAndOperator**](LifecycleRuleAndOperator.md) |  | [optional] 

## Methods

### NewLifecycleRuleFilter

`func NewLifecycleRuleFilter() *LifecycleRuleFilter`

NewLifecycleRuleFilter instantiates a new LifecycleRuleFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleRuleFilterWithDefaults

`func NewLifecycleRuleFilterWithDefaults() *LifecycleRuleFilter`

NewLifecycleRuleFilterWithDefaults instantiates a new LifecycleRuleFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *LifecycleRuleFilter) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *LifecycleRuleFilter) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *LifecycleRuleFilter) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *LifecycleRuleFilter) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetTag

`func (o *LifecycleRuleFilter) GetTag() LifecycleRuleFilterTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *LifecycleRuleFilter) GetTagOk() (*LifecycleRuleFilterTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *LifecycleRuleFilter) SetTag(v LifecycleRuleFilterTag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *LifecycleRuleFilter) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetAnd

`func (o *LifecycleRuleFilter) GetAnd() LifecycleRuleAndOperator`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *LifecycleRuleFilter) GetAndOk() (*LifecycleRuleAndOperator, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *LifecycleRuleFilter) SetAnd(v LifecycleRuleAndOperator)`

SetAnd sets And field to given value.

### HasAnd

`func (o *LifecycleRuleFilter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


