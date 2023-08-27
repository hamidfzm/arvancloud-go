# V1ScopedResourceSelectorRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | **string** | Represents a scope&#39;s relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. | 
**ScopeName** | **string** | The name of the scope that the selector applies to. | 
**Values** | Pointer to **[]string** | An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch. | [optional] 

## Methods

### NewV1ScopedResourceSelectorRequirement

`func NewV1ScopedResourceSelectorRequirement(operator string, scopeName string, ) *V1ScopedResourceSelectorRequirement`

NewV1ScopedResourceSelectorRequirement instantiates a new V1ScopedResourceSelectorRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ScopedResourceSelectorRequirementWithDefaults

`func NewV1ScopedResourceSelectorRequirementWithDefaults() *V1ScopedResourceSelectorRequirement`

NewV1ScopedResourceSelectorRequirementWithDefaults instantiates a new V1ScopedResourceSelectorRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *V1ScopedResourceSelectorRequirement) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *V1ScopedResourceSelectorRequirement) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *V1ScopedResourceSelectorRequirement) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetScopeName

`func (o *V1ScopedResourceSelectorRequirement) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *V1ScopedResourceSelectorRequirement) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *V1ScopedResourceSelectorRequirement) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.


### GetValues

`func (o *V1ScopedResourceSelectorRequirement) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *V1ScopedResourceSelectorRequirement) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *V1ScopedResourceSelectorRequirement) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *V1ScopedResourceSelectorRequirement) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


