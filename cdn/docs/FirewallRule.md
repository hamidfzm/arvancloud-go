# FirewallRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionDetails** | Pointer to [**NullableFirewallActionDetails**](FirewallActionDetails.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**FilterExpr** | **string** | Wireshark-like filter expression | 
**Action** | **string** |  | 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 

## Methods

### NewFirewallRule

`func NewFirewallRule(name string, filterExpr string, action string, ) *FirewallRule`

NewFirewallRule instantiates a new FirewallRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleWithDefaults

`func NewFirewallRuleWithDefaults() *FirewallRule`

NewFirewallRuleWithDefaults instantiates a new FirewallRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionDetails

`func (o *FirewallRule) GetActionDetails() FirewallActionDetails`

GetActionDetails returns the ActionDetails field if non-nil, zero value otherwise.

### GetActionDetailsOk

`func (o *FirewallRule) GetActionDetailsOk() (*FirewallActionDetails, bool)`

GetActionDetailsOk returns a tuple with the ActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDetails

`func (o *FirewallRule) SetActionDetails(v FirewallActionDetails)`

SetActionDetails sets ActionDetails field to given value.

### HasActionDetails

`func (o *FirewallRule) HasActionDetails() bool`

HasActionDetails returns a boolean if a field has been set.

### SetActionDetailsNil

`func (o *FirewallRule) SetActionDetailsNil(b bool)`

 SetActionDetailsNil sets the value for ActionDetails to be an explicit nil

### UnsetActionDetails
`func (o *FirewallRule) UnsetActionDetails()`

UnsetActionDetails ensures that no value is present for ActionDetails, not even an explicit nil
### GetId

`func (o *FirewallRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FirewallRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirewallRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirewallRule) SetName(v string)`

SetName sets Name field to given value.


### GetFilterExpr

`func (o *FirewallRule) GetFilterExpr() string`

GetFilterExpr returns the FilterExpr field if non-nil, zero value otherwise.

### GetFilterExprOk

`func (o *FirewallRule) GetFilterExprOk() (*string, bool)`

GetFilterExprOk returns a tuple with the FilterExpr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpr

`func (o *FirewallRule) SetFilterExpr(v string)`

SetFilterExpr sets FilterExpr field to given value.


### GetAction

`func (o *FirewallRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FirewallRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FirewallRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetIsEnabled

`func (o *FirewallRule) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FirewallRule) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FirewallRule) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FirewallRule) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetNote

`func (o *FirewallRule) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *FirewallRule) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *FirewallRule) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *FirewallRule) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


