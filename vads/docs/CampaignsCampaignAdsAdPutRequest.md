# CampaignsCampaignAdsAdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | Pointer to **int32** | Weight of ad.      *                  When exists similar type ad(s,) use weight for choose ad to insert in VAST XML | [optional] 
**FilterDevice** | Pointer to **[]string** | Set device filter for show ads | [optional] 
**FilterBrowser** | Pointer to **[]string** | Set browser filter for show ads | [optional] 
**FilterPlatform** | Pointer to **[]string** | Set platform filter for show ads | [optional] 
**Quota** | Pointer to [**[]CampaignsCampaignAdsPostRequestQuotaInner**](CampaignsCampaignAdsPostRequestQuotaInner.md) | Limited condition(s) of trigger event that stop present ad in campaign | [optional] 
**QuotaType** | Pointer to **string** | Quota type | [optional] 

## Methods

### NewCampaignsCampaignAdsAdPutRequest

`func NewCampaignsCampaignAdsAdPutRequest() *CampaignsCampaignAdsAdPutRequest`

NewCampaignsCampaignAdsAdPutRequest instantiates a new CampaignsCampaignAdsAdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignsCampaignAdsAdPutRequestWithDefaults

`func NewCampaignsCampaignAdsAdPutRequestWithDefaults() *CampaignsCampaignAdsAdPutRequest`

NewCampaignsCampaignAdsAdPutRequestWithDefaults instantiates a new CampaignsCampaignAdsAdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *CampaignsCampaignAdsAdPutRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CampaignsCampaignAdsAdPutRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CampaignsCampaignAdsAdPutRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CampaignsCampaignAdsAdPutRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetFilterDevice

`func (o *CampaignsCampaignAdsAdPutRequest) GetFilterDevice() []string`

GetFilterDevice returns the FilterDevice field if non-nil, zero value otherwise.

### GetFilterDeviceOk

`func (o *CampaignsCampaignAdsAdPutRequest) GetFilterDeviceOk() (*[]string, bool)`

GetFilterDeviceOk returns a tuple with the FilterDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDevice

`func (o *CampaignsCampaignAdsAdPutRequest) SetFilterDevice(v []string)`

SetFilterDevice sets FilterDevice field to given value.

### HasFilterDevice

`func (o *CampaignsCampaignAdsAdPutRequest) HasFilterDevice() bool`

HasFilterDevice returns a boolean if a field has been set.

### GetFilterBrowser

`func (o *CampaignsCampaignAdsAdPutRequest) GetFilterBrowser() []string`

GetFilterBrowser returns the FilterBrowser field if non-nil, zero value otherwise.

### GetFilterBrowserOk

`func (o *CampaignsCampaignAdsAdPutRequest) GetFilterBrowserOk() (*[]string, bool)`

GetFilterBrowserOk returns a tuple with the FilterBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBrowser

`func (o *CampaignsCampaignAdsAdPutRequest) SetFilterBrowser(v []string)`

SetFilterBrowser sets FilterBrowser field to given value.

### HasFilterBrowser

`func (o *CampaignsCampaignAdsAdPutRequest) HasFilterBrowser() bool`

HasFilterBrowser returns a boolean if a field has been set.

### GetFilterPlatform

`func (o *CampaignsCampaignAdsAdPutRequest) GetFilterPlatform() []string`

GetFilterPlatform returns the FilterPlatform field if non-nil, zero value otherwise.

### GetFilterPlatformOk

`func (o *CampaignsCampaignAdsAdPutRequest) GetFilterPlatformOk() (*[]string, bool)`

GetFilterPlatformOk returns a tuple with the FilterPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPlatform

`func (o *CampaignsCampaignAdsAdPutRequest) SetFilterPlatform(v []string)`

SetFilterPlatform sets FilterPlatform field to given value.

### HasFilterPlatform

`func (o *CampaignsCampaignAdsAdPutRequest) HasFilterPlatform() bool`

HasFilterPlatform returns a boolean if a field has been set.

### GetQuota

`func (o *CampaignsCampaignAdsAdPutRequest) GetQuota() []CampaignsCampaignAdsPostRequestQuotaInner`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *CampaignsCampaignAdsAdPutRequest) GetQuotaOk() (*[]CampaignsCampaignAdsPostRequestQuotaInner, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *CampaignsCampaignAdsAdPutRequest) SetQuota(v []CampaignsCampaignAdsPostRequestQuotaInner)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *CampaignsCampaignAdsAdPutRequest) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetQuotaType

`func (o *CampaignsCampaignAdsAdPutRequest) GetQuotaType() string`

GetQuotaType returns the QuotaType field if non-nil, zero value otherwise.

### GetQuotaTypeOk

`func (o *CampaignsCampaignAdsAdPutRequest) GetQuotaTypeOk() (*string, bool)`

GetQuotaTypeOk returns a tuple with the QuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaType

`func (o *CampaignsCampaignAdsAdPutRequest) SetQuotaType(v string)`

SetQuotaType sets QuotaType field to given value.

### HasQuotaType

`func (o *CampaignsCampaignAdsAdPutRequest) HasQuotaType() bool`

HasQuotaType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


