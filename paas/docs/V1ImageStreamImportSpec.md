# V1ImageStreamImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]V1ImageImportSpec**](V1ImageImportSpec.md) | Images are a list of individual images to import. | [optional] 
**Import** | **bool** | Import indicates whether to perform an import - if so, the specified tags are set on the spec and status of the image stream defined by the type meta. | 
**Repository** | Pointer to [**V1RepositoryImportSpec**](V1RepositoryImportSpec.md) |  | [optional] 

## Methods

### NewV1ImageStreamImportSpec

`func NewV1ImageStreamImportSpec(import_ bool, ) *V1ImageStreamImportSpec`

NewV1ImageStreamImportSpec instantiates a new V1ImageStreamImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageStreamImportSpecWithDefaults

`func NewV1ImageStreamImportSpecWithDefaults() *V1ImageStreamImportSpec`

NewV1ImageStreamImportSpecWithDefaults instantiates a new V1ImageStreamImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *V1ImageStreamImportSpec) GetImages() []V1ImageImportSpec`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *V1ImageStreamImportSpec) GetImagesOk() (*[]V1ImageImportSpec, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *V1ImageStreamImportSpec) SetImages(v []V1ImageImportSpec)`

SetImages sets Images field to given value.

### HasImages

`func (o *V1ImageStreamImportSpec) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetImport

`func (o *V1ImageStreamImportSpec) GetImport() bool`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *V1ImageStreamImportSpec) GetImportOk() (*bool, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *V1ImageStreamImportSpec) SetImport(v bool)`

SetImport sets Import field to given value.


### GetRepository

`func (o *V1ImageStreamImportSpec) GetRepository() V1RepositoryImportSpec`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *V1ImageStreamImportSpec) GetRepositoryOk() (*V1RepositoryImportSpec, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *V1ImageStreamImportSpec) SetRepository(v V1RepositoryImportSpec)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *V1ImageStreamImportSpec) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


