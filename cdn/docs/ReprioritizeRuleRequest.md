# ReprioritizeRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | **string** | ID of the rule you want to move | 
**AfterRuleId** | Pointer to **string** | ID of the rule you want to be prior to the selected rule | [optional] 
**BeforeRuleId** | Pointer to **string** | ID of the rule you want to follow the selected rule | [optional] 

## Methods

### NewReprioritizeRuleRequest

`func NewReprioritizeRuleRequest(ruleId string, ) *ReprioritizeRuleRequest`

NewReprioritizeRuleRequest instantiates a new ReprioritizeRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReprioritizeRuleRequestWithDefaults

`func NewReprioritizeRuleRequestWithDefaults() *ReprioritizeRuleRequest`

NewReprioritizeRuleRequestWithDefaults instantiates a new ReprioritizeRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *ReprioritizeRuleRequest) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ReprioritizeRuleRequest) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ReprioritizeRuleRequest) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetAfterRuleId

`func (o *ReprioritizeRuleRequest) GetAfterRuleId() string`

GetAfterRuleId returns the AfterRuleId field if non-nil, zero value otherwise.

### GetAfterRuleIdOk

`func (o *ReprioritizeRuleRequest) GetAfterRuleIdOk() (*string, bool)`

GetAfterRuleIdOk returns a tuple with the AfterRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterRuleId

`func (o *ReprioritizeRuleRequest) SetAfterRuleId(v string)`

SetAfterRuleId sets AfterRuleId field to given value.

### HasAfterRuleId

`func (o *ReprioritizeRuleRequest) HasAfterRuleId() bool`

HasAfterRuleId returns a boolean if a field has been set.

### GetBeforeRuleId

`func (o *ReprioritizeRuleRequest) GetBeforeRuleId() string`

GetBeforeRuleId returns the BeforeRuleId field if non-nil, zero value otherwise.

### GetBeforeRuleIdOk

`func (o *ReprioritizeRuleRequest) GetBeforeRuleIdOk() (*string, bool)`

GetBeforeRuleIdOk returns a tuple with the BeforeRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeRuleId

`func (o *ReprioritizeRuleRequest) SetBeforeRuleId(v string)`

SetBeforeRuleId sets BeforeRuleId field to given value.

### HasBeforeRuleId

`func (o *ReprioritizeRuleRequest) HasBeforeRuleId() bool`

HasBeforeRuleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


