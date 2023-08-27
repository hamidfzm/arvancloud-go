# ErrorLogChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to [**ErrorLogChartStatistics**](ErrorLogChartStatistics.md) |  | [optional] 
**Charts** | Pointer to [**ErrorLogChartCharts**](ErrorLogChartCharts.md) |  | [optional] 

## Methods

### NewErrorLogChart

`func NewErrorLogChart() *ErrorLogChart`

NewErrorLogChart instantiates a new ErrorLogChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLogChartWithDefaults

`func NewErrorLogChartWithDefaults() *ErrorLogChart`

NewErrorLogChartWithDefaults instantiates a new ErrorLogChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *ErrorLogChart) GetStatistics() ErrorLogChartStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *ErrorLogChart) GetStatisticsOk() (*ErrorLogChartStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *ErrorLogChart) SetStatistics(v ErrorLogChartStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *ErrorLogChart) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *ErrorLogChart) GetCharts() ErrorLogChartCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *ErrorLogChart) GetChartsOk() (*ErrorLogChartCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *ErrorLogChart) SetCharts(v ErrorLogChartCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *ErrorLogChart) HasCharts() bool`

HasCharts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


