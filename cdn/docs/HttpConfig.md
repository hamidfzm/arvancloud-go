# HttpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** |  | 
**Port** | **int32** |  | 
**Path** | **string** | path for test | 
**AllowInsecure** | **bool** |  | 
**ExpectedResponse** | [**ExpectedResponse**](ExpectedResponse.md) |  | 
**Headers** | **map[string]string** |  | 
**FollowRedirects** | Pointer to **bool** |  | [optional] [readonly] 
**Timeout** | **int32** | In milliseconds | 

## Methods

### NewHttpConfig

`func NewHttpConfig(method string, port int32, path string, allowInsecure bool, expectedResponse ExpectedResponse, headers map[string]string, timeout int32, ) *HttpConfig`

NewHttpConfig instantiates a new HttpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpConfigWithDefaults

`func NewHttpConfigWithDefaults() *HttpConfig`

NewHttpConfigWithDefaults instantiates a new HttpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *HttpConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HttpConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HttpConfig) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPort

`func (o *HttpConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HttpConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HttpConfig) SetPort(v int32)`

SetPort sets Port field to given value.


### GetPath

`func (o *HttpConfig) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HttpConfig) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HttpConfig) SetPath(v string)`

SetPath sets Path field to given value.


### GetAllowInsecure

`func (o *HttpConfig) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *HttpConfig) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *HttpConfig) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.


### GetExpectedResponse

`func (o *HttpConfig) GetExpectedResponse() ExpectedResponse`

GetExpectedResponse returns the ExpectedResponse field if non-nil, zero value otherwise.

### GetExpectedResponseOk

`func (o *HttpConfig) GetExpectedResponseOk() (*ExpectedResponse, bool)`

GetExpectedResponseOk returns a tuple with the ExpectedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResponse

`func (o *HttpConfig) SetExpectedResponse(v ExpectedResponse)`

SetExpectedResponse sets ExpectedResponse field to given value.


### GetHeaders

`func (o *HttpConfig) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpConfig) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpConfig) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.


### GetFollowRedirects

`func (o *HttpConfig) GetFollowRedirects() bool`

GetFollowRedirects returns the FollowRedirects field if non-nil, zero value otherwise.

### GetFollowRedirectsOk

`func (o *HttpConfig) GetFollowRedirectsOk() (*bool, bool)`

GetFollowRedirectsOk returns a tuple with the FollowRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowRedirects

`func (o *HttpConfig) SetFollowRedirects(v bool)`

SetFollowRedirects sets FollowRedirects field to given value.

### HasFollowRedirects

`func (o *HttpConfig) HasFollowRedirects() bool`

HasFollowRedirects returns a boolean if a field has been set.

### GetTimeout

`func (o *HttpConfig) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HttpConfig) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HttpConfig) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


