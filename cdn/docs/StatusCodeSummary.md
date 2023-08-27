# StatusCodeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to **map[string]interface{}** |  | [optional] 
**Charts** | Pointer to [**StatusCodeSummaryCharts**](StatusCodeSummaryCharts.md) |  | [optional] 

## Methods

### NewStatusCodeSummary

`func NewStatusCodeSummary() *StatusCodeSummary`

NewStatusCodeSummary instantiates a new StatusCodeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCodeSummaryWithDefaults

`func NewStatusCodeSummaryWithDefaults() *StatusCodeSummary`

NewStatusCodeSummaryWithDefaults instantiates a new StatusCodeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *StatusCodeSummary) GetStatistics() map[string]interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *StatusCodeSummary) GetStatisticsOk() (*map[string]interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *StatusCodeSummary) SetStatistics(v map[string]interface{})`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *StatusCodeSummary) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *StatusCodeSummary) GetCharts() StatusCodeSummaryCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *StatusCodeSummary) GetChartsOk() (*StatusCodeSummaryCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *StatusCodeSummary) SetCharts(v StatusCodeSummaryCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *StatusCodeSummary) HasCharts() bool`

HasCharts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


