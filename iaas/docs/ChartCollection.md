# ChartCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**Chart**](Chart.md) |  | [optional] 
**Disk** | Pointer to [**Chart**](Chart.md) |  | [optional] 
**Network** | Pointer to [**Chart**](Chart.md) |  | [optional] 
**Ram** | Pointer to [**Chart**](Chart.md) |  | [optional] 
**Statistics** | Pointer to **[]string** |  | [optional] 

## Methods

### NewChartCollection

`func NewChartCollection() *ChartCollection`

NewChartCollection instantiates a new ChartCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartCollectionWithDefaults

`func NewChartCollectionWithDefaults() *ChartCollection`

NewChartCollectionWithDefaults instantiates a new ChartCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ChartCollection) GetCpu() Chart`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ChartCollection) GetCpuOk() (*Chart, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ChartCollection) SetCpu(v Chart)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ChartCollection) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDisk

`func (o *ChartCollection) GetDisk() Chart`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *ChartCollection) GetDiskOk() (*Chart, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *ChartCollection) SetDisk(v Chart)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *ChartCollection) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetNetwork

`func (o *ChartCollection) GetNetwork() Chart`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ChartCollection) GetNetworkOk() (*Chart, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ChartCollection) SetNetwork(v Chart)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ChartCollection) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRam

`func (o *ChartCollection) GetRam() Chart`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ChartCollection) GetRamOk() (*Chart, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ChartCollection) SetRam(v Chart)`

SetRam sets Ram field to given value.

### HasRam

`func (o *ChartCollection) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetStatistics

`func (o *ChartCollection) GetStatistics() []string`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *ChartCollection) GetStatisticsOk() (*[]string, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *ChartCollection) SetStatistics(v []string)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *ChartCollection) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


