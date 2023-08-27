# CachingPurge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purge** | **string** | tags is deprecated | 
**PurgeUrls** | Pointer to **[]string** | URLs to be purged from cache. Required if purge value is set to individual. | [optional] 
**PurgeTags** | Pointer to **[]string** | Tags to be purged from cache. Required if purge value is set to tags. Each tag must be 32 characters or less. Only ASCII characters are acceptable.  | [optional] 

## Methods

### NewCachingPurge

`func NewCachingPurge(purge string, ) *CachingPurge`

NewCachingPurge instantiates a new CachingPurge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCachingPurgeWithDefaults

`func NewCachingPurgeWithDefaults() *CachingPurge`

NewCachingPurgeWithDefaults instantiates a new CachingPurge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurge

`func (o *CachingPurge) GetPurge() string`

GetPurge returns the Purge field if non-nil, zero value otherwise.

### GetPurgeOk

`func (o *CachingPurge) GetPurgeOk() (*string, bool)`

GetPurgeOk returns a tuple with the Purge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurge

`func (o *CachingPurge) SetPurge(v string)`

SetPurge sets Purge field to given value.


### GetPurgeUrls

`func (o *CachingPurge) GetPurgeUrls() []string`

GetPurgeUrls returns the PurgeUrls field if non-nil, zero value otherwise.

### GetPurgeUrlsOk

`func (o *CachingPurge) GetPurgeUrlsOk() (*[]string, bool)`

GetPurgeUrlsOk returns a tuple with the PurgeUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeUrls

`func (o *CachingPurge) SetPurgeUrls(v []string)`

SetPurgeUrls sets PurgeUrls field to given value.

### HasPurgeUrls

`func (o *CachingPurge) HasPurgeUrls() bool`

HasPurgeUrls returns a boolean if a field has been set.

### GetPurgeTags

`func (o *CachingPurge) GetPurgeTags() []string`

GetPurgeTags returns the PurgeTags field if non-nil, zero value otherwise.

### GetPurgeTagsOk

`func (o *CachingPurge) GetPurgeTagsOk() (*[]string, bool)`

GetPurgeTagsOk returns a tuple with the PurgeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeTags

`func (o *CachingPurge) SetPurgeTags(v []string)`

SetPurgeTags sets PurgeTags field to given value.

### HasPurgeTags

`func (o *CachingPurge) HasPurgeTags() bool`

HasPurgeTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


