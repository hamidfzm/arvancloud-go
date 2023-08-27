# FeatureUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureId** | Pointer to **string** |  | [optional] 
**Pricing** | Pointer to [**NullableFeaturePricing**](FeaturePricing.md) |  | [optional] 
**EstimatedCost** | Pointer to [**EstimatedCost**](EstimatedCost.md) |  | [optional] 
**Usage** | Pointer to **float32** |  | [optional] 

## Methods

### NewFeatureUsage

`func NewFeatureUsage() *FeatureUsage`

NewFeatureUsage instantiates a new FeatureUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureUsageWithDefaults

`func NewFeatureUsageWithDefaults() *FeatureUsage`

NewFeatureUsageWithDefaults instantiates a new FeatureUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureId

`func (o *FeatureUsage) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *FeatureUsage) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *FeatureUsage) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *FeatureUsage) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetPricing

`func (o *FeatureUsage) GetPricing() FeaturePricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *FeatureUsage) GetPricingOk() (*FeaturePricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *FeatureUsage) SetPricing(v FeaturePricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *FeatureUsage) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *FeatureUsage) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *FeatureUsage) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
### GetEstimatedCost

`func (o *FeatureUsage) GetEstimatedCost() EstimatedCost`

GetEstimatedCost returns the EstimatedCost field if non-nil, zero value otherwise.

### GetEstimatedCostOk

`func (o *FeatureUsage) GetEstimatedCostOk() (*EstimatedCost, bool)`

GetEstimatedCostOk returns a tuple with the EstimatedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCost

`func (o *FeatureUsage) SetEstimatedCost(v EstimatedCost)`

SetEstimatedCost sets EstimatedCost field to given value.

### HasEstimatedCost

`func (o *FeatureUsage) HasEstimatedCost() bool`

HasEstimatedCost returns a boolean if a field has been set.

### GetUsage

`func (o *FeatureUsage) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *FeatureUsage) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *FeatureUsage) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *FeatureUsage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


