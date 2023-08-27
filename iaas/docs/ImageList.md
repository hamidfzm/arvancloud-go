# ImageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | Pointer to **bool** |  | [optional] 
**Images** | Pointer to [**[]ImgDistributionItem**](ImgDistributionItem.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewImageList

`func NewImageList() *ImageList`

NewImageList instantiates a new ImageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageListWithDefaults

`func NewImageListWithDefaults() *ImageList`

NewImageListWithDefaults instantiates a new ImageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplay

`func (o *ImageList) GetDisplay() bool`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ImageList) GetDisplayOk() (*bool, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ImageList) SetDisplay(v bool)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *ImageList) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.

### GetImages

`func (o *ImageList) GetImages() []ImgDistributionItem`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ImageList) GetImagesOk() (*[]ImgDistributionItem, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ImageList) SetImages(v []ImgDistributionItem)`

SetImages sets Images field to given value.

### HasImages

`func (o *ImageList) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetName

`func (o *ImageList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageList) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


