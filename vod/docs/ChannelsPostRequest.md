# ChannelsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of channel | 
**Description** | Pointer to **string** | Description of channel | [optional] 
**SecureLinkEnabled** | Pointer to **bool** | Enable or disable secure link for all videos in channel | [optional] 
**SecureLinkKey** | Pointer to **string** | Key for generate secure links | [optional] 
**SecureLinkWithIp** | Pointer to **bool** | IP can be considered as an optional parameter | [optional] 
**AdsEnabled** | Pointer to **bool** | Enable or disable Ads for all videos in channel | [optional] 
**PresentType** | Pointer to **string** | Ads present method | [optional] 
**CampaignId** | Pointer to **string** | Created CampaignId in Ads | [optional] 

## Methods

### NewChannelsPostRequest

`func NewChannelsPostRequest(title string, ) *ChannelsPostRequest`

NewChannelsPostRequest instantiates a new ChannelsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelsPostRequestWithDefaults

`func NewChannelsPostRequestWithDefaults() *ChannelsPostRequest`

NewChannelsPostRequestWithDefaults instantiates a new ChannelsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ChannelsPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChannelsPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChannelsPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *ChannelsPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelsPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelsPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelsPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecureLinkEnabled

`func (o *ChannelsPostRequest) GetSecureLinkEnabled() bool`

GetSecureLinkEnabled returns the SecureLinkEnabled field if non-nil, zero value otherwise.

### GetSecureLinkEnabledOk

`func (o *ChannelsPostRequest) GetSecureLinkEnabledOk() (*bool, bool)`

GetSecureLinkEnabledOk returns a tuple with the SecureLinkEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkEnabled

`func (o *ChannelsPostRequest) SetSecureLinkEnabled(v bool)`

SetSecureLinkEnabled sets SecureLinkEnabled field to given value.

### HasSecureLinkEnabled

`func (o *ChannelsPostRequest) HasSecureLinkEnabled() bool`

HasSecureLinkEnabled returns a boolean if a field has been set.

### GetSecureLinkKey

`func (o *ChannelsPostRequest) GetSecureLinkKey() string`

GetSecureLinkKey returns the SecureLinkKey field if non-nil, zero value otherwise.

### GetSecureLinkKeyOk

`func (o *ChannelsPostRequest) GetSecureLinkKeyOk() (*string, bool)`

GetSecureLinkKeyOk returns a tuple with the SecureLinkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkKey

`func (o *ChannelsPostRequest) SetSecureLinkKey(v string)`

SetSecureLinkKey sets SecureLinkKey field to given value.

### HasSecureLinkKey

`func (o *ChannelsPostRequest) HasSecureLinkKey() bool`

HasSecureLinkKey returns a boolean if a field has been set.

### GetSecureLinkWithIp

`func (o *ChannelsPostRequest) GetSecureLinkWithIp() bool`

GetSecureLinkWithIp returns the SecureLinkWithIp field if non-nil, zero value otherwise.

### GetSecureLinkWithIpOk

`func (o *ChannelsPostRequest) GetSecureLinkWithIpOk() (*bool, bool)`

GetSecureLinkWithIpOk returns a tuple with the SecureLinkWithIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureLinkWithIp

`func (o *ChannelsPostRequest) SetSecureLinkWithIp(v bool)`

SetSecureLinkWithIp sets SecureLinkWithIp field to given value.

### HasSecureLinkWithIp

`func (o *ChannelsPostRequest) HasSecureLinkWithIp() bool`

HasSecureLinkWithIp returns a boolean if a field has been set.

### GetAdsEnabled

`func (o *ChannelsPostRequest) GetAdsEnabled() bool`

GetAdsEnabled returns the AdsEnabled field if non-nil, zero value otherwise.

### GetAdsEnabledOk

`func (o *ChannelsPostRequest) GetAdsEnabledOk() (*bool, bool)`

GetAdsEnabledOk returns a tuple with the AdsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdsEnabled

`func (o *ChannelsPostRequest) SetAdsEnabled(v bool)`

SetAdsEnabled sets AdsEnabled field to given value.

### HasAdsEnabled

`func (o *ChannelsPostRequest) HasAdsEnabled() bool`

HasAdsEnabled returns a boolean if a field has been set.

### GetPresentType

`func (o *ChannelsPostRequest) GetPresentType() string`

GetPresentType returns the PresentType field if non-nil, zero value otherwise.

### GetPresentTypeOk

`func (o *ChannelsPostRequest) GetPresentTypeOk() (*string, bool)`

GetPresentTypeOk returns a tuple with the PresentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentType

`func (o *ChannelsPostRequest) SetPresentType(v string)`

SetPresentType sets PresentType field to given value.

### HasPresentType

`func (o *ChannelsPostRequest) HasPresentType() bool`

HasPresentType returns a boolean if a field has been set.

### GetCampaignId

`func (o *ChannelsPostRequest) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ChannelsPostRequest) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ChannelsPostRequest) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *ChannelsPostRequest) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


