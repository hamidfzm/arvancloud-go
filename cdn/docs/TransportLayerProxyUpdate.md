# TransportLayerProxyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**ProxyProtocol** | **string** |  | 
**BalanceAlgorithm** | **string** |  | 
**Servers** | Pointer to [**[]TransportLayerProxyServer**](TransportLayerProxyServer.md) |  | [optional] 
**FirewallDefaultAction** | Pointer to **string** |  | [optional] 
**Firewalls** | Pointer to [**[]TransportLayerProxyFirewall**](TransportLayerProxyFirewall.md) |  | [optional] 

## Methods

### NewTransportLayerProxyUpdate

`func NewTransportLayerProxyUpdate(appName string, proxyProtocol string, balanceAlgorithm string, ) *TransportLayerProxyUpdate`

NewTransportLayerProxyUpdate instantiates a new TransportLayerProxyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportLayerProxyUpdateWithDefaults

`func NewTransportLayerProxyUpdateWithDefaults() *TransportLayerProxyUpdate`

NewTransportLayerProxyUpdateWithDefaults instantiates a new TransportLayerProxyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *TransportLayerProxyUpdate) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *TransportLayerProxyUpdate) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *TransportLayerProxyUpdate) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetDescription

`func (o *TransportLayerProxyUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransportLayerProxyUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransportLayerProxyUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransportLayerProxyUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *TransportLayerProxyUpdate) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TransportLayerProxyUpdate) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TransportLayerProxyUpdate) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TransportLayerProxyUpdate) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPort

`func (o *TransportLayerProxyUpdate) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TransportLayerProxyUpdate) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TransportLayerProxyUpdate) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *TransportLayerProxyUpdate) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProxyProtocol

`func (o *TransportLayerProxyUpdate) GetProxyProtocol() string`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *TransportLayerProxyUpdate) GetProxyProtocolOk() (*string, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *TransportLayerProxyUpdate) SetProxyProtocol(v string)`

SetProxyProtocol sets ProxyProtocol field to given value.


### GetBalanceAlgorithm

`func (o *TransportLayerProxyUpdate) GetBalanceAlgorithm() string`

GetBalanceAlgorithm returns the BalanceAlgorithm field if non-nil, zero value otherwise.

### GetBalanceAlgorithmOk

`func (o *TransportLayerProxyUpdate) GetBalanceAlgorithmOk() (*string, bool)`

GetBalanceAlgorithmOk returns a tuple with the BalanceAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAlgorithm

`func (o *TransportLayerProxyUpdate) SetBalanceAlgorithm(v string)`

SetBalanceAlgorithm sets BalanceAlgorithm field to given value.


### GetServers

`func (o *TransportLayerProxyUpdate) GetServers() []TransportLayerProxyServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *TransportLayerProxyUpdate) GetServersOk() (*[]TransportLayerProxyServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *TransportLayerProxyUpdate) SetServers(v []TransportLayerProxyServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *TransportLayerProxyUpdate) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetFirewallDefaultAction

`func (o *TransportLayerProxyUpdate) GetFirewallDefaultAction() string`

GetFirewallDefaultAction returns the FirewallDefaultAction field if non-nil, zero value otherwise.

### GetFirewallDefaultActionOk

`func (o *TransportLayerProxyUpdate) GetFirewallDefaultActionOk() (*string, bool)`

GetFirewallDefaultActionOk returns a tuple with the FirewallDefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallDefaultAction

`func (o *TransportLayerProxyUpdate) SetFirewallDefaultAction(v string)`

SetFirewallDefaultAction sets FirewallDefaultAction field to given value.

### HasFirewallDefaultAction

`func (o *TransportLayerProxyUpdate) HasFirewallDefaultAction() bool`

HasFirewallDefaultAction returns a boolean if a field has been set.

### GetFirewalls

`func (o *TransportLayerProxyUpdate) GetFirewalls() []TransportLayerProxyFirewall`

GetFirewalls returns the Firewalls field if non-nil, zero value otherwise.

### GetFirewallsOk

`func (o *TransportLayerProxyUpdate) GetFirewallsOk() (*[]TransportLayerProxyFirewall, bool)`

GetFirewallsOk returns a tuple with the Firewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewalls

`func (o *TransportLayerProxyUpdate) SetFirewalls(v []TransportLayerProxyFirewall)`

SetFirewalls sets Firewalls field to given value.

### HasFirewalls

`func (o *TransportLayerProxyUpdate) HasFirewalls() bool`

HasFirewalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


