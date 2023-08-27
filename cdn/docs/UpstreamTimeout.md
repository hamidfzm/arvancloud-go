# UpstreamTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectTimeout** | Pointer to **int32** | Seconds to timeout | [optional] [default to 30]
**ReadTimeout** | Pointer to **int32** | Seconds to timeout | [optional] [default to 100]
**SendTimeout** | Pointer to **int32** | Seconds to timeout | [optional] [default to 300]

## Methods

### NewUpstreamTimeout

`func NewUpstreamTimeout() *UpstreamTimeout`

NewUpstreamTimeout instantiates a new UpstreamTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpstreamTimeoutWithDefaults

`func NewUpstreamTimeoutWithDefaults() *UpstreamTimeout`

NewUpstreamTimeoutWithDefaults instantiates a new UpstreamTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectTimeout

`func (o *UpstreamTimeout) GetConnectTimeout() int32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *UpstreamTimeout) GetConnectTimeoutOk() (*int32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *UpstreamTimeout) SetConnectTimeout(v int32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *UpstreamTimeout) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetReadTimeout

`func (o *UpstreamTimeout) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *UpstreamTimeout) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *UpstreamTimeout) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *UpstreamTimeout) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetSendTimeout

`func (o *UpstreamTimeout) GetSendTimeout() int32`

GetSendTimeout returns the SendTimeout field if non-nil, zero value otherwise.

### GetSendTimeoutOk

`func (o *UpstreamTimeout) GetSendTimeoutOk() (*int32, bool)`

GetSendTimeoutOk returns a tuple with the SendTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTimeout

`func (o *UpstreamTimeout) SetSendTimeout(v int32)`

SetSendTimeout sets SendTimeout field to given value.

### HasSendTimeout

`func (o *UpstreamTimeout) HasSendTimeout() bool`

HasSendTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


