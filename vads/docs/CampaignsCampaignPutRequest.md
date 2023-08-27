# CampaignsCampaignPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of campaign | [optional] 
**Description** | Pointer to **string** | Description of campaign | [optional] 
**SkipType** | Pointer to **string** | Skip type - If ad skip_type is follow_campaign then this will be use | [optional] 
**SkipOffset** | Pointer to **int32** | Skip offset in seconds (required if skip type is allow) | [optional] 
**Active** | Pointer to **bool** | If false then vast not working | [optional] 

## Methods

### NewCampaignsCampaignPutRequest

`func NewCampaignsCampaignPutRequest() *CampaignsCampaignPutRequest`

NewCampaignsCampaignPutRequest instantiates a new CampaignsCampaignPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignsCampaignPutRequestWithDefaults

`func NewCampaignsCampaignPutRequestWithDefaults() *CampaignsCampaignPutRequest`

NewCampaignsCampaignPutRequestWithDefaults instantiates a new CampaignsCampaignPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CampaignsCampaignPutRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CampaignsCampaignPutRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CampaignsCampaignPutRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CampaignsCampaignPutRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *CampaignsCampaignPutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignsCampaignPutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignsCampaignPutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignsCampaignPutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSkipType

`func (o *CampaignsCampaignPutRequest) GetSkipType() string`

GetSkipType returns the SkipType field if non-nil, zero value otherwise.

### GetSkipTypeOk

`func (o *CampaignsCampaignPutRequest) GetSkipTypeOk() (*string, bool)`

GetSkipTypeOk returns a tuple with the SkipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipType

`func (o *CampaignsCampaignPutRequest) SetSkipType(v string)`

SetSkipType sets SkipType field to given value.

### HasSkipType

`func (o *CampaignsCampaignPutRequest) HasSkipType() bool`

HasSkipType returns a boolean if a field has been set.

### GetSkipOffset

`func (o *CampaignsCampaignPutRequest) GetSkipOffset() int32`

GetSkipOffset returns the SkipOffset field if non-nil, zero value otherwise.

### GetSkipOffsetOk

`func (o *CampaignsCampaignPutRequest) GetSkipOffsetOk() (*int32, bool)`

GetSkipOffsetOk returns a tuple with the SkipOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOffset

`func (o *CampaignsCampaignPutRequest) SetSkipOffset(v int32)`

SetSkipOffset sets SkipOffset field to given value.

### HasSkipOffset

`func (o *CampaignsCampaignPutRequest) HasSkipOffset() bool`

HasSkipOffset returns a boolean if a field has been set.

### GetActive

`func (o *CampaignsCampaignPutRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CampaignsCampaignPutRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CampaignsCampaignPutRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CampaignsCampaignPutRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


