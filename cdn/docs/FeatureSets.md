# FeatureSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**Plans** | Pointer to [**[]PlanInfo**](PlanInfo.md) |  | [optional] 
**FeatureSets** | Pointer to [**[]FeatureSet**](FeatureSet.md) |  | [optional] 

## Methods

### NewFeatureSets

`func NewFeatureSets() *FeatureSets`

NewFeatureSets instantiates a new FeatureSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureSetsWithDefaults

`func NewFeatureSetsWithDefaults() *FeatureSets`

NewFeatureSetsWithDefaults instantiates a new FeatureSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *FeatureSets) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FeatureSets) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FeatureSets) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FeatureSets) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPlans

`func (o *FeatureSets) GetPlans() []PlanInfo`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *FeatureSets) GetPlansOk() (*[]PlanInfo, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *FeatureSets) SetPlans(v []PlanInfo)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *FeatureSets) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetFeatureSets

`func (o *FeatureSets) GetFeatureSets() []FeatureSet`

GetFeatureSets returns the FeatureSets field if non-nil, zero value otherwise.

### GetFeatureSetsOk

`func (o *FeatureSets) GetFeatureSetsOk() (*[]FeatureSet, bool)`

GetFeatureSetsOk returns a tuple with the FeatureSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSets

`func (o *FeatureSets) SetFeatureSets(v []FeatureSet)`

SetFeatureSets sets FeatureSets field to given value.

### HasFeatureSets

`func (o *FeatureSets) HasFeatureSets() bool`

HasFeatureSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


