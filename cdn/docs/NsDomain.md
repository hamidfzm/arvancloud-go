# NsDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NsDomain** | Pointer to **[]string** | Current NS records of the domain | [optional] 
**NsKeys** | Pointer to **[]string** | Desired NS values for the domain | [optional] 

## Methods

### NewNsDomain

`func NewNsDomain() *NsDomain`

NewNsDomain instantiates a new NsDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsDomainWithDefaults

`func NewNsDomainWithDefaults() *NsDomain`

NewNsDomainWithDefaults instantiates a new NsDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNsDomain

`func (o *NsDomain) GetNsDomain() []string`

GetNsDomain returns the NsDomain field if non-nil, zero value otherwise.

### GetNsDomainOk

`func (o *NsDomain) GetNsDomainOk() (*[]string, bool)`

GetNsDomainOk returns a tuple with the NsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsDomain

`func (o *NsDomain) SetNsDomain(v []string)`

SetNsDomain sets NsDomain field to given value.

### HasNsDomain

`func (o *NsDomain) HasNsDomain() bool`

HasNsDomain returns a boolean if a field has been set.

### GetNsKeys

`func (o *NsDomain) GetNsKeys() []string`

GetNsKeys returns the NsKeys field if non-nil, zero value otherwise.

### GetNsKeysOk

`func (o *NsDomain) GetNsKeysOk() (*[]string, bool)`

GetNsKeysOk returns a tuple with the NsKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsKeys

`func (o *NsDomain) SetNsKeys(v []string)`

SetNsKeys sets NsKeys field to given value.

### HasNsKeys

`func (o *NsDomain) HasNsKeys() bool`

HasNsKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


