# V1ImageStreamStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]V1NamedTagEventList**](V1NamedTagEventList.md) | Tags are a historical record of images associated with each tag. The first entry in the TagEvent array is the currently tagged image. | [optional] 
**DockerImageRepository** | **string** | DockerImageRepository represents the effective location this stream may be accessed at. May be empty until the server determines where the repository is located | 
**PublicDockerImageRepository** | Pointer to **string** | PublicDockerImageRepository represents the public location from where the image can be pulled outside the cluster. This field may be empty if the administrator has not exposed the integrated registry externally. | [optional] 

## Methods

### NewV1ImageStreamStatus

`func NewV1ImageStreamStatus(dockerImageRepository string, ) *V1ImageStreamStatus`

NewV1ImageStreamStatus instantiates a new V1ImageStreamStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageStreamStatusWithDefaults

`func NewV1ImageStreamStatusWithDefaults() *V1ImageStreamStatus`

NewV1ImageStreamStatusWithDefaults instantiates a new V1ImageStreamStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *V1ImageStreamStatus) GetTags() []V1NamedTagEventList`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *V1ImageStreamStatus) GetTagsOk() (*[]V1NamedTagEventList, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *V1ImageStreamStatus) SetTags(v []V1NamedTagEventList)`

SetTags sets Tags field to given value.

### HasTags

`func (o *V1ImageStreamStatus) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDockerImageRepository

`func (o *V1ImageStreamStatus) GetDockerImageRepository() string`

GetDockerImageRepository returns the DockerImageRepository field if non-nil, zero value otherwise.

### GetDockerImageRepositoryOk

`func (o *V1ImageStreamStatus) GetDockerImageRepositoryOk() (*string, bool)`

GetDockerImageRepositoryOk returns a tuple with the DockerImageRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageRepository

`func (o *V1ImageStreamStatus) SetDockerImageRepository(v string)`

SetDockerImageRepository sets DockerImageRepository field to given value.


### GetPublicDockerImageRepository

`func (o *V1ImageStreamStatus) GetPublicDockerImageRepository() string`

GetPublicDockerImageRepository returns the PublicDockerImageRepository field if non-nil, zero value otherwise.

### GetPublicDockerImageRepositoryOk

`func (o *V1ImageStreamStatus) GetPublicDockerImageRepositoryOk() (*string, bool)`

GetPublicDockerImageRepositoryOk returns a tuple with the PublicDockerImageRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDockerImageRepository

`func (o *V1ImageStreamStatus) SetPublicDockerImageRepository(v string)`

SetPublicDockerImageRepository sets PublicDockerImageRepository field to given value.

### HasPublicDockerImageRepository

`func (o *V1ImageStreamStatus) HasPublicDockerImageRepository() bool`

HasPublicDockerImageRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


