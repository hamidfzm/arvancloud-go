# AdsAdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of ad | [optional] 
**Description** | Pointer to **string** | Description of ad | [optional] 
**AdType** | Pointer to **string** | Ad type | [optional] 
**PlayTime** | Pointer to **int32** | Specify mid-roll time in seconds | [optional] 
**SkipType** | Pointer to **string** | Skip type | [optional] 
**SkipOffset** | Pointer to **int32** | Skip offset in seconds (required if skip type is allow) | [optional] 
**ClickThrough** | Pointer to **string** | Click URL when user click on ad | [optional] 

## Methods

### NewAdsAdPutRequest

`func NewAdsAdPutRequest() *AdsAdPutRequest`

NewAdsAdPutRequest instantiates a new AdsAdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdsAdPutRequestWithDefaults

`func NewAdsAdPutRequestWithDefaults() *AdsAdPutRequest`

NewAdsAdPutRequestWithDefaults instantiates a new AdsAdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AdsAdPutRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdsAdPutRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdsAdPutRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdsAdPutRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *AdsAdPutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdsAdPutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdsAdPutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdsAdPutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdType

`func (o *AdsAdPutRequest) GetAdType() string`

GetAdType returns the AdType field if non-nil, zero value otherwise.

### GetAdTypeOk

`func (o *AdsAdPutRequest) GetAdTypeOk() (*string, bool)`

GetAdTypeOk returns a tuple with the AdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdType

`func (o *AdsAdPutRequest) SetAdType(v string)`

SetAdType sets AdType field to given value.

### HasAdType

`func (o *AdsAdPutRequest) HasAdType() bool`

HasAdType returns a boolean if a field has been set.

### GetPlayTime

`func (o *AdsAdPutRequest) GetPlayTime() int32`

GetPlayTime returns the PlayTime field if non-nil, zero value otherwise.

### GetPlayTimeOk

`func (o *AdsAdPutRequest) GetPlayTimeOk() (*int32, bool)`

GetPlayTimeOk returns a tuple with the PlayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayTime

`func (o *AdsAdPutRequest) SetPlayTime(v int32)`

SetPlayTime sets PlayTime field to given value.

### HasPlayTime

`func (o *AdsAdPutRequest) HasPlayTime() bool`

HasPlayTime returns a boolean if a field has been set.

### GetSkipType

`func (o *AdsAdPutRequest) GetSkipType() string`

GetSkipType returns the SkipType field if non-nil, zero value otherwise.

### GetSkipTypeOk

`func (o *AdsAdPutRequest) GetSkipTypeOk() (*string, bool)`

GetSkipTypeOk returns a tuple with the SkipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipType

`func (o *AdsAdPutRequest) SetSkipType(v string)`

SetSkipType sets SkipType field to given value.

### HasSkipType

`func (o *AdsAdPutRequest) HasSkipType() bool`

HasSkipType returns a boolean if a field has been set.

### GetSkipOffset

`func (o *AdsAdPutRequest) GetSkipOffset() int32`

GetSkipOffset returns the SkipOffset field if non-nil, zero value otherwise.

### GetSkipOffsetOk

`func (o *AdsAdPutRequest) GetSkipOffsetOk() (*int32, bool)`

GetSkipOffsetOk returns a tuple with the SkipOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOffset

`func (o *AdsAdPutRequest) SetSkipOffset(v int32)`

SetSkipOffset sets SkipOffset field to given value.

### HasSkipOffset

`func (o *AdsAdPutRequest) HasSkipOffset() bool`

HasSkipOffset returns a boolean if a field has been set.

### GetClickThrough

`func (o *AdsAdPutRequest) GetClickThrough() string`

GetClickThrough returns the ClickThrough field if non-nil, zero value otherwise.

### GetClickThroughOk

`func (o *AdsAdPutRequest) GetClickThroughOk() (*string, bool)`

GetClickThroughOk returns a tuple with the ClickThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickThrough

`func (o *AdsAdPutRequest) SetClickThrough(v string)`

SetClickThrough sets ClickThrough field to given value.

### HasClickThrough

`func (o *AdsAdPutRequest) HasClickThrough() bool`

HasClickThrough returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


