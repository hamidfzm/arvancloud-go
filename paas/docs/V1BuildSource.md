# V1BuildSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binary** | Pointer to [**V1BinaryBuildSource**](V1BinaryBuildSource.md) |  | [optional] 
**ConfigMaps** | Pointer to [**[]V1ConfigMapBuildSource**](V1ConfigMapBuildSource.md) | configMaps represents a list of configMaps and their destinations that will be used for the build. | [optional] 
**ContextDir** | Pointer to **string** | contextDir specifies the sub-directory where the source code for the application exists. This allows to have buildable sources in directory other than root of repository. | [optional] 
**Dockerfile** | Pointer to **string** | dockerfile is the raw contents of a Dockerfile which should be built. When this option is specified, the FROM may be modified based on your strategy base image and additional ENV stanzas from your strategy environment will be added after the FROM, but before the rest of your Dockerfile stanzas. The Dockerfile source type may be used with other options like git - in those cases the Git repo will have any innate Dockerfile replaced in the context dir. | [optional] 
**Git** | Pointer to [**V1GitBuildSource**](V1GitBuildSource.md) |  | [optional] 
**Images** | Pointer to [**[]V1ImageSource**](V1ImageSource.md) | images describes a set of images to be used to provide source for the build | [optional] 
**Secrets** | Pointer to [**[]V1SecretBuildSource**](V1SecretBuildSource.md) | secrets represents a list of secrets and their destinations that will be used only for the build. | [optional] 
**SourceSecret** | Pointer to [**V1LocalObjectReference**](V1LocalObjectReference.md) |  | [optional] 
**Type** | **string** | type of build input to accept | 

## Methods

### NewV1BuildSource

`func NewV1BuildSource(type_ string, ) *V1BuildSource`

NewV1BuildSource instantiates a new V1BuildSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildSourceWithDefaults

`func NewV1BuildSourceWithDefaults() *V1BuildSource`

NewV1BuildSourceWithDefaults instantiates a new V1BuildSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinary

`func (o *V1BuildSource) GetBinary() V1BinaryBuildSource`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *V1BuildSource) GetBinaryOk() (*V1BinaryBuildSource, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *V1BuildSource) SetBinary(v V1BinaryBuildSource)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *V1BuildSource) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetConfigMaps

`func (o *V1BuildSource) GetConfigMaps() []V1ConfigMapBuildSource`

GetConfigMaps returns the ConfigMaps field if non-nil, zero value otherwise.

### GetConfigMapsOk

`func (o *V1BuildSource) GetConfigMapsOk() (*[]V1ConfigMapBuildSource, bool)`

GetConfigMapsOk returns a tuple with the ConfigMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMaps

`func (o *V1BuildSource) SetConfigMaps(v []V1ConfigMapBuildSource)`

SetConfigMaps sets ConfigMaps field to given value.

### HasConfigMaps

`func (o *V1BuildSource) HasConfigMaps() bool`

HasConfigMaps returns a boolean if a field has been set.

### GetContextDir

`func (o *V1BuildSource) GetContextDir() string`

GetContextDir returns the ContextDir field if non-nil, zero value otherwise.

### GetContextDirOk

`func (o *V1BuildSource) GetContextDirOk() (*string, bool)`

GetContextDirOk returns a tuple with the ContextDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextDir

`func (o *V1BuildSource) SetContextDir(v string)`

SetContextDir sets ContextDir field to given value.

### HasContextDir

`func (o *V1BuildSource) HasContextDir() bool`

HasContextDir returns a boolean if a field has been set.

### GetDockerfile

`func (o *V1BuildSource) GetDockerfile() string`

GetDockerfile returns the Dockerfile field if non-nil, zero value otherwise.

### GetDockerfileOk

`func (o *V1BuildSource) GetDockerfileOk() (*string, bool)`

GetDockerfileOk returns a tuple with the Dockerfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfile

`func (o *V1BuildSource) SetDockerfile(v string)`

SetDockerfile sets Dockerfile field to given value.

### HasDockerfile

`func (o *V1BuildSource) HasDockerfile() bool`

HasDockerfile returns a boolean if a field has been set.

### GetGit

`func (o *V1BuildSource) GetGit() V1GitBuildSource`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *V1BuildSource) GetGitOk() (*V1GitBuildSource, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *V1BuildSource) SetGit(v V1GitBuildSource)`

SetGit sets Git field to given value.

### HasGit

`func (o *V1BuildSource) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetImages

`func (o *V1BuildSource) GetImages() []V1ImageSource`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *V1BuildSource) GetImagesOk() (*[]V1ImageSource, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *V1BuildSource) SetImages(v []V1ImageSource)`

SetImages sets Images field to given value.

### HasImages

`func (o *V1BuildSource) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetSecrets

`func (o *V1BuildSource) GetSecrets() []V1SecretBuildSource`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *V1BuildSource) GetSecretsOk() (*[]V1SecretBuildSource, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *V1BuildSource) SetSecrets(v []V1SecretBuildSource)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *V1BuildSource) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetSourceSecret

`func (o *V1BuildSource) GetSourceSecret() V1LocalObjectReference`

GetSourceSecret returns the SourceSecret field if non-nil, zero value otherwise.

### GetSourceSecretOk

`func (o *V1BuildSource) GetSourceSecretOk() (*V1LocalObjectReference, bool)`

GetSourceSecretOk returns a tuple with the SourceSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSecret

`func (o *V1BuildSource) SetSourceSecret(v V1LocalObjectReference)`

SetSourceSecret sets SourceSecret field to given value.

### HasSourceSecret

`func (o *V1BuildSource) HasSourceSecret() bool`

HasSourceSecret returns a boolean if a field has been set.

### GetType

`func (o *V1BuildSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1BuildSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1BuildSource) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


