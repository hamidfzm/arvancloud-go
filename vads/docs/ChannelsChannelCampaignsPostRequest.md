# ChannelsChannelCampaignsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of campaign | 
**Description** | Pointer to **string** | Description of campaign | [optional] 
**SkipType** | **string** | Skip type - If ad skip_type is follow_campaign then this will be use | 
**SkipOffset** | Pointer to **int32** | Skip offset in seconds (required if skip type is allow) | [optional] 
**Active** | **bool** | If false then vast not working | 

## Methods

### NewChannelsChannelCampaignsPostRequest

`func NewChannelsChannelCampaignsPostRequest(title string, skipType string, active bool, ) *ChannelsChannelCampaignsPostRequest`

NewChannelsChannelCampaignsPostRequest instantiates a new ChannelsChannelCampaignsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelsChannelCampaignsPostRequestWithDefaults

`func NewChannelsChannelCampaignsPostRequestWithDefaults() *ChannelsChannelCampaignsPostRequest`

NewChannelsChannelCampaignsPostRequestWithDefaults instantiates a new ChannelsChannelCampaignsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ChannelsChannelCampaignsPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChannelsChannelCampaignsPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChannelsChannelCampaignsPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ChannelsChannelCampaignsPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelsChannelCampaignsPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelsChannelCampaignsPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelsChannelCampaignsPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSkipType

`func (o *ChannelsChannelCampaignsPostRequest) GetSkipType() string`

GetSkipType returns the SkipType field if non-nil, zero value otherwise.

### GetSkipTypeOk

`func (o *ChannelsChannelCampaignsPostRequest) GetSkipTypeOk() (*string, bool)`

GetSkipTypeOk returns a tuple with the SkipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipType

`func (o *ChannelsChannelCampaignsPostRequest) SetSkipType(v string)`

SetSkipType sets SkipType field to given value.


### GetSkipOffset

`func (o *ChannelsChannelCampaignsPostRequest) GetSkipOffset() int32`

GetSkipOffset returns the SkipOffset field if non-nil, zero value otherwise.

### GetSkipOffsetOk

`func (o *ChannelsChannelCampaignsPostRequest) GetSkipOffsetOk() (*int32, bool)`

GetSkipOffsetOk returns a tuple with the SkipOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOffset

`func (o *ChannelsChannelCampaignsPostRequest) SetSkipOffset(v int32)`

SetSkipOffset sets SkipOffset field to given value.

### HasSkipOffset

`func (o *ChannelsChannelCampaignsPostRequest) HasSkipOffset() bool`

HasSkipOffset returns a boolean if a field has been set.

### GetActive

`func (o *ChannelsChannelCampaignsPostRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ChannelsChannelCampaignsPostRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ChannelsChannelCampaignsPostRequest) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


