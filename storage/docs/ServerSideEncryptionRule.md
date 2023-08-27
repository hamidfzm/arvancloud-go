# ServerSideEncryptionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyServerSideEncryptionByDefault** | Pointer to [**ServerSideEncryptionRuleApplyServerSideEncryptionByDefault**](ServerSideEncryptionRuleApplyServerSideEncryptionByDefault.md) |  | [optional] 
**BucketKeyEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewServerSideEncryptionRule

`func NewServerSideEncryptionRule() *ServerSideEncryptionRule`

NewServerSideEncryptionRule instantiates a new ServerSideEncryptionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerSideEncryptionRuleWithDefaults

`func NewServerSideEncryptionRuleWithDefaults() *ServerSideEncryptionRule`

NewServerSideEncryptionRuleWithDefaults instantiates a new ServerSideEncryptionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyServerSideEncryptionByDefault

`func (o *ServerSideEncryptionRule) GetApplyServerSideEncryptionByDefault() ServerSideEncryptionRuleApplyServerSideEncryptionByDefault`

GetApplyServerSideEncryptionByDefault returns the ApplyServerSideEncryptionByDefault field if non-nil, zero value otherwise.

### GetApplyServerSideEncryptionByDefaultOk

`func (o *ServerSideEncryptionRule) GetApplyServerSideEncryptionByDefaultOk() (*ServerSideEncryptionRuleApplyServerSideEncryptionByDefault, bool)`

GetApplyServerSideEncryptionByDefaultOk returns a tuple with the ApplyServerSideEncryptionByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyServerSideEncryptionByDefault

`func (o *ServerSideEncryptionRule) SetApplyServerSideEncryptionByDefault(v ServerSideEncryptionRuleApplyServerSideEncryptionByDefault)`

SetApplyServerSideEncryptionByDefault sets ApplyServerSideEncryptionByDefault field to given value.

### HasApplyServerSideEncryptionByDefault

`func (o *ServerSideEncryptionRule) HasApplyServerSideEncryptionByDefault() bool`

HasApplyServerSideEncryptionByDefault returns a boolean if a field has been set.

### GetBucketKeyEnabled

`func (o *ServerSideEncryptionRule) GetBucketKeyEnabled() bool`

GetBucketKeyEnabled returns the BucketKeyEnabled field if non-nil, zero value otherwise.

### GetBucketKeyEnabledOk

`func (o *ServerSideEncryptionRule) GetBucketKeyEnabledOk() (*bool, bool)`

GetBucketKeyEnabledOk returns a tuple with the BucketKeyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketKeyEnabled

`func (o *ServerSideEncryptionRule) SetBucketKeyEnabled(v bool)`

SetBucketKeyEnabled sets BucketKeyEnabled field to given value.

### HasBucketKeyEnabled

`func (o *ServerSideEncryptionRule) HasBucketKeyEnabled() bool`

HasBucketKeyEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


