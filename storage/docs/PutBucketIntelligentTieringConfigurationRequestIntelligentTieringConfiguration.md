# PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**IntelligentTieringId**](IntelligentTieringId.md) |  | 
**Filter** | Pointer to [**IntelligentTieringFilter**](IntelligentTieringFilter.md) |  | [optional] 
**Status** | [**IntelligentTieringStatus**](IntelligentTieringStatus.md) |  | 
**Tierings** | [**TieringList**](TieringList.md) |  | 

## Methods

### NewPutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration

`func NewPutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration(id IntelligentTieringId, status IntelligentTieringStatus, tierings TieringList, ) *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration`

NewPutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration instantiates a new PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfigurationWithDefaults

`func NewPutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfigurationWithDefaults() *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration`

NewPutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfigurationWithDefaults instantiates a new PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) GetId() IntelligentTieringId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) GetIdOk() (*IntelligentTieringId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) SetId(v IntelligentTieringId)`

SetId sets Id field to given value.


### GetFilter

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) GetFilter() IntelligentTieringFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) GetFilterOk() (*IntelligentTieringFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) SetFilter(v IntelligentTieringFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetStatus

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) GetStatus() IntelligentTieringStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) GetStatusOk() (*IntelligentTieringStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) SetStatus(v IntelligentTieringStatus)`

SetStatus sets Status field to given value.


### GetTierings

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) GetTierings() TieringList`

GetTierings returns the Tierings field if non-nil, zero value otherwise.

### GetTieringsOk

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) GetTieringsOk() (*TieringList, bool)`

GetTieringsOk returns a tuple with the Tierings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierings

`func (o *PutBucketIntelligentTieringConfigurationRequestIntelligentTieringConfiguration) SetTierings(v TieringList)`

SetTierings sets Tierings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


