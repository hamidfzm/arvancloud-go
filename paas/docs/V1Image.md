# V1Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**DockerImageConfig** | Pointer to **string** | DockerImageConfig is a JSON blob that the runtime uses to set up the container. This is a part of manifest schema v2. | [optional] 
**DockerImageLayers** | [**[]V1ImageLayer**](V1ImageLayer.md) | DockerImageLayers represents the layers in the image. May not be set if the image does not define that data. | 
**DockerImageManifest** | Pointer to **string** | DockerImageManifest is the raw JSON of the manifest | [optional] 
**DockerImageManifestMediaType** | Pointer to **string** | DockerImageManifestMediaType specifies the mediaType of manifest. This is a part of manifest schema v2. | [optional] 
**DockerImageMetadata** | Pointer to **string** | DockerImageMetadata contains metadata about this image | [optional] 
**DockerImageMetadataVersion** | Pointer to **string** | DockerImageMetadataVersion conveys the version of the object, which if empty defaults to \&quot;1.0\&quot; | [optional] 
**DockerImageReference** | Pointer to **string** | DockerImageReference is the string that can be used to pull this image. | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Signatures** | Pointer to [**[]V1ImageSignature**](V1ImageSignature.md) | Signatures holds all signatures of the image. | [optional] 

## Methods

### NewV1Image

`func NewV1Image(dockerImageLayers []V1ImageLayer, ) *V1Image`

NewV1Image instantiates a new V1Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageWithDefaults

`func NewV1ImageWithDefaults() *V1Image`

NewV1ImageWithDefaults instantiates a new V1Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1Image) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1Image) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1Image) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1Image) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDockerImageConfig

`func (o *V1Image) GetDockerImageConfig() string`

GetDockerImageConfig returns the DockerImageConfig field if non-nil, zero value otherwise.

### GetDockerImageConfigOk

`func (o *V1Image) GetDockerImageConfigOk() (*string, bool)`

GetDockerImageConfigOk returns a tuple with the DockerImageConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageConfig

`func (o *V1Image) SetDockerImageConfig(v string)`

SetDockerImageConfig sets DockerImageConfig field to given value.

### HasDockerImageConfig

`func (o *V1Image) HasDockerImageConfig() bool`

HasDockerImageConfig returns a boolean if a field has been set.

### GetDockerImageLayers

`func (o *V1Image) GetDockerImageLayers() []V1ImageLayer`

GetDockerImageLayers returns the DockerImageLayers field if non-nil, zero value otherwise.

### GetDockerImageLayersOk

`func (o *V1Image) GetDockerImageLayersOk() (*[]V1ImageLayer, bool)`

GetDockerImageLayersOk returns a tuple with the DockerImageLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageLayers

`func (o *V1Image) SetDockerImageLayers(v []V1ImageLayer)`

SetDockerImageLayers sets DockerImageLayers field to given value.


### GetDockerImageManifest

`func (o *V1Image) GetDockerImageManifest() string`

GetDockerImageManifest returns the DockerImageManifest field if non-nil, zero value otherwise.

### GetDockerImageManifestOk

`func (o *V1Image) GetDockerImageManifestOk() (*string, bool)`

GetDockerImageManifestOk returns a tuple with the DockerImageManifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageManifest

`func (o *V1Image) SetDockerImageManifest(v string)`

SetDockerImageManifest sets DockerImageManifest field to given value.

### HasDockerImageManifest

`func (o *V1Image) HasDockerImageManifest() bool`

HasDockerImageManifest returns a boolean if a field has been set.

### GetDockerImageManifestMediaType

`func (o *V1Image) GetDockerImageManifestMediaType() string`

GetDockerImageManifestMediaType returns the DockerImageManifestMediaType field if non-nil, zero value otherwise.

### GetDockerImageManifestMediaTypeOk

`func (o *V1Image) GetDockerImageManifestMediaTypeOk() (*string, bool)`

GetDockerImageManifestMediaTypeOk returns a tuple with the DockerImageManifestMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageManifestMediaType

`func (o *V1Image) SetDockerImageManifestMediaType(v string)`

SetDockerImageManifestMediaType sets DockerImageManifestMediaType field to given value.

### HasDockerImageManifestMediaType

`func (o *V1Image) HasDockerImageManifestMediaType() bool`

HasDockerImageManifestMediaType returns a boolean if a field has been set.

### GetDockerImageMetadata

`func (o *V1Image) GetDockerImageMetadata() string`

GetDockerImageMetadata returns the DockerImageMetadata field if non-nil, zero value otherwise.

### GetDockerImageMetadataOk

`func (o *V1Image) GetDockerImageMetadataOk() (*string, bool)`

GetDockerImageMetadataOk returns a tuple with the DockerImageMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageMetadata

`func (o *V1Image) SetDockerImageMetadata(v string)`

SetDockerImageMetadata sets DockerImageMetadata field to given value.

### HasDockerImageMetadata

`func (o *V1Image) HasDockerImageMetadata() bool`

HasDockerImageMetadata returns a boolean if a field has been set.

### GetDockerImageMetadataVersion

`func (o *V1Image) GetDockerImageMetadataVersion() string`

GetDockerImageMetadataVersion returns the DockerImageMetadataVersion field if non-nil, zero value otherwise.

### GetDockerImageMetadataVersionOk

`func (o *V1Image) GetDockerImageMetadataVersionOk() (*string, bool)`

GetDockerImageMetadataVersionOk returns a tuple with the DockerImageMetadataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageMetadataVersion

`func (o *V1Image) SetDockerImageMetadataVersion(v string)`

SetDockerImageMetadataVersion sets DockerImageMetadataVersion field to given value.

### HasDockerImageMetadataVersion

`func (o *V1Image) HasDockerImageMetadataVersion() bool`

HasDockerImageMetadataVersion returns a boolean if a field has been set.

### GetDockerImageReference

`func (o *V1Image) GetDockerImageReference() string`

GetDockerImageReference returns the DockerImageReference field if non-nil, zero value otherwise.

### GetDockerImageReferenceOk

`func (o *V1Image) GetDockerImageReferenceOk() (*string, bool)`

GetDockerImageReferenceOk returns a tuple with the DockerImageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageReference

`func (o *V1Image) SetDockerImageReference(v string)`

SetDockerImageReference sets DockerImageReference field to given value.

### HasDockerImageReference

`func (o *V1Image) HasDockerImageReference() bool`

HasDockerImageReference returns a boolean if a field has been set.

### GetKind

`func (o *V1Image) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1Image) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1Image) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1Image) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1Image) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1Image) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1Image) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1Image) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSignatures

`func (o *V1Image) GetSignatures() []V1ImageSignature`

GetSignatures returns the Signatures field if non-nil, zero value otherwise.

### GetSignaturesOk

`func (o *V1Image) GetSignaturesOk() (*[]V1ImageSignature, bool)`

GetSignaturesOk returns a tuple with the Signatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatures

`func (o *V1Image) SetSignatures(v []V1ImageSignature)`

SetSignatures sets Signatures field to given value.

### HasSignatures

`func (o *V1Image) HasSignatures() bool`

HasSignatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


