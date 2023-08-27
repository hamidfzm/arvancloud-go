# TransportLayerProxyStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Domain** | **string** |  | 
**Port** | **int32** |  | 
**ProxyProtocol** | **string** |  | 
**BalanceAlgorithm** | **string** |  | 
**Servers** | Pointer to [**[]TransportLayerProxyServer**](TransportLayerProxyServer.md) |  | [optional] 
**FirewallDefaultAction** | **string** |  | 
**Firewalls** | Pointer to [**[]TransportLayerProxyFirewall**](TransportLayerProxyFirewall.md) |  | [optional] 

## Methods

### NewTransportLayerProxyStore

`func NewTransportLayerProxyStore(appName string, domain string, port int32, proxyProtocol string, balanceAlgorithm string, firewallDefaultAction string, ) *TransportLayerProxyStore`

NewTransportLayerProxyStore instantiates a new TransportLayerProxyStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportLayerProxyStoreWithDefaults

`func NewTransportLayerProxyStoreWithDefaults() *TransportLayerProxyStore`

NewTransportLayerProxyStoreWithDefaults instantiates a new TransportLayerProxyStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *TransportLayerProxyStore) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *TransportLayerProxyStore) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *TransportLayerProxyStore) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetDescription

`func (o *TransportLayerProxyStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransportLayerProxyStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransportLayerProxyStore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransportLayerProxyStore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *TransportLayerProxyStore) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TransportLayerProxyStore) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TransportLayerProxyStore) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPort

`func (o *TransportLayerProxyStore) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TransportLayerProxyStore) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TransportLayerProxyStore) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProxyProtocol

`func (o *TransportLayerProxyStore) GetProxyProtocol() string`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *TransportLayerProxyStore) GetProxyProtocolOk() (*string, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *TransportLayerProxyStore) SetProxyProtocol(v string)`

SetProxyProtocol sets ProxyProtocol field to given value.


### GetBalanceAlgorithm

`func (o *TransportLayerProxyStore) GetBalanceAlgorithm() string`

GetBalanceAlgorithm returns the BalanceAlgorithm field if non-nil, zero value otherwise.

### GetBalanceAlgorithmOk

`func (o *TransportLayerProxyStore) GetBalanceAlgorithmOk() (*string, bool)`

GetBalanceAlgorithmOk returns a tuple with the BalanceAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAlgorithm

`func (o *TransportLayerProxyStore) SetBalanceAlgorithm(v string)`

SetBalanceAlgorithm sets BalanceAlgorithm field to given value.


### GetServers

`func (o *TransportLayerProxyStore) GetServers() []TransportLayerProxyServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *TransportLayerProxyStore) GetServersOk() (*[]TransportLayerProxyServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *TransportLayerProxyStore) SetServers(v []TransportLayerProxyServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *TransportLayerProxyStore) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetFirewallDefaultAction

`func (o *TransportLayerProxyStore) GetFirewallDefaultAction() string`

GetFirewallDefaultAction returns the FirewallDefaultAction field if non-nil, zero value otherwise.

### GetFirewallDefaultActionOk

`func (o *TransportLayerProxyStore) GetFirewallDefaultActionOk() (*string, bool)`

GetFirewallDefaultActionOk returns a tuple with the FirewallDefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallDefaultAction

`func (o *TransportLayerProxyStore) SetFirewallDefaultAction(v string)`

SetFirewallDefaultAction sets FirewallDefaultAction field to given value.


### GetFirewalls

`func (o *TransportLayerProxyStore) GetFirewalls() []TransportLayerProxyFirewall`

GetFirewalls returns the Firewalls field if non-nil, zero value otherwise.

### GetFirewallsOk

`func (o *TransportLayerProxyStore) GetFirewallsOk() (*[]TransportLayerProxyFirewall, bool)`

GetFirewallsOk returns a tuple with the Firewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewalls

`func (o *TransportLayerProxyStore) SetFirewalls(v []TransportLayerProxyFirewall)`

SetFirewalls sets Firewalls field to given value.

### HasFirewalls

`func (o *TransportLayerProxyStore) HasFirewalls() bool`

HasFirewalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


