# TransportLayerProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**AppName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Ip** | Pointer to **NullableString** |  | [optional] 
**ProxyProtocol** | Pointer to **string** |  | [optional] 
**BalanceAlgorithm** | Pointer to **string** |  | [optional] 
**Servers** | Pointer to [**[]TransportLayerProxyServersInner**](TransportLayerProxyServersInner.md) |  | [optional] 
**FirewallDefaultAction** | Pointer to **string** |  | [optional] 
**Firewalls** | Pointer to [**[]TransportLayerProxyFirewallsInner**](TransportLayerProxyFirewallsInner.md) |  | [optional] 

## Methods

### NewTransportLayerProxy

`func NewTransportLayerProxy() *TransportLayerProxy`

NewTransportLayerProxy instantiates a new TransportLayerProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportLayerProxyWithDefaults

`func NewTransportLayerProxyWithDefaults() *TransportLayerProxy`

NewTransportLayerProxyWithDefaults instantiates a new TransportLayerProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransportLayerProxy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransportLayerProxy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransportLayerProxy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransportLayerProxy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransportLayerProxy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransportLayerProxy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransportLayerProxy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransportLayerProxy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *TransportLayerProxy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransportLayerProxy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransportLayerProxy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransportLayerProxy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAppName

`func (o *TransportLayerProxy) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *TransportLayerProxy) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *TransportLayerProxy) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *TransportLayerProxy) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetDescription

`func (o *TransportLayerProxy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransportLayerProxy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransportLayerProxy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransportLayerProxy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *TransportLayerProxy) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TransportLayerProxy) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TransportLayerProxy) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TransportLayerProxy) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPort

`func (o *TransportLayerProxy) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TransportLayerProxy) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TransportLayerProxy) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *TransportLayerProxy) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetIp

`func (o *TransportLayerProxy) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TransportLayerProxy) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TransportLayerProxy) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *TransportLayerProxy) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *TransportLayerProxy) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *TransportLayerProxy) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetProxyProtocol

`func (o *TransportLayerProxy) GetProxyProtocol() string`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *TransportLayerProxy) GetProxyProtocolOk() (*string, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *TransportLayerProxy) SetProxyProtocol(v string)`

SetProxyProtocol sets ProxyProtocol field to given value.

### HasProxyProtocol

`func (o *TransportLayerProxy) HasProxyProtocol() bool`

HasProxyProtocol returns a boolean if a field has been set.

### GetBalanceAlgorithm

`func (o *TransportLayerProxy) GetBalanceAlgorithm() string`

GetBalanceAlgorithm returns the BalanceAlgorithm field if non-nil, zero value otherwise.

### GetBalanceAlgorithmOk

`func (o *TransportLayerProxy) GetBalanceAlgorithmOk() (*string, bool)`

GetBalanceAlgorithmOk returns a tuple with the BalanceAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAlgorithm

`func (o *TransportLayerProxy) SetBalanceAlgorithm(v string)`

SetBalanceAlgorithm sets BalanceAlgorithm field to given value.

### HasBalanceAlgorithm

`func (o *TransportLayerProxy) HasBalanceAlgorithm() bool`

HasBalanceAlgorithm returns a boolean if a field has been set.

### GetServers

`func (o *TransportLayerProxy) GetServers() []TransportLayerProxyServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *TransportLayerProxy) GetServersOk() (*[]TransportLayerProxyServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *TransportLayerProxy) SetServers(v []TransportLayerProxyServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *TransportLayerProxy) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetFirewallDefaultAction

`func (o *TransportLayerProxy) GetFirewallDefaultAction() string`

GetFirewallDefaultAction returns the FirewallDefaultAction field if non-nil, zero value otherwise.

### GetFirewallDefaultActionOk

`func (o *TransportLayerProxy) GetFirewallDefaultActionOk() (*string, bool)`

GetFirewallDefaultActionOk returns a tuple with the FirewallDefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallDefaultAction

`func (o *TransportLayerProxy) SetFirewallDefaultAction(v string)`

SetFirewallDefaultAction sets FirewallDefaultAction field to given value.

### HasFirewallDefaultAction

`func (o *TransportLayerProxy) HasFirewallDefaultAction() bool`

HasFirewallDefaultAction returns a boolean if a field has been set.

### GetFirewalls

`func (o *TransportLayerProxy) GetFirewalls() []TransportLayerProxyFirewallsInner`

GetFirewalls returns the Firewalls field if non-nil, zero value otherwise.

### GetFirewallsOk

`func (o *TransportLayerProxy) GetFirewallsOk() (*[]TransportLayerProxyFirewallsInner, bool)`

GetFirewallsOk returns a tuple with the Firewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewalls

`func (o *TransportLayerProxy) SetFirewalls(v []TransportLayerProxyFirewallsInner)`

SetFirewalls sets Firewalls field to given value.

### HasFirewalls

`func (o *TransportLayerProxy) HasFirewalls() bool`

HasFirewalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


