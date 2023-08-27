# AttackReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to [**AttackReportStatistics**](AttackReportStatistics.md) |  | [optional] 
**Charts** | Pointer to [**AttackReportCharts**](AttackReportCharts.md) |  | [optional] 

## Methods

### NewAttackReport

`func NewAttackReport() *AttackReport`

NewAttackReport instantiates a new AttackReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackReportWithDefaults

`func NewAttackReportWithDefaults() *AttackReport`

NewAttackReportWithDefaults instantiates a new AttackReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *AttackReport) GetStatistics() AttackReportStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *AttackReport) GetStatisticsOk() (*AttackReportStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *AttackReport) SetStatistics(v AttackReportStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *AttackReport) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *AttackReport) GetCharts() AttackReportCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *AttackReport) GetChartsOk() (*AttackReportCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *AttackReport) SetCharts(v AttackReportCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *AttackReport) HasCharts() bool`

HasCharts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


