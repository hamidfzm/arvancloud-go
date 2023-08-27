# StatusCodeReportChartsStatusCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**Series** | Pointer to [**[]StatusCodeReportChartsStatusCodeSeriesInner**](StatusCodeReportChartsStatusCodeSeriesInner.md) |  | [optional] 

## Methods

### NewStatusCodeReportChartsStatusCode

`func NewStatusCodeReportChartsStatusCode() *StatusCodeReportChartsStatusCode`

NewStatusCodeReportChartsStatusCode instantiates a new StatusCodeReportChartsStatusCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCodeReportChartsStatusCodeWithDefaults

`func NewStatusCodeReportChartsStatusCodeWithDefaults() *StatusCodeReportChartsStatusCode`

NewStatusCodeReportChartsStatusCodeWithDefaults instantiates a new StatusCodeReportChartsStatusCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StatusCodeReportChartsStatusCode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusCodeReportChartsStatusCode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusCodeReportChartsStatusCode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatusCodeReportChartsStatusCode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategories

`func (o *StatusCodeReportChartsStatusCode) GetCategories() []time.Time`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *StatusCodeReportChartsStatusCode) GetCategoriesOk() (*[]time.Time, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *StatusCodeReportChartsStatusCode) SetCategories(v []time.Time)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *StatusCodeReportChartsStatusCode) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSeries

`func (o *StatusCodeReportChartsStatusCode) GetSeries() []StatusCodeReportChartsStatusCodeSeriesInner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *StatusCodeReportChartsStatusCode) GetSeriesOk() (*[]StatusCodeReportChartsStatusCodeSeriesInner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *StatusCodeReportChartsStatusCode) SetSeries(v []StatusCodeReportChartsStatusCodeSeriesInner)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *StatusCodeReportChartsStatusCode) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


