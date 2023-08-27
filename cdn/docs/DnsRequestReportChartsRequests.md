# DnsRequestReportChartsRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**Series** | Pointer to [**[]DnsRequestReportChartsRequestsSeriesInner**](DnsRequestReportChartsRequestsSeriesInner.md) |  | [optional] 

## Methods

### NewDnsRequestReportChartsRequests

`func NewDnsRequestReportChartsRequests() *DnsRequestReportChartsRequests`

NewDnsRequestReportChartsRequests instantiates a new DnsRequestReportChartsRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRequestReportChartsRequestsWithDefaults

`func NewDnsRequestReportChartsRequestsWithDefaults() *DnsRequestReportChartsRequests`

NewDnsRequestReportChartsRequestsWithDefaults instantiates a new DnsRequestReportChartsRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DnsRequestReportChartsRequests) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DnsRequestReportChartsRequests) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DnsRequestReportChartsRequests) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DnsRequestReportChartsRequests) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCategories

`func (o *DnsRequestReportChartsRequests) GetCategories() []time.Time`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *DnsRequestReportChartsRequests) GetCategoriesOk() (*[]time.Time, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *DnsRequestReportChartsRequests) SetCategories(v []time.Time)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *DnsRequestReportChartsRequests) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSeries

`func (o *DnsRequestReportChartsRequests) GetSeries() []DnsRequestReportChartsRequestsSeriesInner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *DnsRequestReportChartsRequests) GetSeriesOk() (*[]DnsRequestReportChartsRequestsSeriesInner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *DnsRequestReportChartsRequests) SetSeries(v []DnsRequestReportChartsRequestsSeriesInner)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *DnsRequestReportChartsRequests) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


