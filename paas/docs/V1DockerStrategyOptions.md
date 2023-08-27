# V1DockerStrategyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildArgs** | Pointer to [**[]V1EnvVar**](V1EnvVar.md) | Args contains any build arguments that are to be passed to Docker.  See https://docs.docker.com/engine/reference/builder/#/arg for more details | [optional] 
**NoCache** | Pointer to **bool** | noCache overrides the docker-strategy noCache option in the build config | [optional] 

## Methods

### NewV1DockerStrategyOptions

`func NewV1DockerStrategyOptions() *V1DockerStrategyOptions`

NewV1DockerStrategyOptions instantiates a new V1DockerStrategyOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DockerStrategyOptionsWithDefaults

`func NewV1DockerStrategyOptionsWithDefaults() *V1DockerStrategyOptions`

NewV1DockerStrategyOptionsWithDefaults instantiates a new V1DockerStrategyOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildArgs

`func (o *V1DockerStrategyOptions) GetBuildArgs() []V1EnvVar`

GetBuildArgs returns the BuildArgs field if non-nil, zero value otherwise.

### GetBuildArgsOk

`func (o *V1DockerStrategyOptions) GetBuildArgsOk() (*[]V1EnvVar, bool)`

GetBuildArgsOk returns a tuple with the BuildArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildArgs

`func (o *V1DockerStrategyOptions) SetBuildArgs(v []V1EnvVar)`

SetBuildArgs sets BuildArgs field to given value.

### HasBuildArgs

`func (o *V1DockerStrategyOptions) HasBuildArgs() bool`

HasBuildArgs returns a boolean if a field has been set.

### GetNoCache

`func (o *V1DockerStrategyOptions) GetNoCache() bool`

GetNoCache returns the NoCache field if non-nil, zero value otherwise.

### GetNoCacheOk

`func (o *V1DockerStrategyOptions) GetNoCacheOk() (*bool, bool)`

GetNoCacheOk returns a tuple with the NoCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCache

`func (o *V1DockerStrategyOptions) SetNoCache(v bool)`

SetNoCache sets NoCache field to given value.

### HasNoCache

`func (o *V1DockerStrategyOptions) HasNoCache() bool`

HasNoCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


