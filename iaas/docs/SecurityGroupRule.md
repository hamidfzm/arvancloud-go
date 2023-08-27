# SecurityGroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**EtherType** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**PortEnd** | Pointer to **int64** |  | [optional] 
**PortStart** | Pointer to **int64** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewSecurityGroupRule

`func NewSecurityGroupRule() *SecurityGroupRule`

NewSecurityGroupRule instantiates a new SecurityGroupRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRuleWithDefaults

`func NewSecurityGroupRuleWithDefaults() *SecurityGroupRule`

NewSecurityGroupRuleWithDefaults instantiates a new SecurityGroupRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SecurityGroupRule) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityGroupRule) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityGroupRule) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SecurityGroupRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityGroupRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroupRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroupRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityGroupRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *SecurityGroupRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SecurityGroupRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SecurityGroupRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SecurityGroupRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEtherType

`func (o *SecurityGroupRule) GetEtherType() string`

GetEtherType returns the EtherType field if non-nil, zero value otherwise.

### GetEtherTypeOk

`func (o *SecurityGroupRule) GetEtherTypeOk() (*string, bool)`

GetEtherTypeOk returns a tuple with the EtherType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherType

`func (o *SecurityGroupRule) SetEtherType(v string)`

SetEtherType sets EtherType field to given value.

### HasEtherType

`func (o *SecurityGroupRule) HasEtherType() bool`

HasEtherType returns a boolean if a field has been set.

### GetGroupId

`func (o *SecurityGroupRule) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SecurityGroupRule) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SecurityGroupRule) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SecurityGroupRule) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *SecurityGroupRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityGroupRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *SecurityGroupRule) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SecurityGroupRule) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SecurityGroupRule) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SecurityGroupRule) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPortEnd

`func (o *SecurityGroupRule) GetPortEnd() int64`

GetPortEnd returns the PortEnd field if non-nil, zero value otherwise.

### GetPortEndOk

`func (o *SecurityGroupRule) GetPortEndOk() (*int64, bool)`

GetPortEndOk returns a tuple with the PortEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortEnd

`func (o *SecurityGroupRule) SetPortEnd(v int64)`

SetPortEnd sets PortEnd field to given value.

### HasPortEnd

`func (o *SecurityGroupRule) HasPortEnd() bool`

HasPortEnd returns a boolean if a field has been set.

### GetPortStart

`func (o *SecurityGroupRule) GetPortStart() int64`

GetPortStart returns the PortStart field if non-nil, zero value otherwise.

### GetPortStartOk

`func (o *SecurityGroupRule) GetPortStartOk() (*int64, bool)`

GetPortStartOk returns a tuple with the PortStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStart

`func (o *SecurityGroupRule) SetPortStart(v int64)`

SetPortStart sets PortStart field to given value.

### HasPortStart

`func (o *SecurityGroupRule) HasPortStart() bool`

HasPortStart returns a boolean if a field has been set.

### GetProtocol

`func (o *SecurityGroupRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SecurityGroupRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SecurityGroupRule) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SecurityGroupRule) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SecurityGroupRule) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SecurityGroupRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


