# V1CustomBuildStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildAPIVersion** | Pointer to **string** | buildAPIVersion is the requested API version for the Build object serialized and passed to the custom builder | [optional] 
**Env** | Pointer to [**[]V1EnvVar**](V1EnvVar.md) | env contains additional environment variables you want to pass into a builder container. | [optional] 
**ExposeDockerSocket** | Pointer to **bool** | exposeDockerSocket will allow running Docker commands (and build Docker images) from inside the Docker container. | [optional] 
**ForcePull** | Pointer to **bool** | forcePull describes if the controller should configure the build pod to always pull the images for the builder or only pull if it is not present locally | [optional] 
**From** | [**V1ObjectReference**](V1ObjectReference.md) |  | 
**PullSecret** | Pointer to [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] 
**Secrets** | Pointer to [**[]V1SecretSpec**](V1SecretSpec.md) | secrets is a list of additional secrets that will be included in the build pod | [optional] 

## Methods

### NewV1CustomBuildStrategy

`func NewV1CustomBuildStrategy(from V1ObjectReference, ) *V1CustomBuildStrategy`

NewV1CustomBuildStrategy instantiates a new V1CustomBuildStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CustomBuildStrategyWithDefaults

`func NewV1CustomBuildStrategyWithDefaults() *V1CustomBuildStrategy`

NewV1CustomBuildStrategyWithDefaults instantiates a new V1CustomBuildStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildAPIVersion

`func (o *V1CustomBuildStrategy) GetBuildAPIVersion() string`

GetBuildAPIVersion returns the BuildAPIVersion field if non-nil, zero value otherwise.

### GetBuildAPIVersionOk

`func (o *V1CustomBuildStrategy) GetBuildAPIVersionOk() (*string, bool)`

GetBuildAPIVersionOk returns a tuple with the BuildAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildAPIVersion

`func (o *V1CustomBuildStrategy) SetBuildAPIVersion(v string)`

SetBuildAPIVersion sets BuildAPIVersion field to given value.

### HasBuildAPIVersion

`func (o *V1CustomBuildStrategy) HasBuildAPIVersion() bool`

HasBuildAPIVersion returns a boolean if a field has been set.

### GetEnv

`func (o *V1CustomBuildStrategy) GetEnv() []V1EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *V1CustomBuildStrategy) GetEnvOk() (*[]V1EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *V1CustomBuildStrategy) SetEnv(v []V1EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *V1CustomBuildStrategy) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetExposeDockerSocket

`func (o *V1CustomBuildStrategy) GetExposeDockerSocket() bool`

GetExposeDockerSocket returns the ExposeDockerSocket field if non-nil, zero value otherwise.

### GetExposeDockerSocketOk

`func (o *V1CustomBuildStrategy) GetExposeDockerSocketOk() (*bool, bool)`

GetExposeDockerSocketOk returns a tuple with the ExposeDockerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeDockerSocket

`func (o *V1CustomBuildStrategy) SetExposeDockerSocket(v bool)`

SetExposeDockerSocket sets ExposeDockerSocket field to given value.

### HasExposeDockerSocket

`func (o *V1CustomBuildStrategy) HasExposeDockerSocket() bool`

HasExposeDockerSocket returns a boolean if a field has been set.

### GetForcePull

`func (o *V1CustomBuildStrategy) GetForcePull() bool`

GetForcePull returns the ForcePull field if non-nil, zero value otherwise.

### GetForcePullOk

`func (o *V1CustomBuildStrategy) GetForcePullOk() (*bool, bool)`

GetForcePullOk returns a tuple with the ForcePull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePull

`func (o *V1CustomBuildStrategy) SetForcePull(v bool)`

SetForcePull sets ForcePull field to given value.

### HasForcePull

`func (o *V1CustomBuildStrategy) HasForcePull() bool`

HasForcePull returns a boolean if a field has been set.

### GetFrom

`func (o *V1CustomBuildStrategy) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1CustomBuildStrategy) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1CustomBuildStrategy) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.


### GetPullSecret

`func (o *V1CustomBuildStrategy) GetPullSecret() V1LocalObjectReference`

GetPullSecret returns the PullSecret field if non-nil, zero value otherwise.

### GetPullSecretOk

`func (o *V1CustomBuildStrategy) GetPullSecretOk() (*V1LocalObjectReference, bool)`

GetPullSecretOk returns a tuple with the PullSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullSecret

`func (o *V1CustomBuildStrategy) SetPullSecret(v V1LocalObjectReference)`

SetPullSecret sets PullSecret field to given value.

### HasPullSecret

`func (o *V1CustomBuildStrategy) HasPullSecret() bool`

HasPullSecret returns a boolean if a field has been set.

### GetSecrets

`func (o *V1CustomBuildStrategy) GetSecrets() []V1SecretSpec`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *V1CustomBuildStrategy) GetSecretsOk() (*[]V1SecretSpec, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *V1CustomBuildStrategy) SetSecrets(v []V1SecretSpec)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *V1CustomBuildStrategy) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


