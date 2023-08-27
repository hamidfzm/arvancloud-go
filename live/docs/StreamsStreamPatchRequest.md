# StreamsStreamPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the stream | [optional] 
**Description** | Pointer to **string** | Description of the stream | [optional] 
**Type** | Pointer to **string** | Stream latency type | [optional] 
**Mode** | Pointer to **string** | Way to receive input stream | [optional] 
**InputUrl** | Pointer to **string** | Public URL of stream input in PULL mode | [optional] 
**Slug** | Pointer to **string** | Unique slug for your stream. Only contain lower case letters and digits | [optional] 
**Timeshift** | Pointer to **int32** | Timeshift (DVR) in minutes to watch the earlier content | [optional] 
**FpsMode** | Pointer to **string** | stream fps mode | [optional] 
**Fps** | Pointer to **int32** | Input steam frame per second | [optional] 
**ArchiveEnabled** | Pointer to **bool** | Set this value to True to archive the stream as a VoD | [optional] [default to false]
**CatchupEnabled** | Pointer to **bool** | Set this value to True to enable catchup for the stream | [optional] [default to false]
**CatchupPeriod** | Pointer to **int32** | Hours of catchup must be available for the stream | [optional] 
**ArchiveMode** | Pointer to **string** | Way to archive stream | [optional] 
**ChannelId** | Pointer to **string** | Channel Id to save archive | [optional] 
**WatermarkId** | Pointer to **string** | If you want to use watermark for a video, use this ID | [optional] 
**WatermarkArea** | Pointer to **string** | Area of the watermark if watermark_id presents | [optional] 
**ConvertInfo** | Pointer to [**[]StreamsPostRequestConvertInfoInner**](StreamsPostRequestConvertInfoInner.md) | Array of convert details | [optional] 
**SecureLinkEnabled** | Pointer to **bool** | Enable or disable secure link for all videos in channel | [optional] 
**SecureLinkKey** | Pointer to **string** | Key for generate secure links | [optional] 
**SecureLinkWithIp** | Pointer to **bool** | IP can be considered as an optional parameter | [optional] 
**AdsEnabled** | Pointer to **bool** | Enable or disable Ads | [optional] 
**PresentType** | Pointer to **string** | Ads present method | [optional] 
**CampaignId** | Pointer to **string** | Created CampaignId in Ads | [optional] 

## Methods

### NewStreamsStreamPatchRequest

`func NewStreamsStreamPatchRequest() *StreamsStreamPatchRequest`

NewStreamsStreamPatchRequest instantiates a new StreamsStreamPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsStreamPatchRequestWithDefaults

`func NewStreamsStreamPatchRequestWithDefaults() *StreamsStreamPatchRequest`

NewStreamsStreamPatchRequestWithDefaults instantiates a new StreamsStreamPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *StreamsStreamPatchRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StreamsStreamPatchRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StreamsStreamPatchRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StreamsStreamPatchRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *StreamsStreamPatchRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StreamsStreamPatchRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StreamsStreamPatchRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StreamsStreamPatchRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *StreamsStreamPatchRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamsStreamPatchRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamsStreamPatchRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StreamsStreamPatchRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMode

`func (o *StreamsStreamPatchRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *StreamsStreamPatchRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *StreamsStreamPatchRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *StreamsStreamPatchRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetInputUrl

`func (o *StreamsStreamPatchRequest) GetInputUrl() string`

GetInputUrl returns the InputUrl field if non-nil, zero value otherwise.

### GetInputUrlOk

`func (o *StreamsStreamPatchRequest) GetInputUrlOk() (*string, bool)`

GetInputUrlOk returns a tuple with the InputUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputUrl

`func (o *StreamsStreamPatchRequest) SetInputUrl(v string)`

SetInputUrl sets InputUrl field to given value.

### HasInputUrl

`func (o *StreamsStreamPatchRequest) HasInputUrl() bool`

HasInputUrl returns a boolean if a field has been set.

### GetSlug

`func (o *StreamsStreamPatchRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StreamsStreamPatchRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StreamsStreamPatchRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *StreamsStreamPatchRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTimeshift

`func (o *StreamsStreamPatchRequest) GetTimeshift() int32`

GetTimeshift returns the Timeshift field if non-nil, zero value otherwise.

### GetTimeshiftOk

`func (o *StreamsStreamPatchRequest) GetTimeshiftOk() (*int32, bool)`

GetTimeshiftOk returns a tuple with the Timeshift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeshift

`func (o *StreamsStreamPatchRequest) SetTimeshift(v int32)`

SetTimeshift sets Timeshift field to given value.

### HasTimeshift

`func (o *StreamsStreamPatchRequest) HasTimeshift() bool`

HasTimeshift returns a boolean if a field has been set.

### GetFpsMode

`func (o *StreamsStreamPatchRequest) GetFpsMode() string`

GetFpsMode returns the FpsMode field if non-nil, zero value otherwise.

### GetFpsModeOk

`func (o *StreamsStreamPatchRequest) GetFpsModeOk() (*string, bool)`

GetFpsModeOk returns a tuple with the FpsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpsMode

`func (o *StreamsStreamPatchRequest) SetFpsMode(v string)`

SetFpsMode sets FpsMode field to given value.

### HasFpsMode

`func (o *StreamsStreamPatchRequest) HasFpsMode() bool`

HasFpsMode returns a boolean if a field has been set.

### GetFps

`func (o *StreamsStreamPatchRequest) GetFps() int32`

GetFps returns the Fps field if non-nil, zero value otherwise.

### GetFpsOk

`func (o *StreamsStreamPatchRequest) GetFpsOk() (*int32, bool)`

GetFpsOk returns a tuple with the Fps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFps

`func (o *StreamsStreamPatchRequest) SetFps(v int32)`

SetFps sets Fps field to given value.

### HasFps

`func (o *StreamsStreamPatchRequest) HasFps() bool`

HasFps returns a boolean if a field has been set.

### GetArchiveEnabled

`func (o *StreamsStreamPatchRequest) GetArchiveEnabled() bool`

GetArchiveEnabled returns the ArchiveEnabled field if non-nil, zero value otherwise.

### GetArchiveEnabledOk

`func (o *StreamsStreamPatchRequest) GetArchiveEnabledOk() (*bool, bool)`

GetArchiveEnabledOk returns a tuple with the ArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveEnabled

`func (o *StreamsStreamPatchRequest) SetArchiveEnabled(v bool)`

SetArchiveEnabled sets ArchiveEnabled field to given value.

### HasArchiveEnabled

`func (o *StreamsStreamPatchRequest) HasArchiveEnabled() bool`

HasArchiveEnabled returns a boolean if a field has been set.

### GetCatchupEnabled

`func (o *StreamsStreamPatchRequest) GetCatchupEnabled() bool`

GetCatchupEnabled returns the CatchupEnabled field if non-nil, zero value otherwise.

### GetCatchupEnabledOk

`func (o *StreamsStreamPatchRequest) GetCatchupEnabledOk() (*bool, bool)`

GetCatchupEnabledOk returns a tuple with the CatchupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchupEnabled

`func (o *StreamsStreamPatchRequest) SetCatchupEnabled(v bool)`

SetCatchupEnabled sets CatchupEnabled field to given value.

### HasCatchupEnabled

`func (o *StreamsStreamPatchRequest) HasCatchupEnabled() bool`

HasCatchupEnabled returns a boolean if a field has been set.

### GetCatchupPeriod

`func (o *StreamsStreamPatchRequest) GetCatchupPeriod() int32`

GetCatchupPeriod returns the CatchupPeriod field if non-nil, zero value otherwise.

### GetCatchupPeriodOk

`func (o *StreamsStreamPatchRequest) GetCatchupPeriodOk() (*int32, bool)`

GetCatchupPeriodOk returns a tuple with the CatchupPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchupPeriod

`func (o *StreamsStreamPatchRequest) SetCatchupPeriod(v int32)`

SetCatchupPeriod sets CatchupPeriod field to given value.

### HasCatchupPeriod

`func (o *StreamsStreamPatchRequest) HasCatchupPeriod() bool`

HasCatchupPeriod returns a boolean if a field has been set.

### GetArchiveMode

`func (o *StreamsStreamPatchRequest) GetArchiveMode() string`

GetArchiveMode returns the ArchiveMode field if non-nil, zero value otherwise.

### GetArchiveModeOk

`func (o *StreamsStreamPatchRequest) GetArchiveModeOk() (*string, bool)`

GetArchiveModeOk returns a tuple with the ArchiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveMode

`func (o *StreamsStreamPatchRequest) SetArchiveMode(v string)`

SetArchiveMode sets ArchiveMode field to given value.

### HasArchiveMode

`func (o *StreamsStreamPatchRequest) HasArchiveMode() bool`

HasArchiveMode returns a boolean if a field has been set.

### GetChannelId

`func (o *StreamsStreamPatchRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *StreamsStreamPatchRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *StreamsStreamPatchRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *StreamsStreamPatchRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetWatermarkId

`func (o *StreamsStreamPatchRequest) GetWatermarkId() string`

GetWatermarkId returns the WatermarkId field if non-nil, zero value otherwise.

### GetWatermarkIdOk

`func (o *StreamsStreamPatchRequest) GetWatermarkIdOk() (*string, bool)`

GetWatermarkIdOk returns a tuple with the WatermarkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkId

`func (o *StreamsStreamPatchRequest) SetWatermarkId(v string)`

SetWatermarkId sets WatermarkId field to given value.

### HasWatermarkId

`func (o *StreamsStreamPatchRequest) HasWatermarkId() bool`

HasWatermarkId returns a boolean if a field has been set.

### GetWatermarkArea

`func (o *StreamsStreamPatchRequest) GetWatermarkArea() string`

GetWatermarkArea returns the WatermarkArea field if non-nil, zero value otherwise.

### GetWatermarkAreaOk

`func (o *StreamsStreamPatchRequest) GetWatermarkAreaOk() (*string, bool)`

GetWatermarkAreaOk returns a tuple with the WatermarkArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermarkArea

`func (o *StreamsStreamPatchRequest) SetWatermarkArea(v string)`

SetWatermarkArea sets WatermarkArea field to given value.

### HasWatermarkArea

`func (o *StreamsStreamPatchRequest) HasWatermarkArea() bool`

HasWatermarkArea returns a boolean if a field has been set.

### GetConvertInfo

`func (o *StreamsStreamPatchRequest) GetConvertInfo() []StreamsPostRequestConvertInfoInner`

GetConvertInfo returns the ConvertInfo field if non-nil, zero value otherwise.

### GetConvertInfoOk

`func (o *StreamsStreamPatchRequest) GetConvertInfoOk() (*[]StreamsPostRequestConvertInfoInner, bool)`

GetConvertInfoOk returns a tuple with the ConvertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertInfo

`func (o *StreamsStreamPatchRequest) SetConvertInfo(v []StreamsPostRequestConvertInfoInner)`

SetConvertInfo sets ConvertInfo field to given value.

### HasConvertInfo

`func (o *StreamsStreamPatchRequest) HasConvertInfo() bool`

HasConvertInfo returns a boolean if a field has been set.

### GetSecureLinkEnabled

`func (o *StreamsStreamPatchRequest) GetSecureLinkEnabled() bool`

GetSecureLinkEnabled returns the SecureLinkEnabled field if non-nil, zero value otherwise.

### GetSecureLinkEnabledOk

`func (o *StreamsStreamPatchRequest) GetSecureLinkEnabledOk() (*bool, bool)`

GetSecureLinkEnabledOk returns a tuple with the SecureLinkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkEnabled

`func (o *StreamsStreamPatchRequest) SetSecureLinkEnabled(v bool)`

SetSecureLinkEnabled sets SecureLinkEnabled field to given value.

### HasSecureLinkEnabled

`func (o *StreamsStreamPatchRequest) HasSecureLinkEnabled() bool`

HasSecureLinkEnabled returns a boolean if a field has been set.

### GetSecureLinkKey

`func (o *StreamsStreamPatchRequest) GetSecureLinkKey() string`

GetSecureLinkKey returns the SecureLinkKey field if non-nil, zero value otherwise.

### GetSecureLinkKeyOk

`func (o *StreamsStreamPatchRequest) GetSecureLinkKeyOk() (*string, bool)`

GetSecureLinkKeyOk returns a tuple with the SecureLinkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkKey

`func (o *StreamsStreamPatchRequest) SetSecureLinkKey(v string)`

SetSecureLinkKey sets SecureLinkKey field to given value.

### HasSecureLinkKey

`func (o *StreamsStreamPatchRequest) HasSecureLinkKey() bool`

HasSecureLinkKey returns a boolean if a field has been set.

### GetSecureLinkWithIp

`func (o *StreamsStreamPatchRequest) GetSecureLinkWithIp() bool`

GetSecureLinkWithIp returns the SecureLinkWithIp field if non-nil, zero value otherwise.

### GetSecureLinkWithIpOk

`func (o *StreamsStreamPatchRequest) GetSecureLinkWithIpOk() (*bool, bool)`

GetSecureLinkWithIpOk returns a tuple with the SecureLinkWithIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkWithIp

`func (o *StreamsStreamPatchRequest) SetSecureLinkWithIp(v bool)`

SetSecureLinkWithIp sets SecureLinkWithIp field to given value.

### HasSecureLinkWithIp

`func (o *StreamsStreamPatchRequest) HasSecureLinkWithIp() bool`

HasSecureLinkWithIp returns a boolean if a field has been set.

### GetAdsEnabled

`func (o *StreamsStreamPatchRequest) GetAdsEnabled() bool`

GetAdsEnabled returns the AdsEnabled field if non-nil, zero value otherwise.

### GetAdsEnabledOk

`func (o *StreamsStreamPatchRequest) GetAdsEnabledOk() (*bool, bool)`

GetAdsEnabledOk returns a tuple with the AdsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdsEnabled

`func (o *StreamsStreamPatchRequest) SetAdsEnabled(v bool)`

SetAdsEnabled sets AdsEnabled field to given value.

### HasAdsEnabled

`func (o *StreamsStreamPatchRequest) HasAdsEnabled() bool`

HasAdsEnabled returns a boolean if a field has been set.

### GetPresentType

`func (o *StreamsStreamPatchRequest) GetPresentType() string`

GetPresentType returns the PresentType field if non-nil, zero value otherwise.

### GetPresentTypeOk

`func (o *StreamsStreamPatchRequest) GetPresentTypeOk() (*string, bool)`

GetPresentTypeOk returns a tuple with the PresentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentType

`func (o *StreamsStreamPatchRequest) SetPresentType(v string)`

SetPresentType sets PresentType field to given value.

### HasPresentType

`func (o *StreamsStreamPatchRequest) HasPresentType() bool`

HasPresentType returns a boolean if a field has been set.

### GetCampaignId

`func (o *StreamsStreamPatchRequest) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *StreamsStreamPatchRequest) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *StreamsStreamPatchRequest) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *StreamsStreamPatchRequest) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


