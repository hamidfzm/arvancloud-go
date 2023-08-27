# ServerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**Documents** | Pointer to [**[]ImgDoc**](ImgDoc.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**MinDisk** | Pointer to **int32** |  | [optional] 
**MinRam** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewServerImage

`func NewServerImage() *ServerImage`

NewServerImage instantiates a new ServerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerImageWithDefaults

`func NewServerImageWithDefaults() *ServerImage`

NewServerImageWithDefaults instantiates a new ServerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ServerImage) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServerImage) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServerImage) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ServerImage) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDocuments

`func (o *ServerImage) GetDocuments() []ImgDoc`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ServerImage) GetDocumentsOk() (*[]ImgDoc, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ServerImage) SetDocuments(v []ImgDoc)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *ServerImage) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetId

`func (o *ServerImage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerImage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerImage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerImage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ServerImage) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServerImage) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServerImage) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServerImage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMinDisk

`func (o *ServerImage) GetMinDisk() int32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ServerImage) GetMinDiskOk() (*int32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ServerImage) SetMinDisk(v int32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ServerImage) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMinRam

`func (o *ServerImage) GetMinRam() int32`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *ServerImage) GetMinRamOk() (*int32, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *ServerImage) SetMinRam(v int32)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *ServerImage) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### GetName

`func (o *ServerImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *ServerImage) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *ServerImage) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *ServerImage) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *ServerImage) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetOsVersion

`func (o *ServerImage) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ServerImage) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ServerImage) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *ServerImage) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetProgress

`func (o *ServerImage) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ServerImage) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ServerImage) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ServerImage) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSize

`func (o *ServerImage) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ServerImage) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ServerImage) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ServerImage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *ServerImage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerImage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerImage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerImage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsername

`func (o *ServerImage) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServerImage) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServerImage) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServerImage) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


