# CreateSecurityGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSecurityGroupRequest

`func NewCreateSecurityGroupRequest() *CreateSecurityGroupRequest`

NewCreateSecurityGroupRequest instantiates a new CreateSecurityGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupRequestWithDefaults

`func NewCreateSecurityGroupRequestWithDefaults() *CreateSecurityGroupRequest`

NewCreateSecurityGroupRequestWithDefaults instantiates a new CreateSecurityGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateSecurityGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSecurityGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSecurityGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSecurityGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateSecurityGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSecurityGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSecurityGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSecurityGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


