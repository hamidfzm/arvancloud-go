# V1ImageStreamImportStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]V1ImageImportStatus**](V1ImageImportStatus.md) | Images is set with the result of importing spec.images | [optional] 
**Import** | Pointer to [**V1ImageStream**](V1ImageStream.md) |  | [optional] 
**Repository** | Pointer to [**V1RepositoryImportStatus**](V1RepositoryImportStatus.md) |  | [optional] 

## Methods

### NewV1ImageStreamImportStatus

`func NewV1ImageStreamImportStatus() *V1ImageStreamImportStatus`

NewV1ImageStreamImportStatus instantiates a new V1ImageStreamImportStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageStreamImportStatusWithDefaults

`func NewV1ImageStreamImportStatusWithDefaults() *V1ImageStreamImportStatus`

NewV1ImageStreamImportStatusWithDefaults instantiates a new V1ImageStreamImportStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *V1ImageStreamImportStatus) GetImages() []V1ImageImportStatus`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *V1ImageStreamImportStatus) GetImagesOk() (*[]V1ImageImportStatus, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *V1ImageStreamImportStatus) SetImages(v []V1ImageImportStatus)`

SetImages sets Images field to given value.

### HasImages

`func (o *V1ImageStreamImportStatus) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetImport

`func (o *V1ImageStreamImportStatus) GetImport() V1ImageStream`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *V1ImageStreamImportStatus) GetImportOk() (*V1ImageStream, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *V1ImageStreamImportStatus) SetImport(v V1ImageStream)`

SetImport sets Import field to given value.

### HasImport

`func (o *V1ImageStreamImportStatus) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetRepository

`func (o *V1ImageStreamImportStatus) GetRepository() V1RepositoryImportStatus`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *V1ImageStreamImportStatus) GetRepositoryOk() (*V1RepositoryImportStatus, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *V1ImageStreamImportStatus) SetRepository(v V1RepositoryImportStatus)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *V1ImageStreamImportStatus) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


