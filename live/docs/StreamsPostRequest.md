# StreamsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the stream | 
**Description** | Pointer to **string** | Description of the stream | [optional] 
**Type** | **string** | Stream latency type | 
**Mode** | **string** | Way to receive input stream | 
**InputUrl** | Pointer to **string** | Public URL of stream input in PULL mode | [optional] 
**Slug** | **string** | Unique slug for your stream. Only contain lower case letters and digits | 
**Timeshift** | Pointer to **int32** | Timeshift (DVR) in minutes to watch the earlier content | [optional] 
**FpsMode** | Pointer to **string** | stream fps mode | [optional] 
**Fps** | **int32** | Input steam frame per second | 
**ArchiveEnabled** | Pointer to **bool** | Set this value to True to archive the stream as a VoD | [optional] [default to false]
**CatchupEnabled** | Pointer to **bool** | Set this value to True to enable catchup for the stream | [optional] [default to false]
**CatchupPeriod** | Pointer to **int32** | Hours of catchup must be available for the stream | [optional] 
**ArchiveMode** | Pointer to **string** | Way to archive stream | [optional] 
**ChannelId** | Pointer to **string** | Channel Id to save archive | [optional] 
**WatermarkId** | Pointer to **string** | If you want to use watermark for a video, use this ID | [optional] 
**WatermarkArea** | Pointer to **string** | Area of the watermark if watermark_id presents | [optional] 
**ConvertInfo** | [**[]StreamsPostRequestConvertInfoInner**](StreamsPostRequestConvertInfoInner.md) | Array of convert details | 
**SecureLinkEnabled** | Pointer to **bool** | Enable or disable secure link for all videos in channel | [optional] 
**SecureLinkKey** | Pointer to **string** | Key for generate secure links | [optional] 
**SecureLinkWithIp** | Pointer to **bool** | IP can be considered as an optional parameter | [optional] 
**AdsEnabled** | Pointer to **bool** | Enable or disable Ads for created stream | [optional] 
**PresentType** | Pointer to **string** | Ads present method | [optional] 
**CampaignId** | Pointer to **string** | Created CampaignId in Ads | [optional] 

## Methods

### NewStreamsPostRequest

`func NewStreamsPostRequest(title string, type_ string, mode string, slug string, fps int32, convertInfo []StreamsPostRequestConvertInfoInner, ) *StreamsPostRequest`

NewStreamsPostRequest instantiates a new StreamsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsPostRequestWithDefaults

`func NewStreamsPostRequestWithDefaults() *StreamsPostRequest`

NewStreamsPostRequestWithDefaults instantiates a new StreamsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *StreamsPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StreamsPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StreamsPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *StreamsPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamsPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamsPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StreamsPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *StreamsPostRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamsPostRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamsPostRequest) SetType(v string)`

SetType sets Type field to given value.


### GetMode

`func (o *StreamsPostRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *StreamsPostRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *StreamsPostRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetInputUrl

`func (o *StreamsPostRequest) GetInputUrl() string`

GetInputUrl returns the InputUrl field if non-nil, zero value otherwise.

### GetInputUrlOk

`func (o *StreamsPostRequest) GetInputUrlOk() (*string, bool)`

GetInputUrlOk returns a tuple with the InputUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputUrl

`func (o *StreamsPostRequest) SetInputUrl(v string)`

SetInputUrl sets InputUrl field to given value.

### HasInputUrl

`func (o *StreamsPostRequest) HasInputUrl() bool`

HasInputUrl returns a boolean if a field has been set.

### GetSlug

`func (o *StreamsPostRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StreamsPostRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StreamsPostRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTimeshift

`func (o *StreamsPostRequest) GetTimeshift() int32`

GetTimeshift returns the Timeshift field if non-nil, zero value otherwise.

### GetTimeshiftOk

`func (o *StreamsPostRequest) GetTimeshiftOk() (*int32, bool)`

GetTimeshiftOk returns a tuple with the Timeshift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeshift

`func (o *StreamsPostRequest) SetTimeshift(v int32)`

SetTimeshift sets Timeshift field to given value.

### HasTimeshift

`func (o *StreamsPostRequest) HasTimeshift() bool`

HasTimeshift returns a boolean if a field has been set.

### GetFpsMode

`func (o *StreamsPostRequest) GetFpsMode() string`

GetFpsMode returns the FpsMode field if non-nil, zero value otherwise.

### GetFpsModeOk

`func (o *StreamsPostRequest) GetFpsModeOk() (*string, bool)`

GetFpsModeOk returns a tuple with the FpsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpsMode

`func (o *StreamsPostRequest) SetFpsMode(v string)`

SetFpsMode sets FpsMode field to given value.

### HasFpsMode

`func (o *StreamsPostRequest) HasFpsMode() bool`

HasFpsMode returns a boolean if a field has been set.

### GetFps

`func (o *StreamsPostRequest) GetFps() int32`

GetFps returns the Fps field if non-nil, zero value otherwise.

### GetFpsOk

`func (o *StreamsPostRequest) GetFpsOk() (*int32, bool)`

GetFpsOk returns a tuple with the Fps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFps

`func (o *StreamsPostRequest) SetFps(v int32)`

SetFps sets Fps field to given value.


### GetArchiveEnabled

`func (o *StreamsPostRequest) GetArchiveEnabled() bool`

GetArchiveEnabled returns the ArchiveEnabled field if non-nil, zero value otherwise.

### GetArchiveEnabledOk

`func (o *StreamsPostRequest) GetArchiveEnabledOk() (*bool, bool)`

GetArchiveEnabledOk returns a tuple with the ArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveEnabled

`func (o *StreamsPostRequest) SetArchiveEnabled(v bool)`

SetArchiveEnabled sets ArchiveEnabled field to given value.

### HasArchiveEnabled

`func (o *StreamsPostRequest) HasArchiveEnabled() bool`

HasArchiveEnabled returns a boolean if a field has been set.

### GetCatchupEnabled

`func (o *StreamsPostRequest) GetCatchupEnabled() bool`

GetCatchupEnabled returns the CatchupEnabled field if non-nil, zero value otherwise.

### GetCatchupEnabledOk

`func (o *StreamsPostRequest) GetCatchupEnabledOk() (*bool, bool)`

GetCatchupEnabledOk returns a tuple with the CatchupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchupEnabled

`func (o *StreamsPostRequest) SetCatchupEnabled(v bool)`

SetCatchupEnabled sets CatchupEnabled field to given value.

### HasCatchupEnabled

`func (o *StreamsPostRequest) HasCatchupEnabled() bool`

HasCatchupEnabled returns a boolean if a field has been set.

### GetCatchupPeriod

`func (o *StreamsPostRequest) GetCatchupPeriod() int32`

GetCatchupPeriod returns the CatchupPeriod field if non-nil, zero value otherwise.

### GetCatchupPeriodOk

`func (o *StreamsPostRequest) GetCatchupPeriodOk() (*int32, bool)`

GetCatchupPeriodOk returns a tuple with the CatchupPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchupPeriod

`func (o *StreamsPostRequest) SetCatchupPeriod(v int32)`

SetCatchupPeriod sets CatchupPeriod field to given value.

### HasCatchupPeriod

`func (o *StreamsPostRequest) HasCatchupPeriod() bool`

HasCatchupPeriod returns a boolean if a field has been set.

### GetArchiveMode

`func (o *StreamsPostRequest) GetArchiveMode() string`

GetArchiveMode returns the ArchiveMode field if non-nil, zero value otherwise.

### GetArchiveModeOk

`func (o *StreamsPostRequest) GetArchiveModeOk() (*string, bool)`

GetArchiveModeOk returns a tuple with the ArchiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveMode

`func (o *StreamsPostRequest) SetArchiveMode(v string)`

SetArchiveMode sets ArchiveMode field to given value.

### HasArchiveMode

`func (o *StreamsPostRequest) HasArchiveMode() bool`

HasArchiveMode returns a boolean if a field has been set.

### GetChannelId

`func (o *StreamsPostRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *StreamsPostRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *StreamsPostRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *StreamsPostRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetWatermarkId

`func (o *StreamsPostRequest) GetWatermarkId() string`

GetWatermarkId returns the WatermarkId field if non-nil, zero value otherwise.

### GetWatermarkIdOk

`func (o *StreamsPostRequest) GetWatermarkIdOk() (*string, bool)`

GetWatermarkIdOk returns a tuple with the WatermarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkId

`func (o *StreamsPostRequest) SetWatermarkId(v string)`

SetWatermarkId sets WatermarkId field to given value.

### HasWatermarkId

`func (o *StreamsPostRequest) HasWatermarkId() bool`

HasWatermarkId returns a boolean if a field has been set.

### GetWatermarkArea

`func (o *StreamsPostRequest) GetWatermarkArea() string`

GetWatermarkArea returns the WatermarkArea field if non-nil, zero value otherwise.

### GetWatermarkAreaOk

`func (o *StreamsPostRequest) GetWatermarkAreaOk() (*string, bool)`

GetWatermarkAreaOk returns a tuple with the WatermarkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkArea

`func (o *StreamsPostRequest) SetWatermarkArea(v string)`

SetWatermarkArea sets WatermarkArea field to given value.

### HasWatermarkArea

`func (o *StreamsPostRequest) HasWatermarkArea() bool`

HasWatermarkArea returns a boolean if a field has been set.

### GetConvertInfo

`func (o *StreamsPostRequest) GetConvertInfo() []StreamsPostRequestConvertInfoInner`

GetConvertInfo returns the ConvertInfo field if non-nil, zero value otherwise.

### GetConvertInfoOk

`func (o *StreamsPostRequest) GetConvertInfoOk() (*[]StreamsPostRequestConvertInfoInner, bool)`

GetConvertInfoOk returns a tuple with the ConvertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertInfo

`func (o *StreamsPostRequest) SetConvertInfo(v []StreamsPostRequestConvertInfoInner)`

SetConvertInfo sets ConvertInfo field to given value.


### GetSecureLinkEnabled

`func (o *StreamsPostRequest) GetSecureLinkEnabled() bool`

GetSecureLinkEnabled returns the SecureLinkEnabled field if non-nil, zero value otherwise.

### GetSecureLinkEnabledOk

`func (o *StreamsPostRequest) GetSecureLinkEnabledOk() (*bool, bool)`

GetSecureLinkEnabledOk returns a tuple with the SecureLinkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkEnabled

`func (o *StreamsPostRequest) SetSecureLinkEnabled(v bool)`

SetSecureLinkEnabled sets SecureLinkEnabled field to given value.

### HasSecureLinkEnabled

`func (o *StreamsPostRequest) HasSecureLinkEnabled() bool`

HasSecureLinkEnabled returns a boolean if a field has been set.

### GetSecureLinkKey

`func (o *StreamsPostRequest) GetSecureLinkKey() string`

GetSecureLinkKey returns the SecureLinkKey field if non-nil, zero value otherwise.

### GetSecureLinkKeyOk

`func (o *StreamsPostRequest) GetSecureLinkKeyOk() (*string, bool)`

GetSecureLinkKeyOk returns a tuple with the SecureLinkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkKey

`func (o *StreamsPostRequest) SetSecureLinkKey(v string)`

SetSecureLinkKey sets SecureLinkKey field to given value.

### HasSecureLinkKey

`func (o *StreamsPostRequest) HasSecureLinkKey() bool`

HasSecureLinkKey returns a boolean if a field has been set.

### GetSecureLinkWithIp

`func (o *StreamsPostRequest) GetSecureLinkWithIp() bool`

GetSecureLinkWithIp returns the SecureLinkWithIp field if non-nil, zero value otherwise.

### GetSecureLinkWithIpOk

`func (o *StreamsPostRequest) GetSecureLinkWithIpOk() (*bool, bool)`

GetSecureLinkWithIpOk returns a tuple with the SecureLinkWithIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkWithIp

`func (o *StreamsPostRequest) SetSecureLinkWithIp(v bool)`

SetSecureLinkWithIp sets SecureLinkWithIp field to given value.

### HasSecureLinkWithIp

`func (o *StreamsPostRequest) HasSecureLinkWithIp() bool`

HasSecureLinkWithIp returns a boolean if a field has been set.

### GetAdsEnabled

`func (o *StreamsPostRequest) GetAdsEnabled() bool`

GetAdsEnabled returns the AdsEnabled field if non-nil, zero value otherwise.

### GetAdsEnabledOk

`func (o *StreamsPostRequest) GetAdsEnabledOk() (*bool, bool)`

GetAdsEnabledOk returns a tuple with the AdsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdsEnabled

`func (o *StreamsPostRequest) SetAdsEnabled(v bool)`

SetAdsEnabled sets AdsEnabled field to given value.

### HasAdsEnabled

`func (o *StreamsPostRequest) HasAdsEnabled() bool`

HasAdsEnabled returns a boolean if a field has been set.

### GetPresentType

`func (o *StreamsPostRequest) GetPresentType() string`

GetPresentType returns the PresentType field if non-nil, zero value otherwise.

### GetPresentTypeOk

`func (o *StreamsPostRequest) GetPresentTypeOk() (*string, bool)`

GetPresentTypeOk returns a tuple with the PresentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentType

`func (o *StreamsPostRequest) SetPresentType(v string)`

SetPresentType sets PresentType field to given value.

### HasPresentType

`func (o *StreamsPostRequest) HasPresentType() bool`

HasPresentType returns a boolean if a field has been set.

### GetCampaignId

`func (o *StreamsPostRequest) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *StreamsPostRequest) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *StreamsPostRequest) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *StreamsPostRequest) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


