# V1ResourceQuotaSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hard** | Pointer to **map[string]interface{}** | hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/ | [optional] 
**ScopeSelector** | Pointer to [**V1ScopeSelector**](V1ScopeSelector.md) |  | [optional] 
**Scopes** | Pointer to [**[]V1ResourceQuotaScope**](V1ResourceQuotaScope.md) | A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects. | [optional] 

## Methods

### NewV1ResourceQuotaSpec

`func NewV1ResourceQuotaSpec() *V1ResourceQuotaSpec`

NewV1ResourceQuotaSpec instantiates a new V1ResourceQuotaSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ResourceQuotaSpecWithDefaults

`func NewV1ResourceQuotaSpecWithDefaults() *V1ResourceQuotaSpec`

NewV1ResourceQuotaSpecWithDefaults instantiates a new V1ResourceQuotaSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHard

`func (o *V1ResourceQuotaSpec) GetHard() map[string]interface{}`

GetHard returns the Hard field if non-nil, zero value otherwise.

### GetHardOk

`func (o *V1ResourceQuotaSpec) GetHardOk() (*map[string]interface{}, bool)`

GetHardOk returns a tuple with the Hard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHard

`func (o *V1ResourceQuotaSpec) SetHard(v map[string]interface{})`

SetHard sets Hard field to given value.

### HasHard

`func (o *V1ResourceQuotaSpec) HasHard() bool`

HasHard returns a boolean if a field has been set.

### GetScopeSelector

`func (o *V1ResourceQuotaSpec) GetScopeSelector() V1ScopeSelector`

GetScopeSelector returns the ScopeSelector field if non-nil, zero value otherwise.

### GetScopeSelectorOk

`func (o *V1ResourceQuotaSpec) GetScopeSelectorOk() (*V1ScopeSelector, bool)`

GetScopeSelectorOk returns a tuple with the ScopeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSelector

`func (o *V1ResourceQuotaSpec) SetScopeSelector(v V1ScopeSelector)`

SetScopeSelector sets ScopeSelector field to given value.

### HasScopeSelector

`func (o *V1ResourceQuotaSpec) HasScopeSelector() bool`

HasScopeSelector returns a boolean if a field has been set.

### GetScopes

`func (o *V1ResourceQuotaSpec) GetScopes() []V1ResourceQuotaScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *V1ResourceQuotaSpec) GetScopesOk() (*[]V1ResourceQuotaScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *V1ResourceQuotaSpec) SetScopes(v []V1ResourceQuotaScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *V1ResourceQuotaSpec) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


