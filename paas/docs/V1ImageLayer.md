# V1ImageLayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaType** | **string** | MediaType of the referenced object. | 
**Name** | **string** | Name of the layer as defined by the underlying store. | 
**Size** | **int64** | Size of the layer in bytes as defined by the underlying store. | 

## Methods

### NewV1ImageLayer

`func NewV1ImageLayer(mediaType string, name string, size int64, ) *V1ImageLayer`

NewV1ImageLayer instantiates a new V1ImageLayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageLayerWithDefaults

`func NewV1ImageLayerWithDefaults() *V1ImageLayer`

NewV1ImageLayerWithDefaults instantiates a new V1ImageLayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaType

`func (o *V1ImageLayer) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *V1ImageLayer) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *V1ImageLayer) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetName

`func (o *V1ImageLayer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ImageLayer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ImageLayer) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *V1ImageLayer) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *V1ImageLayer) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *V1ImageLayer) SetSize(v int64)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


