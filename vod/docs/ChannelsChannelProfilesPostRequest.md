# ChannelsChannelProfilesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the profile | 
**Description** | Pointer to **string** | Description of the profile | [optional] 
**ConvertMode** | **string** | Convert mode | 
**ThumbnailTime** | Pointer to **int32** | Screenshot time for generate thumbnail for video in seconds | [optional] 
**WatermarkId** | Pointer to **string** | If you want to use watermark for a video, use this ID | [optional] 
**WatermarkArea** | Pointer to **string** | Area of the watermark if watermark_id presents | [optional] 
**ConvertInfo** | Pointer to [**[]ChannelsChannelProfilesPostRequestConvertInfoInner**](ChannelsChannelProfilesPostRequestConvertInfoInner.md) | Array of convert details | [optional] 
**Options** | Pointer to [**[]ChannelsChannelProfilesPostRequestOptionsInner**](ChannelsChannelProfilesPostRequestOptionsInner.md) | Array of option details | [optional] 

## Methods

### NewChannelsChannelProfilesPostRequest

`func NewChannelsChannelProfilesPostRequest(title string, convertMode string, ) *ChannelsChannelProfilesPostRequest`

NewChannelsChannelProfilesPostRequest instantiates a new ChannelsChannelProfilesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelsChannelProfilesPostRequestWithDefaults

`func NewChannelsChannelProfilesPostRequestWithDefaults() *ChannelsChannelProfilesPostRequest`

NewChannelsChannelProfilesPostRequestWithDefaults instantiates a new ChannelsChannelProfilesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ChannelsChannelProfilesPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChannelsChannelProfilesPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChannelsChannelProfilesPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ChannelsChannelProfilesPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelsChannelProfilesPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelsChannelProfilesPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelsChannelProfilesPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConvertMode

`func (o *ChannelsChannelProfilesPostRequest) GetConvertMode() string`

GetConvertMode returns the ConvertMode field if non-nil, zero value otherwise.

### GetConvertModeOk

`func (o *ChannelsChannelProfilesPostRequest) GetConvertModeOk() (*string, bool)`

GetConvertModeOk returns a tuple with the ConvertMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertMode

`func (o *ChannelsChannelProfilesPostRequest) SetConvertMode(v string)`

SetConvertMode sets ConvertMode field to given value.


### GetThumbnailTime

`func (o *ChannelsChannelProfilesPostRequest) GetThumbnailTime() int32`

GetThumbnailTime returns the ThumbnailTime field if non-nil, zero value otherwise.

### GetThumbnailTimeOk

`func (o *ChannelsChannelProfilesPostRequest) GetThumbnailTimeOk() (*int32, bool)`

GetThumbnailTimeOk returns a tuple with the ThumbnailTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailTime

`func (o *ChannelsChannelProfilesPostRequest) SetThumbnailTime(v int32)`

SetThumbnailTime sets ThumbnailTime field to given value.

### HasThumbnailTime

`func (o *ChannelsChannelProfilesPostRequest) HasThumbnailTime() bool`

HasThumbnailTime returns a boolean if a field has been set.

### GetWatermarkId

`func (o *ChannelsChannelProfilesPostRequest) GetWatermarkId() string`

GetWatermarkId returns the WatermarkId field if non-nil, zero value otherwise.

### GetWatermarkIdOk

`func (o *ChannelsChannelProfilesPostRequest) GetWatermarkIdOk() (*string, bool)`

GetWatermarkIdOk returns a tuple with the WatermarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkId

`func (o *ChannelsChannelProfilesPostRequest) SetWatermarkId(v string)`

SetWatermarkId sets WatermarkId field to given value.

### HasWatermarkId

`func (o *ChannelsChannelProfilesPostRequest) HasWatermarkId() bool`

HasWatermarkId returns a boolean if a field has been set.

### GetWatermarkArea

`func (o *ChannelsChannelProfilesPostRequest) GetWatermarkArea() string`

GetWatermarkArea returns the WatermarkArea field if non-nil, zero value otherwise.

### GetWatermarkAreaOk

`func (o *ChannelsChannelProfilesPostRequest) GetWatermarkAreaOk() (*string, bool)`

GetWatermarkAreaOk returns a tuple with the WatermarkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkArea

`func (o *ChannelsChannelProfilesPostRequest) SetWatermarkArea(v string)`

SetWatermarkArea sets WatermarkArea field to given value.

### HasWatermarkArea

`func (o *ChannelsChannelProfilesPostRequest) HasWatermarkArea() bool`

HasWatermarkArea returns a boolean if a field has been set.

### GetConvertInfo

`func (o *ChannelsChannelProfilesPostRequest) GetConvertInfo() []ChannelsChannelProfilesPostRequestConvertInfoInner`

GetConvertInfo returns the ConvertInfo field if non-nil, zero value otherwise.

### GetConvertInfoOk

`func (o *ChannelsChannelProfilesPostRequest) GetConvertInfoOk() (*[]ChannelsChannelProfilesPostRequestConvertInfoInner, bool)`

GetConvertInfoOk returns a tuple with the ConvertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertInfo

`func (o *ChannelsChannelProfilesPostRequest) SetConvertInfo(v []ChannelsChannelProfilesPostRequestConvertInfoInner)`

SetConvertInfo sets ConvertInfo field to given value.

### HasConvertInfo

`func (o *ChannelsChannelProfilesPostRequest) HasConvertInfo() bool`

HasConvertInfo returns a boolean if a field has been set.

### GetOptions

`func (o *ChannelsChannelProfilesPostRequest) GetOptions() []ChannelsChannelProfilesPostRequestOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ChannelsChannelProfilesPostRequest) GetOptionsOk() (*[]ChannelsChannelProfilesPostRequestOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ChannelsChannelProfilesPostRequest) SetOptions(v []ChannelsChannelProfilesPostRequestOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ChannelsChannelProfilesPostRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


