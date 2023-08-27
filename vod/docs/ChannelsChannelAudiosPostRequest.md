# ChannelsChannelAudiosPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the audio | 
**Description** | Pointer to **string** | Description of the audio | [optional] 
**AudioUrl** | Pointer to **string** | Public URL of audio | [optional] 
**FileId** | Pointer to **string** | ID of the audio file | [optional] 
**ConvertMode** | **string** | Convert mode | 
**ParallelConvert** | Pointer to **bool** | Set this convert parallel when any audio(s) is converting. Parallel limit is 3 | [optional] [default to false]
**ConvertInfo** | Pointer to [**[]ChannelsChannelAudiosPostRequestConvertInfoInner**](ChannelsChannelAudiosPostRequestConvertInfoInner.md) | Array of convert details | [optional] 

## Methods

### NewChannelsChannelAudiosPostRequest

`func NewChannelsChannelAudiosPostRequest(title string, convertMode string, ) *ChannelsChannelAudiosPostRequest`

NewChannelsChannelAudiosPostRequest instantiates a new ChannelsChannelAudiosPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelsChannelAudiosPostRequestWithDefaults

`func NewChannelsChannelAudiosPostRequestWithDefaults() *ChannelsChannelAudiosPostRequest`

NewChannelsChannelAudiosPostRequestWithDefaults instantiates a new ChannelsChannelAudiosPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ChannelsChannelAudiosPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChannelsChannelAudiosPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChannelsChannelAudiosPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ChannelsChannelAudiosPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelsChannelAudiosPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelsChannelAudiosPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelsChannelAudiosPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAudioUrl

`func (o *ChannelsChannelAudiosPostRequest) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *ChannelsChannelAudiosPostRequest) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *ChannelsChannelAudiosPostRequest) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.

### HasAudioUrl

`func (o *ChannelsChannelAudiosPostRequest) HasAudioUrl() bool`

HasAudioUrl returns a boolean if a field has been set.

### GetFileId

`func (o *ChannelsChannelAudiosPostRequest) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ChannelsChannelAudiosPostRequest) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ChannelsChannelAudiosPostRequest) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *ChannelsChannelAudiosPostRequest) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetConvertMode

`func (o *ChannelsChannelAudiosPostRequest) GetConvertMode() string`

GetConvertMode returns the ConvertMode field if non-nil, zero value otherwise.

### GetConvertModeOk

`func (o *ChannelsChannelAudiosPostRequest) GetConvertModeOk() (*string, bool)`

GetConvertModeOk returns a tuple with the ConvertMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertMode

`func (o *ChannelsChannelAudiosPostRequest) SetConvertMode(v string)`

SetConvertMode sets ConvertMode field to given value.


### GetParallelConvert

`func (o *ChannelsChannelAudiosPostRequest) GetParallelConvert() bool`

GetParallelConvert returns the ParallelConvert field if non-nil, zero value otherwise.

### GetParallelConvertOk

`func (o *ChannelsChannelAudiosPostRequest) GetParallelConvertOk() (*bool, bool)`

GetParallelConvertOk returns a tuple with the ParallelConvert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelConvert

`func (o *ChannelsChannelAudiosPostRequest) SetParallelConvert(v bool)`

SetParallelConvert sets ParallelConvert field to given value.

### HasParallelConvert

`func (o *ChannelsChannelAudiosPostRequest) HasParallelConvert() bool`

HasParallelConvert returns a boolean if a field has been set.

### GetConvertInfo

`func (o *ChannelsChannelAudiosPostRequest) GetConvertInfo() []ChannelsChannelAudiosPostRequestConvertInfoInner`

GetConvertInfo returns the ConvertInfo field if non-nil, zero value otherwise.

### GetConvertInfoOk

`func (o *ChannelsChannelAudiosPostRequest) GetConvertInfoOk() (*[]ChannelsChannelAudiosPostRequestConvertInfoInner, bool)`

GetConvertInfoOk returns a tuple with the ConvertInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertInfo

`func (o *ChannelsChannelAudiosPostRequest) SetConvertInfo(v []ChannelsChannelAudiosPostRequestConvertInfoInner)`

SetConvertInfo sets ConvertInfo field to given value.

### HasConvertInfo

`func (o *ChannelsChannelAudiosPostRequest) HasConvertInfo() bool`

HasConvertInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


