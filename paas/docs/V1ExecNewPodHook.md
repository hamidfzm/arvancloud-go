# V1ExecNewPodHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **[]string** | Command is the action command and its arguments. | 
**ContainerName** | **string** | ContainerName is the name of a container in the deployment pod template whose Docker image will be used for the hook pod&#39;s container. | 
**Env** | Pointer to [**[]V1EnvVar**](V1EnvVar.md) | Env is a set of environment variables to supply to the hook pod&#39;s container. | [optional] 
**Volumes** | Pointer to **[]string** | Volumes is a list of named volumes from the pod template which should be copied to the hook pod. Volumes names not found in pod spec are ignored. An empty list means no volumes will be copied. | [optional] 

## Methods

### NewV1ExecNewPodHook

`func NewV1ExecNewPodHook(command []string, containerName string, ) *V1ExecNewPodHook`

NewV1ExecNewPodHook instantiates a new V1ExecNewPodHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ExecNewPodHookWithDefaults

`func NewV1ExecNewPodHookWithDefaults() *V1ExecNewPodHook`

NewV1ExecNewPodHookWithDefaults instantiates a new V1ExecNewPodHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *V1ExecNewPodHook) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V1ExecNewPodHook) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V1ExecNewPodHook) SetCommand(v []string)`

SetCommand sets Command field to given value.


### GetContainerName

`func (o *V1ExecNewPodHook) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *V1ExecNewPodHook) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *V1ExecNewPodHook) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.


### GetEnv

`func (o *V1ExecNewPodHook) GetEnv() []V1EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *V1ExecNewPodHook) GetEnvOk() (*[]V1EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *V1ExecNewPodHook) SetEnv(v []V1EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *V1ExecNewPodHook) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetVolumes

`func (o *V1ExecNewPodHook) GetVolumes() []string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *V1ExecNewPodHook) GetVolumesOk() (*[]string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *V1ExecNewPodHook) SetVolumes(v []string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *V1ExecNewPodHook) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


