# PageRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache200** | Pointer to **string** |  | [optional] [default to "30m"]
**CacheAny** | Pointer to **string** |  | [optional] [default to "0s"]
**CacheCookie** | Pointer to **string** | Cookie variables to consider in cache (comma separaterd values) | [optional] [default to ""]
**CacheDeviceType** | Pointer to **bool** |  | [optional] [default to false]
**CacheArgs** | Pointer to **bool** |  | [optional] [default to true]
**CacheArg** | Pointer to **string** | Query string arguments to consider in cache (&amp; seperated values) | [optional] [default to ""]
**CacheScheme** | Pointer to **bool** |  | [optional] [default to true]
**CacheBrowser** | Pointer to **string** |  | [optional] [default to "default"]
**CacheIgnoreSc** | Pointer to **bool** | Ignore default behavior in caching set-cookie header | [optional] [default to false]
**CacheIgnoreVary** | Pointer to **bool** | Ignore default behavior in caching vary header | [optional] [default to true]
**CacheIgnoreCc** | Pointer to **bool** |  | [optional] [default to true]
**CorsHeader** | Pointer to **string** |  | [optional] [default to "-"]
**RewriteUrl** | Pointer to **string** |  | [optional] [default to "-"]
**SlinkSecret** | Pointer to **string** |  | [optional] [default to ""]
**SlinkMd5** | Pointer to **[]string** |  | [optional] [default to ["remote_addr","file","expires"]]
**LoadBalancer** | Pointer to **NullableString** | Name or ID of the load balancer | [optional] 
**ClusterStatus** | Pointer to **bool** |  | [optional] [default to false]
**ImageResize** | Pointer to [**PageRuleImageResize**](PageRuleImageResize.md) |  | [optional] 
**ClusterId** | Pointer to **NullableString** |  | [optional] 
**UpstreamTimeout** | Pointer to [**UpstreamTimeout**](UpstreamTimeout.md) |  | [optional] 
**ReqCustomHeaders** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**ResCustomHeaders** | Pointer to **[]map[string]interface{}** |  | [optional] [default to []]
**ReqHideHeaders** | Pointer to **[]string** |  | [optional] [default to []]
**ResHideHeaders** | Pointer to **[]string** |  | [optional] [default to []]
**CustomHostHeader** | Pointer to **string** |  | [optional] [default to ""]
**Redirect** | Pointer to [**PageRuleRedirect**](PageRuleRedirect.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**DomainId** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int32** | Order of the page-rule | [optional] 
**UrlType** | Pointer to **string** | This flag is deprecated in favor of is_protected flag | [optional] [default to "default"]
**IsProtected** | Pointer to **bool** | Protected records cannot be modified or deleted by user. | [optional] [readonly] [default to false]
**Url** | Pointer to **string** | URL pattern of target pages | [optional] 
**CacheLevel** | Pointer to **string** |  | [optional] [default to "query_string"]
**WafStatus** | Pointer to **bool** |  | [optional] [default to true]
**FwStatus** | Pointer to **bool** | Shows whether firewall is enabled or not | [optional] [default to true]
**Acceleration** | Pointer to [**Acceleration**](Acceleration.md) |  | [optional] 
**SlinkStatus** | Pointer to **bool** | Secure link is enabled or not | [optional] [default to false]
**Status** | Pointer to **bool** | Is the page-rule enabled? | [optional] [default to true]
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewPageRule

`func NewPageRule() *PageRule`

NewPageRule instantiates a new PageRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageRuleWithDefaults

`func NewPageRuleWithDefaults() *PageRule`

NewPageRuleWithDefaults instantiates a new PageRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache200

`func (o *PageRule) GetCache200() string`

GetCache200 returns the Cache200 field if non-nil, zero value otherwise.

### GetCache200Ok

`func (o *PageRule) GetCache200Ok() (*string, bool)`

GetCache200Ok returns a tuple with the Cache200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache200

`func (o *PageRule) SetCache200(v string)`

SetCache200 sets Cache200 field to given value.

### HasCache200

`func (o *PageRule) HasCache200() bool`

HasCache200 returns a boolean if a field has been set.

### GetCacheAny

`func (o *PageRule) GetCacheAny() string`

GetCacheAny returns the CacheAny field if non-nil, zero value otherwise.

### GetCacheAnyOk

`func (o *PageRule) GetCacheAnyOk() (*string, bool)`

GetCacheAnyOk returns a tuple with the CacheAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheAny

`func (o *PageRule) SetCacheAny(v string)`

SetCacheAny sets CacheAny field to given value.

### HasCacheAny

`func (o *PageRule) HasCacheAny() bool`

HasCacheAny returns a boolean if a field has been set.

### GetCacheCookie

`func (o *PageRule) GetCacheCookie() string`

GetCacheCookie returns the CacheCookie field if non-nil, zero value otherwise.

### GetCacheCookieOk

`func (o *PageRule) GetCacheCookieOk() (*string, bool)`

GetCacheCookieOk returns a tuple with the CacheCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCookie

`func (o *PageRule) SetCacheCookie(v string)`

SetCacheCookie sets CacheCookie field to given value.

### HasCacheCookie

`func (o *PageRule) HasCacheCookie() bool`

HasCacheCookie returns a boolean if a field has been set.

### GetCacheDeviceType

`func (o *PageRule) GetCacheDeviceType() bool`

GetCacheDeviceType returns the CacheDeviceType field if non-nil, zero value otherwise.

### GetCacheDeviceTypeOk

`func (o *PageRule) GetCacheDeviceTypeOk() (*bool, bool)`

GetCacheDeviceTypeOk returns a tuple with the CacheDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDeviceType

`func (o *PageRule) SetCacheDeviceType(v bool)`

SetCacheDeviceType sets CacheDeviceType field to given value.

### HasCacheDeviceType

`func (o *PageRule) HasCacheDeviceType() bool`

HasCacheDeviceType returns a boolean if a field has been set.

### GetCacheArgs

`func (o *PageRule) GetCacheArgs() bool`

GetCacheArgs returns the CacheArgs field if non-nil, zero value otherwise.

### GetCacheArgsOk

`func (o *PageRule) GetCacheArgsOk() (*bool, bool)`

GetCacheArgsOk returns a tuple with the CacheArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheArgs

`func (o *PageRule) SetCacheArgs(v bool)`

SetCacheArgs sets CacheArgs field to given value.

### HasCacheArgs

`func (o *PageRule) HasCacheArgs() bool`

HasCacheArgs returns a boolean if a field has been set.

### GetCacheArg

`func (o *PageRule) GetCacheArg() string`

GetCacheArg returns the CacheArg field if non-nil, zero value otherwise.

### GetCacheArgOk

`func (o *PageRule) GetCacheArgOk() (*string, bool)`

GetCacheArgOk returns a tuple with the CacheArg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheArg

`func (o *PageRule) SetCacheArg(v string)`

SetCacheArg sets CacheArg field to given value.

### HasCacheArg

`func (o *PageRule) HasCacheArg() bool`

HasCacheArg returns a boolean if a field has been set.

### GetCacheScheme

`func (o *PageRule) GetCacheScheme() bool`

GetCacheScheme returns the CacheScheme field if non-nil, zero value otherwise.

### GetCacheSchemeOk

`func (o *PageRule) GetCacheSchemeOk() (*bool, bool)`

GetCacheSchemeOk returns a tuple with the CacheScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheScheme

`func (o *PageRule) SetCacheScheme(v bool)`

SetCacheScheme sets CacheScheme field to given value.

### HasCacheScheme

`func (o *PageRule) HasCacheScheme() bool`

HasCacheScheme returns a boolean if a field has been set.

### GetCacheBrowser

`func (o *PageRule) GetCacheBrowser() string`

GetCacheBrowser returns the CacheBrowser field if non-nil, zero value otherwise.

### GetCacheBrowserOk

`func (o *PageRule) GetCacheBrowserOk() (*string, bool)`

GetCacheBrowserOk returns a tuple with the CacheBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheBrowser

`func (o *PageRule) SetCacheBrowser(v string)`

SetCacheBrowser sets CacheBrowser field to given value.

### HasCacheBrowser

`func (o *PageRule) HasCacheBrowser() bool`

HasCacheBrowser returns a boolean if a field has been set.

### GetCacheIgnoreSc

`func (o *PageRule) GetCacheIgnoreSc() bool`

GetCacheIgnoreSc returns the CacheIgnoreSc field if non-nil, zero value otherwise.

### GetCacheIgnoreScOk

`func (o *PageRule) GetCacheIgnoreScOk() (*bool, bool)`

GetCacheIgnoreScOk returns a tuple with the CacheIgnoreSc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheIgnoreSc

`func (o *PageRule) SetCacheIgnoreSc(v bool)`

SetCacheIgnoreSc sets CacheIgnoreSc field to given value.

### HasCacheIgnoreSc

`func (o *PageRule) HasCacheIgnoreSc() bool`

HasCacheIgnoreSc returns a boolean if a field has been set.

### GetCacheIgnoreVary

`func (o *PageRule) GetCacheIgnoreVary() bool`

GetCacheIgnoreVary returns the CacheIgnoreVary field if non-nil, zero value otherwise.

### GetCacheIgnoreVaryOk

`func (o *PageRule) GetCacheIgnoreVaryOk() (*bool, bool)`

GetCacheIgnoreVaryOk returns a tuple with the CacheIgnoreVary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheIgnoreVary

`func (o *PageRule) SetCacheIgnoreVary(v bool)`

SetCacheIgnoreVary sets CacheIgnoreVary field to given value.

### HasCacheIgnoreVary

`func (o *PageRule) HasCacheIgnoreVary() bool`

HasCacheIgnoreVary returns a boolean if a field has been set.

### GetCacheIgnoreCc

`func (o *PageRule) GetCacheIgnoreCc() bool`

GetCacheIgnoreCc returns the CacheIgnoreCc field if non-nil, zero value otherwise.

### GetCacheIgnoreCcOk

`func (o *PageRule) GetCacheIgnoreCcOk() (*bool, bool)`

GetCacheIgnoreCcOk returns a tuple with the CacheIgnoreCc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheIgnoreCc

`func (o *PageRule) SetCacheIgnoreCc(v bool)`

SetCacheIgnoreCc sets CacheIgnoreCc field to given value.

### HasCacheIgnoreCc

`func (o *PageRule) HasCacheIgnoreCc() bool`

HasCacheIgnoreCc returns a boolean if a field has been set.

### GetCorsHeader

`func (o *PageRule) GetCorsHeader() string`

GetCorsHeader returns the CorsHeader field if non-nil, zero value otherwise.

### GetCorsHeaderOk

`func (o *PageRule) GetCorsHeaderOk() (*string, bool)`

GetCorsHeaderOk returns a tuple with the CorsHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsHeader

`func (o *PageRule) SetCorsHeader(v string)`

SetCorsHeader sets CorsHeader field to given value.

### HasCorsHeader

`func (o *PageRule) HasCorsHeader() bool`

HasCorsHeader returns a boolean if a field has been set.

### GetRewriteUrl

`func (o *PageRule) GetRewriteUrl() string`

GetRewriteUrl returns the RewriteUrl field if non-nil, zero value otherwise.

### GetRewriteUrlOk

`func (o *PageRule) GetRewriteUrlOk() (*string, bool)`

GetRewriteUrlOk returns a tuple with the RewriteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteUrl

`func (o *PageRule) SetRewriteUrl(v string)`

SetRewriteUrl sets RewriteUrl field to given value.

### HasRewriteUrl

`func (o *PageRule) HasRewriteUrl() bool`

HasRewriteUrl returns a boolean if a field has been set.

### GetSlinkSecret

`func (o *PageRule) GetSlinkSecret() string`

GetSlinkSecret returns the SlinkSecret field if non-nil, zero value otherwise.

### GetSlinkSecretOk

`func (o *PageRule) GetSlinkSecretOk() (*string, bool)`

GetSlinkSecretOk returns a tuple with the SlinkSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlinkSecret

`func (o *PageRule) SetSlinkSecret(v string)`

SetSlinkSecret sets SlinkSecret field to given value.

### HasSlinkSecret

`func (o *PageRule) HasSlinkSecret() bool`

HasSlinkSecret returns a boolean if a field has been set.

### GetSlinkMd5

`func (o *PageRule) GetSlinkMd5() []string`

GetSlinkMd5 returns the SlinkMd5 field if non-nil, zero value otherwise.

### GetSlinkMd5Ok

`func (o *PageRule) GetSlinkMd5Ok() (*[]string, bool)`

GetSlinkMd5Ok returns a tuple with the SlinkMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlinkMd5

`func (o *PageRule) SetSlinkMd5(v []string)`

SetSlinkMd5 sets SlinkMd5 field to given value.

### HasSlinkMd5

`func (o *PageRule) HasSlinkMd5() bool`

HasSlinkMd5 returns a boolean if a field has been set.

### SetSlinkMd5Nil

`func (o *PageRule) SetSlinkMd5Nil(b bool)`

 SetSlinkMd5Nil sets the value for SlinkMd5 to be an explicit nil

### UnsetSlinkMd5
`func (o *PageRule) UnsetSlinkMd5()`

UnsetSlinkMd5 ensures that no value is present for SlinkMd5, not even an explicit nil
### GetLoadBalancer

`func (o *PageRule) GetLoadBalancer() string`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *PageRule) GetLoadBalancerOk() (*string, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *PageRule) SetLoadBalancer(v string)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *PageRule) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### SetLoadBalancerNil

`func (o *PageRule) SetLoadBalancerNil(b bool)`

 SetLoadBalancerNil sets the value for LoadBalancer to be an explicit nil

### UnsetLoadBalancer
`func (o *PageRule) UnsetLoadBalancer()`

UnsetLoadBalancer ensures that no value is present for LoadBalancer, not even an explicit nil
### GetClusterStatus

`func (o *PageRule) GetClusterStatus() bool`

GetClusterStatus returns the ClusterStatus field if non-nil, zero value otherwise.

### GetClusterStatusOk

`func (o *PageRule) GetClusterStatusOk() (*bool, bool)`

GetClusterStatusOk returns a tuple with the ClusterStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStatus

`func (o *PageRule) SetClusterStatus(v bool)`

SetClusterStatus sets ClusterStatus field to given value.

### HasClusterStatus

`func (o *PageRule) HasClusterStatus() bool`

HasClusterStatus returns a boolean if a field has been set.

### GetImageResize

`func (o *PageRule) GetImageResize() PageRuleImageResize`

GetImageResize returns the ImageResize field if non-nil, zero value otherwise.

### GetImageResizeOk

`func (o *PageRule) GetImageResizeOk() (*PageRuleImageResize, bool)`

GetImageResizeOk returns a tuple with the ImageResize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageResize

`func (o *PageRule) SetImageResize(v PageRuleImageResize)`

SetImageResize sets ImageResize field to given value.

### HasImageResize

`func (o *PageRule) HasImageResize() bool`

HasImageResize returns a boolean if a field has been set.

### GetClusterId

`func (o *PageRule) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PageRule) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PageRule) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *PageRule) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### SetClusterIdNil

`func (o *PageRule) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *PageRule) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetUpstreamTimeout

`func (o *PageRule) GetUpstreamTimeout() UpstreamTimeout`

GetUpstreamTimeout returns the UpstreamTimeout field if non-nil, zero value otherwise.

### GetUpstreamTimeoutOk

`func (o *PageRule) GetUpstreamTimeoutOk() (*UpstreamTimeout, bool)`

GetUpstreamTimeoutOk returns a tuple with the UpstreamTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamTimeout

`func (o *PageRule) SetUpstreamTimeout(v UpstreamTimeout)`

SetUpstreamTimeout sets UpstreamTimeout field to given value.

### HasUpstreamTimeout

`func (o *PageRule) HasUpstreamTimeout() bool`

HasUpstreamTimeout returns a boolean if a field has been set.

### GetReqCustomHeaders

`func (o *PageRule) GetReqCustomHeaders() []map[string]interface{}`

GetReqCustomHeaders returns the ReqCustomHeaders field if non-nil, zero value otherwise.

### GetReqCustomHeadersOk

`func (o *PageRule) GetReqCustomHeadersOk() (*[]map[string]interface{}, bool)`

GetReqCustomHeadersOk returns a tuple with the ReqCustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqCustomHeaders

`func (o *PageRule) SetReqCustomHeaders(v []map[string]interface{})`

SetReqCustomHeaders sets ReqCustomHeaders field to given value.

### HasReqCustomHeaders

`func (o *PageRule) HasReqCustomHeaders() bool`

HasReqCustomHeaders returns a boolean if a field has been set.

### GetResCustomHeaders

`func (o *PageRule) GetResCustomHeaders() []map[string]interface{}`

GetResCustomHeaders returns the ResCustomHeaders field if non-nil, zero value otherwise.

### GetResCustomHeadersOk

`func (o *PageRule) GetResCustomHeadersOk() (*[]map[string]interface{}, bool)`

GetResCustomHeadersOk returns a tuple with the ResCustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResCustomHeaders

`func (o *PageRule) SetResCustomHeaders(v []map[string]interface{})`

SetResCustomHeaders sets ResCustomHeaders field to given value.

### HasResCustomHeaders

`func (o *PageRule) HasResCustomHeaders() bool`

HasResCustomHeaders returns a boolean if a field has been set.

### GetReqHideHeaders

`func (o *PageRule) GetReqHideHeaders() []string`

GetReqHideHeaders returns the ReqHideHeaders field if non-nil, zero value otherwise.

### GetReqHideHeadersOk

`func (o *PageRule) GetReqHideHeadersOk() (*[]string, bool)`

GetReqHideHeadersOk returns a tuple with the ReqHideHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqHideHeaders

`func (o *PageRule) SetReqHideHeaders(v []string)`

SetReqHideHeaders sets ReqHideHeaders field to given value.

### HasReqHideHeaders

`func (o *PageRule) HasReqHideHeaders() bool`

HasReqHideHeaders returns a boolean if a field has been set.

### GetResHideHeaders

`func (o *PageRule) GetResHideHeaders() []string`

GetResHideHeaders returns the ResHideHeaders field if non-nil, zero value otherwise.

### GetResHideHeadersOk

`func (o *PageRule) GetResHideHeadersOk() (*[]string, bool)`

GetResHideHeadersOk returns a tuple with the ResHideHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResHideHeaders

`func (o *PageRule) SetResHideHeaders(v []string)`

SetResHideHeaders sets ResHideHeaders field to given value.

### HasResHideHeaders

`func (o *PageRule) HasResHideHeaders() bool`

HasResHideHeaders returns a boolean if a field has been set.

### GetCustomHostHeader

`func (o *PageRule) GetCustomHostHeader() string`

GetCustomHostHeader returns the CustomHostHeader field if non-nil, zero value otherwise.

### GetCustomHostHeaderOk

`func (o *PageRule) GetCustomHostHeaderOk() (*string, bool)`

GetCustomHostHeaderOk returns a tuple with the CustomHostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHostHeader

`func (o *PageRule) SetCustomHostHeader(v string)`

SetCustomHostHeader sets CustomHostHeader field to given value.

### HasCustomHostHeader

`func (o *PageRule) HasCustomHostHeader() bool`

HasCustomHostHeader returns a boolean if a field has been set.

### GetRedirect

`func (o *PageRule) GetRedirect() PageRuleRedirect`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *PageRule) GetRedirectOk() (*PageRuleRedirect, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *PageRule) SetRedirect(v PageRuleRedirect)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *PageRule) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetId

`func (o *PageRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PageRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PageRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PageRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomainId

`func (o *PageRule) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *PageRule) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *PageRule) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *PageRule) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetSeq

`func (o *PageRule) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *PageRule) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *PageRule) SetSeq(v int32)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *PageRule) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetUrlType

`func (o *PageRule) GetUrlType() string`

GetUrlType returns the UrlType field if non-nil, zero value otherwise.

### GetUrlTypeOk

`func (o *PageRule) GetUrlTypeOk() (*string, bool)`

GetUrlTypeOk returns a tuple with the UrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlType

`func (o *PageRule) SetUrlType(v string)`

SetUrlType sets UrlType field to given value.

### HasUrlType

`func (o *PageRule) HasUrlType() bool`

HasUrlType returns a boolean if a field has been set.

### GetIsProtected

`func (o *PageRule) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *PageRule) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *PageRule) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *PageRule) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetUrl

`func (o *PageRule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PageRule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PageRule) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PageRule) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCacheLevel

`func (o *PageRule) GetCacheLevel() string`

GetCacheLevel returns the CacheLevel field if non-nil, zero value otherwise.

### GetCacheLevelOk

`func (o *PageRule) GetCacheLevelOk() (*string, bool)`

GetCacheLevelOk returns a tuple with the CacheLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLevel

`func (o *PageRule) SetCacheLevel(v string)`

SetCacheLevel sets CacheLevel field to given value.

### HasCacheLevel

`func (o *PageRule) HasCacheLevel() bool`

HasCacheLevel returns a boolean if a field has been set.

### GetWafStatus

`func (o *PageRule) GetWafStatus() bool`

GetWafStatus returns the WafStatus field if non-nil, zero value otherwise.

### GetWafStatusOk

`func (o *PageRule) GetWafStatusOk() (*bool, bool)`

GetWafStatusOk returns a tuple with the WafStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafStatus

`func (o *PageRule) SetWafStatus(v bool)`

SetWafStatus sets WafStatus field to given value.

### HasWafStatus

`func (o *PageRule) HasWafStatus() bool`

HasWafStatus returns a boolean if a field has been set.

### GetFwStatus

`func (o *PageRule) GetFwStatus() bool`

GetFwStatus returns the FwStatus field if non-nil, zero value otherwise.

### GetFwStatusOk

`func (o *PageRule) GetFwStatusOk() (*bool, bool)`

GetFwStatusOk returns a tuple with the FwStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwStatus

`func (o *PageRule) SetFwStatus(v bool)`

SetFwStatus sets FwStatus field to given value.

### HasFwStatus

`func (o *PageRule) HasFwStatus() bool`

HasFwStatus returns a boolean if a field has been set.

### GetAcceleration

`func (o *PageRule) GetAcceleration() Acceleration`

GetAcceleration returns the Acceleration field if non-nil, zero value otherwise.

### GetAccelerationOk

`func (o *PageRule) GetAccelerationOk() (*Acceleration, bool)`

GetAccelerationOk returns a tuple with the Acceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceleration

`func (o *PageRule) SetAcceleration(v Acceleration)`

SetAcceleration sets Acceleration field to given value.

### HasAcceleration

`func (o *PageRule) HasAcceleration() bool`

HasAcceleration returns a boolean if a field has been set.

### GetSlinkStatus

`func (o *PageRule) GetSlinkStatus() bool`

GetSlinkStatus returns the SlinkStatus field if non-nil, zero value otherwise.

### GetSlinkStatusOk

`func (o *PageRule) GetSlinkStatusOk() (*bool, bool)`

GetSlinkStatusOk returns a tuple with the SlinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlinkStatus

`func (o *PageRule) SetSlinkStatus(v bool)`

SetSlinkStatus sets SlinkStatus field to given value.

### HasSlinkStatus

`func (o *PageRule) HasSlinkStatus() bool`

HasSlinkStatus returns a boolean if a field has been set.

### GetStatus

`func (o *PageRule) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PageRule) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PageRule) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PageRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PageRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PageRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PageRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PageRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PageRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PageRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PageRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PageRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


