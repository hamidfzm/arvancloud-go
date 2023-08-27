# FeatureDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Plans** | Pointer to [**FeatureDefinitionPlans**](FeatureDefinitionPlans.md) |  | [optional] 
**Meta** | Pointer to [**FeatureDefinitionMeta**](FeatureDefinitionMeta.md) |  | [optional] 

## Methods

### NewFeatureDefinition

`func NewFeatureDefinition() *FeatureDefinition`

NewFeatureDefinition instantiates a new FeatureDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureDefinitionWithDefaults

`func NewFeatureDefinitionWithDefaults() *FeatureDefinition`

NewFeatureDefinitionWithDefaults instantiates a new FeatureDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeatureDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeatureDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeatureDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeatureDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlans

`func (o *FeatureDefinition) GetPlans() FeatureDefinitionPlans`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *FeatureDefinition) GetPlansOk() (*FeatureDefinitionPlans, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *FeatureDefinition) SetPlans(v FeatureDefinitionPlans)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *FeatureDefinition) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### GetMeta

`func (o *FeatureDefinition) GetMeta() FeatureDefinitionMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FeatureDefinition) GetMetaOk() (*FeatureDefinitionMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FeatureDefinition) SetMeta(v FeatureDefinitionMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FeatureDefinition) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


