# V1DockerBuildStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildArgs** | Pointer to [**[]V1EnvVar**](V1EnvVar.md) | buildArgs contains build arguments that will be resolved in the Dockerfile.  See https://docs.docker.com/engine/reference/builder/#/arg for more details. | [optional] 
**DockerfilePath** | Pointer to **string** | dockerfilePath is the path of the Dockerfile that will be used to build the Docker image, relative to the root of the context (contextDir). | [optional] 
**Env** | Pointer to [**[]V1EnvVar**](V1EnvVar.md) | env contains additional environment variables you want to pass into a builder container. | [optional] 
**ForcePull** | Pointer to **bool** | forcePull describes if the builder should pull the images from registry prior to building. | [optional] 
**From** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 
**ImageOptimizationPolicy** | Pointer to [**V1ImageOptimizationPolicy**](V1ImageOptimizationPolicy.md) |  | [optional] 
**NoCache** | Pointer to **bool** | noCache if set to true indicates that the docker build must be executed with the --no-cache&#x3D;true flag | [optional] 
**PullSecret** | Pointer to [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] 

## Methods

### NewV1DockerBuildStrategy

`func NewV1DockerBuildStrategy() *V1DockerBuildStrategy`

NewV1DockerBuildStrategy instantiates a new V1DockerBuildStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DockerBuildStrategyWithDefaults

`func NewV1DockerBuildStrategyWithDefaults() *V1DockerBuildStrategy`

NewV1DockerBuildStrategyWithDefaults instantiates a new V1DockerBuildStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildArgs

`func (o *V1DockerBuildStrategy) GetBuildArgs() []V1EnvVar`

GetBuildArgs returns the BuildArgs field if non-nil, zero value otherwise.

### GetBuildArgsOk

`func (o *V1DockerBuildStrategy) GetBuildArgsOk() (*[]V1EnvVar, bool)`

GetBuildArgsOk returns a tuple with the BuildArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildArgs

`func (o *V1DockerBuildStrategy) SetBuildArgs(v []V1EnvVar)`

SetBuildArgs sets BuildArgs field to given value.

### HasBuildArgs

`func (o *V1DockerBuildStrategy) HasBuildArgs() bool`

HasBuildArgs returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *V1DockerBuildStrategy) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *V1DockerBuildStrategy) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *V1DockerBuildStrategy) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *V1DockerBuildStrategy) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### GetEnv

`func (o *V1DockerBuildStrategy) GetEnv() []V1EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *V1DockerBuildStrategy) GetEnvOk() (*[]V1EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *V1DockerBuildStrategy) SetEnv(v []V1EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *V1DockerBuildStrategy) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetForcePull

`func (o *V1DockerBuildStrategy) GetForcePull() bool`

GetForcePull returns the ForcePull field if non-nil, zero value otherwise.

### GetForcePullOk

`func (o *V1DockerBuildStrategy) GetForcePullOk() (*bool, bool)`

GetForcePullOk returns a tuple with the ForcePull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePull

`func (o *V1DockerBuildStrategy) SetForcePull(v bool)`

SetForcePull sets ForcePull field to given value.

### HasForcePull

`func (o *V1DockerBuildStrategy) HasForcePull() bool`

HasForcePull returns a boolean if a field has been set.

### GetFrom

`func (o *V1DockerBuildStrategy) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1DockerBuildStrategy) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1DockerBuildStrategy) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.

### HasFrom

`func (o *V1DockerBuildStrategy) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetImageOptimizationPolicy

`func (o *V1DockerBuildStrategy) GetImageOptimizationPolicy() V1ImageOptimizationPolicy`

GetImageOptimizationPolicy returns the ImageOptimizationPolicy field if non-nil, zero value otherwise.

### GetImageOptimizationPolicyOk

`func (o *V1DockerBuildStrategy) GetImageOptimizationPolicyOk() (*V1ImageOptimizationPolicy, bool)`

GetImageOptimizationPolicyOk returns a tuple with the ImageOptimizationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptimizationPolicy

`func (o *V1DockerBuildStrategy) SetImageOptimizationPolicy(v V1ImageOptimizationPolicy)`

SetImageOptimizationPolicy sets ImageOptimizationPolicy field to given value.

### HasImageOptimizationPolicy

`func (o *V1DockerBuildStrategy) HasImageOptimizationPolicy() bool`

HasImageOptimizationPolicy returns a boolean if a field has been set.

### GetNoCache

`func (o *V1DockerBuildStrategy) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *V1DockerBuildStrategy) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *V1DockerBuildStrategy) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *V1DockerBuildStrategy) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.

### GetPullSecret

`func (o *V1DockerBuildStrategy) GetPullSecret() V1LocalObjectReference`

GetPullSecret returns the PullSecret field if non-nil, zero value otherwise.

### GetPullSecretOk

`func (o *V1DockerBuildStrategy) GetPullSecretOk() (*V1LocalObjectReference, bool)`

GetPullSecretOk returns a tuple with the PullSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullSecret

`func (o *V1DockerBuildStrategy) SetPullSecret(v V1LocalObjectReference)`

SetPullSecret sets PullSecret field to given value.

### HasPullSecret

`func (o *V1DockerBuildStrategy) HasPullSecret() bool`

HasPullSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


