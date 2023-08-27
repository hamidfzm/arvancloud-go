# TrafficChartsRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**Series** | Pointer to [**[]TrafficChartsRequestsSeriesInner**](TrafficChartsRequestsSeriesInner.md) |  | [optional] 

## Methods

### NewTrafficChartsRequests

`func NewTrafficChartsRequests() *TrafficChartsRequests`

NewTrafficChartsRequests instantiates a new TrafficChartsRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficChartsRequestsWithDefaults

`func NewTrafficChartsRequestsWithDefaults() *TrafficChartsRequests`

NewTrafficChartsRequestsWithDefaults instantiates a new TrafficChartsRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TrafficChartsRequests) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TrafficChartsRequests) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TrafficChartsRequests) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TrafficChartsRequests) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCategories

`func (o *TrafficChartsRequests) GetCategories() []time.Time`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *TrafficChartsRequests) GetCategoriesOk() (*[]time.Time, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *TrafficChartsRequests) SetCategories(v []time.Time)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *TrafficChartsRequests) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSeries

`func (o *TrafficChartsRequests) GetSeries() []TrafficChartsRequestsSeriesInner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *TrafficChartsRequests) GetSeriesOk() (*[]TrafficChartsRequestsSeriesInner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *TrafficChartsRequests) SetSeries(v []TrafficChartsRequestsSeriesInner)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *TrafficChartsRequests) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


