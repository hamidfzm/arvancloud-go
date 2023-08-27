# ErrorLogChartChartsStatusCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**Series** | Pointer to [**[]ErrorLogChartChartsStatusCodeSeriesInner**](ErrorLogChartChartsStatusCodeSeriesInner.md) |  | [optional] 

## Methods

### NewErrorLogChartChartsStatusCode

`func NewErrorLogChartChartsStatusCode() *ErrorLogChartChartsStatusCode`

NewErrorLogChartChartsStatusCode instantiates a new ErrorLogChartChartsStatusCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLogChartChartsStatusCodeWithDefaults

`func NewErrorLogChartChartsStatusCodeWithDefaults() *ErrorLogChartChartsStatusCode`

NewErrorLogChartChartsStatusCodeWithDefaults instantiates a new ErrorLogChartChartsStatusCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ErrorLogChartChartsStatusCode) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorLogChartChartsStatusCode) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorLogChartChartsStatusCode) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ErrorLogChartChartsStatusCode) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCategories

`func (o *ErrorLogChartChartsStatusCode) GetCategories() []time.Time`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ErrorLogChartChartsStatusCode) GetCategoriesOk() (*[]time.Time, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ErrorLogChartChartsStatusCode) SetCategories(v []time.Time)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ErrorLogChartChartsStatusCode) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSeries

`func (o *ErrorLogChartChartsStatusCode) GetSeries() []ErrorLogChartChartsStatusCodeSeriesInner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ErrorLogChartChartsStatusCode) GetSeriesOk() (*[]ErrorLogChartChartsStatusCodeSeriesInner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ErrorLogChartChartsStatusCode) SetSeries(v []ErrorLogChartChartsStatusCodeSeriesInner)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ErrorLogChartChartsStatusCode) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


