# TargetGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grantee** | Pointer to [**TargetGrantGrantee**](TargetGrantGrantee.md) |  | [optional] 
**Permission** | Pointer to [**BucketLogsPermission**](BucketLogsPermission.md) |  | [optional] 

## Methods

### NewTargetGrant

`func NewTargetGrant() *TargetGrant`

NewTargetGrant instantiates a new TargetGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGrantWithDefaults

`func NewTargetGrantWithDefaults() *TargetGrant`

NewTargetGrantWithDefaults instantiates a new TargetGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantee

`func (o *TargetGrant) GetGrantee() TargetGrantGrantee`

GetGrantee returns the Grantee field if non-nil, zero value otherwise.

### GetGranteeOk

`func (o *TargetGrant) GetGranteeOk() (*TargetGrantGrantee, bool)`

GetGranteeOk returns a tuple with the Grantee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantee

`func (o *TargetGrant) SetGrantee(v TargetGrantGrantee)`

SetGrantee sets Grantee field to given value.

### HasGrantee

`func (o *TargetGrant) HasGrantee() bool`

HasGrantee returns a boolean if a field has been set.

### GetPermission

`func (o *TargetGrant) GetPermission() BucketLogsPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *TargetGrant) GetPermissionOk() (*BucketLogsPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *TargetGrant) SetPermission(v BucketLogsPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *TargetGrant) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


