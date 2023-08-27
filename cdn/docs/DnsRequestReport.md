# DnsRequestReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to [**DnsRequestReportStatistics**](DnsRequestReportStatistics.md) |  | [optional] 
**Charts** | Pointer to [**DnsRequestReportCharts**](DnsRequestReportCharts.md) |  | [optional] 

## Methods

### NewDnsRequestReport

`func NewDnsRequestReport() *DnsRequestReport`

NewDnsRequestReport instantiates a new DnsRequestReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRequestReportWithDefaults

`func NewDnsRequestReportWithDefaults() *DnsRequestReport`

NewDnsRequestReportWithDefaults instantiates a new DnsRequestReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *DnsRequestReport) GetStatistics() DnsRequestReportStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *DnsRequestReport) GetStatisticsOk() (*DnsRequestReportStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *DnsRequestReport) SetStatistics(v DnsRequestReportStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *DnsRequestReport) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *DnsRequestReport) GetCharts() DnsRequestReportCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *DnsRequestReport) GetChartsOk() (*DnsRequestReportCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *DnsRequestReport) SetCharts(v DnsRequestReportCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *DnsRequestReport) HasCharts() bool`

HasCharts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


