# V1ImageImportStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**V1Image**](V1Image.md) |  | [optional] 
**Status** | [**V1Status**](V1Status.md) |  | 
**Tag** | Pointer to **string** | Tag is the tag this image was located under, if any | [optional] 

## Methods

### NewV1ImageImportStatus

`func NewV1ImageImportStatus(status V1Status, ) *V1ImageImportStatus`

NewV1ImageImportStatus instantiates a new V1ImageImportStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageImportStatusWithDefaults

`func NewV1ImageImportStatusWithDefaults() *V1ImageImportStatus`

NewV1ImageImportStatusWithDefaults instantiates a new V1ImageImportStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *V1ImageImportStatus) GetImage() V1Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V1ImageImportStatus) GetImageOk() (*V1Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V1ImageImportStatus) SetImage(v V1Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *V1ImageImportStatus) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetStatus

`func (o *V1ImageImportStatus) GetStatus() V1Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ImageImportStatus) GetStatusOk() (*V1Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ImageImportStatus) SetStatus(v V1Status)`

SetStatus sets Status field to given value.


### GetTag

`func (o *V1ImageImportStatus) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V1ImageImportStatus) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V1ImageImportStatus) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V1ImageImportStatus) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


