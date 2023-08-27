# FirewallRuleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionDetails** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**FilterExpr** | Pointer to **string** | Wireshark-like filter expression | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 

## Methods

### NewFirewallRuleView

`func NewFirewallRuleView() *FirewallRuleView`

NewFirewallRuleView instantiates a new FirewallRuleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleViewWithDefaults

`func NewFirewallRuleViewWithDefaults() *FirewallRuleView`

NewFirewallRuleViewWithDefaults instantiates a new FirewallRuleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionDetails

`func (o *FirewallRuleView) GetActionDetails() map[string]interface{}`

GetActionDetails returns the ActionDetails field if non-nil, zero value otherwise.

### GetActionDetailsOk

`func (o *FirewallRuleView) GetActionDetailsOk() (*map[string]interface{}, bool)`

GetActionDetailsOk returns a tuple with the ActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDetails

`func (o *FirewallRuleView) SetActionDetails(v map[string]interface{})`

SetActionDetails sets ActionDetails field to given value.

### HasActionDetails

`func (o *FirewallRuleView) HasActionDetails() bool`

HasActionDetails returns a boolean if a field has been set.

### GetId

`func (o *FirewallRuleView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRuleView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRuleView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallRuleView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FirewallRuleView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallRuleView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallRuleView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirewallRuleView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilterExpr

`func (o *FirewallRuleView) GetFilterExpr() string`

GetFilterExpr returns the FilterExpr field if non-nil, zero value otherwise.

### GetFilterExprOk

`func (o *FirewallRuleView) GetFilterExprOk() (*string, bool)`

GetFilterExprOk returns a tuple with the FilterExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpr

`func (o *FirewallRuleView) SetFilterExpr(v string)`

SetFilterExpr sets FilterExpr field to given value.

### HasFilterExpr

`func (o *FirewallRuleView) HasFilterExpr() bool`

HasFilterExpr returns a boolean if a field has been set.

### GetAction

`func (o *FirewallRuleView) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallRuleView) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallRuleView) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *FirewallRuleView) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIsEnabled

`func (o *FirewallRuleView) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FirewallRuleView) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FirewallRuleView) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FirewallRuleView) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetNote

`func (o *FirewallRuleView) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *FirewallRuleView) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *FirewallRuleView) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *FirewallRuleView) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


