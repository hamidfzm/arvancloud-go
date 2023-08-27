# PutBucketEncryptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerSideEncryptionConfiguration** | [**ServerSideEncryptionConfiguration**](ServerSideEncryptionConfiguration.md) |  | 

## Methods

### NewPutBucketEncryptionRequest

`func NewPutBucketEncryptionRequest(serverSideEncryptionConfiguration ServerSideEncryptionConfiguration, ) *PutBucketEncryptionRequest`

NewPutBucketEncryptionRequest instantiates a new PutBucketEncryptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketEncryptionRequestWithDefaults

`func NewPutBucketEncryptionRequestWithDefaults() *PutBucketEncryptionRequest`

NewPutBucketEncryptionRequestWithDefaults instantiates a new PutBucketEncryptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerSideEncryptionConfiguration

`func (o *PutBucketEncryptionRequest) GetServerSideEncryptionConfiguration() ServerSideEncryptionConfiguration`

GetServerSideEncryptionConfiguration returns the ServerSideEncryptionConfiguration field if non-nil, zero value otherwise.

### GetServerSideEncryptionConfigurationOk

`func (o *PutBucketEncryptionRequest) GetServerSideEncryptionConfigurationOk() (*ServerSideEncryptionConfiguration, bool)`

GetServerSideEncryptionConfigurationOk returns a tuple with the ServerSideEncryptionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideEncryptionConfiguration

`func (o *PutBucketEncryptionRequest) SetServerSideEncryptionConfiguration(v ServerSideEncryptionConfiguration)`

SetServerSideEncryptionConfiguration sets ServerSideEncryptionConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


