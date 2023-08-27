# AttachPortSecurityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupId** | Pointer to **string** |  | [optional] 

## Methods

### NewAttachPortSecurityRequest

`func NewAttachPortSecurityRequest() *AttachPortSecurityRequest`

NewAttachPortSecurityRequest instantiates a new AttachPortSecurityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachPortSecurityRequestWithDefaults

`func NewAttachPortSecurityRequestWithDefaults() *AttachPortSecurityRequest`

NewAttachPortSecurityRequestWithDefaults instantiates a new AttachPortSecurityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupId

`func (o *AttachPortSecurityRequest) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *AttachPortSecurityRequest) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *AttachPortSecurityRequest) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.

### HasSecurityGroupId

`func (o *AttachPortSecurityRequest) HasSecurityGroupId() bool`

HasSecurityGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


