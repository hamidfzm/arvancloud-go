# AccessControlPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grants** | Pointer to [**Grants**](Grants.md) |  | [optional] 
**Owner** | Pointer to [**Owner**](Owner.md) |  | [optional] 

## Methods

### NewAccessControlPolicy

`func NewAccessControlPolicy() *AccessControlPolicy`

NewAccessControlPolicy instantiates a new AccessControlPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlPolicyWithDefaults

`func NewAccessControlPolicyWithDefaults() *AccessControlPolicy`

NewAccessControlPolicyWithDefaults instantiates a new AccessControlPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrants

`func (o *AccessControlPolicy) GetGrants() Grants`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *AccessControlPolicy) GetGrantsOk() (*Grants, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *AccessControlPolicy) SetGrants(v Grants)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *AccessControlPolicy) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetOwner

`func (o *AccessControlPolicy) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccessControlPolicy) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccessControlPolicy) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AccessControlPolicy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


