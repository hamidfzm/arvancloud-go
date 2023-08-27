# V1JenkinsPipelineBuildStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to [**[]V1EnvVar**](V1EnvVar.md) | env contains additional environment variables you want to pass into a build pipeline. | [optional] 
**Jenkinsfile** | Pointer to **string** | Jenkinsfile defines the optional raw contents of a Jenkinsfile which defines a Jenkins pipeline build. | [optional] 
**JenkinsfilePath** | Pointer to **string** | JenkinsfilePath is the optional path of the Jenkinsfile that will be used to configure the pipeline relative to the root of the context (contextDir). If both JenkinsfilePath &amp; Jenkinsfile are both not specified, this defaults to Jenkinsfile in the root of the specified contextDir. | [optional] 

## Methods

### NewV1JenkinsPipelineBuildStrategy

`func NewV1JenkinsPipelineBuildStrategy() *V1JenkinsPipelineBuildStrategy`

NewV1JenkinsPipelineBuildStrategy instantiates a new V1JenkinsPipelineBuildStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1JenkinsPipelineBuildStrategyWithDefaults

`func NewV1JenkinsPipelineBuildStrategyWithDefaults() *V1JenkinsPipelineBuildStrategy`

NewV1JenkinsPipelineBuildStrategyWithDefaults instantiates a new V1JenkinsPipelineBuildStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *V1JenkinsPipelineBuildStrategy) GetEnv() []V1EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *V1JenkinsPipelineBuildStrategy) GetEnvOk() (*[]V1EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *V1JenkinsPipelineBuildStrategy) SetEnv(v []V1EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *V1JenkinsPipelineBuildStrategy) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetJenkinsfile

`func (o *V1JenkinsPipelineBuildStrategy) GetJenkinsfile() string`

GetJenkinsfile returns the Jenkinsfile field if non-nil, zero value otherwise.

### GetJenkinsfileOk

`func (o *V1JenkinsPipelineBuildStrategy) GetJenkinsfileOk() (*string, bool)`

GetJenkinsfileOk returns a tuple with the Jenkinsfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsfile

`func (o *V1JenkinsPipelineBuildStrategy) SetJenkinsfile(v string)`

SetJenkinsfile sets Jenkinsfile field to given value.

### HasJenkinsfile

`func (o *V1JenkinsPipelineBuildStrategy) HasJenkinsfile() bool`

HasJenkinsfile returns a boolean if a field has been set.

### GetJenkinsfilePath

`func (o *V1JenkinsPipelineBuildStrategy) GetJenkinsfilePath() string`

GetJenkinsfilePath returns the JenkinsfilePath field if non-nil, zero value otherwise.

### GetJenkinsfilePathOk

`func (o *V1JenkinsPipelineBuildStrategy) GetJenkinsfilePathOk() (*string, bool)`

GetJenkinsfilePathOk returns a tuple with the JenkinsfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsfilePath

`func (o *V1JenkinsPipelineBuildStrategy) SetJenkinsfilePath(v string)`

SetJenkinsfilePath sets JenkinsfilePath field to given value.

### HasJenkinsfilePath

`func (o *V1JenkinsPipelineBuildStrategy) HasJenkinsfilePath() bool`

HasJenkinsfilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


