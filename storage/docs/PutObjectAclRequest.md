# PutObjectAclRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControlPolicy** | Pointer to [**PutBucketAclRequestAccessControlPolicy**](PutBucketAclRequestAccessControlPolicy.md) |  | [optional] 

## Methods

### NewPutObjectAclRequest

`func NewPutObjectAclRequest() *PutObjectAclRequest`

NewPutObjectAclRequest instantiates a new PutObjectAclRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectAclRequestWithDefaults

`func NewPutObjectAclRequestWithDefaults() *PutObjectAclRequest`

NewPutObjectAclRequestWithDefaults instantiates a new PutObjectAclRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControlPolicy

`func (o *PutObjectAclRequest) GetAccessControlPolicy() PutBucketAclRequestAccessControlPolicy`

GetAccessControlPolicy returns the AccessControlPolicy field if non-nil, zero value otherwise.

### GetAccessControlPolicyOk

`func (o *PutObjectAclRequest) GetAccessControlPolicyOk() (*PutBucketAclRequestAccessControlPolicy, bool)`

GetAccessControlPolicyOk returns a tuple with the AccessControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlPolicy

`func (o *PutObjectAclRequest) SetAccessControlPolicy(v PutBucketAclRequestAccessControlPolicy)`

SetAccessControlPolicy sets AccessControlPolicy field to given value.

### HasAccessControlPolicy

`func (o *PutObjectAclRequest) HasAccessControlPolicy() bool`

HasAccessControlPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


