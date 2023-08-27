# StatusCodeReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to [**StatusCodeReportStatistics**](StatusCodeReportStatistics.md) |  | [optional] 
**Charts** | Pointer to [**StatusCodeReportCharts**](StatusCodeReportCharts.md) |  | [optional] 

## Methods

### NewStatusCodeReport

`func NewStatusCodeReport() *StatusCodeReport`

NewStatusCodeReport instantiates a new StatusCodeReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCodeReportWithDefaults

`func NewStatusCodeReportWithDefaults() *StatusCodeReport`

NewStatusCodeReportWithDefaults instantiates a new StatusCodeReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *StatusCodeReport) GetStatistics() StatusCodeReportStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *StatusCodeReport) GetStatisticsOk() (*StatusCodeReportStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *StatusCodeReport) SetStatistics(v StatusCodeReportStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *StatusCodeReport) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *StatusCodeReport) GetCharts() StatusCodeReportCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *StatusCodeReport) GetChartsOk() (*StatusCodeReportCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *StatusCodeReport) SetCharts(v StatusCodeReportCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *StatusCodeReport) HasCharts() bool`

HasCharts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


