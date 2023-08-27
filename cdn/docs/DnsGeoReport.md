# DnsGeoReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to **[]interface{}** |  | [optional] 
**Charts** | Pointer to [**DnsGeoReportCharts**](DnsGeoReportCharts.md) |  | [optional] 
**Lists** | Pointer to [**[]DnsGeoReportListsInner**](DnsGeoReportListsInner.md) |  | [optional] 

## Methods

### NewDnsGeoReport

`func NewDnsGeoReport() *DnsGeoReport`

NewDnsGeoReport instantiates a new DnsGeoReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsGeoReportWithDefaults

`func NewDnsGeoReportWithDefaults() *DnsGeoReport`

NewDnsGeoReportWithDefaults instantiates a new DnsGeoReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *DnsGeoReport) GetStatistics() []interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *DnsGeoReport) GetStatisticsOk() (*[]interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *DnsGeoReport) SetStatistics(v []interface{})`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *DnsGeoReport) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *DnsGeoReport) GetCharts() DnsGeoReportCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *DnsGeoReport) GetChartsOk() (*DnsGeoReportCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *DnsGeoReport) SetCharts(v DnsGeoReportCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *DnsGeoReport) HasCharts() bool`

HasCharts returns a boolean if a field has been set.

### GetLists

`func (o *DnsGeoReport) GetLists() []DnsGeoReportListsInner`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *DnsGeoReport) GetListsOk() (*[]DnsGeoReportListsInner, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *DnsGeoReport) SetLists(v []DnsGeoReportListsInner)`

SetLists sets Lists field to given value.

### HasLists

`func (o *DnsGeoReport) HasLists() bool`

HasLists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


