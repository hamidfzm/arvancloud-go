# FeaturePricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeTier** | Pointer to **int32** |  | [optional] 
**Flat** | Pointer to [**NullableFeaturePrice**](FeaturePrice.md) |  | [optional] 
**PerUnit** | Pointer to [**NullableFeaturePrice**](FeaturePrice.md) |  | [optional] 

## Methods

### NewFeaturePricing

`func NewFeaturePricing() *FeaturePricing`

NewFeaturePricing instantiates a new FeaturePricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturePricingWithDefaults

`func NewFeaturePricingWithDefaults() *FeaturePricing`

NewFeaturePricingWithDefaults instantiates a new FeaturePricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeTier

`func (o *FeaturePricing) GetFreeTier() int32`

GetFreeTier returns the FreeTier field if non-nil, zero value otherwise.

### GetFreeTierOk

`func (o *FeaturePricing) GetFreeTierOk() (*int32, bool)`

GetFreeTierOk returns a tuple with the FreeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTier

`func (o *FeaturePricing) SetFreeTier(v int32)`

SetFreeTier sets FreeTier field to given value.

### HasFreeTier

`func (o *FeaturePricing) HasFreeTier() bool`

HasFreeTier returns a boolean if a field has been set.

### GetFlat

`func (o *FeaturePricing) GetFlat() FeaturePrice`

GetFlat returns the Flat field if non-nil, zero value otherwise.

### GetFlatOk

`func (o *FeaturePricing) GetFlatOk() (*FeaturePrice, bool)`

GetFlatOk returns a tuple with the Flat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlat

`func (o *FeaturePricing) SetFlat(v FeaturePrice)`

SetFlat sets Flat field to given value.

### HasFlat

`func (o *FeaturePricing) HasFlat() bool`

HasFlat returns a boolean if a field has been set.

### SetFlatNil

`func (o *FeaturePricing) SetFlatNil(b bool)`

 SetFlatNil sets the value for Flat to be an explicit nil

### UnsetFlat
`func (o *FeaturePricing) UnsetFlat()`

UnsetFlat ensures that no value is present for Flat, not even an explicit nil
### GetPerUnit

`func (o *FeaturePricing) GetPerUnit() FeaturePrice`

GetPerUnit returns the PerUnit field if non-nil, zero value otherwise.

### GetPerUnitOk

`func (o *FeaturePricing) GetPerUnitOk() (*FeaturePrice, bool)`

GetPerUnitOk returns a tuple with the PerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerUnit

`func (o *FeaturePricing) SetPerUnit(v FeaturePrice)`

SetPerUnit sets PerUnit field to given value.

### HasPerUnit

`func (o *FeaturePricing) HasPerUnit() bool`

HasPerUnit returns a boolean if a field has been set.

### SetPerUnitNil

`func (o *FeaturePricing) SetPerUnitNil(b bool)`

 SetPerUnitNil sets the value for PerUnit to be an explicit nil

### UnsetPerUnit
`func (o *FeaturePricing) UnsetPerUnit()`

UnsetPerUnit ensures that no value is present for PerUnit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


