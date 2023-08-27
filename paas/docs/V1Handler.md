# V1Handler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exec** | Pointer to [**V1ExecAction**](V1ExecAction.md) |  | [optional] 
**HttpGet** | Pointer to [**V1HTTPGetAction**](V1HTTPGetAction.md) |  | [optional] 
**TcpSocket** | Pointer to [**V1TCPSocketAction**](V1TCPSocketAction.md) |  | [optional] 

## Methods

### NewV1Handler

`func NewV1Handler() *V1Handler`

NewV1Handler instantiates a new V1Handler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HandlerWithDefaults

`func NewV1HandlerWithDefaults() *V1Handler`

NewV1HandlerWithDefaults instantiates a new V1Handler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExec

`func (o *V1Handler) GetExec() V1ExecAction`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *V1Handler) GetExecOk() (*V1ExecAction, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *V1Handler) SetExec(v V1ExecAction)`

SetExec sets Exec field to given value.

### HasExec

`func (o *V1Handler) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetHttpGet

`func (o *V1Handler) GetHttpGet() V1HTTPGetAction`

GetHttpGet returns the HttpGet field if non-nil, zero value otherwise.

### GetHttpGetOk

`func (o *V1Handler) GetHttpGetOk() (*V1HTTPGetAction, bool)`

GetHttpGetOk returns a tuple with the HttpGet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpGet

`func (o *V1Handler) SetHttpGet(v V1HTTPGetAction)`

SetHttpGet sets HttpGet field to given value.

### HasHttpGet

`func (o *V1Handler) HasHttpGet() bool`

HasHttpGet returns a boolean if a field has been set.

### GetTcpSocket

`func (o *V1Handler) GetTcpSocket() V1TCPSocketAction`

GetTcpSocket returns the TcpSocket field if non-nil, zero value otherwise.

### GetTcpSocketOk

`func (o *V1Handler) GetTcpSocketOk() (*V1TCPSocketAction, bool)`

GetTcpSocketOk returns a tuple with the TcpSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSocket

`func (o *V1Handler) SetTcpSocket(v V1TCPSocketAction)`

SetTcpSocket sets TcpSocket field to given value.

### HasTcpSocket

`func (o *V1Handler) HasTcpSocket() bool`

HasTcpSocket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


