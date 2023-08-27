# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStateUp** | Pointer to **bool** |  | [optional] 
**AvailabilityZoneHints** | Pointer to **[]string** |  | [optional] 
**AvailabilityZones** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DhcpIp** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Ipv4AddressScope** | Pointer to **string** |  | [optional] 
**Ipv6AddressScope** | Pointer to **string** |  | [optional] 
**Mtu** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PortSecurityEnabled** | Pointer to **bool** |  | [optional] 
**QosPolicyId** | Pointer to **string** |  | [optional] 
**RevisionNumber** | Pointer to **int64** |  | [optional] 
**Routerexternal** | Pointer to **bool** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Subnets** | Pointer to [**[]Subnet**](Subnet.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStateUp

`func (o *Network) GetAdminStateUp() bool`

GetAdminStateUp returns the AdminStateUp field if non-nil, zero value otherwise.

### GetAdminStateUpOk

`func (o *Network) GetAdminStateUpOk() (*bool, bool)`

GetAdminStateUpOk returns a tuple with the AdminStateUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStateUp

`func (o *Network) SetAdminStateUp(v bool)`

SetAdminStateUp sets AdminStateUp field to given value.

### HasAdminStateUp

`func (o *Network) HasAdminStateUp() bool`

HasAdminStateUp returns a boolean if a field has been set.

### GetAvailabilityZoneHints

`func (o *Network) GetAvailabilityZoneHints() []string`

GetAvailabilityZoneHints returns the AvailabilityZoneHints field if non-nil, zero value otherwise.

### GetAvailabilityZoneHintsOk

`func (o *Network) GetAvailabilityZoneHintsOk() (*[]string, bool)`

GetAvailabilityZoneHintsOk returns a tuple with the AvailabilityZoneHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneHints

`func (o *Network) SetAvailabilityZoneHints(v []string)`

SetAvailabilityZoneHints sets AvailabilityZoneHints field to given value.

### HasAvailabilityZoneHints

`func (o *Network) HasAvailabilityZoneHints() bool`

HasAvailabilityZoneHints returns a boolean if a field has been set.

### GetAvailabilityZones

`func (o *Network) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *Network) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *Network) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *Network) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Network) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Network) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Network) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Network) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Network) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Network) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Network) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Network) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDhcpIp

`func (o *Network) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *Network) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *Network) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *Network) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### GetId

`func (o *Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Network) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Network) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpv4AddressScope

`func (o *Network) GetIpv4AddressScope() string`

GetIpv4AddressScope returns the Ipv4AddressScope field if non-nil, zero value otherwise.

### GetIpv4AddressScopeOk

`func (o *Network) GetIpv4AddressScopeOk() (*string, bool)`

GetIpv4AddressScopeOk returns a tuple with the Ipv4AddressScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddressScope

`func (o *Network) SetIpv4AddressScope(v string)`

SetIpv4AddressScope sets Ipv4AddressScope field to given value.

### HasIpv4AddressScope

`func (o *Network) HasIpv4AddressScope() bool`

HasIpv4AddressScope returns a boolean if a field has been set.

### GetIpv6AddressScope

`func (o *Network) GetIpv6AddressScope() string`

GetIpv6AddressScope returns the Ipv6AddressScope field if non-nil, zero value otherwise.

### GetIpv6AddressScopeOk

`func (o *Network) GetIpv6AddressScopeOk() (*string, bool)`

GetIpv6AddressScopeOk returns a tuple with the Ipv6AddressScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6AddressScope

`func (o *Network) SetIpv6AddressScope(v string)`

SetIpv6AddressScope sets Ipv6AddressScope field to given value.

### HasIpv6AddressScope

`func (o *Network) HasIpv6AddressScope() bool`

HasIpv6AddressScope returns a boolean if a field has been set.

### GetMtu

`func (o *Network) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *Network) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *Network) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *Network) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Network) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Network) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortSecurityEnabled

`func (o *Network) GetPortSecurityEnabled() bool`

GetPortSecurityEnabled returns the PortSecurityEnabled field if non-nil, zero value otherwise.

### GetPortSecurityEnabledOk

`func (o *Network) GetPortSecurityEnabledOk() (*bool, bool)`

GetPortSecurityEnabledOk returns a tuple with the PortSecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSecurityEnabled

`func (o *Network) SetPortSecurityEnabled(v bool)`

SetPortSecurityEnabled sets PortSecurityEnabled field to given value.

### HasPortSecurityEnabled

`func (o *Network) HasPortSecurityEnabled() bool`

HasPortSecurityEnabled returns a boolean if a field has been set.

### GetQosPolicyId

`func (o *Network) GetQosPolicyId() string`

GetQosPolicyId returns the QosPolicyId field if non-nil, zero value otherwise.

### GetQosPolicyIdOk

`func (o *Network) GetQosPolicyIdOk() (*string, bool)`

GetQosPolicyIdOk returns a tuple with the QosPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosPolicyId

`func (o *Network) SetQosPolicyId(v string)`

SetQosPolicyId sets QosPolicyId field to given value.

### HasQosPolicyId

`func (o *Network) HasQosPolicyId() bool`

HasQosPolicyId returns a boolean if a field has been set.

### GetRevisionNumber

`func (o *Network) GetRevisionNumber() int64`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *Network) GetRevisionNumberOk() (*int64, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *Network) SetRevisionNumber(v int64)`

SetRevisionNumber sets RevisionNumber field to given value.

### HasRevisionNumber

`func (o *Network) HasRevisionNumber() bool`

HasRevisionNumber returns a boolean if a field has been set.

### GetRouterexternal

`func (o *Network) GetRouterexternal() bool`

GetRouterexternal returns the Routerexternal field if non-nil, zero value otherwise.

### GetRouterexternalOk

`func (o *Network) GetRouterexternalOk() (*bool, bool)`

GetRouterexternalOk returns a tuple with the Routerexternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterexternal

`func (o *Network) SetRouterexternal(v bool)`

SetRouterexternal sets Routerexternal field to given value.

### HasRouterexternal

`func (o *Network) HasRouterexternal() bool`

HasRouterexternal returns a boolean if a field has been set.

### GetShared

`func (o *Network) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *Network) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *Network) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *Network) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetStatus

`func (o *Network) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Network) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Network) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Network) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubnets

`func (o *Network) GetSubnets() []Subnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *Network) GetSubnetsOk() (*[]Subnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *Network) SetSubnets(v []Subnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *Network) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetTags

`func (o *Network) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Network) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Network) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Network) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTenantId

`func (o *Network) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Network) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Network) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Network) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Network) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Network) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Network) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Network) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


