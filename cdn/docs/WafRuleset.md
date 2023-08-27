# WafRuleset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]WafRulesetRulesInner**](WafRulesetRulesInner.md) |  | [optional] 

## Methods

### NewWafRuleset

`func NewWafRuleset() *WafRuleset`

NewWafRuleset instantiates a new WafRuleset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRulesetWithDefaults

`func NewWafRulesetWithDefaults() *WafRuleset`

NewWafRulesetWithDefaults instantiates a new WafRuleset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafRuleset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafRuleset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafRuleset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafRuleset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WafRuleset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafRuleset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafRuleset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafRuleset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *WafRuleset) GetRules() []WafRulesetRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *WafRuleset) GetRulesOk() (*[]WafRulesetRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *WafRuleset) SetRules(v []WafRulesetRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *WafRuleset) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


