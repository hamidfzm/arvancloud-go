# DomainWafPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Params** | Pointer to **map[string]interface{}** | parameters of the package | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Provider** | Pointer to [**WafPackageProvider**](WafPackageProvider.md) |  | [optional] 
**ParamsSchema** | Pointer to **map[string]interface{}** | JSON-schema of parameters of the package | [optional] 
**DisabledRules** | Pointer to **[]string** | It will be filled by default disabled rules when it&#39;s not provided | [optional] 
**DisabledRulesets** | Pointer to **[]string** | It will be filled by default disabled rulesets when it&#39;s not provided | [optional] 

## Methods

### NewDomainWafPackage

`func NewDomainWafPackage() *DomainWafPackage`

NewDomainWafPackage instantiates a new DomainWafPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWafPackageWithDefaults

`func NewDomainWafPackageWithDefaults() *DomainWafPackage`

NewDomainWafPackageWithDefaults instantiates a new DomainWafPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParams

`func (o *DomainWafPackage) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *DomainWafPackage) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *DomainWafPackage) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *DomainWafPackage) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetIsEnabled

`func (o *DomainWafPackage) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DomainWafPackage) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DomainWafPackage) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DomainWafPackage) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetId

`func (o *DomainWafPackage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainWafPackage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainWafPackage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainWafPackage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DomainWafPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainWafPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainWafPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainWafPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *DomainWafPackage) GetProvider() WafPackageProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DomainWafPackage) GetProviderOk() (*WafPackageProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DomainWafPackage) SetProvider(v WafPackageProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DomainWafPackage) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetParamsSchema

`func (o *DomainWafPackage) GetParamsSchema() map[string]interface{}`

GetParamsSchema returns the ParamsSchema field if non-nil, zero value otherwise.

### GetParamsSchemaOk

`func (o *DomainWafPackage) GetParamsSchemaOk() (*map[string]interface{}, bool)`

GetParamsSchemaOk returns a tuple with the ParamsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSchema

`func (o *DomainWafPackage) SetParamsSchema(v map[string]interface{})`

SetParamsSchema sets ParamsSchema field to given value.

### HasParamsSchema

`func (o *DomainWafPackage) HasParamsSchema() bool`

HasParamsSchema returns a boolean if a field has been set.

### GetDisabledRules

`func (o *DomainWafPackage) GetDisabledRules() []string`

GetDisabledRules returns the DisabledRules field if non-nil, zero value otherwise.

### GetDisabledRulesOk

`func (o *DomainWafPackage) GetDisabledRulesOk() (*[]string, bool)`

GetDisabledRulesOk returns a tuple with the DisabledRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledRules

`func (o *DomainWafPackage) SetDisabledRules(v []string)`

SetDisabledRules sets DisabledRules field to given value.

### HasDisabledRules

`func (o *DomainWafPackage) HasDisabledRules() bool`

HasDisabledRules returns a boolean if a field has been set.

### GetDisabledRulesets

`func (o *DomainWafPackage) GetDisabledRulesets() []string`

GetDisabledRulesets returns the DisabledRulesets field if non-nil, zero value otherwise.

### GetDisabledRulesetsOk

`func (o *DomainWafPackage) GetDisabledRulesetsOk() (*[]string, bool)`

GetDisabledRulesetsOk returns a tuple with the DisabledRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledRulesets

`func (o *DomainWafPackage) SetDisabledRulesets(v []string)`

SetDisabledRulesets sets DisabledRulesets field to given value.

### HasDisabledRulesets

`func (o *DomainWafPackage) HasDisabledRulesets() bool`

HasDisabledRulesets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


