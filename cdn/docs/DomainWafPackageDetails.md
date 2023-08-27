# DomainWafPackageDetails

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

### NewDomainWafPackageDetails

`func NewDomainWafPackageDetails() *DomainWafPackageDetails`

NewDomainWafPackageDetails instantiates a new DomainWafPackageDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWafPackageDetailsWithDefaults

`func NewDomainWafPackageDetailsWithDefaults() *DomainWafPackageDetails`

NewDomainWafPackageDetailsWithDefaults instantiates a new DomainWafPackageDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainWafPackageDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainWafPackageDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainWafPackageDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DomainWafPackageDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DomainWafPackageDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainWafPackageDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainWafPackageDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainWafPackageDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *DomainWafPackageDetails) GetProvider() WafPackageProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DomainWafPackageDetails) GetProviderOk() (*WafPackageProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DomainWafPackageDetails) SetProvider(v WafPackageProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DomainWafPackageDetails) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetParamsSchema

`func (o *DomainWafPackageDetails) GetParamsSchema() map[string]interface{}`

GetParamsSchema returns the ParamsSchema field if non-nil, zero value otherwise.

### GetParamsSchemaOk

`func (o *DomainWafPackageDetails) GetParamsSchemaOk() (*map[string]interface{}, bool)`

GetParamsSchemaOk returns a tuple with the ParamsSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamsSchema

`func (o *DomainWafPackageDetails) SetParamsSchema(v map[string]interface{})`

SetParamsSchema sets ParamsSchema field to given value.

### HasParamsSchema

`func (o *DomainWafPackageDetails) HasParamsSchema() bool`

HasParamsSchema returns a boolean if a field has been set.

### GetDisabledRules

`func (o *DomainWafPackageDetails) GetDisabledRules() []string`

GetDisabledRules returns the DisabledRules field if non-nil, zero value otherwise.

### GetDisabledRulesOk

`func (o *DomainWafPackageDetails) GetDisabledRulesOk() (*[]string, bool)`

GetDisabledRulesOk returns a tuple with the DisabledRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledRules

`func (o *DomainWafPackageDetails) SetDisabledRules(v []string)`

SetDisabledRules sets DisabledRules field to given value.

### HasDisabledRules

`func (o *DomainWafPackageDetails) HasDisabledRules() bool`

HasDisabledRules returns a boolean if a field has been set.

### GetDisabledRulesets

`func (o *DomainWafPackageDetails) GetDisabledRulesets() []string`

GetDisabledRulesets returns the DisabledRulesets field if non-nil, zero value otherwise.

### GetDisabledRulesetsOk

`func (o *DomainWafPackageDetails) GetDisabledRulesetsOk() (*[]string, bool)`

GetDisabledRulesetsOk returns a tuple with the DisabledRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledRulesets

`func (o *DomainWafPackageDetails) SetDisabledRulesets(v []string)`

SetDisabledRulesets sets DisabledRulesets field to given value.

### HasDisabledRulesets

`func (o *DomainWafPackageDetails) HasDisabledRulesets() bool`

HasDisabledRulesets returns a boolean if a field has been set.

### GetRulesets

`func (o *DomainWafPackageDetails) GetRulesets() []WafRuleset`

GetRulesets returns the Rulesets field if non-nil, zero value otherwise.

### GetRulesetsOk

`func (o *DomainWafPackageDetails) GetRulesetsOk() (*[]WafRuleset, bool)`

GetRulesetsOk returns a tuple with the Rulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesets

`func (o *DomainWafPackageDetails) SetRulesets(v []WafRuleset)`

SetRulesets sets Rulesets field to given value.

### HasRulesets

`func (o *DomainWafPackageDetails) HasRulesets() bool`

HasRulesets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


