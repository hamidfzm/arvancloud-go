# CampaignsCampaignAdsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdId** | **string** | The id of ad      *                  Notice: Both ad &amp; campaign most belongs to same channel. | 
**Weight** | **int32** | Weight of ad.      *                  When exists similar type ad(s,) use weight for choose ad to insert in VAST XML | 
**FilterDevice** | Pointer to **[]string** | Set device filter for show ads | [optional] 
**FilterBrowser** | Pointer to **[]string** | Set browser filter for show ads | [optional] 
**FilterPlatform** | Pointer to **[]string** | Set platform filter for show ads | [optional] 
**Quota** | Pointer to [**[]CampaignsCampaignAdsPostRequestQuotaInner**](CampaignsCampaignAdsPostRequestQuotaInner.md) | Limited condition(s) of trigger event that stop present ad in campaign | [optional] 
**QuotaType** | Pointer to **string** | Quota type | [optional] 

## Methods

### NewCampaignsCampaignAdsPostRequest

`func NewCampaignsCampaignAdsPostRequest(adId string, weight int32, ) *CampaignsCampaignAdsPostRequest`

NewCampaignsCampaignAdsPostRequest instantiates a new CampaignsCampaignAdsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignsCampaignAdsPostRequestWithDefaults

`func NewCampaignsCampaignAdsPostRequestWithDefaults() *CampaignsCampaignAdsPostRequest`

NewCampaignsCampaignAdsPostRequestWithDefaults instantiates a new CampaignsCampaignAdsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdId

`func (o *CampaignsCampaignAdsPostRequest) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *CampaignsCampaignAdsPostRequest) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *CampaignsCampaignAdsPostRequest) SetAdId(v string)`

SetAdId sets AdId field to given value.


### GetWeight

`func (o *CampaignsCampaignAdsPostRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CampaignsCampaignAdsPostRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CampaignsCampaignAdsPostRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetFilterDevice

`func (o *CampaignsCampaignAdsPostRequest) GetFilterDevice() []string`

GetFilterDevice returns the FilterDevice field if non-nil, zero value otherwise.

### GetFilterDeviceOk

`func (o *CampaignsCampaignAdsPostRequest) GetFilterDeviceOk() (*[]string, bool)`

GetFilterDeviceOk returns a tuple with the FilterDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDevice

`func (o *CampaignsCampaignAdsPostRequest) SetFilterDevice(v []string)`

SetFilterDevice sets FilterDevice field to given value.

### HasFilterDevice

`func (o *CampaignsCampaignAdsPostRequest) HasFilterDevice() bool`

HasFilterDevice returns a boolean if a field has been set.

### GetFilterBrowser

`func (o *CampaignsCampaignAdsPostRequest) GetFilterBrowser() []string`

GetFilterBrowser returns the FilterBrowser field if non-nil, zero value otherwise.

### GetFilterBrowserOk

`func (o *CampaignsCampaignAdsPostRequest) GetFilterBrowserOk() (*[]string, bool)`

GetFilterBrowserOk returns a tuple with the FilterBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBrowser

`func (o *CampaignsCampaignAdsPostRequest) SetFilterBrowser(v []string)`

SetFilterBrowser sets FilterBrowser field to given value.

### HasFilterBrowser

`func (o *CampaignsCampaignAdsPostRequest) HasFilterBrowser() bool`

HasFilterBrowser returns a boolean if a field has been set.

### GetFilterPlatform

`func (o *CampaignsCampaignAdsPostRequest) GetFilterPlatform() []string`

GetFilterPlatform returns the FilterPlatform field if non-nil, zero value otherwise.

### GetFilterPlatformOk

`func (o *CampaignsCampaignAdsPostRequest) GetFilterPlatformOk() (*[]string, bool)`

GetFilterPlatformOk returns a tuple with the FilterPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPlatform

`func (o *CampaignsCampaignAdsPostRequest) SetFilterPlatform(v []string)`

SetFilterPlatform sets FilterPlatform field to given value.

### HasFilterPlatform

`func (o *CampaignsCampaignAdsPostRequest) HasFilterPlatform() bool`

HasFilterPlatform returns a boolean if a field has been set.

### GetQuota

`func (o *CampaignsCampaignAdsPostRequest) GetQuota() []CampaignsCampaignAdsPostRequestQuotaInner`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CampaignsCampaignAdsPostRequest) GetQuotaOk() (*[]CampaignsCampaignAdsPostRequestQuotaInner, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CampaignsCampaignAdsPostRequest) SetQuota(v []CampaignsCampaignAdsPostRequestQuotaInner)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CampaignsCampaignAdsPostRequest) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetQuotaType

`func (o *CampaignsCampaignAdsPostRequest) GetQuotaType() string`

GetQuotaType returns the QuotaType field if non-nil, zero value otherwise.

### GetQuotaTypeOk

`func (o *CampaignsCampaignAdsPostRequest) GetQuotaTypeOk() (*string, bool)`

GetQuotaTypeOk returns a tuple with the QuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaType

`func (o *CampaignsCampaignAdsPostRequest) SetQuotaType(v string)`

SetQuotaType sets QuotaType field to given value.

### HasQuotaType

`func (o *CampaignsCampaignAdsPostRequest) HasQuotaType() bool`

HasQuotaType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


