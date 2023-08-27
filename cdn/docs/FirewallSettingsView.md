# FirewallSettingsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultActionDetails** | Pointer to **map[string]interface{}** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**DefaultAction** | Pointer to **string** |  | [optional] 
**VerifySni** | Pointer to **bool** | True to verify that SNI and hostname are equal | [optional] [default to false]

## Methods

### NewFirewallSettingsView

`func NewFirewallSettingsView() *FirewallSettingsView`

NewFirewallSettingsView instantiates a new FirewallSettingsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSettingsViewWithDefaults

`func NewFirewallSettingsViewWithDefaults() *FirewallSettingsView`

NewFirewallSettingsViewWithDefaults instantiates a new FirewallSettingsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultActionDetails

`func (o *FirewallSettingsView) GetDefaultActionDetails() map[string]interface{}`

GetDefaultActionDetails returns the DefaultActionDetails field if non-nil, zero value otherwise.

### GetDefaultActionDetailsOk

`func (o *FirewallSettingsView) GetDefaultActionDetailsOk() (*map[string]interface{}, bool)`

GetDefaultActionDetailsOk returns a tuple with the DefaultActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultActionDetails

`func (o *FirewallSettingsView) SetDefaultActionDetails(v map[string]interface{})`

SetDefaultActionDetails sets DefaultActionDetails field to given value.

### HasDefaultActionDetails

`func (o *FirewallSettingsView) HasDefaultActionDetails() bool`

HasDefaultActionDetails returns a boolean if a field has been set.

### SetDefaultActionDetailsNil

`func (o *FirewallSettingsView) SetDefaultActionDetailsNil(b bool)`

 SetDefaultActionDetailsNil sets the value for DefaultActionDetails to be an explicit nil

### UnsetDefaultActionDetails
`func (o *FirewallSettingsView) UnsetDefaultActionDetails()`

UnsetDefaultActionDetails ensures that no value is present for DefaultActionDetails, not even an explicit nil
### GetIsEnabled

`func (o *FirewallSettingsView) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *FirewallSettingsView) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *FirewallSettingsView) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *FirewallSettingsView) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDefaultAction

`func (o *FirewallSettingsView) GetDefaultAction() string`

GetDefaultAction returns the DefaultAction field if non-nil, zero value otherwise.

### GetDefaultActionOk

`func (o *FirewallSettingsView) GetDefaultActionOk() (*string, bool)`

GetDefaultActionOk returns a tuple with the DefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAction

`func (o *FirewallSettingsView) SetDefaultAction(v string)`

SetDefaultAction sets DefaultAction field to given value.

### HasDefaultAction

`func (o *FirewallSettingsView) HasDefaultAction() bool`

HasDefaultAction returns a boolean if a field has been set.

### GetVerifySni

`func (o *FirewallSettingsView) GetVerifySni() bool`

GetVerifySni returns the VerifySni field if non-nil, zero value otherwise.

### GetVerifySniOk

`func (o *FirewallSettingsView) GetVerifySniOk() (*bool, bool)`

GetVerifySniOk returns a tuple with the VerifySni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySni

`func (o *FirewallSettingsView) SetVerifySni(v bool)`

SetVerifySni sets VerifySni field to given value.

### HasVerifySni

`func (o *FirewallSettingsView) HasVerifySni() bool`

HasVerifySni returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


