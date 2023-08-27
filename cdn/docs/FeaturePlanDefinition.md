# FeaturePlanDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**FeaturePlanDefinitionMeta**](FeaturePlanDefinitionMeta.md) |  | [optional] 
**UsageLimit** | Pointer to [**NullableUsageLimit**](UsageLimit.md) |  | [optional] 
**Pricing** | Pointer to [**NullableFeaturePricing**](FeaturePricing.md) |  | [optional] 

## Methods

### NewFeaturePlanDefinition

`func NewFeaturePlanDefinition() *FeaturePlanDefinition`

NewFeaturePlanDefinition instantiates a new FeaturePlanDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturePlanDefinitionWithDefaults

`func NewFeaturePlanDefinitionWithDefaults() *FeaturePlanDefinition`

NewFeaturePlanDefinitionWithDefaults instantiates a new FeaturePlanDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *FeaturePlanDefinition) GetMeta() FeaturePlanDefinitionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FeaturePlanDefinition) GetMetaOk() (*FeaturePlanDefinitionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FeaturePlanDefinition) SetMeta(v FeaturePlanDefinitionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FeaturePlanDefinition) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUsageLimit

`func (o *FeaturePlanDefinition) GetUsageLimit() UsageLimit`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *FeaturePlanDefinition) GetUsageLimitOk() (*UsageLimit, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *FeaturePlanDefinition) SetUsageLimit(v UsageLimit)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *FeaturePlanDefinition) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimitNil

`func (o *FeaturePlanDefinition) SetUsageLimitNil(b bool)`

 SetUsageLimitNil sets the value for UsageLimit to be an explicit nil

### UnsetUsageLimit
`func (o *FeaturePlanDefinition) UnsetUsageLimit()`

UnsetUsageLimit ensures that no value is present for UsageLimit, not even an explicit nil
### GetPricing

`func (o *FeaturePlanDefinition) GetPricing() FeaturePricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *FeaturePlanDefinition) GetPricingOk() (*FeaturePricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *FeaturePlanDefinition) SetPricing(v FeaturePricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *FeaturePlanDefinition) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *FeaturePlanDefinition) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *FeaturePlanDefinition) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


