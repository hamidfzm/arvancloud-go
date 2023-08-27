# LoadBalancerSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** |  | [optional] 
**NextUpstreamTcp** | Pointer to **string** |  | [optional] [default to "off"]
**Protocol** | Pointer to **string** |  | [optional] 
**Keepalive** | Pointer to **string** |  | [optional] [default to "off"]
**MaxFails** | Pointer to **float32** | Zero will disable failing strategy. | [optional] [default to 0]
**FailTimeout** | Pointer to **string** | Human friendly time duration. | [optional] [default to "10s"]

## Methods

### NewLoadBalancerSetting

`func NewLoadBalancerSetting() *LoadBalancerSetting`

NewLoadBalancerSetting instantiates a new LoadBalancerSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerSettingWithDefaults

`func NewLoadBalancerSettingWithDefaults() *LoadBalancerSetting`

NewLoadBalancerSettingWithDefaults instantiates a new LoadBalancerSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *LoadBalancerSetting) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoadBalancerSetting) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoadBalancerSetting) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoadBalancerSetting) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNextUpstreamTcp

`func (o *LoadBalancerSetting) GetNextUpstreamTcp() string`

GetNextUpstreamTcp returns the NextUpstreamTcp field if non-nil, zero value otherwise.

### GetNextUpstreamTcpOk

`func (o *LoadBalancerSetting) GetNextUpstreamTcpOk() (*string, bool)`

GetNextUpstreamTcpOk returns a tuple with the NextUpstreamTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpstreamTcp

`func (o *LoadBalancerSetting) SetNextUpstreamTcp(v string)`

SetNextUpstreamTcp sets NextUpstreamTcp field to given value.

### HasNextUpstreamTcp

`func (o *LoadBalancerSetting) HasNextUpstreamTcp() bool`

HasNextUpstreamTcp returns a boolean if a field has been set.

### GetProtocol

`func (o *LoadBalancerSetting) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerSetting) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerSetting) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LoadBalancerSetting) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetKeepalive

`func (o *LoadBalancerSetting) GetKeepalive() string`

GetKeepalive returns the Keepalive field if non-nil, zero value otherwise.

### GetKeepaliveOk

`func (o *LoadBalancerSetting) GetKeepaliveOk() (*string, bool)`

GetKeepaliveOk returns a tuple with the Keepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepalive

`func (o *LoadBalancerSetting) SetKeepalive(v string)`

SetKeepalive sets Keepalive field to given value.

### HasKeepalive

`func (o *LoadBalancerSetting) HasKeepalive() bool`

HasKeepalive returns a boolean if a field has been set.

### GetMaxFails

`func (o *LoadBalancerSetting) GetMaxFails() float32`

GetMaxFails returns the MaxFails field if non-nil, zero value otherwise.

### GetMaxFailsOk

`func (o *LoadBalancerSetting) GetMaxFailsOk() (*float32, bool)`

GetMaxFailsOk returns a tuple with the MaxFails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFails

`func (o *LoadBalancerSetting) SetMaxFails(v float32)`

SetMaxFails sets MaxFails field to given value.

### HasMaxFails

`func (o *LoadBalancerSetting) HasMaxFails() bool`

HasMaxFails returns a boolean if a field has been set.

### GetFailTimeout

`func (o *LoadBalancerSetting) GetFailTimeout() string`

GetFailTimeout returns the FailTimeout field if non-nil, zero value otherwise.

### GetFailTimeoutOk

`func (o *LoadBalancerSetting) GetFailTimeoutOk() (*string, bool)`

GetFailTimeoutOk returns a tuple with the FailTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailTimeout

`func (o *LoadBalancerSetting) SetFailTimeout(v string)`

SetFailTimeout sets FailTimeout field to given value.

### HasFailTimeout

`func (o *LoadBalancerSetting) HasFailTimeout() bool`

HasFailTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


