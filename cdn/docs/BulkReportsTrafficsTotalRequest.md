# BulkReportsTrafficsTotalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to **[]string** | List of domains&#39; IDs | [optional] 
**ExcludeSubdomains** | Pointer to **bool** | Whether to include sub-domains or report only root domain traffic | [optional] 

## Methods

### NewBulkReportsTrafficsTotalRequest

`func NewBulkReportsTrafficsTotalRequest() *BulkReportsTrafficsTotalRequest`

NewBulkReportsTrafficsTotalRequest instantiates a new BulkReportsTrafficsTotalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkReportsTrafficsTotalRequestWithDefaults

`func NewBulkReportsTrafficsTotalRequestWithDefaults() *BulkReportsTrafficsTotalRequest`

NewBulkReportsTrafficsTotalRequestWithDefaults instantiates a new BulkReportsTrafficsTotalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *BulkReportsTrafficsTotalRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *BulkReportsTrafficsTotalRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *BulkReportsTrafficsTotalRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *BulkReportsTrafficsTotalRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetExcludeSubdomains

`func (o *BulkReportsTrafficsTotalRequest) GetExcludeSubdomains() bool`

GetExcludeSubdomains returns the ExcludeSubdomains field if non-nil, zero value otherwise.

### GetExcludeSubdomainsOk

`func (o *BulkReportsTrafficsTotalRequest) GetExcludeSubdomainsOk() (*bool, bool)`

GetExcludeSubdomainsOk returns a tuple with the ExcludeSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSubdomains

`func (o *BulkReportsTrafficsTotalRequest) SetExcludeSubdomains(v bool)`

SetExcludeSubdomains sets ExcludeSubdomains field to given value.

### HasExcludeSubdomains

`func (o *BulkReportsTrafficsTotalRequest) HasExcludeSubdomains() bool`

HasExcludeSubdomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


