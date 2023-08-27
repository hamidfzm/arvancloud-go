# BaseFirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**FilterExpr** | Pointer to **string** | Wireshark-like filter expression | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 

## Methods

### NewBaseFirewallRule

`func NewBaseFirewallRule() *BaseFirewallRule`

NewBaseFirewallRule instantiates a new BaseFirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseFirewallRuleWithDefaults

`func NewBaseFirewallRuleWithDefaults() *BaseFirewallRule`

NewBaseFirewallRuleWithDefaults instantiates a new BaseFirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseFirewallRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseFirewallRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseFirewallRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseFirewallRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BaseFirewallRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseFirewallRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseFirewallRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseFirewallRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilterExpr

`func (o *BaseFirewallRule) GetFilterExpr() string`

GetFilterExpr returns the FilterExpr field if non-nil, zero value otherwise.

### GetFilterExprOk

`func (o *BaseFirewallRule) GetFilterExprOk() (*string, bool)`

GetFilterExprOk returns a tuple with the FilterExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpr

`func (o *BaseFirewallRule) SetFilterExpr(v string)`

SetFilterExpr sets FilterExpr field to given value.

### HasFilterExpr

`func (o *BaseFirewallRule) HasFilterExpr() bool`

HasFilterExpr returns a boolean if a field has been set.

### GetAction

`func (o *BaseFirewallRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BaseFirewallRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BaseFirewallRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BaseFirewallRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIsEnabled

`func (o *BaseFirewallRule) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *BaseFirewallRule) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *BaseFirewallRule) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *BaseFirewallRule) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetNote

`func (o *BaseFirewallRule) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BaseFirewallRule) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BaseFirewallRule) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BaseFirewallRule) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


