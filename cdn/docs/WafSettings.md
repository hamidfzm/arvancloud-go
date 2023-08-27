# WafSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**Mode** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]DomainWafPackage**](DomainWafPackage.md) | Pacakges and their configurations that are used to configure WAF. | [optional] [readonly] 

## Methods

### NewWafSettings

`func NewWafSettings() *WafSettings`

NewWafSettings instantiates a new WafSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafSettingsWithDefaults

`func NewWafSettingsWithDefaults() *WafSettings`

NewWafSettingsWithDefaults instantiates a new WafSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *WafSettings) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *WafSettings) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *WafSettings) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *WafSettings) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMode

`func (o *WafSettings) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WafSettings) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WafSettings) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WafSettings) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPackages

`func (o *WafSettings) GetPackages() []DomainWafPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *WafSettings) GetPackagesOk() (*[]DomainWafPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *WafSettings) SetPackages(v []DomainWafPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *WafSettings) HasPackages() bool`

HasPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


