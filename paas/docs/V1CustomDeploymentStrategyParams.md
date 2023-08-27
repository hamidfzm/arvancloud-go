# V1CustomDeploymentStrategyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **[]string** | Command is optional and overrides CMD in the container Image. | [optional] 
**Environment** | Pointer to [**[]V1EnvVar**](V1EnvVar.md) | Environment holds the environment which will be given to the container for Image. | [optional] 
**Image** | Pointer to **string** | Image specifies a Docker image which can carry out a deployment. | [optional] 

## Methods

### NewV1CustomDeploymentStrategyParams

`func NewV1CustomDeploymentStrategyParams() *V1CustomDeploymentStrategyParams`

NewV1CustomDeploymentStrategyParams instantiates a new V1CustomDeploymentStrategyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CustomDeploymentStrategyParamsWithDefaults

`func NewV1CustomDeploymentStrategyParamsWithDefaults() *V1CustomDeploymentStrategyParams`

NewV1CustomDeploymentStrategyParamsWithDefaults instantiates a new V1CustomDeploymentStrategyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *V1CustomDeploymentStrategyParams) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V1CustomDeploymentStrategyParams) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V1CustomDeploymentStrategyParams) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V1CustomDeploymentStrategyParams) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEnvironment

`func (o *V1CustomDeploymentStrategyParams) GetEnvironment() []V1EnvVar`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *V1CustomDeploymentStrategyParams) GetEnvironmentOk() (*[]V1EnvVar, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *V1CustomDeploymentStrategyParams) SetEnvironment(v []V1EnvVar)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *V1CustomDeploymentStrategyParams) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetImage

`func (o *V1CustomDeploymentStrategyParams) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V1CustomDeploymentStrategyParams) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V1CustomDeploymentStrategyParams) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *V1CustomDeploymentStrategyParams) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


