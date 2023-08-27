# GetObjectAclOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**Owner**](Owner.md) |  | [optional] 
**Grants** | Pointer to [**Grants**](Grants.md) |  | [optional] 

## Methods

### NewGetObjectAclOutput

`func NewGetObjectAclOutput() *GetObjectAclOutput`

NewGetObjectAclOutput instantiates a new GetObjectAclOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectAclOutputWithDefaults

`func NewGetObjectAclOutputWithDefaults() *GetObjectAclOutput`

NewGetObjectAclOutputWithDefaults instantiates a new GetObjectAclOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *GetObjectAclOutput) GetOwner() Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetObjectAclOutput) GetOwnerOk() (*Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetObjectAclOutput) SetOwner(v Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetObjectAclOutput) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGrants

`func (o *GetObjectAclOutput) GetGrants() Grants`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *GetObjectAclOutput) GetGrantsOk() (*Grants, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *GetObjectAclOutput) SetGrants(v Grants)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *GetObjectAclOutput) HasGrants() bool`

HasGrants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


