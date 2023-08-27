# V1ImageStreamMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Image** | [**V1Image**](V1Image.md) |  | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Tag** | **string** | Tag is a string value this image can be located with inside the stream. | 

## Methods

### NewV1ImageStreamMapping

`func NewV1ImageStreamMapping(image V1Image, tag string, ) *V1ImageStreamMapping`

NewV1ImageStreamMapping instantiates a new V1ImageStreamMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageStreamMappingWithDefaults

`func NewV1ImageStreamMappingWithDefaults() *V1ImageStreamMapping`

NewV1ImageStreamMappingWithDefaults instantiates a new V1ImageStreamMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1ImageStreamMapping) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ImageStreamMapping) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ImageStreamMapping) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ImageStreamMapping) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetImage

`func (o *V1ImageStreamMapping) GetImage() V1Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V1ImageStreamMapping) GetImageOk() (*V1Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V1ImageStreamMapping) SetImage(v V1Image)`

SetImage sets Image field to given value.


### GetKind

`func (o *V1ImageStreamMapping) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ImageStreamMapping) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ImageStreamMapping) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ImageStreamMapping) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1ImageStreamMapping) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ImageStreamMapping) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ImageStreamMapping) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ImageStreamMapping) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTag

`func (o *V1ImageStreamMapping) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V1ImageStreamMapping) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V1ImageStreamMapping) SetTag(v string)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


