# V1ScopeSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchExpressions** | Pointer to [**[]V1ScopedResourceSelectorRequirement**](V1ScopedResourceSelectorRequirement.md) | A list of scope selector requirements by scope of the resources. | [optional] 

## Methods

### NewV1ScopeSelector

`func NewV1ScopeSelector() *V1ScopeSelector`

NewV1ScopeSelector instantiates a new V1ScopeSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ScopeSelectorWithDefaults

`func NewV1ScopeSelectorWithDefaults() *V1ScopeSelector`

NewV1ScopeSelectorWithDefaults instantiates a new V1ScopeSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchExpressions

`func (o *V1ScopeSelector) GetMatchExpressions() []V1ScopedResourceSelectorRequirement`

GetMatchExpressions returns the MatchExpressions field if non-nil, zero value otherwise.

### GetMatchExpressionsOk

`func (o *V1ScopeSelector) GetMatchExpressionsOk() (*[]V1ScopedResourceSelectorRequirement, bool)`

GetMatchExpressionsOk returns a tuple with the MatchExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpressions

`func (o *V1ScopeSelector) SetMatchExpressions(v []V1ScopedResourceSelectorRequirement)`

SetMatchExpressions sets MatchExpressions field to given value.

### HasMatchExpressions

`func (o *V1ScopeSelector) HasMatchExpressions() bool`

HasMatchExpressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


