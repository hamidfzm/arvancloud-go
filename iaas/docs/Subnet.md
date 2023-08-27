# Subnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocationPools** | Pointer to [**[]AllocationPool**](AllocationPool.md) | Sub-ranges of CIDR available for dynamic allocation to ports. See AllocationPool. | [optional] 
**Cidr** | Pointer to **string** | CIDR representing IP range for this subnet, based on IP version. | [optional] 
**Description** | Pointer to **string** | Description for the subnet. | [optional] 
**DnsNameservers** | Pointer to **[]string** | DNS name servers used by hosts in this subnet. | [optional] 
**EnableDhcp** | Pointer to **bool** | Specifies whether DHCP is enabled for this subnet or not. | [optional] 
**GatewayIp** | Pointer to **string** | Default gateway used by devices in this subnet. | [optional] 
**HostRoutes** | Pointer to [**[]HostRoute**](HostRoute.md) | Routes that should be used by devices with IPs from this subnet (not including local subnet route). | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpVersion** | Pointer to **string** | IP version, either &#x60;4&#39; or &#x60;6&#39;. | [optional] 
**Ipv6AddressMode** | Pointer to **string** | The IPv6 address modes specifies mechanisms for assigning IPv6 IP addresses. | [optional] 
**Ipv6RaMode** | Pointer to **string** | The IPv6 router advertisement specifies whether the networking service should transmit ICMPv6 packets. | [optional] 
**Name** | Pointer to **string** | Human-readable name for the subnet. Might not be unique. | [optional] 
**NetworkId** | Pointer to **string** | UUID of the parent network. | [optional] 
**ProjectId** | Pointer to **string** | ProjectID is the project owner of the subnet. | [optional] 
**RevisionNumber** | Pointer to **int64** |  | [optional] 
**Servers** | Pointer to [**[]NetworkServer**](NetworkServer.md) |  | [optional] 
**ServiceTypes** | Pointer to **[]string** |  | [optional] 
**SubnetpoolId** | Pointer to **string** | SubnetPoolID is the id of the subnet pool associated with the subnet. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TenantId** | Pointer to **string** | TenantID is the project owner of the subnet. | [optional] 

## Methods

### NewSubnet

`func NewSubnet() *Subnet`

NewSubnet instantiates a new Subnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetWithDefaults

`func NewSubnetWithDefaults() *Subnet`

NewSubnetWithDefaults instantiates a new Subnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocationPools

`func (o *Subnet) GetAllocationPools() []AllocationPool`

GetAllocationPools returns the AllocationPools field if non-nil, zero value otherwise.

### GetAllocationPoolsOk

`func (o *Subnet) GetAllocationPoolsOk() (*[]AllocationPool, bool)`

GetAllocationPoolsOk returns a tuple with the AllocationPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationPools

`func (o *Subnet) SetAllocationPools(v []AllocationPool)`

SetAllocationPools sets AllocationPools field to given value.

### HasAllocationPools

`func (o *Subnet) HasAllocationPools() bool`

HasAllocationPools returns a boolean if a field has been set.

### GetCidr

`func (o *Subnet) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Subnet) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Subnet) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *Subnet) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetDescription

`func (o *Subnet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Subnet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Subnet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Subnet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsNameservers

`func (o *Subnet) GetDnsNameservers() []string`

GetDnsNameservers returns the DnsNameservers field if non-nil, zero value otherwise.

### GetDnsNameserversOk

`func (o *Subnet) GetDnsNameserversOk() (*[]string, bool)`

GetDnsNameserversOk returns a tuple with the DnsNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameservers

`func (o *Subnet) SetDnsNameservers(v []string)`

SetDnsNameservers sets DnsNameservers field to given value.

### HasDnsNameservers

`func (o *Subnet) HasDnsNameservers() bool`

HasDnsNameservers returns a boolean if a field has been set.

### GetEnableDhcp

`func (o *Subnet) GetEnableDhcp() bool`

GetEnableDhcp returns the EnableDhcp field if non-nil, zero value otherwise.

### GetEnableDhcpOk

`func (o *Subnet) GetEnableDhcpOk() (*bool, bool)`

GetEnableDhcpOk returns a tuple with the EnableDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcp

`func (o *Subnet) SetEnableDhcp(v bool)`

SetEnableDhcp sets EnableDhcp field to given value.

### HasEnableDhcp

`func (o *Subnet) HasEnableDhcp() bool`

HasEnableDhcp returns a boolean if a field has been set.

### GetGatewayIp

`func (o *Subnet) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *Subnet) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *Subnet) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *Subnet) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetHostRoutes

`func (o *Subnet) GetHostRoutes() []HostRoute`

GetHostRoutes returns the HostRoutes field if non-nil, zero value otherwise.

### GetHostRoutesOk

`func (o *Subnet) GetHostRoutesOk() (*[]HostRoute, bool)`

GetHostRoutesOk returns a tuple with the HostRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRoutes

`func (o *Subnet) SetHostRoutes(v []HostRoute)`

SetHostRoutes sets HostRoutes field to given value.

### HasHostRoutes

`func (o *Subnet) HasHostRoutes() bool`

HasHostRoutes returns a boolean if a field has been set.

### GetId

`func (o *Subnet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subnet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subnet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subnet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpVersion

`func (o *Subnet) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *Subnet) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *Subnet) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *Subnet) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetIpv6AddressMode

`func (o *Subnet) GetIpv6AddressMode() string`

GetIpv6AddressMode returns the Ipv6AddressMode field if non-nil, zero value otherwise.

### GetIpv6AddressModeOk

`func (o *Subnet) GetIpv6AddressModeOk() (*string, bool)`

GetIpv6AddressModeOk returns a tuple with the Ipv6AddressMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6AddressMode

`func (o *Subnet) SetIpv6AddressMode(v string)`

SetIpv6AddressMode sets Ipv6AddressMode field to given value.

### HasIpv6AddressMode

`func (o *Subnet) HasIpv6AddressMode() bool`

HasIpv6AddressMode returns a boolean if a field has been set.

### GetIpv6RaMode

`func (o *Subnet) GetIpv6RaMode() string`

GetIpv6RaMode returns the Ipv6RaMode field if non-nil, zero value otherwise.

### GetIpv6RaModeOk

`func (o *Subnet) GetIpv6RaModeOk() (*string, bool)`

GetIpv6RaModeOk returns a tuple with the Ipv6RaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6RaMode

`func (o *Subnet) SetIpv6RaMode(v string)`

SetIpv6RaMode sets Ipv6RaMode field to given value.

### HasIpv6RaMode

`func (o *Subnet) HasIpv6RaMode() bool`

HasIpv6RaMode returns a boolean if a field has been set.

### GetName

`func (o *Subnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Subnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkId

`func (o *Subnet) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *Subnet) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *Subnet) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *Subnet) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetProjectId

`func (o *Subnet) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Subnet) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Subnet) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Subnet) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRevisionNumber

`func (o *Subnet) GetRevisionNumber() int64`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *Subnet) GetRevisionNumberOk() (*int64, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *Subnet) SetRevisionNumber(v int64)`

SetRevisionNumber sets RevisionNumber field to given value.

### HasRevisionNumber

`func (o *Subnet) HasRevisionNumber() bool`

HasRevisionNumber returns a boolean if a field has been set.

### GetServers

`func (o *Subnet) GetServers() []NetworkServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Subnet) GetServersOk() (*[]NetworkServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Subnet) SetServers(v []NetworkServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Subnet) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetServiceTypes

`func (o *Subnet) GetServiceTypes() []string`

GetServiceTypes returns the ServiceTypes field if non-nil, zero value otherwise.

### GetServiceTypesOk

`func (o *Subnet) GetServiceTypesOk() (*[]string, bool)`

GetServiceTypesOk returns a tuple with the ServiceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypes

`func (o *Subnet) SetServiceTypes(v []string)`

SetServiceTypes sets ServiceTypes field to given value.

### HasServiceTypes

`func (o *Subnet) HasServiceTypes() bool`

HasServiceTypes returns a boolean if a field has been set.

### GetSubnetpoolId

`func (o *Subnet) GetSubnetpoolId() string`

GetSubnetpoolId returns the SubnetpoolId field if non-nil, zero value otherwise.

### GetSubnetpoolIdOk

`func (o *Subnet) GetSubnetpoolIdOk() (*string, bool)`

GetSubnetpoolIdOk returns a tuple with the SubnetpoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetpoolId

`func (o *Subnet) SetSubnetpoolId(v string)`

SetSubnetpoolId sets SubnetpoolId field to given value.

### HasSubnetpoolId

`func (o *Subnet) HasSubnetpoolId() bool`

HasSubnetpoolId returns a boolean if a field has been set.

### GetTags

`func (o *Subnet) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Subnet) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Subnet) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Subnet) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTenantId

`func (o *Subnet) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Subnet) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Subnet) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Subnet) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


