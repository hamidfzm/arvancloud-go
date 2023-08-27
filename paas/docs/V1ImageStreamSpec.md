# V1ImageStreamSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to [**[]V1TagReference**](V1TagReference.md) | tags map arbitrary string values to specific image locators | [optional] 
**DockerImageRepository** | Pointer to **string** | dockerImageRepository is optional, if specified this stream is backed by a Docker repository on this server Deprecated: This field is deprecated as of v3.7 and will be removed in a future release. Specify the source for the tags to be imported in each tag via the spec.tags.from reference instead. | [optional] 
**LookupPolicy** | Pointer to [**V1ImageLookupPolicy**](V1ImageLookupPolicy.md) |  | [optional] 

## Methods

### NewV1ImageStreamSpec

`func NewV1ImageStreamSpec() *V1ImageStreamSpec`

NewV1ImageStreamSpec instantiates a new V1ImageStreamSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageStreamSpecWithDefaults

`func NewV1ImageStreamSpecWithDefaults() *V1ImageStreamSpec`

NewV1ImageStreamSpecWithDefaults instantiates a new V1ImageStreamSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *V1ImageStreamSpec) GetTags() []V1TagReference`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *V1ImageStreamSpec) GetTagsOk() (*[]V1TagReference, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *V1ImageStreamSpec) SetTags(v []V1TagReference)`

SetTags sets Tags field to given value.

### HasTags

`func (o *V1ImageStreamSpec) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDockerImageRepository

`func (o *V1ImageStreamSpec) GetDockerImageRepository() string`

GetDockerImageRepository returns the DockerImageRepository field if non-nil, zero value otherwise.

### GetDockerImageRepositoryOk

`func (o *V1ImageStreamSpec) GetDockerImageRepositoryOk() (*string, bool)`

GetDockerImageRepositoryOk returns a tuple with the DockerImageRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageRepository

`func (o *V1ImageStreamSpec) SetDockerImageRepository(v string)`

SetDockerImageRepository sets DockerImageRepository field to given value.

### HasDockerImageRepository

`func (o *V1ImageStreamSpec) HasDockerImageRepository() bool`

HasDockerImageRepository returns a boolean if a field has been set.

### GetLookupPolicy

`func (o *V1ImageStreamSpec) GetLookupPolicy() V1ImageLookupPolicy`

GetLookupPolicy returns the LookupPolicy field if non-nil, zero value otherwise.

### GetLookupPolicyOk

`func (o *V1ImageStreamSpec) GetLookupPolicyOk() (*V1ImageLookupPolicy, bool)`

GetLookupPolicyOk returns a tuple with the LookupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupPolicy

`func (o *V1ImageStreamSpec) SetLookupPolicy(v V1ImageLookupPolicy)`

SetLookupPolicy sets LookupPolicy field to given value.

### HasLookupPolicy

`func (o *V1ImageStreamSpec) HasLookupPolicy() bool`

HasLookupPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


