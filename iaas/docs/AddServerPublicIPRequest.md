# AddServerPublicIPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroup** | Pointer to **string** |  | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddServerPublicIPRequest

`func NewAddServerPublicIPRequest() *AddServerPublicIPRequest`

NewAddServerPublicIPRequest instantiates a new AddServerPublicIPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServerPublicIPRequestWithDefaults

`func NewAddServerPublicIPRequestWithDefaults() *AddServerPublicIPRequest`

NewAddServerPublicIPRequestWithDefaults instantiates a new AddServerPublicIPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroup

`func (o *AddServerPublicIPRequest) GetSecurityGroup() string`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *AddServerPublicIPRequest) GetSecurityGroupOk() (*string, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *AddServerPublicIPRequest) SetSecurityGroup(v string)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *AddServerPublicIPRequest) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *AddServerPublicIPRequest) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *AddServerPublicIPRequest) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *AddServerPublicIPRequest) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *AddServerPublicIPRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


