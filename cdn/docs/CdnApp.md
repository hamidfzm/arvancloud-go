# CdnApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Categories** | Pointer to [**[]ApplicationCategory**](ApplicationCategory.md) |  | [optional] 
**Rank** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** | absolute link to logo image | [optional] 
**Pictures** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**SupportEmail** | Pointer to **string** |  | [optional] 
**InstallJson** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**LikeStats** | Pointer to [**CdnAppLikeStats**](CdnAppLikeStats.md) |  | [optional] 
**LikeByUser** | Pointer to **bool** | True means she likes, False means she dislikes. null means she did not vote. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCdnApp

`func NewCdnApp() *CdnApp`

NewCdnApp instantiates a new CdnApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnAppWithDefaults

`func NewCdnAppWithDefaults() *CdnApp`

NewCdnAppWithDefaults instantiates a new CdnApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CdnApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CdnApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CdnApp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CdnApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategories

`func (o *CdnApp) GetCategories() []ApplicationCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CdnApp) GetCategoriesOk() (*[]ApplicationCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CdnApp) SetCategories(v []ApplicationCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CdnApp) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetRank

`func (o *CdnApp) GetRank() float32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *CdnApp) GetRankOk() (*float32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *CdnApp) SetRank(v float32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *CdnApp) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetName

`func (o *CdnApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdnApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdnApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CdnApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *CdnApp) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CdnApp) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CdnApp) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CdnApp) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetShortDescription

`func (o *CdnApp) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CdnApp) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CdnApp) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CdnApp) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetDescription

`func (o *CdnApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CdnApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CdnApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CdnApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogo

`func (o *CdnApp) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CdnApp) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CdnApp) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CdnApp) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPictures

`func (o *CdnApp) GetPictures() []string`

GetPictures returns the Pictures field if non-nil, zero value otherwise.

### GetPicturesOk

`func (o *CdnApp) GetPicturesOk() (*[]string, bool)`

GetPicturesOk returns a tuple with the Pictures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictures

`func (o *CdnApp) SetPictures(v []string)`

SetPictures sets Pictures field to given value.

### HasPictures

`func (o *CdnApp) HasPictures() bool`

HasPictures returns a boolean if a field has been set.

### GetVendor

`func (o *CdnApp) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CdnApp) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CdnApp) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CdnApp) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetSupportEmail

`func (o *CdnApp) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *CdnApp) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *CdnApp) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *CdnApp) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### GetInstallJson

`func (o *CdnApp) GetInstallJson() map[string]interface{}`

GetInstallJson returns the InstallJson field if non-nil, zero value otherwise.

### GetInstallJsonOk

`func (o *CdnApp) GetInstallJsonOk() (*map[string]interface{}, bool)`

GetInstallJsonOk returns a tuple with the InstallJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallJson

`func (o *CdnApp) SetInstallJson(v map[string]interface{})`

SetInstallJson sets InstallJson field to given value.

### HasInstallJson

`func (o *CdnApp) HasInstallJson() bool`

HasInstallJson returns a boolean if a field has been set.

### GetStatus

`func (o *CdnApp) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CdnApp) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CdnApp) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CdnApp) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLikeStats

`func (o *CdnApp) GetLikeStats() CdnAppLikeStats`

GetLikeStats returns the LikeStats field if non-nil, zero value otherwise.

### GetLikeStatsOk

`func (o *CdnApp) GetLikeStatsOk() (*CdnAppLikeStats, bool)`

GetLikeStatsOk returns a tuple with the LikeStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeStats

`func (o *CdnApp) SetLikeStats(v CdnAppLikeStats)`

SetLikeStats sets LikeStats field to given value.

### HasLikeStats

`func (o *CdnApp) HasLikeStats() bool`

HasLikeStats returns a boolean if a field has been set.

### GetLikeByUser

`func (o *CdnApp) GetLikeByUser() bool`

GetLikeByUser returns the LikeByUser field if non-nil, zero value otherwise.

### GetLikeByUserOk

`func (o *CdnApp) GetLikeByUserOk() (*bool, bool)`

GetLikeByUserOk returns a tuple with the LikeByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeByUser

`func (o *CdnApp) SetLikeByUser(v bool)`

SetLikeByUser sets LikeByUser field to given value.

### HasLikeByUser

`func (o *CdnApp) HasLikeByUser() bool`

HasLikeByUser returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CdnApp) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CdnApp) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CdnApp) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CdnApp) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CdnApp) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CdnApp) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CdnApp) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CdnApp) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


