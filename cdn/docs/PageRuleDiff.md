# PageRuleDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL pattern of target pages | [optional] 
**CacheLevel** | Pointer to **string** |  | [optional] 
**WafStatus** | Pointer to **bool** |  | [optional] 
**FwStatus** | Pointer to **bool** | Shows whether firewall is enabled or not | [optional] 
**Acceleration** | Pointer to [**Acceleration**](Acceleration.md) |  | [optional] 
**SlinkStatus** | Pointer to **bool** | Secure link is enabled or not | [optional] 
**Status** | Pointer to **bool** | Is the page-rule enabled? | [optional] 
**Cache200** | Pointer to **string** |  | [optional] 
**CacheAny** | Pointer to **string** |  | [optional] 
**CacheCookie** | Pointer to **string** | Cookie variables to consider in cache (comma separaterd values) | [optional] 
**CacheArgs** | Pointer to **bool** |  | [optional] 
**CacheArg** | Pointer to **string** | Query string arguments to consider in cache (&amp; seperated values) | [optional] [default to ""]
**CacheScheme** | Pointer to **bool** |  | [optional] 
**CacheBrowser** | Pointer to **string** |  | [optional] 
**CacheIgnoreSc** | Pointer to **bool** | Ignore default behavior in caching set-cookie header | [optional] 
**CacheIgnoreVary** | Pointer to **bool** | Ignore default behavior in caching vary header | [optional] 
**CacheIgnoreCc** | Pointer to **bool** |  | [optional] 
**CorsHeader** | Pointer to **string** |  | [optional] 
**RewriteUrl** | Pointer to **string** |  | [optional] 
**SlinkSecret** | Pointer to **string** |  | [optional] 
**SlinkMd5** | Pointer to **[]string** |  | [optional] 
**ClusterStatus** | Pointer to **bool** |  | [optional] 
**ClusterId** | Pointer to **NullableString** |  | [optional] 
**UpstreamTimeout** | Pointer to [**UpstreamTimeout**](UpstreamTimeout.md) |  | [optional] 
**ReqCustomHeaders** | Pointer to [**[]PageRuleDiffReqCustomHeadersInner**](PageRuleDiffReqCustomHeadersInner.md) |  | [optional] [default to []]
**ResCustomHeaders** | Pointer to [**[]PageRuleDiffReqCustomHeadersInner**](PageRuleDiffReqCustomHeadersInner.md) |  | [optional] [default to []]
**ReqHideHeaders** | Pointer to **[]string** |  | [optional] [default to []]
**ResHideHeaders** | Pointer to **[]string** |  | [optional] 
**CustomHostHeader** | Pointer to **string** |  | [optional] 
**Redirect** | Pointer to [**PageRuleDiffRedirect**](PageRuleDiffRedirect.md) |  | [optional] 

## Methods

### NewPageRuleDiff

`func NewPageRuleDiff() *PageRuleDiff`

NewPageRuleDiff instantiates a new PageRuleDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageRuleDiffWithDefaults

`func NewPageRuleDiffWithDefaults() *PageRuleDiff`

NewPageRuleDiffWithDefaults instantiates a new PageRuleDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PageRuleDiff) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PageRuleDiff) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PageRuleDiff) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PageRuleDiff) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCacheLevel

`func (o *PageRuleDiff) GetCacheLevel() string`

GetCacheLevel returns the CacheLevel field if non-nil, zero value otherwise.

### GetCacheLevelOk

`func (o *PageRuleDiff) GetCacheLevelOk() (*string, bool)`

GetCacheLevelOk returns a tuple with the CacheLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLevel

`func (o *PageRuleDiff) SetCacheLevel(v string)`

SetCacheLevel sets CacheLevel field to given value.

### HasCacheLevel

`func (o *PageRuleDiff) HasCacheLevel() bool`

HasCacheLevel returns a boolean if a field has been set.

### GetWafStatus

`func (o *PageRuleDiff) GetWafStatus() bool`

GetWafStatus returns the WafStatus field if non-nil, zero value otherwise.

### GetWafStatusOk

`func (o *PageRuleDiff) GetWafStatusOk() (*bool, bool)`

GetWafStatusOk returns a tuple with the WafStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus

`func (o *PageRuleDiff) SetWafStatus(v bool)`

SetWafStatus sets WafStatus field to given value.

### HasWafStatus

`func (o *PageRuleDiff) HasWafStatus() bool`

HasWafStatus returns a boolean if a field has been set.

### GetFwStatus

`func (o *PageRuleDiff) GetFwStatus() bool`

GetFwStatus returns the FwStatus field if non-nil, zero value otherwise.

### GetFwStatusOk

`func (o *PageRuleDiff) GetFwStatusOk() (*bool, bool)`

GetFwStatusOk returns a tuple with the FwStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwStatus

`func (o *PageRuleDiff) SetFwStatus(v bool)`

SetFwStatus sets FwStatus field to given value.

### HasFwStatus

`func (o *PageRuleDiff) HasFwStatus() bool`

HasFwStatus returns a boolean if a field has been set.

### GetAcceleration

`func (o *PageRuleDiff) GetAcceleration() Acceleration`

GetAcceleration returns the Acceleration field if non-nil, zero value otherwise.

### GetAccelerationOk

`func (o *PageRuleDiff) GetAccelerationOk() (*Acceleration, bool)`

GetAccelerationOk returns a tuple with the Acceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleration

`func (o *PageRuleDiff) SetAcceleration(v Acceleration)`

SetAcceleration sets Acceleration field to given value.

### HasAcceleration

`func (o *PageRuleDiff) HasAcceleration() bool`

HasAcceleration returns a boolean if a field has been set.

### GetSlinkStatus

`func (o *PageRuleDiff) GetSlinkStatus() bool`

GetSlinkStatus returns the SlinkStatus field if non-nil, zero value otherwise.

### GetSlinkStatusOk

`func (o *PageRuleDiff) GetSlinkStatusOk() (*bool, bool)`

GetSlinkStatusOk returns a tuple with the SlinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlinkStatus

`func (o *PageRuleDiff) SetSlinkStatus(v bool)`

SetSlinkStatus sets SlinkStatus field to given value.

### HasSlinkStatus

`func (o *PageRuleDiff) HasSlinkStatus() bool`

HasSlinkStatus returns a boolean if a field has been set.

### GetStatus

`func (o *PageRuleDiff) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PageRuleDiff) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PageRuleDiff) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PageRuleDiff) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCache200

`func (o *PageRuleDiff) GetCache200() string`

GetCache200 returns the Cache200 field if non-nil, zero value otherwise.

### GetCache200Ok

`func (o *PageRuleDiff) GetCache200Ok() (*string, bool)`

GetCache200Ok returns a tuple with the Cache200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache200

`func (o *PageRuleDiff) SetCache200(v string)`

SetCache200 sets Cache200 field to given value.

### HasCache200

`func (o *PageRuleDiff) HasCache200() bool`

HasCache200 returns a boolean if a field has been set.

### GetCacheAny

`func (o *PageRuleDiff) GetCacheAny() string`

GetCacheAny returns the CacheAny field if non-nil, zero value otherwise.

### GetCacheAnyOk

`func (o *PageRuleDiff) GetCacheAnyOk() (*string, bool)`

GetCacheAnyOk returns a tuple with the CacheAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAny

`func (o *PageRuleDiff) SetCacheAny(v string)`

SetCacheAny sets CacheAny field to given value.

### HasCacheAny

`func (o *PageRuleDiff) HasCacheAny() bool`

HasCacheAny returns a boolean if a field has been set.

### GetCacheCookie

`func (o *PageRuleDiff) GetCacheCookie() string`

GetCacheCookie returns the CacheCookie field if non-nil, zero value otherwise.

### GetCacheCookieOk

`func (o *PageRuleDiff) GetCacheCookieOk() (*string, bool)`

GetCacheCookieOk returns a tuple with the CacheCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCookie

`func (o *PageRuleDiff) SetCacheCookie(v string)`

SetCacheCookie sets CacheCookie field to given value.

### HasCacheCookie

`func (o *PageRuleDiff) HasCacheCookie() bool`

HasCacheCookie returns a boolean if a field has been set.

### GetCacheArgs

`func (o *PageRuleDiff) GetCacheArgs() bool`

GetCacheArgs returns the CacheArgs field if non-nil, zero value otherwise.

### GetCacheArgsOk

`func (o *PageRuleDiff) GetCacheArgsOk() (*bool, bool)`

GetCacheArgsOk returns a tuple with the CacheArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheArgs

`func (o *PageRuleDiff) SetCacheArgs(v bool)`

SetCacheArgs sets CacheArgs field to given value.

### HasCacheArgs

`func (o *PageRuleDiff) HasCacheArgs() bool`

HasCacheArgs returns a boolean if a field has been set.

### GetCacheArg

`func (o *PageRuleDiff) GetCacheArg() string`

GetCacheArg returns the CacheArg field if non-nil, zero value otherwise.

### GetCacheArgOk

`func (o *PageRuleDiff) GetCacheArgOk() (*string, bool)`

GetCacheArgOk returns a tuple with the CacheArg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheArg

`func (o *PageRuleDiff) SetCacheArg(v string)`

SetCacheArg sets CacheArg field to given value.

### HasCacheArg

`func (o *PageRuleDiff) HasCacheArg() bool`

HasCacheArg returns a boolean if a field has been set.

### GetCacheScheme

`func (o *PageRuleDiff) GetCacheScheme() bool`

GetCacheScheme returns the CacheScheme field if non-nil, zero value otherwise.

### GetCacheSchemeOk

`func (o *PageRuleDiff) GetCacheSchemeOk() (*bool, bool)`

GetCacheSchemeOk returns a tuple with the CacheScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheScheme

`func (o *PageRuleDiff) SetCacheScheme(v bool)`

SetCacheScheme sets CacheScheme field to given value.

### HasCacheScheme

`func (o *PageRuleDiff) HasCacheScheme() bool`

HasCacheScheme returns a boolean if a field has been set.

### GetCacheBrowser

`func (o *PageRuleDiff) GetCacheBrowser() string`

GetCacheBrowser returns the CacheBrowser field if non-nil, zero value otherwise.

### GetCacheBrowserOk

`func (o *PageRuleDiff) GetCacheBrowserOk() (*string, bool)`

GetCacheBrowserOk returns a tuple with the CacheBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheBrowser

`func (o *PageRuleDiff) SetCacheBrowser(v string)`

SetCacheBrowser sets CacheBrowser field to given value.

### HasCacheBrowser

`func (o *PageRuleDiff) HasCacheBrowser() bool`

HasCacheBrowser returns a boolean if a field has been set.

### GetCacheIgnoreSc

`func (o *PageRuleDiff) GetCacheIgnoreSc() bool`

GetCacheIgnoreSc returns the CacheIgnoreSc field if non-nil, zero value otherwise.

### GetCacheIgnoreScOk

`func (o *PageRuleDiff) GetCacheIgnoreScOk() (*bool, bool)`

GetCacheIgnoreScOk returns a tuple with the CacheIgnoreSc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheIgnoreSc

`func (o *PageRuleDiff) SetCacheIgnoreSc(v bool)`

SetCacheIgnoreSc sets CacheIgnoreSc field to given value.

### HasCacheIgnoreSc

`func (o *PageRuleDiff) HasCacheIgnoreSc() bool`

HasCacheIgnoreSc returns a boolean if a field has been set.

### GetCacheIgnoreVary

`func (o *PageRuleDiff) GetCacheIgnoreVary() bool`

GetCacheIgnoreVary returns the CacheIgnoreVary field if non-nil, zero value otherwise.

### GetCacheIgnoreVaryOk

`func (o *PageRuleDiff) GetCacheIgnoreVaryOk() (*bool, bool)`

GetCacheIgnoreVaryOk returns a tuple with the CacheIgnoreVary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheIgnoreVary

`func (o *PageRuleDiff) SetCacheIgnoreVary(v bool)`

SetCacheIgnoreVary sets CacheIgnoreVary field to given value.

### HasCacheIgnoreVary

`func (o *PageRuleDiff) HasCacheIgnoreVary() bool`

HasCacheIgnoreVary returns a boolean if a field has been set.

### GetCacheIgnoreCc

`func (o *PageRuleDiff) GetCacheIgnoreCc() bool`

GetCacheIgnoreCc returns the CacheIgnoreCc field if non-nil, zero value otherwise.

### GetCacheIgnoreCcOk

`func (o *PageRuleDiff) GetCacheIgnoreCcOk() (*bool, bool)`

GetCacheIgnoreCcOk returns a tuple with the CacheIgnoreCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheIgnoreCc

`func (o *PageRuleDiff) SetCacheIgnoreCc(v bool)`

SetCacheIgnoreCc sets CacheIgnoreCc field to given value.

### HasCacheIgnoreCc

`func (o *PageRuleDiff) HasCacheIgnoreCc() bool`

HasCacheIgnoreCc returns a boolean if a field has been set.

### GetCorsHeader

`func (o *PageRuleDiff) GetCorsHeader() string`

GetCorsHeader returns the CorsHeader field if non-nil, zero value otherwise.

### GetCorsHeaderOk

`func (o *PageRuleDiff) GetCorsHeaderOk() (*string, bool)`

GetCorsHeaderOk returns a tuple with the CorsHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsHeader

`func (o *PageRuleDiff) SetCorsHeader(v string)`

SetCorsHeader sets CorsHeader field to given value.

### HasCorsHeader

`func (o *PageRuleDiff) HasCorsHeader() bool`

HasCorsHeader returns a boolean if a field has been set.

### GetRewriteUrl

`func (o *PageRuleDiff) GetRewriteUrl() string`

GetRewriteUrl returns the RewriteUrl field if non-nil, zero value otherwise.

### GetRewriteUrlOk

`func (o *PageRuleDiff) GetRewriteUrlOk() (*string, bool)`

GetRewriteUrlOk returns a tuple with the RewriteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteUrl

`func (o *PageRuleDiff) SetRewriteUrl(v string)`

SetRewriteUrl sets RewriteUrl field to given value.

### HasRewriteUrl

`func (o *PageRuleDiff) HasRewriteUrl() bool`

HasRewriteUrl returns a boolean if a field has been set.

### GetSlinkSecret

`func (o *PageRuleDiff) GetSlinkSecret() string`

GetSlinkSecret returns the SlinkSecret field if non-nil, zero value otherwise.

### GetSlinkSecretOk

`func (o *PageRuleDiff) GetSlinkSecretOk() (*string, bool)`

GetSlinkSecretOk returns a tuple with the SlinkSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlinkSecret

`func (o *PageRuleDiff) SetSlinkSecret(v string)`

SetSlinkSecret sets SlinkSecret field to given value.

### HasSlinkSecret

`func (o *PageRuleDiff) HasSlinkSecret() bool`

HasSlinkSecret returns a boolean if a field has been set.

### GetSlinkMd5

`func (o *PageRuleDiff) GetSlinkMd5() []string`

GetSlinkMd5 returns the SlinkMd5 field if non-nil, zero value otherwise.

### GetSlinkMd5Ok

`func (o *PageRuleDiff) GetSlinkMd5Ok() (*[]string, bool)`

GetSlinkMd5Ok returns a tuple with the SlinkMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlinkMd5

`func (o *PageRuleDiff) SetSlinkMd5(v []string)`

SetSlinkMd5 sets SlinkMd5 field to given value.

### HasSlinkMd5

`func (o *PageRuleDiff) HasSlinkMd5() bool`

HasSlinkMd5 returns a boolean if a field has been set.

### SetSlinkMd5Nil

`func (o *PageRuleDiff) SetSlinkMd5Nil(b bool)`

 SetSlinkMd5Nil sets the value for SlinkMd5 to be an explicit nil

### UnsetSlinkMd5
`func (o *PageRuleDiff) UnsetSlinkMd5()`

UnsetSlinkMd5 ensures that no value is present for SlinkMd5, not even an explicit nil
### GetClusterStatus

`func (o *PageRuleDiff) GetClusterStatus() bool`

GetClusterStatus returns the ClusterStatus field if non-nil, zero value otherwise.

### GetClusterStatusOk

`func (o *PageRuleDiff) GetClusterStatusOk() (*bool, bool)`

GetClusterStatusOk returns a tuple with the ClusterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStatus

`func (o *PageRuleDiff) SetClusterStatus(v bool)`

SetClusterStatus sets ClusterStatus field to given value.

### HasClusterStatus

`func (o *PageRuleDiff) HasClusterStatus() bool`

HasClusterStatus returns a boolean if a field has been set.

### GetClusterId

`func (o *PageRuleDiff) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PageRuleDiff) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PageRuleDiff) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *PageRuleDiff) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *PageRuleDiff) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *PageRuleDiff) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetUpstreamTimeout

`func (o *PageRuleDiff) GetUpstreamTimeout() UpstreamTimeout`

GetUpstreamTimeout returns the UpstreamTimeout field if non-nil, zero value otherwise.

### GetUpstreamTimeoutOk

`func (o *PageRuleDiff) GetUpstreamTimeoutOk() (*UpstreamTimeout, bool)`

GetUpstreamTimeoutOk returns a tuple with the UpstreamTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamTimeout

`func (o *PageRuleDiff) SetUpstreamTimeout(v UpstreamTimeout)`

SetUpstreamTimeout sets UpstreamTimeout field to given value.

### HasUpstreamTimeout

`func (o *PageRuleDiff) HasUpstreamTimeout() bool`

HasUpstreamTimeout returns a boolean if a field has been set.

### GetReqCustomHeaders

`func (o *PageRuleDiff) GetReqCustomHeaders() []PageRuleDiffReqCustomHeadersInner`

GetReqCustomHeaders returns the ReqCustomHeaders field if non-nil, zero value otherwise.

### GetReqCustomHeadersOk

`func (o *PageRuleDiff) GetReqCustomHeadersOk() (*[]PageRuleDiffReqCustomHeadersInner, bool)`

GetReqCustomHeadersOk returns a tuple with the ReqCustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqCustomHeaders

`func (o *PageRuleDiff) SetReqCustomHeaders(v []PageRuleDiffReqCustomHeadersInner)`

SetReqCustomHeaders sets ReqCustomHeaders field to given value.

### HasReqCustomHeaders

`func (o *PageRuleDiff) HasReqCustomHeaders() bool`

HasReqCustomHeaders returns a boolean if a field has been set.

### GetResCustomHeaders

`func (o *PageRuleDiff) GetResCustomHeaders() []PageRuleDiffReqCustomHeadersInner`

GetResCustomHeaders returns the ResCustomHeaders field if non-nil, zero value otherwise.

### GetResCustomHeadersOk

`func (o *PageRuleDiff) GetResCustomHeadersOk() (*[]PageRuleDiffReqCustomHeadersInner, bool)`

GetResCustomHeadersOk returns a tuple with the ResCustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResCustomHeaders

`func (o *PageRuleDiff) SetResCustomHeaders(v []PageRuleDiffReqCustomHeadersInner)`

SetResCustomHeaders sets ResCustomHeaders field to given value.

### HasResCustomHeaders

`func (o *PageRuleDiff) HasResCustomHeaders() bool`

HasResCustomHeaders returns a boolean if a field has been set.

### GetReqHideHeaders

`func (o *PageRuleDiff) GetReqHideHeaders() []string`

GetReqHideHeaders returns the ReqHideHeaders field if non-nil, zero value otherwise.

### GetReqHideHeadersOk

`func (o *PageRuleDiff) GetReqHideHeadersOk() (*[]string, bool)`

GetReqHideHeadersOk returns a tuple with the ReqHideHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqHideHeaders

`func (o *PageRuleDiff) SetReqHideHeaders(v []string)`

SetReqHideHeaders sets ReqHideHeaders field to given value.

### HasReqHideHeaders

`func (o *PageRuleDiff) HasReqHideHeaders() bool`

HasReqHideHeaders returns a boolean if a field has been set.

### GetResHideHeaders

`func (o *PageRuleDiff) GetResHideHeaders() []string`

GetResHideHeaders returns the ResHideHeaders field if non-nil, zero value otherwise.

### GetResHideHeadersOk

`func (o *PageRuleDiff) GetResHideHeadersOk() (*[]string, bool)`

GetResHideHeadersOk returns a tuple with the ResHideHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResHideHeaders

`func (o *PageRuleDiff) SetResHideHeaders(v []string)`

SetResHideHeaders sets ResHideHeaders field to given value.

### HasResHideHeaders

`func (o *PageRuleDiff) HasResHideHeaders() bool`

HasResHideHeaders returns a boolean if a field has been set.

### GetCustomHostHeader

`func (o *PageRuleDiff) GetCustomHostHeader() string`

GetCustomHostHeader returns the CustomHostHeader field if non-nil, zero value otherwise.

### GetCustomHostHeaderOk

`func (o *PageRuleDiff) GetCustomHostHeaderOk() (*string, bool)`

GetCustomHostHeaderOk returns a tuple with the CustomHostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHostHeader

`func (o *PageRuleDiff) SetCustomHostHeader(v string)`

SetCustomHostHeader sets CustomHostHeader field to given value.

### HasCustomHostHeader

`func (o *PageRuleDiff) HasCustomHostHeader() bool`

HasCustomHostHeader returns a boolean if a field has been set.

### GetRedirect

`func (o *PageRuleDiff) GetRedirect() PageRuleDiffRedirect`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *PageRuleDiff) GetRedirectOk() (*PageRuleDiffRedirect, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *PageRuleDiff) SetRedirect(v PageRuleDiffRedirect)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *PageRuleDiff) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


