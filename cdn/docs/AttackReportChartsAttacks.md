# AttackReportChartsAttacks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**Series** | Pointer to [**[]AttackReportChartsAttacksSeriesInner**](AttackReportChartsAttacksSeriesInner.md) |  | [optional] 

## Methods

### NewAttackReportChartsAttacks

`func NewAttackReportChartsAttacks() *AttackReportChartsAttacks`

NewAttackReportChartsAttacks instantiates a new AttackReportChartsAttacks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackReportChartsAttacksWithDefaults

`func NewAttackReportChartsAttacksWithDefaults() *AttackReportChartsAttacks`

NewAttackReportChartsAttacksWithDefaults instantiates a new AttackReportChartsAttacks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AttackReportChartsAttacks) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttackReportChartsAttacks) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttackReportChartsAttacks) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AttackReportChartsAttacks) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCategories

`func (o *AttackReportChartsAttacks) GetCategories() []time.Time`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AttackReportChartsAttacks) GetCategoriesOk() (*[]time.Time, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AttackReportChartsAttacks) SetCategories(v []time.Time)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AttackReportChartsAttacks) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSeries

`func (o *AttackReportChartsAttacks) GetSeries() []AttackReportChartsAttacksSeriesInner`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *AttackReportChartsAttacks) GetSeriesOk() (*[]AttackReportChartsAttacksSeriesInner, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *AttackReportChartsAttacks) SetSeries(v []AttackReportChartsAttacksSeriesInner)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *AttackReportChartsAttacks) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


