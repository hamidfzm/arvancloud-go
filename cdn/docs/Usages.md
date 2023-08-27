# Usages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureUsages** | Pointer to [**[]FeatureUsage**](FeatureUsage.md) |  | [optional] 
**EstimatedCost** | Pointer to [**EstimatedCost**](EstimatedCost.md) |  | [optional] 

## Methods

### NewUsages

`func NewUsages() *Usages`

NewUsages instantiates a new Usages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagesWithDefaults

`func NewUsagesWithDefaults() *Usages`

NewUsagesWithDefaults instantiates a new Usages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureUsages

`func (o *Usages) GetFeatureUsages() []FeatureUsage`

GetFeatureUsages returns the FeatureUsages field if non-nil, zero value otherwise.

### GetFeatureUsagesOk

`func (o *Usages) GetFeatureUsagesOk() (*[]FeatureUsage, bool)`

GetFeatureUsagesOk returns a tuple with the FeatureUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureUsages

`func (o *Usages) SetFeatureUsages(v []FeatureUsage)`

SetFeatureUsages sets FeatureUsages field to given value.

### HasFeatureUsages

`func (o *Usages) HasFeatureUsages() bool`

HasFeatureUsages returns a boolean if a field has been set.

### GetEstimatedCost

`func (o *Usages) GetEstimatedCost() EstimatedCost`

GetEstimatedCost returns the EstimatedCost field if non-nil, zero value otherwise.

### GetEstimatedCostOk

`func (o *Usages) GetEstimatedCostOk() (*EstimatedCost, bool)`

GetEstimatedCostOk returns a tuple with the EstimatedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedCost

`func (o *Usages) SetEstimatedCost(v EstimatedCost)`

SetEstimatedCost sets EstimatedCost field to given value.

### HasEstimatedCost

`func (o *Usages) HasEstimatedCost() bool`

HasEstimatedCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


