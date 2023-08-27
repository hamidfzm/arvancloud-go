# DomainStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Name of the domain | 
**DomainType** | Pointer to **string** | If you want to register a subdomain, you can use cname setup by sending partial type | [optional] [default to "full"]
**PlanLevel** | Pointer to **int32** | - &#x60;0&#x60; - Traffic - &#x60;1&#x60; - Basic - &#x60;2&#x60; - Growth - &#x60;3&#x60; - Professional - &#x60;4&#x60; - Enterprise  | [optional] 

## Methods

### NewDomainStore

`func NewDomainStore(domain string, ) *DomainStore`

NewDomainStore instantiates a new DomainStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainStoreWithDefaults

`func NewDomainStoreWithDefaults() *DomainStore`

NewDomainStoreWithDefaults instantiates a new DomainStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainStore) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainStore) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainStore) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDomainType

`func (o *DomainStore) GetDomainType() string`

GetDomainType returns the DomainType field if non-nil, zero value otherwise.

### GetDomainTypeOk

`func (o *DomainStore) GetDomainTypeOk() (*string, bool)`

GetDomainTypeOk returns a tuple with the DomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainType

`func (o *DomainStore) SetDomainType(v string)`

SetDomainType sets DomainType field to given value.

### HasDomainType

`func (o *DomainStore) HasDomainType() bool`

HasDomainType returns a boolean if a field has been set.

### GetPlanLevel

`func (o *DomainStore) GetPlanLevel() int32`

GetPlanLevel returns the PlanLevel field if non-nil, zero value otherwise.

### GetPlanLevelOk

`func (o *DomainStore) GetPlanLevelOk() (*int32, bool)`

GetPlanLevelOk returns a tuple with the PlanLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanLevel

`func (o *DomainStore) SetPlanLevel(v int32)`

SetPlanLevel sets PlanLevel field to given value.

### HasPlanLevel

`func (o *DomainStore) HasPlanLevel() bool`

HasPlanLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


