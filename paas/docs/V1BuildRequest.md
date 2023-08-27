# V1BuildRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Binary** | Pointer to [**V1BinaryBuildSource**](V1BinaryBuildSource.md) |  | [optional] 
**DockerStrategyOptions** | Pointer to [**V1DockerStrategyOptions**](V1DockerStrategyOptions.md) |  | [optional] 
**Env** | Pointer to [**[]V1EnvVar**](V1EnvVar.md) | env contains additional environment variables you want to pass into a builder container. | [optional] 
**From** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**LastVersion** | Pointer to **int64** | lastVersion (optional) is the LastVersion of the BuildConfig that was used to generate the build. If the BuildConfig in the generator doesn&#39;t match, a build will not be generated. | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Revision** | Pointer to [**V1SourceRevision**](V1SourceRevision.md) |  | [optional] 
**SourceStrategyOptions** | Pointer to [**V1SourceStrategyOptions**](V1SourceStrategyOptions.md) |  | [optional] 
**TriggeredBy** | [**[]V1BuildTriggerCause**](V1BuildTriggerCause.md) | triggeredBy describes which triggers started the most recent update to the build configuration and contains information about those triggers. | 
**TriggeredByImage** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 

## Methods

### NewV1BuildRequest

`func NewV1BuildRequest(triggeredBy []V1BuildTriggerCause, ) *V1BuildRequest`

NewV1BuildRequest instantiates a new V1BuildRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildRequestWithDefaults

`func NewV1BuildRequestWithDefaults() *V1BuildRequest`

NewV1BuildRequestWithDefaults instantiates a new V1BuildRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1BuildRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1BuildRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1BuildRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1BuildRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetBinary

`func (o *V1BuildRequest) GetBinary() V1BinaryBuildSource`

GetBinary returns the Binary field if non-nil, zero value otherwise.

### GetBinaryOk

`func (o *V1BuildRequest) GetBinaryOk() (*V1BinaryBuildSource, bool)`

GetBinaryOk returns a tuple with the Binary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinary

`func (o *V1BuildRequest) SetBinary(v V1BinaryBuildSource)`

SetBinary sets Binary field to given value.

### HasBinary

`func (o *V1BuildRequest) HasBinary() bool`

HasBinary returns a boolean if a field has been set.

### GetDockerStrategyOptions

`func (o *V1BuildRequest) GetDockerStrategyOptions() V1DockerStrategyOptions`

GetDockerStrategyOptions returns the DockerStrategyOptions field if non-nil, zero value otherwise.

### GetDockerStrategyOptionsOk

`func (o *V1BuildRequest) GetDockerStrategyOptionsOk() (*V1DockerStrategyOptions, bool)`

GetDockerStrategyOptionsOk returns a tuple with the DockerStrategyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerStrategyOptions

`func (o *V1BuildRequest) SetDockerStrategyOptions(v V1DockerStrategyOptions)`

SetDockerStrategyOptions sets DockerStrategyOptions field to given value.

### HasDockerStrategyOptions

`func (o *V1BuildRequest) HasDockerStrategyOptions() bool`

HasDockerStrategyOptions returns a boolean if a field has been set.

### GetEnv

`func (o *V1BuildRequest) GetEnv() []V1EnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *V1BuildRequest) GetEnvOk() (*[]V1EnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *V1BuildRequest) SetEnv(v []V1EnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *V1BuildRequest) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetFrom

`func (o *V1BuildRequest) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1BuildRequest) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1BuildRequest) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.

### HasFrom

`func (o *V1BuildRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetKind

`func (o *V1BuildRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1BuildRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1BuildRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1BuildRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastVersion

`func (o *V1BuildRequest) GetLastVersion() int64`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *V1BuildRequest) GetLastVersionOk() (*int64, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *V1BuildRequest) SetLastVersion(v int64)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *V1BuildRequest) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *V1BuildRequest) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1BuildRequest) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1BuildRequest) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1BuildRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRevision

`func (o *V1BuildRequest) GetRevision() V1SourceRevision`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1BuildRequest) GetRevisionOk() (*V1SourceRevision, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1BuildRequest) SetRevision(v V1SourceRevision)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1BuildRequest) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSourceStrategyOptions

`func (o *V1BuildRequest) GetSourceStrategyOptions() V1SourceStrategyOptions`

GetSourceStrategyOptions returns the SourceStrategyOptions field if non-nil, zero value otherwise.

### GetSourceStrategyOptionsOk

`func (o *V1BuildRequest) GetSourceStrategyOptionsOk() (*V1SourceStrategyOptions, bool)`

GetSourceStrategyOptionsOk returns a tuple with the SourceStrategyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStrategyOptions

`func (o *V1BuildRequest) SetSourceStrategyOptions(v V1SourceStrategyOptions)`

SetSourceStrategyOptions sets SourceStrategyOptions field to given value.

### HasSourceStrategyOptions

`func (o *V1BuildRequest) HasSourceStrategyOptions() bool`

HasSourceStrategyOptions returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *V1BuildRequest) GetTriggeredBy() []V1BuildTriggerCause`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *V1BuildRequest) GetTriggeredByOk() (*[]V1BuildTriggerCause, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *V1BuildRequest) SetTriggeredBy(v []V1BuildTriggerCause)`

SetTriggeredBy sets TriggeredBy field to given value.


### GetTriggeredByImage

`func (o *V1BuildRequest) GetTriggeredByImage() V1ObjectReference`

GetTriggeredByImage returns the TriggeredByImage field if non-nil, zero value otherwise.

### GetTriggeredByImageOk

`func (o *V1BuildRequest) GetTriggeredByImageOk() (*V1ObjectReference, bool)`

GetTriggeredByImageOk returns a tuple with the TriggeredByImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredByImage

`func (o *V1BuildRequest) SetTriggeredByImage(v V1ObjectReference)`

SetTriggeredByImage sets TriggeredByImage field to given value.

### HasTriggeredByImage

`func (o *V1BuildRequest) HasTriggeredByImage() bool`

HasTriggeredByImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


