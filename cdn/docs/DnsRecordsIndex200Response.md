# DnsRecordsIndex200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DnsRecordGeneric**](DnsRecordGeneric.md) |  | [optional] 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginatedResponseMeta**](PaginatedResponseMeta.md) |  | [optional] 

## Methods

### NewDnsRecordsIndex200Response

`func NewDnsRecordsIndex200Response() *DnsRecordsIndex200Response`

NewDnsRecordsIndex200Response instantiates a new DnsRecordsIndex200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordsIndex200ResponseWithDefaults

`func NewDnsRecordsIndex200ResponseWithDefaults() *DnsRecordsIndex200Response`

NewDnsRecordsIndex200ResponseWithDefaults instantiates a new DnsRecordsIndex200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DnsRecordsIndex200Response) GetData() []DnsRecordGeneric`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DnsRecordsIndex200Response) GetDataOk() (*[]DnsRecordGeneric, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DnsRecordsIndex200Response) SetData(v []DnsRecordGeneric)`

SetData sets Data field to given value.

### HasData

`func (o *DnsRecordsIndex200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *DnsRecordsIndex200Response) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DnsRecordsIndex200Response) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DnsRecordsIndex200Response) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DnsRecordsIndex200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *DnsRecordsIndex200Response) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DnsRecordsIndex200Response) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DnsRecordsIndex200Response) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DnsRecordsIndex200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


