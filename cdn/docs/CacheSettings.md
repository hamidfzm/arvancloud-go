# CacheSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheDeveloperMode** | Pointer to **bool** |  | [optional] 
**CacheConsistentUptime** | Pointer to **bool** |  | [optional] 
**CacheStatus** | Pointer to **string** |  | [optional] 
**CachePage200** | Pointer to **string** |  | [optional] 
**CachePageAny** | Pointer to **string** |  | [optional] 
**CacheBrowser** | Pointer to **string** |  | [optional] 
**CacheScheme** | Pointer to **bool** | To consider scheme (HTTP/HTTPs) in cache | [optional] 
**CacheIgnoreSc** | Pointer to **bool** | Ignore default behavior in caching set-cookie header | [optional] 
**CacheCookie** | Pointer to **string** | Cookie variables to consider in cache (comma separaterd values) | [optional] 
**CacheArgs** | Pointer to **bool** | To consider query args or not | [optional] 
**CacheArg** | Pointer to **string** | Query string arguments to consider in cache (&amp; seperated values) | [optional] [default to ""]

## Methods

### NewCacheSettings

`func NewCacheSettings() *CacheSettings`

NewCacheSettings instantiates a new CacheSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingsWithDefaults

`func NewCacheSettingsWithDefaults() *CacheSettings`

NewCacheSettingsWithDefaults instantiates a new CacheSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheDeveloperMode

`func (o *CacheSettings) GetCacheDeveloperMode() bool`

GetCacheDeveloperMode returns the CacheDeveloperMode field if non-nil, zero value otherwise.

### GetCacheDeveloperModeOk

`func (o *CacheSettings) GetCacheDeveloperModeOk() (*bool, bool)`

GetCacheDeveloperModeOk returns a tuple with the CacheDeveloperMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDeveloperMode

`func (o *CacheSettings) SetCacheDeveloperMode(v bool)`

SetCacheDeveloperMode sets CacheDeveloperMode field to given value.

### HasCacheDeveloperMode

`func (o *CacheSettings) HasCacheDeveloperMode() bool`

HasCacheDeveloperMode returns a boolean if a field has been set.

### GetCacheConsistentUptime

`func (o *CacheSettings) GetCacheConsistentUptime() bool`

GetCacheConsistentUptime returns the CacheConsistentUptime field if non-nil, zero value otherwise.

### GetCacheConsistentUptimeOk

`func (o *CacheSettings) GetCacheConsistentUptimeOk() (*bool, bool)`

GetCacheConsistentUptimeOk returns a tuple with the CacheConsistentUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheConsistentUptime

`func (o *CacheSettings) SetCacheConsistentUptime(v bool)`

SetCacheConsistentUptime sets CacheConsistentUptime field to given value.

### HasCacheConsistentUptime

`func (o *CacheSettings) HasCacheConsistentUptime() bool`

HasCacheConsistentUptime returns a boolean if a field has been set.

### GetCacheStatus

`func (o *CacheSettings) GetCacheStatus() string`

GetCacheStatus returns the CacheStatus field if non-nil, zero value otherwise.

### GetCacheStatusOk

`func (o *CacheSettings) GetCacheStatusOk() (*string, bool)`

GetCacheStatusOk returns a tuple with the CacheStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheStatus

`func (o *CacheSettings) SetCacheStatus(v string)`

SetCacheStatus sets CacheStatus field to given value.

### HasCacheStatus

`func (o *CacheSettings) HasCacheStatus() bool`

HasCacheStatus returns a boolean if a field has been set.

### GetCachePage200

`func (o *CacheSettings) GetCachePage200() string`

GetCachePage200 returns the CachePage200 field if non-nil, zero value otherwise.

### GetCachePage200Ok

`func (o *CacheSettings) GetCachePage200Ok() (*string, bool)`

GetCachePage200Ok returns a tuple with the CachePage200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePage200

`func (o *CacheSettings) SetCachePage200(v string)`

SetCachePage200 sets CachePage200 field to given value.

### HasCachePage200

`func (o *CacheSettings) HasCachePage200() bool`

HasCachePage200 returns a boolean if a field has been set.

### GetCachePageAny

`func (o *CacheSettings) GetCachePageAny() string`

GetCachePageAny returns the CachePageAny field if non-nil, zero value otherwise.

### GetCachePageAnyOk

`func (o *CacheSettings) GetCachePageAnyOk() (*string, bool)`

GetCachePageAnyOk returns a tuple with the CachePageAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePageAny

`func (o *CacheSettings) SetCachePageAny(v string)`

SetCachePageAny sets CachePageAny field to given value.

### HasCachePageAny

`func (o *CacheSettings) HasCachePageAny() bool`

HasCachePageAny returns a boolean if a field has been set.

### GetCacheBrowser

`func (o *CacheSettings) GetCacheBrowser() string`

GetCacheBrowser returns the CacheBrowser field if non-nil, zero value otherwise.

### GetCacheBrowserOk

`func (o *CacheSettings) GetCacheBrowserOk() (*string, bool)`

GetCacheBrowserOk returns a tuple with the CacheBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheBrowser

`func (o *CacheSettings) SetCacheBrowser(v string)`

SetCacheBrowser sets CacheBrowser field to given value.

### HasCacheBrowser

`func (o *CacheSettings) HasCacheBrowser() bool`

HasCacheBrowser returns a boolean if a field has been set.

### GetCacheScheme

`func (o *CacheSettings) GetCacheScheme() bool`

GetCacheScheme returns the CacheScheme field if non-nil, zero value otherwise.

### GetCacheSchemeOk

`func (o *CacheSettings) GetCacheSchemeOk() (*bool, bool)`

GetCacheSchemeOk returns a tuple with the CacheScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheScheme

`func (o *CacheSettings) SetCacheScheme(v bool)`

SetCacheScheme sets CacheScheme field to given value.

### HasCacheScheme

`func (o *CacheSettings) HasCacheScheme() bool`

HasCacheScheme returns a boolean if a field has been set.

### GetCacheIgnoreSc

`func (o *CacheSettings) GetCacheIgnoreSc() bool`

GetCacheIgnoreSc returns the CacheIgnoreSc field if non-nil, zero value otherwise.

### GetCacheIgnoreScOk

`func (o *CacheSettings) GetCacheIgnoreScOk() (*bool, bool)`

GetCacheIgnoreScOk returns a tuple with the CacheIgnoreSc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheIgnoreSc

`func (o *CacheSettings) SetCacheIgnoreSc(v bool)`

SetCacheIgnoreSc sets CacheIgnoreSc field to given value.

### HasCacheIgnoreSc

`func (o *CacheSettings) HasCacheIgnoreSc() bool`

HasCacheIgnoreSc returns a boolean if a field has been set.

### GetCacheCookie

`func (o *CacheSettings) GetCacheCookie() string`

GetCacheCookie returns the CacheCookie field if non-nil, zero value otherwise.

### GetCacheCookieOk

`func (o *CacheSettings) GetCacheCookieOk() (*string, bool)`

GetCacheCookieOk returns a tuple with the CacheCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCookie

`func (o *CacheSettings) SetCacheCookie(v string)`

SetCacheCookie sets CacheCookie field to given value.

### HasCacheCookie

`func (o *CacheSettings) HasCacheCookie() bool`

HasCacheCookie returns a boolean if a field has been set.

### GetCacheArgs

`func (o *CacheSettings) GetCacheArgs() bool`

GetCacheArgs returns the CacheArgs field if non-nil, zero value otherwise.

### GetCacheArgsOk

`func (o *CacheSettings) GetCacheArgsOk() (*bool, bool)`

GetCacheArgsOk returns a tuple with the CacheArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheArgs

`func (o *CacheSettings) SetCacheArgs(v bool)`

SetCacheArgs sets CacheArgs field to given value.

### HasCacheArgs

`func (o *CacheSettings) HasCacheArgs() bool`

HasCacheArgs returns a boolean if a field has been set.

### GetCacheArg

`func (o *CacheSettings) GetCacheArg() string`

GetCacheArg returns the CacheArg field if non-nil, zero value otherwise.

### GetCacheArgOk

`func (o *CacheSettings) GetCacheArgOk() (*string, bool)`

GetCacheArgOk returns a tuple with the CacheArg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheArg

`func (o *CacheSettings) SetCacheArg(v string)`

SetCacheArg sets CacheArg field to given value.

### HasCacheArg

`func (o *CacheSettings) HasCacheArg() bool`

HasCacheArg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


