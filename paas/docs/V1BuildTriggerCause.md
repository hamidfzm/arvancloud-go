# V1BuildTriggerCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BitbucketWebHook** | Pointer to [**V1BitbucketWebHookCause**](V1BitbucketWebHookCause.md) |  | [optional] 
**GenericWebHook** | Pointer to [**V1GenericWebHookCause**](V1GenericWebHookCause.md) |  | [optional] 
**GithubWebHook** | Pointer to [**V1GitHubWebHookCause**](V1GitHubWebHookCause.md) |  | [optional] 
**GitlabWebHook** | Pointer to [**V1GitLabWebHookCause**](V1GitLabWebHookCause.md) |  | [optional] 
**ImageChangeBuild** | Pointer to [**V1ImageChangeCause**](V1ImageChangeCause.md) |  | [optional] 
**Message** | Pointer to **string** | message is used to store a human readable message for why the build was triggered. E.g.: \&quot;Manually triggered by user\&quot;, \&quot;Configuration change\&quot;,etc. | [optional] 

## Methods

### NewV1BuildTriggerCause

`func NewV1BuildTriggerCause() *V1BuildTriggerCause`

NewV1BuildTriggerCause instantiates a new V1BuildTriggerCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildTriggerCauseWithDefaults

`func NewV1BuildTriggerCauseWithDefaults() *V1BuildTriggerCause`

NewV1BuildTriggerCauseWithDefaults instantiates a new V1BuildTriggerCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBitbucketWebHook

`func (o *V1BuildTriggerCause) GetBitbucketWebHook() V1BitbucketWebHookCause`

GetBitbucketWebHook returns the BitbucketWebHook field if non-nil, zero value otherwise.

### GetBitbucketWebHookOk

`func (o *V1BuildTriggerCause) GetBitbucketWebHookOk() (*V1BitbucketWebHookCause, bool)`

GetBitbucketWebHookOk returns a tuple with the BitbucketWebHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitbucketWebHook

`func (o *V1BuildTriggerCause) SetBitbucketWebHook(v V1BitbucketWebHookCause)`

SetBitbucketWebHook sets BitbucketWebHook field to given value.

### HasBitbucketWebHook

`func (o *V1BuildTriggerCause) HasBitbucketWebHook() bool`

HasBitbucketWebHook returns a boolean if a field has been set.

### GetGenericWebHook

`func (o *V1BuildTriggerCause) GetGenericWebHook() V1GenericWebHookCause`

GetGenericWebHook returns the GenericWebHook field if non-nil, zero value otherwise.

### GetGenericWebHookOk

`func (o *V1BuildTriggerCause) GetGenericWebHookOk() (*V1GenericWebHookCause, bool)`

GetGenericWebHookOk returns a tuple with the GenericWebHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericWebHook

`func (o *V1BuildTriggerCause) SetGenericWebHook(v V1GenericWebHookCause)`

SetGenericWebHook sets GenericWebHook field to given value.

### HasGenericWebHook

`func (o *V1BuildTriggerCause) HasGenericWebHook() bool`

HasGenericWebHook returns a boolean if a field has been set.

### GetGithubWebHook

`func (o *V1BuildTriggerCause) GetGithubWebHook() V1GitHubWebHookCause`

GetGithubWebHook returns the GithubWebHook field if non-nil, zero value otherwise.

### GetGithubWebHookOk

`func (o *V1BuildTriggerCause) GetGithubWebHookOk() (*V1GitHubWebHookCause, bool)`

GetGithubWebHookOk returns a tuple with the GithubWebHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubWebHook

`func (o *V1BuildTriggerCause) SetGithubWebHook(v V1GitHubWebHookCause)`

SetGithubWebHook sets GithubWebHook field to given value.

### HasGithubWebHook

`func (o *V1BuildTriggerCause) HasGithubWebHook() bool`

HasGithubWebHook returns a boolean if a field has been set.

### GetGitlabWebHook

`func (o *V1BuildTriggerCause) GetGitlabWebHook() V1GitLabWebHookCause`

GetGitlabWebHook returns the GitlabWebHook field if non-nil, zero value otherwise.

### GetGitlabWebHookOk

`func (o *V1BuildTriggerCause) GetGitlabWebHookOk() (*V1GitLabWebHookCause, bool)`

GetGitlabWebHookOk returns a tuple with the GitlabWebHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabWebHook

`func (o *V1BuildTriggerCause) SetGitlabWebHook(v V1GitLabWebHookCause)`

SetGitlabWebHook sets GitlabWebHook field to given value.

### HasGitlabWebHook

`func (o *V1BuildTriggerCause) HasGitlabWebHook() bool`

HasGitlabWebHook returns a boolean if a field has been set.

### GetImageChangeBuild

`func (o *V1BuildTriggerCause) GetImageChangeBuild() V1ImageChangeCause`

GetImageChangeBuild returns the ImageChangeBuild field if non-nil, zero value otherwise.

### GetImageChangeBuildOk

`func (o *V1BuildTriggerCause) GetImageChangeBuildOk() (*V1ImageChangeCause, bool)`

GetImageChangeBuildOk returns a tuple with the ImageChangeBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageChangeBuild

`func (o *V1BuildTriggerCause) SetImageChangeBuild(v V1ImageChangeCause)`

SetImageChangeBuild sets ImageChangeBuild field to given value.

### HasImageChangeBuild

`func (o *V1BuildTriggerCause) HasImageChangeBuild() bool`

HasImageChangeBuild returns a boolean if a field has been set.

### GetMessage

`func (o *V1BuildTriggerCause) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1BuildTriggerCause) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1BuildTriggerCause) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1BuildTriggerCause) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


