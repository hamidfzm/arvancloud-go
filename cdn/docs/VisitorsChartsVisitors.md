# VisitorsChartsVisitors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**Series** | Pointer to [**[]VisitorsChartsVisitorsSeriesInner**](VisitorsChartsVisitorsSeriesInner.md) |  | [optional] 

## Methods

### NewVisitorsChartsVisitors

`func NewVisitorsChartsVisitors() *VisitorsChartsVisitors`

NewVisitorsChartsVisitors instantiates a new VisitorsChartsVisitors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisitorsChartsVisitorsWithDefaults

`func NewVisitorsChartsVisitorsWithDefaults() *VisitorsChartsVisitors`

NewVisitorsChartsVisitorsWithDefaults instantiates a new VisitorsChartsVisitors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *VisitorsChartsVisitors) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VisitorsChartsVisitors) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VisitorsChartsVisitors) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VisitorsChartsVisitors) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCategories

`func (o *VisitorsChartsVisitors) GetCategories() []time.Time`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *VisitorsChartsVisitors) GetCategoriesOk() (*[]time.Time, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *VisitorsChartsVisitors) SetCategories(v []time.Time)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *VisitorsChartsVisitors) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSeries

`func (o *VisitorsChartsVisitors) GetSeries() []VisitorsChartsVisitorsSeriesInner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *VisitorsChartsVisitors) GetSeriesOk() (*[]VisitorsChartsVisitorsSeriesInner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *VisitorsChartsVisitors) SetSeries(v []VisitorsChartsVisitorsSeriesInner)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *VisitorsChartsVisitors) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


