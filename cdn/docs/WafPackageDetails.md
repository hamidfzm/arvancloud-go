# WafPackageDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Provider** | Pointer to [**WafPackageProvider**](WafPackageProvider.md) |  | [optional] 
**ParamsSchema** | Pointer to **map[string]interface{}** | JSON-schema of parameters of the package | [optional] 
**DisabledRules** | Pointer to **[]string** | It will be filled by default disabled rules when it&#39;s not provided | [optional] 
**DisabledRulesets** | Pointer to **[]string** | It will be filled by default disabled rulesets when it&#39;s not provided | [optional] 
**Rulesets** | Pointer to [**[]WafRuleset**](WafRuleset.md) |  | [optional] 

## Methods

### NewWafPackageDetails

`func NewWafPackageDetails() *WafPackageDetails`

NewWafPackageDetails instantiates a new WafPackageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafPackageDetailsWithDefaults

`func NewWafPackageDetailsWithDefaults() *WafPackageDetails`

NewWafPackageDetailsWithDefaults instantiates a new WafPackageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafPackageDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafPackageDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafPackageDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafPackageDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WafPackageDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WafPackageDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WafPackageDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WafPackageDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *WafPackageDetails) GetProvider() WafPackageProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *WafPackageDetails) GetProviderOk() (*WafPackageProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *WafPackageDetails) SetProvider(v WafPackageProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *WafPackageDetails) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetParamsSchema

`func (o *WafPackageDetails) GetParamsSchema() map[string]interface{}`

GetParamsSchema returns the ParamsSchema field if non-nil, zero value otherwise.

### GetParamsSchemaOk

`func (o *WafPackageDetails) GetParamsSchemaOk() (*map[string]interface{}, bool)`

GetParamsSchemaOk returns a tuple with the ParamsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSchema

`func (o *WafPackageDetails) SetParamsSchema(v map[string]interface{})`

SetParamsSchema sets ParamsSchema field to given value.

### HasParamsSchema

`func (o *WafPackageDetails) HasParamsSchema() bool`

HasParamsSchema returns a boolean if a field has been set.

### GetDisabledRules

`func (o *WafPackageDetails) GetDisabledRules() []string`

GetDisabledRules returns the DisabledRules field if non-nil, zero value otherwise.

### GetDisabledRulesOk

`func (o *WafPackageDetails) GetDisabledRulesOk() (*[]string, bool)`

GetDisabledRulesOk returns a tuple with the DisabledRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledRules

`func (o *WafPackageDetails) SetDisabledRules(v []string)`

SetDisabledRules sets DisabledRules field to given value.

### HasDisabledRules

`func (o *WafPackageDetails) HasDisabledRules() bool`

HasDisabledRules returns a boolean if a field has been set.

### GetDisabledRulesets

`func (o *WafPackageDetails) GetDisabledRulesets() []string`

GetDisabledRulesets returns the DisabledRulesets field if non-nil, zero value otherwise.

### GetDisabledRulesetsOk

`func (o *WafPackageDetails) GetDisabledRulesetsOk() (*[]string, bool)`

GetDisabledRulesetsOk returns a tuple with the DisabledRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledRulesets

`func (o *WafPackageDetails) SetDisabledRulesets(v []string)`

SetDisabledRulesets sets DisabledRulesets field to given value.

### HasDisabledRulesets

`func (o *WafPackageDetails) HasDisabledRulesets() bool`

HasDisabledRulesets returns a boolean if a field has been set.

### GetRulesets

`func (o *WafPackageDetails) GetRulesets() []WafRuleset`

GetRulesets returns the Rulesets field if non-nil, zero value otherwise.

### GetRulesetsOk

`func (o *WafPackageDetails) GetRulesetsOk() (*[]WafRuleset, bool)`

GetRulesetsOk returns a tuple with the Rulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesets

`func (o *WafPackageDetails) SetRulesets(v []WafRuleset)`

SetRulesets sets Rulesets field to given value.

### HasRulesets

`func (o *WafPackageDetails) HasRulesets() bool`

HasRulesets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


