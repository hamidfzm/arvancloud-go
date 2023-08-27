# V1ImageStreamLayers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Blobs** | **map[string]interface{}** | blobs is a map of blob name to metadata about the blob. | 
**Images** | **map[string]interface{}** | images is a map between an image name and the names of the blobs and config that comprise the image. | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 

## Methods

### NewV1ImageStreamLayers

`func NewV1ImageStreamLayers(blobs map[string]interface{}, images map[string]interface{}, ) *V1ImageStreamLayers`

NewV1ImageStreamLayers instantiates a new V1ImageStreamLayers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageStreamLayersWithDefaults

`func NewV1ImageStreamLayersWithDefaults() *V1ImageStreamLayers`

NewV1ImageStreamLayersWithDefaults instantiates a new V1ImageStreamLayers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1ImageStreamLayers) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ImageStreamLayers) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ImageStreamLayers) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ImageStreamLayers) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetBlobs

`func (o *V1ImageStreamLayers) GetBlobs() map[string]interface{}`

GetBlobs returns the Blobs field if non-nil, zero value otherwise.

### GetBlobsOk

`func (o *V1ImageStreamLayers) GetBlobsOk() (*map[string]interface{}, bool)`

GetBlobsOk returns a tuple with the Blobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobs

`func (o *V1ImageStreamLayers) SetBlobs(v map[string]interface{})`

SetBlobs sets Blobs field to given value.


### GetImages

`func (o *V1ImageStreamLayers) GetImages() map[string]interface{}`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *V1ImageStreamLayers) GetImagesOk() (*map[string]interface{}, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *V1ImageStreamLayers) SetImages(v map[string]interface{})`

SetImages sets Images field to given value.


### GetKind

`func (o *V1ImageStreamLayers) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ImageStreamLayers) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ImageStreamLayers) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ImageStreamLayers) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1ImageStreamLayers) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ImageStreamLayers) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ImageStreamLayers) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ImageStreamLayers) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


