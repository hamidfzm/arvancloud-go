# NetworkServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**map[string][]ServerAddress**](array.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Ips** | Pointer to [**[]FullIP**](FullIP.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PublicIp** | Pointer to [**[]PublicIP**](PublicIP.md) |  | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkServer

`func NewNetworkServer() *NetworkServer`

NewNetworkServer instantiates a new NetworkServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServerWithDefaults

`func NewNetworkServerWithDefaults() *NetworkServer`

NewNetworkServerWithDefaults instantiates a new NetworkServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *NetworkServer) GetAddresses() map[string][]ServerAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *NetworkServer) GetAddressesOk() (*map[string][]ServerAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *NetworkServer) SetAddresses(v map[string][]ServerAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *NetworkServer) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetId

`func (o *NetworkServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIps

`func (o *NetworkServer) GetIps() []FullIP`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *NetworkServer) GetIpsOk() (*[]FullIP, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *NetworkServer) SetIps(v []FullIP)`

SetIps sets Ips field to given value.

### HasIps

`func (o *NetworkServer) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetName

`func (o *NetworkServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicIp

`func (o *NetworkServer) GetPublicIp() []PublicIP`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *NetworkServer) GetPublicIpOk() (*[]PublicIP, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *NetworkServer) SetPublicIp(v []PublicIP)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *NetworkServer) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *NetworkServer) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *NetworkServer) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *NetworkServer) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *NetworkServer) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


