# V1RepositoryImportStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalTags** | Pointer to **[]string** | AdditionalTags are tags that exist in the repository but were not imported because a maximum limit of automatic imports was applied. | [optional] 
**Images** | Pointer to [**[]V1ImageImportStatus**](V1ImageImportStatus.md) | Images is a list of images successfully retrieved by the import of the repository. | [optional] 
**Status** | Pointer to [**V1Status**](V1Status.md) |  | [optional] 

## Methods

### NewV1RepositoryImportStatus

`func NewV1RepositoryImportStatus() *V1RepositoryImportStatus`

NewV1RepositoryImportStatus instantiates a new V1RepositoryImportStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RepositoryImportStatusWithDefaults

`func NewV1RepositoryImportStatusWithDefaults() *V1RepositoryImportStatus`

NewV1RepositoryImportStatusWithDefaults instantiates a new V1RepositoryImportStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalTags

`func (o *V1RepositoryImportStatus) GetAdditionalTags() []string`

GetAdditionalTags returns the AdditionalTags field if non-nil, zero value otherwise.

### GetAdditionalTagsOk

`func (o *V1RepositoryImportStatus) GetAdditionalTagsOk() (*[]string, bool)`

GetAdditionalTagsOk returns a tuple with the AdditionalTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalTags

`func (o *V1RepositoryImportStatus) SetAdditionalTags(v []string)`

SetAdditionalTags sets AdditionalTags field to given value.

### HasAdditionalTags

`func (o *V1RepositoryImportStatus) HasAdditionalTags() bool`

HasAdditionalTags returns a boolean if a field has been set.

### GetImages

`func (o *V1RepositoryImportStatus) GetImages() []V1ImageImportStatus`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *V1RepositoryImportStatus) GetImagesOk() (*[]V1ImageImportStatus, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *V1RepositoryImportStatus) SetImages(v []V1ImageImportStatus)`

SetImages sets Images field to given value.

### HasImages

`func (o *V1RepositoryImportStatus) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetStatus

`func (o *V1RepositoryImportStatus) GetStatus() V1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1RepositoryImportStatus) GetStatusOk() (*V1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1RepositoryImportStatus) SetStatus(v V1Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1RepositoryImportStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


