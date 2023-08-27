# ResponseTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to [**ResponseTimeStatistics**](ResponseTimeStatistics.md) |  | [optional] 
**Charts** | Pointer to [**ResponseTimeCharts**](ResponseTimeCharts.md) |  | [optional] 

## Methods

### NewResponseTime

`func NewResponseTime() *ResponseTime`

NewResponseTime instantiates a new ResponseTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTimeWithDefaults

`func NewResponseTimeWithDefaults() *ResponseTime`

NewResponseTimeWithDefaults instantiates a new ResponseTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *ResponseTime) GetStatistics() ResponseTimeStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *ResponseTime) GetStatisticsOk() (*ResponseTimeStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *ResponseTime) SetStatistics(v ResponseTimeStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *ResponseTime) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *ResponseTime) GetCharts() ResponseTimeCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *ResponseTime) GetChartsOk() (*ResponseTimeCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *ResponseTime) SetCharts(v ResponseTimeCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *ResponseTime) HasCharts() bool`

HasCharts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


