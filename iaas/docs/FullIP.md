# FullIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FloatIp** | Pointer to **map[string]interface{}** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**PortId** | Pointer to **string** |  | [optional] 
**PortSecurityEnabled** | Pointer to **bool** |  | [optional] 
**Ptr** | Pointer to **map[string]interface{}** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**SecurityGroups** | Pointer to [**[]PortSecGroupData**](PortSecGroupData.md) |  | [optional] 
**SubnetId** | Pointer to **string** |  | [optional] 
**SubnetName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewFullIP

`func NewFullIP() *FullIP`

NewFullIP instantiates a new FullIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullIPWithDefaults

`func NewFullIPWithDefaults() *FullIP`

NewFullIPWithDefaults instantiates a new FullIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloatIp

`func (o *FullIP) GetFloatIp() map[string]interface{}`

GetFloatIp returns the FloatIp field if non-nil, zero value otherwise.

### GetFloatIpOk

`func (o *FullIP) GetFloatIpOk() (*map[string]interface{}, bool)`

GetFloatIpOk returns a tuple with the FloatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatIp

`func (o *FullIP) SetFloatIp(v map[string]interface{})`

SetFloatIp sets FloatIp field to given value.

### HasFloatIp

`func (o *FullIP) HasFloatIp() bool`

HasFloatIp returns a boolean if a field has been set.

### GetIp

`func (o *FullIP) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *FullIP) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *FullIP) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *FullIP) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMacAddress

`func (o *FullIP) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *FullIP) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *FullIP) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *FullIP) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPortId

`func (o *FullIP) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *FullIP) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *FullIP) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *FullIP) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortSecurityEnabled

`func (o *FullIP) GetPortSecurityEnabled() bool`

GetPortSecurityEnabled returns the PortSecurityEnabled field if non-nil, zero value otherwise.

### GetPortSecurityEnabledOk

`func (o *FullIP) GetPortSecurityEnabledOk() (*bool, bool)`

GetPortSecurityEnabledOk returns a tuple with the PortSecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSecurityEnabled

`func (o *FullIP) SetPortSecurityEnabled(v bool)`

SetPortSecurityEnabled sets PortSecurityEnabled field to given value.

### HasPortSecurityEnabled

`func (o *FullIP) HasPortSecurityEnabled() bool`

HasPortSecurityEnabled returns a boolean if a field has been set.

### GetPtr

`func (o *FullIP) GetPtr() map[string]interface{}`

GetPtr returns the Ptr field if non-nil, zero value otherwise.

### GetPtrOk

`func (o *FullIP) GetPtrOk() (*map[string]interface{}, bool)`

GetPtrOk returns a tuple with the Ptr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtr

`func (o *FullIP) SetPtr(v map[string]interface{})`

SetPtr sets Ptr field to given value.

### HasPtr

`func (o *FullIP) HasPtr() bool`

HasPtr returns a boolean if a field has been set.

### GetPublic

`func (o *FullIP) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *FullIP) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *FullIP) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *FullIP) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *FullIP) GetSecurityGroups() []PortSecGroupData`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *FullIP) GetSecurityGroupsOk() (*[]PortSecGroupData, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *FullIP) SetSecurityGroups(v []PortSecGroupData)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *FullIP) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSubnetId

`func (o *FullIP) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *FullIP) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *FullIP) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *FullIP) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetSubnetName

`func (o *FullIP) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *FullIP) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *FullIP) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *FullIP) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### GetVersion

`func (o *FullIP) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FullIP) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FullIP) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FullIP) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


