# ChannelsChannelVideosPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the video | 
**Description** | Pointer to **string** | Description of the video | [optional] 
**VideoUrl** | Pointer to **string** | Public URL of video | [optional] 
**FileId** | Pointer to **string** | ID of the video file | [optional] 
**ConvertMode** | **string** | Convert mode | 
**ProfileId** | Pointer to **string** | The profile ID that video convert with it (priority is with video properties) | [optional] 
**ParallelConvert** | Pointer to **bool** | Set this convert parallel when any video(s) is converting. Parallel limit is 3 | [optional] [default to false]
**ThumbnailTime** | Pointer to **int32** | Screenshot time for generate thumbnail for video in seconds | [optional] 
**WatermarkId** | Pointer to **string** | If you want to use watermark for a video, use this ID | [optional] 
**WatermarkArea** | Pointer to **string** | Area of the watermark if watermark_id presents | [optional] 
**ConvertInfo** | Pointer to [**[]ChannelsChannelProfilesPostRequestConvertInfoInner**](ChannelsChannelProfilesPostRequestConvertInfoInner.md) | Array of convert details | [optional] 
**Options** | Pointer to [**[]ChannelsChannelProfilesPostRequestOptionsInner**](ChannelsChannelProfilesPostRequestOptionsInner.md) | Array of option details | [optional] 

## Methods

### NewChannelsChannelVideosPostRequest

`func NewChannelsChannelVideosPostRequest(title string, convertMode string, ) *ChannelsChannelVideosPostRequest`

NewChannelsChannelVideosPostRequest instantiates a new ChannelsChannelVideosPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelsChannelVideosPostRequestWithDefaults

`func NewChannelsChannelVideosPostRequestWithDefaults() *ChannelsChannelVideosPostRequest`

NewChannelsChannelVideosPostRequestWithDefaults instantiates a new ChannelsChannelVideosPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ChannelsChannelVideosPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChannelsChannelVideosPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChannelsChannelVideosPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ChannelsChannelVideosPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelsChannelVideosPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelsChannelVideosPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelsChannelVideosPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVideoUrl

`func (o *ChannelsChannelVideosPostRequest) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *ChannelsChannelVideosPostRequest) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *ChannelsChannelVideosPostRequest) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *ChannelsChannelVideosPostRequest) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetFileId

`func (o *ChannelsChannelVideosPostRequest) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ChannelsChannelVideosPostRequest) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ChannelsChannelVideosPostRequest) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *ChannelsChannelVideosPostRequest) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetConvertMode

`func (o *ChannelsChannelVideosPostRequest) GetConvertMode() string`

GetConvertMode returns the ConvertMode field if non-nil, zero value otherwise.

### GetConvertModeOk

`func (o *ChannelsChannelVideosPostRequest) GetConvertModeOk() (*string, bool)`

GetConvertModeOk returns a tuple with the ConvertMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertMode

`func (o *ChannelsChannelVideosPostRequest) SetConvertMode(v string)`

SetConvertMode sets ConvertMode field to given value.


### GetProfileId

`func (o *ChannelsChannelVideosPostRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ChannelsChannelVideosPostRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ChannelsChannelVideosPostRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *ChannelsChannelVideosPostRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetParallelConvert

`func (o *ChannelsChannelVideosPostRequest) GetParallelConvert() bool`

GetParallelConvert returns the ParallelConvert field if non-nil, zero value otherwise.

### GetParallelConvertOk

`func (o *ChannelsChannelVideosPostRequest) GetParallelConvertOk() (*bool, bool)`

GetParallelConvertOk returns a tuple with the ParallelConvert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelConvert

`func (o *ChannelsChannelVideosPostRequest) SetParallelConvert(v bool)`

SetParallelConvert sets ParallelConvert field to given value.

### HasParallelConvert

`func (o *ChannelsChannelVideosPostRequest) HasParallelConvert() bool`

HasParallelConvert returns a boolean if a field has been set.

### GetThumbnailTime

`func (o *ChannelsChannelVideosPostRequest) GetThumbnailTime() int32`

GetThumbnailTime returns the ThumbnailTime field if non-nil, zero value otherwise.

### GetThumbnailTimeOk

`func (o *ChannelsChannelVideosPostRequest) GetThumbnailTimeOk() (*int32, bool)`

GetThumbnailTimeOk returns a tuple with the ThumbnailTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailTime

`func (o *ChannelsChannelVideosPostRequest) SetThumbnailTime(v int32)`

SetThumbnailTime sets ThumbnailTime field to given value.

### HasThumbnailTime

`func (o *ChannelsChannelVideosPostRequest) HasThumbnailTime() bool`

HasThumbnailTime returns a boolean if a field has been set.

### GetWatermarkId

`func (o *ChannelsChannelVideosPostRequest) GetWatermarkId() string`

GetWatermarkId returns the WatermarkId field if non-nil, zero value otherwise.

### GetWatermarkIdOk

`func (o *ChannelsChannelVideosPostRequest) GetWatermarkIdOk() (*string, bool)`

GetWatermarkIdOk returns a tuple with the WatermarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkId

`func (o *ChannelsChannelVideosPostRequest) SetWatermarkId(v string)`

SetWatermarkId sets WatermarkId field to given value.

### HasWatermarkId

`func (o *ChannelsChannelVideosPostRequest) HasWatermarkId() bool`

HasWatermarkId returns a boolean if a field has been set.

### GetWatermarkArea

`func (o *ChannelsChannelVideosPostRequest) GetWatermarkArea() string`

GetWatermarkArea returns the WatermarkArea field if non-nil, zero value otherwise.

### GetWatermarkAreaOk

`func (o *ChannelsChannelVideosPostRequest) GetWatermarkAreaOk() (*string, bool)`

GetWatermarkAreaOk returns a tuple with the WatermarkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkArea

`func (o *ChannelsChannelVideosPostRequest) SetWatermarkArea(v string)`

SetWatermarkArea sets WatermarkArea field to given value.

### HasWatermarkArea

`func (o *ChannelsChannelVideosPostRequest) HasWatermarkArea() bool`

HasWatermarkArea returns a boolean if a field has been set.

### GetConvertInfo

`func (o *ChannelsChannelVideosPostRequest) GetConvertInfo() []ChannelsChannelProfilesPostRequestConvertInfoInner`

GetConvertInfo returns the ConvertInfo field if non-nil, zero value otherwise.

### GetConvertInfoOk

`func (o *ChannelsChannelVideosPostRequest) GetConvertInfoOk() (*[]ChannelsChannelProfilesPostRequestConvertInfoInner, bool)`

GetConvertInfoOk returns a tuple with the ConvertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertInfo

`func (o *ChannelsChannelVideosPostRequest) SetConvertInfo(v []ChannelsChannelProfilesPostRequestConvertInfoInner)`

SetConvertInfo sets ConvertInfo field to given value.

### HasConvertInfo

`func (o *ChannelsChannelVideosPostRequest) HasConvertInfo() bool`

HasConvertInfo returns a boolean if a field has been set.

### GetOptions

`func (o *ChannelsChannelVideosPostRequest) GetOptions() []ChannelsChannelProfilesPostRequestOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ChannelsChannelVideosPostRequest) GetOptionsOk() (*[]ChannelsChannelProfilesPostRequestOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ChannelsChannelVideosPostRequest) SetOptions(v []ChannelsChannelProfilesPostRequestOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ChannelsChannelVideosPostRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


