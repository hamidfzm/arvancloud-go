# V1ImageSourcePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDir** | **string** | destinationDir is the relative directory within the build directory where files copied from the image are placed. | 
**SourcePath** | **string** | sourcePath is the absolute path of the file or directory inside the image to copy to the build directory.  If the source path ends in /. then the content of the directory will be copied, but the directory itself will not be created at the destination. | 

## Methods

### NewV1ImageSourcePath

`func NewV1ImageSourcePath(destinationDir string, sourcePath string, ) *V1ImageSourcePath`

NewV1ImageSourcePath instantiates a new V1ImageSourcePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageSourcePathWithDefaults

`func NewV1ImageSourcePathWithDefaults() *V1ImageSourcePath`

NewV1ImageSourcePathWithDefaults instantiates a new V1ImageSourcePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDir

`func (o *V1ImageSourcePath) GetDestinationDir() string`

GetDestinationDir returns the DestinationDir field if non-nil, zero value otherwise.

### GetDestinationDirOk

`func (o *V1ImageSourcePath) GetDestinationDirOk() (*string, bool)`

GetDestinationDirOk returns a tuple with the DestinationDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDir

`func (o *V1ImageSourcePath) SetDestinationDir(v string)`

SetDestinationDir sets DestinationDir field to given value.


### GetSourcePath

`func (o *V1ImageSourcePath) GetSourcePath() string`

GetSourcePath returns the SourcePath field if non-nil, zero value otherwise.

### GetSourcePathOk

`func (o *V1ImageSourcePath) GetSourcePathOk() (*string, bool)`

GetSourcePathOk returns a tuple with the SourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePath

`func (o *V1ImageSourcePath) SetSourcePath(v string)`

SetSourcePath sets SourcePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


