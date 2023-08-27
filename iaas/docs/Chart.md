# Chart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** |  | [optional] 
**Series** | Pointer to [**[]ChartDataset**](ChartDataset.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewChart

`func NewChart() *Chart`

NewChart instantiates a new Chart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartWithDefaults

`func NewChartWithDefaults() *Chart`

NewChartWithDefaults instantiates a new Chart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *Chart) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Chart) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Chart) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Chart) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSeries

`func (o *Chart) GetSeries() []ChartDataset`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Chart) GetSeriesOk() (*[]ChartDataset, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Chart) SetSeries(v []ChartDataset)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *Chart) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetTitle

`func (o *Chart) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Chart) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Chart) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Chart) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


