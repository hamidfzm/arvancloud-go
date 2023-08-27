# V1SourceBuildStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | Pointer to [**[]V1EnvVar**](V1EnvVar.md) | env contains additional environment variables you want to pass into a builder container. | [optional] 
**ForcePull** | Pointer to **bool** | forcePull describes if the builder should pull the images from registry prior to building. | [optional] 
**From** | [**V1ObjectReference**](V1ObjectReference.md) |  | 
**Incremental** | Pointer to **bool** | incremental flag forces the Source build to do incremental builds if true. | [optional] 
**PullSecret** | Pointer to [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] 
**Scripts** | Pointer to **string** | scripts is the location of Source scripts | [optional] 

## Methods

### NewV1SourceBuildStrategy

`func NewV1SourceBuildStrategy(from V1ObjectReference, ) *V1SourceBuildStrategy`

NewV1SourceBuildStrategy instantiates a new V1SourceBuildStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SourceBuildStrategyWithDefaults

`func NewV1SourceBuildStrategyWithDefaults() *V1SourceBuildStrategy`

NewV1SourceBuildStrategyWithDefaults instantiates a new V1SourceBuildStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnv

`func (o *V1SourceBuildStrategy) GetEnv() []V1EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *V1SourceBuildStrategy) GetEnvOk() (*[]V1EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *V1SourceBuildStrategy) SetEnv(v []V1EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *V1SourceBuildStrategy) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetForcePull

`func (o *V1SourceBuildStrategy) GetForcePull() bool`

GetForcePull returns the ForcePull field if non-nil, zero value otherwise.

### GetForcePullOk

`func (o *V1SourceBuildStrategy) GetForcePullOk() (*bool, bool)`

GetForcePullOk returns a tuple with the ForcePull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePull

`func (o *V1SourceBuildStrategy) SetForcePull(v bool)`

SetForcePull sets ForcePull field to given value.

### HasForcePull

`func (o *V1SourceBuildStrategy) HasForcePull() bool`

HasForcePull returns a boolean if a field has been set.

### GetFrom

`func (o *V1SourceBuildStrategy) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1SourceBuildStrategy) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1SourceBuildStrategy) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.


### GetIncremental

`func (o *V1SourceBuildStrategy) GetIncremental() bool`

GetIncremental returns the Incremental field if non-nil, zero value otherwise.

### GetIncrementalOk

`func (o *V1SourceBuildStrategy) GetIncrementalOk() (*bool, bool)`

GetIncrementalOk returns a tuple with the Incremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncremental

`func (o *V1SourceBuildStrategy) SetIncremental(v bool)`

SetIncremental sets Incremental field to given value.

### HasIncremental

`func (o *V1SourceBuildStrategy) HasIncremental() bool`

HasIncremental returns a boolean if a field has been set.

### GetPullSecret

`func (o *V1SourceBuildStrategy) GetPullSecret() V1LocalObjectReference`

GetPullSecret returns the PullSecret field if non-nil, zero value otherwise.

### GetPullSecretOk

`func (o *V1SourceBuildStrategy) GetPullSecretOk() (*V1LocalObjectReference, bool)`

GetPullSecretOk returns a tuple with the PullSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullSecret

`func (o *V1SourceBuildStrategy) SetPullSecret(v V1LocalObjectReference)`

SetPullSecret sets PullSecret field to given value.

### HasPullSecret

`func (o *V1SourceBuildStrategy) HasPullSecret() bool`

HasPullSecret returns a boolean if a field has been set.

### GetScripts

`func (o *V1SourceBuildStrategy) GetScripts() string`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *V1SourceBuildStrategy) GetScriptsOk() (*string, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *V1SourceBuildStrategy) SetScripts(v string)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *V1SourceBuildStrategy) HasScripts() bool`

HasScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


