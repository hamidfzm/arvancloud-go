# DeprecatedNs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NsDomain** | Pointer to **[]string** | Current NS records of the domain | [optional] 
**NsKeys** | Pointer to **[]string** | Desired NS values for the domain | [optional] 
**NsStatus** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeprecatedNs

`func NewDeprecatedNs() *DeprecatedNs`

NewDeprecatedNs instantiates a new DeprecatedNs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedNsWithDefaults

`func NewDeprecatedNsWithDefaults() *DeprecatedNs`

NewDeprecatedNsWithDefaults instantiates a new DeprecatedNs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNsDomain

`func (o *DeprecatedNs) GetNsDomain() []string`

GetNsDomain returns the NsDomain field if non-nil, zero value otherwise.

### GetNsDomainOk

`func (o *DeprecatedNs) GetNsDomainOk() (*[]string, bool)`

GetNsDomainOk returns a tuple with the NsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsDomain

`func (o *DeprecatedNs) SetNsDomain(v []string)`

SetNsDomain sets NsDomain field to given value.

### HasNsDomain

`func (o *DeprecatedNs) HasNsDomain() bool`

HasNsDomain returns a boolean if a field has been set.

### GetNsKeys

`func (o *DeprecatedNs) GetNsKeys() []string`

GetNsKeys returns the NsKeys field if non-nil, zero value otherwise.

### GetNsKeysOk

`func (o *DeprecatedNs) GetNsKeysOk() (*[]string, bool)`

GetNsKeysOk returns a tuple with the NsKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsKeys

`func (o *DeprecatedNs) SetNsKeys(v []string)`

SetNsKeys sets NsKeys field to given value.

### HasNsKeys

`func (o *DeprecatedNs) HasNsKeys() bool`

HasNsKeys returns a boolean if a field has been set.

### GetNsStatus

`func (o *DeprecatedNs) GetNsStatus() bool`

GetNsStatus returns the NsStatus field if non-nil, zero value otherwise.

### GetNsStatusOk

`func (o *DeprecatedNs) GetNsStatusOk() (*bool, bool)`

GetNsStatusOk returns a tuple with the NsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsStatus

`func (o *DeprecatedNs) SetNsStatus(v bool)`

SetNsStatus sets NsStatus field to given value.

### HasNsStatus

`func (o *DeprecatedNs) HasNsStatus() bool`

HasNsStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


