# ImportImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinDisk** | Pointer to **int64** |  | [optional] 
**MinRam** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewImportImageRequest

`func NewImportImageRequest() *ImportImageRequest`

NewImportImageRequest instantiates a new ImportImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportImageRequestWithDefaults

`func NewImportImageRequestWithDefaults() *ImportImageRequest`

NewImportImageRequestWithDefaults instantiates a new ImportImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinDisk

`func (o *ImportImageRequest) GetMinDisk() int64`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ImportImageRequest) GetMinDiskOk() (*int64, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ImportImageRequest) SetMinDisk(v int64)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ImportImageRequest) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMinRam

`func (o *ImportImageRequest) GetMinRam() int64`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *ImportImageRequest) GetMinRamOk() (*int64, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *ImportImageRequest) SetMinRam(v int64)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *ImportImageRequest) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### GetName

`func (o *ImportImageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportImageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportImageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImportImageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *ImportImageRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImportImageRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImportImageRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ImportImageRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


