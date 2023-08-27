# PutBucketPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | The bucket policy as a JSON document. | 

## Methods

### NewPutBucketPolicyRequest

`func NewPutBucketPolicyRequest(policy string, ) *PutBucketPolicyRequest`

NewPutBucketPolicyRequest instantiates a new PutBucketPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketPolicyRequestWithDefaults

`func NewPutBucketPolicyRequestWithDefaults() *PutBucketPolicyRequest`

NewPutBucketPolicyRequestWithDefaults instantiates a new PutBucketPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *PutBucketPolicyRequest) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PutBucketPolicyRequest) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PutBucketPolicyRequest) SetPolicy(v string)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


