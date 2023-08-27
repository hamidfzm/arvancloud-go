# BulkReportsVisitorsTotalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to **[]string** | List of domains&#39; IDs | [optional] 
**ExcludeSubdomains** | Pointer to **interface{}** | Whether to include sub-domains or report only root domain traffic | [optional] 

## Methods

### NewBulkReportsVisitorsTotalRequest

`func NewBulkReportsVisitorsTotalRequest() *BulkReportsVisitorsTotalRequest`

NewBulkReportsVisitorsTotalRequest instantiates a new BulkReportsVisitorsTotalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkReportsVisitorsTotalRequestWithDefaults

`func NewBulkReportsVisitorsTotalRequestWithDefaults() *BulkReportsVisitorsTotalRequest`

NewBulkReportsVisitorsTotalRequestWithDefaults instantiates a new BulkReportsVisitorsTotalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *BulkReportsVisitorsTotalRequest) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *BulkReportsVisitorsTotalRequest) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *BulkReportsVisitorsTotalRequest) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *BulkReportsVisitorsTotalRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetExcludeSubdomains

`func (o *BulkReportsVisitorsTotalRequest) GetExcludeSubdomains() interface{}`

GetExcludeSubdomains returns the ExcludeSubdomains field if non-nil, zero value otherwise.

### GetExcludeSubdomainsOk

`func (o *BulkReportsVisitorsTotalRequest) GetExcludeSubdomainsOk() (*interface{}, bool)`

GetExcludeSubdomainsOk returns a tuple with the ExcludeSubdomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSubdomains

`func (o *BulkReportsVisitorsTotalRequest) SetExcludeSubdomains(v interface{})`

SetExcludeSubdomains sets ExcludeSubdomains field to given value.

### HasExcludeSubdomains

`func (o *BulkReportsVisitorsTotalRequest) HasExcludeSubdomains() bool`

HasExcludeSubdomains returns a boolean if a field has been set.

### SetExcludeSubdomainsNil

`func (o *BulkReportsVisitorsTotalRequest) SetExcludeSubdomainsNil(b bool)`

 SetExcludeSubdomainsNil sets the value for ExcludeSubdomains to be an explicit nil

### UnsetExcludeSubdomains
`func (o *BulkReportsVisitorsTotalRequest) UnsetExcludeSubdomains()`

UnsetExcludeSubdomains ensures that no value is present for ExcludeSubdomains, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


