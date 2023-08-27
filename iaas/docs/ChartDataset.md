# ChartDataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]float64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewChartDataset

`func NewChartDataset() *ChartDataset`

NewChartDataset instantiates a new ChartDataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartDatasetWithDefaults

`func NewChartDatasetWithDefaults() *ChartDataset`

NewChartDatasetWithDefaults instantiates a new ChartDataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ChartDataset) GetData() []float64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChartDataset) GetDataOk() (*[]float64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChartDataset) SetData(v []float64)`

SetData sets Data field to given value.

### HasData

`func (o *ChartDataset) HasData() bool`

HasData returns a boolean if a field has been set.

### GetName

`func (o *ChartDataset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChartDataset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChartDataset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChartDataset) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


