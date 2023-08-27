# V1BitbucketWebHookCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to [**V1SourceRevision**](V1SourceRevision.md) |  | [optional] 
**Secret** | Pointer to **string** | Secret is the obfuscated webhook secret that triggered a build. | [optional] 

## Methods

### NewV1BitbucketWebHookCause

`func NewV1BitbucketWebHookCause() *V1BitbucketWebHookCause`

NewV1BitbucketWebHookCause instantiates a new V1BitbucketWebHookCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BitbucketWebHookCauseWithDefaults

`func NewV1BitbucketWebHookCauseWithDefaults() *V1BitbucketWebHookCause`

NewV1BitbucketWebHookCauseWithDefaults instantiates a new V1BitbucketWebHookCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *V1BitbucketWebHookCause) GetRevision() V1SourceRevision`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1BitbucketWebHookCause) GetRevisionOk() (*V1SourceRevision, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1BitbucketWebHookCause) SetRevision(v V1SourceRevision)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1BitbucketWebHookCause) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSecret

`func (o *V1BitbucketWebHookCause) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V1BitbucketWebHookCause) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V1BitbucketWebHookCause) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *V1BitbucketWebHookCause) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


