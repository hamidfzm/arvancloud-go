# V1TagEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **string** | Created holds the time the TagEvent was created | 
**DockerImageReference** | **string** | DockerImageReference is the string that can be used to pull this image | 
**Generation** | **int64** | Generation is the spec tag generation that resulted in this tag being updated | 
**Image** | **string** | Image is the image | 

## Methods

### NewV1TagEvent

`func NewV1TagEvent(created string, dockerImageReference string, generation int64, image string, ) *V1TagEvent`

NewV1TagEvent instantiates a new V1TagEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TagEventWithDefaults

`func NewV1TagEventWithDefaults() *V1TagEvent`

NewV1TagEventWithDefaults instantiates a new V1TagEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *V1TagEvent) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *V1TagEvent) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *V1TagEvent) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetDockerImageReference

`func (o *V1TagEvent) GetDockerImageReference() string`

GetDockerImageReference returns the DockerImageReference field if non-nil, zero value otherwise.

### GetDockerImageReferenceOk

`func (o *V1TagEvent) GetDockerImageReferenceOk() (*string, bool)`

GetDockerImageReferenceOk returns a tuple with the DockerImageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageReference

`func (o *V1TagEvent) SetDockerImageReference(v string)`

SetDockerImageReference sets DockerImageReference field to given value.


### GetGeneration

`func (o *V1TagEvent) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *V1TagEvent) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *V1TagEvent) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.


### GetImage

`func (o *V1TagEvent) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V1TagEvent) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V1TagEvent) SetImage(v string)`

SetImage sets Image field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


