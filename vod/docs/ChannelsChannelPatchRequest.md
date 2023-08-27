# ChannelsChannelPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of channel | [optional] 
**Description** | Pointer to **string** | Description of channel | [optional] 
**SecureLinkEnabled** | Pointer to **bool** | Enable or disable secure link for all videos in channel | [optional] 
**SecureLinkKey** | Pointer to **string** | Key for generate secure links | [optional] 
**SecureLinkWithIp** | Pointer to **bool** | IP can be considered as an optional parameter | [optional] 
**AdsEnabled** | Pointer to **bool** | Enable or disable Ads for all videos in channel | [optional] 
**PresentType** | Pointer to **string** | Ads present method | [optional] 
**CampaignId** | Pointer to **string** | Created CampaignId in Ads | [optional] 

## Methods

### NewChannelsChannelPatchRequest

`func NewChannelsChannelPatchRequest() *ChannelsChannelPatchRequest`

NewChannelsChannelPatchRequest instantiates a new ChannelsChannelPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelsChannelPatchRequestWithDefaults

`func NewChannelsChannelPatchRequestWithDefaults() *ChannelsChannelPatchRequest`

NewChannelsChannelPatchRequestWithDefaults instantiates a new ChannelsChannelPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ChannelsChannelPatchRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChannelsChannelPatchRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChannelsChannelPatchRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChannelsChannelPatchRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ChannelsChannelPatchRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelsChannelPatchRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelsChannelPatchRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelsChannelPatchRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecureLinkEnabled

`func (o *ChannelsChannelPatchRequest) GetSecureLinkEnabled() bool`

GetSecureLinkEnabled returns the SecureLinkEnabled field if non-nil, zero value otherwise.

### GetSecureLinkEnabledOk

`func (o *ChannelsChannelPatchRequest) GetSecureLinkEnabledOk() (*bool, bool)`

GetSecureLinkEnabledOk returns a tuple with the SecureLinkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkEnabled

`func (o *ChannelsChannelPatchRequest) SetSecureLinkEnabled(v bool)`

SetSecureLinkEnabled sets SecureLinkEnabled field to given value.

### HasSecureLinkEnabled

`func (o *ChannelsChannelPatchRequest) HasSecureLinkEnabled() bool`

HasSecureLinkEnabled returns a boolean if a field has been set.

### GetSecureLinkKey

`func (o *ChannelsChannelPatchRequest) GetSecureLinkKey() string`

GetSecureLinkKey returns the SecureLinkKey field if non-nil, zero value otherwise.

### GetSecureLinkKeyOk

`func (o *ChannelsChannelPatchRequest) GetSecureLinkKeyOk() (*string, bool)`

GetSecureLinkKeyOk returns a tuple with the SecureLinkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkKey

`func (o *ChannelsChannelPatchRequest) SetSecureLinkKey(v string)`

SetSecureLinkKey sets SecureLinkKey field to given value.

### HasSecureLinkKey

`func (o *ChannelsChannelPatchRequest) HasSecureLinkKey() bool`

HasSecureLinkKey returns a boolean if a field has been set.

### GetSecureLinkWithIp

`func (o *ChannelsChannelPatchRequest) GetSecureLinkWithIp() bool`

GetSecureLinkWithIp returns the SecureLinkWithIp field if non-nil, zero value otherwise.

### GetSecureLinkWithIpOk

`func (o *ChannelsChannelPatchRequest) GetSecureLinkWithIpOk() (*bool, bool)`

GetSecureLinkWithIpOk returns a tuple with the SecureLinkWithIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkWithIp

`func (o *ChannelsChannelPatchRequest) SetSecureLinkWithIp(v bool)`

SetSecureLinkWithIp sets SecureLinkWithIp field to given value.

### HasSecureLinkWithIp

`func (o *ChannelsChannelPatchRequest) HasSecureLinkWithIp() bool`

HasSecureLinkWithIp returns a boolean if a field has been set.

### GetAdsEnabled

`func (o *ChannelsChannelPatchRequest) GetAdsEnabled() bool`

GetAdsEnabled returns the AdsEnabled field if non-nil, zero value otherwise.

### GetAdsEnabledOk

`func (o *ChannelsChannelPatchRequest) GetAdsEnabledOk() (*bool, bool)`

GetAdsEnabledOk returns a tuple with the AdsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdsEnabled

`func (o *ChannelsChannelPatchRequest) SetAdsEnabled(v bool)`

SetAdsEnabled sets AdsEnabled field to given value.

### HasAdsEnabled

`func (o *ChannelsChannelPatchRequest) HasAdsEnabled() bool`

HasAdsEnabled returns a boolean if a field has been set.

### GetPresentType

`func (o *ChannelsChannelPatchRequest) GetPresentType() string`

GetPresentType returns the PresentType field if non-nil, zero value otherwise.

### GetPresentTypeOk

`func (o *ChannelsChannelPatchRequest) GetPresentTypeOk() (*string, bool)`

GetPresentTypeOk returns a tuple with the PresentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentType

`func (o *ChannelsChannelPatchRequest) SetPresentType(v string)`

SetPresentType sets PresentType field to given value.

### HasPresentType

`func (o *ChannelsChannelPatchRequest) HasPresentType() bool`

HasPresentType returns a boolean if a field has been set.

### GetCampaignId

`func (o *ChannelsChannelPatchRequest) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ChannelsChannelPatchRequest) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ChannelsChannelPatchRequest) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *ChannelsChannelPatchRequest) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


