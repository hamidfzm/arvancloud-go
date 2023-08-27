# V1BuildTriggerPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bitbucket** | Pointer to [**V1WebHookTrigger**](V1WebHookTrigger.md) |  | [optional] 
**Generic** | Pointer to [**V1WebHookTrigger**](V1WebHookTrigger.md) |  | [optional] 
**Github** | Pointer to [**V1WebHookTrigger**](V1WebHookTrigger.md) |  | [optional] 
**Gitlab** | Pointer to [**V1WebHookTrigger**](V1WebHookTrigger.md) |  | [optional] 
**ImageChange** | Pointer to [**V1ImageChangeTrigger**](V1ImageChangeTrigger.md) |  | [optional] 
**Type** | **string** | type is the type of build trigger | 

## Methods

### NewV1BuildTriggerPolicy

`func NewV1BuildTriggerPolicy(type_ string, ) *V1BuildTriggerPolicy`

NewV1BuildTriggerPolicy instantiates a new V1BuildTriggerPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildTriggerPolicyWithDefaults

`func NewV1BuildTriggerPolicyWithDefaults() *V1BuildTriggerPolicy`

NewV1BuildTriggerPolicyWithDefaults instantiates a new V1BuildTriggerPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitbucket

`func (o *V1BuildTriggerPolicy) GetBitbucket() V1WebHookTrigger`

GetBitbucket returns the Bitbucket field if non-nil, zero value otherwise.

### GetBitbucketOk

`func (o *V1BuildTriggerPolicy) GetBitbucketOk() (*V1WebHookTrigger, bool)`

GetBitbucketOk returns a tuple with the Bitbucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitbucket

`func (o *V1BuildTriggerPolicy) SetBitbucket(v V1WebHookTrigger)`

SetBitbucket sets Bitbucket field to given value.

### HasBitbucket

`func (o *V1BuildTriggerPolicy) HasBitbucket() bool`

HasBitbucket returns a boolean if a field has been set.

### GetGeneric

`func (o *V1BuildTriggerPolicy) GetGeneric() V1WebHookTrigger`

GetGeneric returns the Generic field if non-nil, zero value otherwise.

### GetGenericOk

`func (o *V1BuildTriggerPolicy) GetGenericOk() (*V1WebHookTrigger, bool)`

GetGenericOk returns a tuple with the Generic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneric

`func (o *V1BuildTriggerPolicy) SetGeneric(v V1WebHookTrigger)`

SetGeneric sets Generic field to given value.

### HasGeneric

`func (o *V1BuildTriggerPolicy) HasGeneric() bool`

HasGeneric returns a boolean if a field has been set.

### GetGithub

`func (o *V1BuildTriggerPolicy) GetGithub() V1WebHookTrigger`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *V1BuildTriggerPolicy) GetGithubOk() (*V1WebHookTrigger, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *V1BuildTriggerPolicy) SetGithub(v V1WebHookTrigger)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *V1BuildTriggerPolicy) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetGitlab

`func (o *V1BuildTriggerPolicy) GetGitlab() V1WebHookTrigger`

GetGitlab returns the Gitlab field if non-nil, zero value otherwise.

### GetGitlabOk

`func (o *V1BuildTriggerPolicy) GetGitlabOk() (*V1WebHookTrigger, bool)`

GetGitlabOk returns a tuple with the Gitlab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlab

`func (o *V1BuildTriggerPolicy) SetGitlab(v V1WebHookTrigger)`

SetGitlab sets Gitlab field to given value.

### HasGitlab

`func (o *V1BuildTriggerPolicy) HasGitlab() bool`

HasGitlab returns a boolean if a field has been set.

### GetImageChange

`func (o *V1BuildTriggerPolicy) GetImageChange() V1ImageChangeTrigger`

GetImageChange returns the ImageChange field if non-nil, zero value otherwise.

### GetImageChangeOk

`func (o *V1BuildTriggerPolicy) GetImageChangeOk() (*V1ImageChangeTrigger, bool)`

GetImageChangeOk returns a tuple with the ImageChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageChange

`func (o *V1BuildTriggerPolicy) SetImageChange(v V1ImageChangeTrigger)`

SetImageChange sets ImageChange field to given value.

### HasImageChange

`func (o *V1BuildTriggerPolicy) HasImageChange() bool`

HasImageChange returns a boolean if a field has been set.

### GetType

`func (o *V1BuildTriggerPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1BuildTriggerPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1BuildTriggerPolicy) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


