# AttackReportMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statistics** | Pointer to [**[]AttackReportMapStatisticsInner**](AttackReportMapStatisticsInner.md) |  | [optional] 
**Charts** | Pointer to [**AttackReportMapCharts**](AttackReportMapCharts.md) |  | [optional] 
**Lists** | Pointer to [**[]AttackReportMapStatisticsInner**](AttackReportMapStatisticsInner.md) |  | [optional] 

## Methods

### NewAttackReportMap

`func NewAttackReportMap() *AttackReportMap`

NewAttackReportMap instantiates a new AttackReportMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackReportMapWithDefaults

`func NewAttackReportMapWithDefaults() *AttackReportMap`

NewAttackReportMapWithDefaults instantiates a new AttackReportMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatistics

`func (o *AttackReportMap) GetStatistics() []AttackReportMapStatisticsInner`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *AttackReportMap) GetStatisticsOk() (*[]AttackReportMapStatisticsInner, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *AttackReportMap) SetStatistics(v []AttackReportMapStatisticsInner)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *AttackReportMap) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetCharts

`func (o *AttackReportMap) GetCharts() AttackReportMapCharts`

GetCharts returns the Charts field if non-nil, zero value otherwise.

### GetChartsOk

`func (o *AttackReportMap) GetChartsOk() (*AttackReportMapCharts, bool)`

GetChartsOk returns a tuple with the Charts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharts

`func (o *AttackReportMap) SetCharts(v AttackReportMapCharts)`

SetCharts sets Charts field to given value.

### HasCharts

`func (o *AttackReportMap) HasCharts() bool`

HasCharts returns a boolean if a field has been set.

### GetLists

`func (o *AttackReportMap) GetLists() []AttackReportMapStatisticsInner`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *AttackReportMap) GetListsOk() (*[]AttackReportMapStatisticsInner, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *AttackReportMap) SetLists(v []AttackReportMapStatisticsInner)`

SetLists sets Lists field to given value.

### HasLists

`func (o *AttackReportMap) HasLists() bool`

HasLists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


