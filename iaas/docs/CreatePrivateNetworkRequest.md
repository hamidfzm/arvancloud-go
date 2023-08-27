# CreatePrivateNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Dhcp** | Pointer to **string** |  | [optional] 
**DnsServers** | Pointer to **string** |  | [optional] 
**EnableDhcp** | Pointer to **bool** |  | [optional] 
**EnableGateway** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NetworkId** | Pointer to **string** |  | [optional] 
**SubnetGateway** | Pointer to **string** |  | [optional] 
**SubnetId** | Pointer to **string** |  | [optional] 
**SubnetIp** | Pointer to **string** |  | [optional] 

## Methods

### NewCreatePrivateNetworkRequest

`func NewCreatePrivateNetworkRequest() *CreatePrivateNetworkRequest`

NewCreatePrivateNetworkRequest instantiates a new CreatePrivateNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePrivateNetworkRequestWithDefaults

`func NewCreatePrivateNetworkRequestWithDefaults() *CreatePrivateNetworkRequest`

NewCreatePrivateNetworkRequestWithDefaults instantiates a new CreatePrivateNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreatePrivateNetworkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePrivateNetworkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePrivateNetworkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePrivateNetworkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDhcp

`func (o *CreatePrivateNetworkRequest) GetDhcp() string`

GetDhcp returns the Dhcp field if non-nil, zero value otherwise.

### GetDhcpOk

`func (o *CreatePrivateNetworkRequest) GetDhcpOk() (*string, bool)`

GetDhcpOk returns a tuple with the Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp

`func (o *CreatePrivateNetworkRequest) SetDhcp(v string)`

SetDhcp sets Dhcp field to given value.

### HasDhcp

`func (o *CreatePrivateNetworkRequest) HasDhcp() bool`

HasDhcp returns a boolean if a field has been set.

### GetDnsServers

`func (o *CreatePrivateNetworkRequest) GetDnsServers() string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *CreatePrivateNetworkRequest) GetDnsServersOk() (*string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *CreatePrivateNetworkRequest) SetDnsServers(v string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *CreatePrivateNetworkRequest) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetEnableDhcp

`func (o *CreatePrivateNetworkRequest) GetEnableDhcp() bool`

GetEnableDhcp returns the EnableDhcp field if non-nil, zero value otherwise.

### GetEnableDhcpOk

`func (o *CreatePrivateNetworkRequest) GetEnableDhcpOk() (*bool, bool)`

GetEnableDhcpOk returns a tuple with the EnableDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcp

`func (o *CreatePrivateNetworkRequest) SetEnableDhcp(v bool)`

SetEnableDhcp sets EnableDhcp field to given value.

### HasEnableDhcp

`func (o *CreatePrivateNetworkRequest) HasEnableDhcp() bool`

HasEnableDhcp returns a boolean if a field has been set.

### GetEnableGateway

`func (o *CreatePrivateNetworkRequest) GetEnableGateway() bool`

GetEnableGateway returns the EnableGateway field if non-nil, zero value otherwise.

### GetEnableGatewayOk

`func (o *CreatePrivateNetworkRequest) GetEnableGatewayOk() (*bool, bool)`

GetEnableGatewayOk returns a tuple with the EnableGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGateway

`func (o *CreatePrivateNetworkRequest) SetEnableGateway(v bool)`

SetEnableGateway sets EnableGateway field to given value.

### HasEnableGateway

`func (o *CreatePrivateNetworkRequest) HasEnableGateway() bool`

HasEnableGateway returns a boolean if a field has been set.

### GetName

`func (o *CreatePrivateNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePrivateNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePrivateNetworkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePrivateNetworkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkId

`func (o *CreatePrivateNetworkRequest) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CreatePrivateNetworkRequest) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CreatePrivateNetworkRequest) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *CreatePrivateNetworkRequest) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSubnetGateway

`func (o *CreatePrivateNetworkRequest) GetSubnetGateway() string`

GetSubnetGateway returns the SubnetGateway field if non-nil, zero value otherwise.

### GetSubnetGatewayOk

`func (o *CreatePrivateNetworkRequest) GetSubnetGatewayOk() (*string, bool)`

GetSubnetGatewayOk returns a tuple with the SubnetGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetGateway

`func (o *CreatePrivateNetworkRequest) SetSubnetGateway(v string)`

SetSubnetGateway sets SubnetGateway field to given value.

### HasSubnetGateway

`func (o *CreatePrivateNetworkRequest) HasSubnetGateway() bool`

HasSubnetGateway returns a boolean if a field has been set.

### GetSubnetId

`func (o *CreatePrivateNetworkRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreatePrivateNetworkRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreatePrivateNetworkRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *CreatePrivateNetworkRequest) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetSubnetIp

`func (o *CreatePrivateNetworkRequest) GetSubnetIp() string`

GetSubnetIp returns the SubnetIp field if non-nil, zero value otherwise.

### GetSubnetIpOk

`func (o *CreatePrivateNetworkRequest) GetSubnetIpOk() (*string, bool)`

GetSubnetIpOk returns a tuple with the SubnetIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIp

`func (o *CreatePrivateNetworkRequest) SetSubnetIp(v string)`

SetSubnetIp sets SubnetIp field to given value.

### HasSubnetIp

`func (o *CreatePrivateNetworkRequest) HasSubnetIp() bool`

HasSubnetIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


