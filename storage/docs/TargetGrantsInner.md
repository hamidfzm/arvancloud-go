# TargetGrantsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grantee** | Pointer to [**TargetGrantGrantee**](TargetGrantGrantee.md) |  | [optional] 
**Permission** | Pointer to [**BucketLogsPermission**](BucketLogsPermission.md) |  | [optional] 

## Methods

### NewTargetGrantsInner

`func NewTargetGrantsInner() *TargetGrantsInner`

NewTargetGrantsInner instantiates a new TargetGrantsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGrantsInnerWithDefaults

`func NewTargetGrantsInnerWithDefaults() *TargetGrantsInner`

NewTargetGrantsInnerWithDefaults instantiates a new TargetGrantsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantee

`func (o *TargetGrantsInner) GetGrantee() TargetGrantGrantee`

GetGrantee returns the Grantee field if non-nil, zero value otherwise.

### GetGranteeOk

`func (o *TargetGrantsInner) GetGranteeOk() (*TargetGrantGrantee, bool)`

GetGranteeOk returns a tuple with the Grantee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantee

`func (o *TargetGrantsInner) SetGrantee(v TargetGrantGrantee)`

SetGrantee sets Grantee field to given value.

### HasGrantee

`func (o *TargetGrantsInner) HasGrantee() bool`

HasGrantee returns a boolean if a field has been set.

### GetPermission

`func (o *TargetGrantsInner) GetPermission() BucketLogsPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *TargetGrantsInner) GetPermissionOk() (*BucketLogsPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *TargetGrantsInner) SetPermission(v BucketLogsPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *TargetGrantsInner) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


