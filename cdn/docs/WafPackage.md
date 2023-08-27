# WafPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Provider** | Pointer to [**WafPackageProvider**](WafPackageProvider.md) |  | [optional] 
**ParamsSchema** | Pointer to **map[string]interface{}** | JSON-schema of parameters of the package | [optional] 
**DisabledRules** | Pointer to **[]string** | It will be filled by default disabled rules when it&#39;s not provided | [optional] 
**DisabledRulesets** | Pointer to **[]string** | It will be filled by default disabled rulesets when it&#39;s not provided | [optional] 

## Methods

### NewWafPackage

`func NewWafPackage() *WafPackage`

NewWafPackage instantiates a new WafPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafPackageWithDefaults

`func NewWafPackageWithDefaults() *WafPackage`

NewWafPackageWithDefaults instantiates a new WafPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafPackage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafPackage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafPackage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafPackage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WafPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *WafPackage) GetProvider() WafPackageProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *WafPackage) GetProviderOk() (*WafPackageProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *WafPackage) SetProvider(v WafPackageProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *WafPackage) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetParamsSchema

`func (o *WafPackage) GetParamsSchema() map[string]interface{}`

GetParamsSchema returns the ParamsSchema field if non-nil, zero value otherwise.

### GetParamsSchemaOk

`func (o *WafPackage) GetParamsSchemaOk() (*map[string]interface{}, bool)`

GetParamsSchemaOk returns a tuple with the ParamsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSchema

`func (o *WafPackage) SetParamsSchema(v map[string]interface{})`

SetParamsSchema sets ParamsSchema field to given value.

### HasParamsSchema

`func (o *WafPackage) HasParamsSchema() bool`

HasParamsSchema returns a boolean if a field has been set.

### GetDisabledRules

`func (o *WafPackage) GetDisabledRules() []string`

GetDisabledRules returns the DisabledRules field if non-nil, zero value otherwise.

### GetDisabledRulesOk

`func (o *WafPackage) GetDisabledRulesOk() (*[]string, bool)`

GetDisabledRulesOk returns a tuple with the DisabledRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledRules

`func (o *WafPackage) SetDisabledRules(v []string)`

SetDisabledRules sets DisabledRules field to given value.

### HasDisabledRules

`func (o *WafPackage) HasDisabledRules() bool`

HasDisabledRules returns a boolean if a field has been set.

### GetDisabledRulesets

`func (o *WafPackage) GetDisabledRulesets() []string`

GetDisabledRulesets returns the DisabledRulesets field if non-nil, zero value otherwise.

### GetDisabledRulesetsOk

`func (o *WafPackage) GetDisabledRulesetsOk() (*[]string, bool)`

GetDisabledRulesetsOk returns a tuple with the DisabledRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledRulesets

`func (o *WafPackage) SetDisabledRulesets(v []string)`

SetDisabledRulesets sets DisabledRulesets field to given value.

### HasDisabledRulesets

`func (o *WafPackage) HasDisabledRulesets() bool`

HasDisabledRulesets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


