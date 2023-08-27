# Visitors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to [**VisitorsStatistics**](VisitorsStatistics.md) |  | [optional] 
**Charts** | Pointer to [**VisitorsCharts**](VisitorsCharts.md) |  | [optional] 

## Methods

### NewVisitors

`func NewVisitors() *Visitors`

NewVisitors instantiates a new Visitors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisitorsWithDefaults

`func NewVisitorsWithDefaults() *Visitors`

NewVisitorsWithDefaults instantiates a new Visitors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *Visitors) GetStatistics() VisitorsStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *Visitors) GetStatisticsOk() (*VisitorsStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *Visitors) SetStatistics(v VisitorsStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *Visitors) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *Visitors) GetCharts() VisitorsCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *Visitors) GetChartsOk() (*VisitorsCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *Visitors) SetCharts(v VisitorsCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *Visitors) HasCharts() bool`

HasCharts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


