# V1GitBuildSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpProxy** | Pointer to **string** | httpProxy is a proxy used to reach the git repository over http | [optional] 
**HttpsProxy** | Pointer to **string** | httpsProxy is a proxy used to reach the git repository over https | [optional] 
**NoProxy** | Pointer to **string** | noProxy is the list of domains for which the proxy should not be used | [optional] 
**Ref** | Pointer to **string** | ref is the branch/tag/ref to build. | [optional] 
**Uri** | **string** | uri points to the source that will be built. The structure of the source will depend on the type of build to run | 

## Methods

### NewV1GitBuildSource

`func NewV1GitBuildSource(uri string, ) *V1GitBuildSource`

NewV1GitBuildSource instantiates a new V1GitBuildSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GitBuildSourceWithDefaults

`func NewV1GitBuildSourceWithDefaults() *V1GitBuildSource`

NewV1GitBuildSourceWithDefaults instantiates a new V1GitBuildSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpProxy

`func (o *V1GitBuildSource) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *V1GitBuildSource) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *V1GitBuildSource) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *V1GitBuildSource) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### GetHttpsProxy

`func (o *V1GitBuildSource) GetHttpsProxy() string`

GetHttpsProxy returns the HttpsProxy field if non-nil, zero value otherwise.

### GetHttpsProxyOk

`func (o *V1GitBuildSource) GetHttpsProxyOk() (*string, bool)`

GetHttpsProxyOk returns a tuple with the HttpsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsProxy

`func (o *V1GitBuildSource) SetHttpsProxy(v string)`

SetHttpsProxy sets HttpsProxy field to given value.

### HasHttpsProxy

`func (o *V1GitBuildSource) HasHttpsProxy() bool`

HasHttpsProxy returns a boolean if a field has been set.

### GetNoProxy

`func (o *V1GitBuildSource) GetNoProxy() string`

GetNoProxy returns the NoProxy field if non-nil, zero value otherwise.

### GetNoProxyOk

`func (o *V1GitBuildSource) GetNoProxyOk() (*string, bool)`

GetNoProxyOk returns a tuple with the NoProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProxy

`func (o *V1GitBuildSource) SetNoProxy(v string)`

SetNoProxy sets NoProxy field to given value.

### HasNoProxy

`func (o *V1GitBuildSource) HasNoProxy() bool`

HasNoProxy returns a boolean if a field has been set.

### GetRef

`func (o *V1GitBuildSource) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *V1GitBuildSource) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *V1GitBuildSource) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *V1GitBuildSource) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetUri

`func (o *V1GitBuildSource) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *V1GitBuildSource) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *V1GitBuildSource) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


