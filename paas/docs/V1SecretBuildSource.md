# V1SecretBuildSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDir** | Pointer to **string** | destinationDir is the directory where the files from the secret should be available for the build time. For the Source build strategy, these will be injected into a container where the assemble script runs. Later, when the script finishes, all files injected will be truncated to zero length. For the Docker build strategy, these will be copied into the build directory, where the Dockerfile is located, so users can ADD or COPY them during docker build. | [optional] 
**Secret** | [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | 

## Methods

### NewV1SecretBuildSource

`func NewV1SecretBuildSource(secret V1LocalObjectReference, ) *V1SecretBuildSource`

NewV1SecretBuildSource instantiates a new V1SecretBuildSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SecretBuildSourceWithDefaults

`func NewV1SecretBuildSourceWithDefaults() *V1SecretBuildSource`

NewV1SecretBuildSourceWithDefaults instantiates a new V1SecretBuildSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDir

`func (o *V1SecretBuildSource) GetDestinationDir() string`

GetDestinationDir returns the DestinationDir field if non-nil, zero value otherwise.

### GetDestinationDirOk

`func (o *V1SecretBuildSource) GetDestinationDirOk() (*string, bool)`

GetDestinationDirOk returns a tuple with the DestinationDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDir

`func (o *V1SecretBuildSource) SetDestinationDir(v string)`

SetDestinationDir sets DestinationDir field to given value.

### HasDestinationDir

`func (o *V1SecretBuildSource) HasDestinationDir() bool`

HasDestinationDir returns a boolean if a field has been set.

### GetSecret

`func (o *V1SecretBuildSource) GetSecret() V1LocalObjectReference`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *V1SecretBuildSource) GetSecretOk() (*V1LocalObjectReference, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *V1SecretBuildSource) SetSecret(v V1LocalObjectReference)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


