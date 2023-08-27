# Grant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grantee** | Pointer to [**GrantGrantee**](GrantGrantee.md) |  | [optional] 
**Permission** | Pointer to [**Permission**](Permission.md) |  | [optional] 

## Methods

### NewGrant

`func NewGrant() *Grant`

NewGrant instantiates a new Grant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantWithDefaults

`func NewGrantWithDefaults() *Grant`

NewGrantWithDefaults instantiates a new Grant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantee

`func (o *Grant) GetGrantee() GrantGrantee`

GetGrantee returns the Grantee field if non-nil, zero value otherwise.

### GetGranteeOk

`func (o *Grant) GetGranteeOk() (*GrantGrantee, bool)`

GetGranteeOk returns a tuple with the Grantee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantee

`func (o *Grant) SetGrantee(v GrantGrantee)`

SetGrantee sets Grantee field to given value.

### HasGrantee

`func (o *Grant) HasGrantee() bool`

HasGrantee returns a boolean if a field has been set.

### GetPermission

`func (o *Grant) GetPermission() Permission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *Grant) GetPermissionOk() (*Permission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *Grant) SetPermission(v Permission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *Grant) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


