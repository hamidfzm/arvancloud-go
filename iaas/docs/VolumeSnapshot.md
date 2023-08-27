# VolumeSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RealSize** | Pointer to **int64** |  | [optional] 
**RealSizeStatus** | Pointer to **bool** |  | [optional] 
**RevertedOn** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VolumeId** | Pointer to **string** |  | [optional] 
**VolumeName** | Pointer to **string** |  | [optional] 

## Methods

### NewVolumeSnapshot

`func NewVolumeSnapshot() *VolumeSnapshot`

NewVolumeSnapshot instantiates a new VolumeSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSnapshotWithDefaults

`func NewVolumeSnapshotWithDefaults() *VolumeSnapshot`

NewVolumeSnapshotWithDefaults instantiates a new VolumeSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *VolumeSnapshot) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeSnapshot) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeSnapshot) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumeSnapshot) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *VolumeSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VolumeSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VolumeSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VolumeSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *VolumeSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageId

`func (o *VolumeSnapshot) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *VolumeSnapshot) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *VolumeSnapshot) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *VolumeSnapshot) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetName

`func (o *VolumeSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRealSize

`func (o *VolumeSnapshot) GetRealSize() int64`

GetRealSize returns the RealSize field if non-nil, zero value otherwise.

### GetRealSizeOk

`func (o *VolumeSnapshot) GetRealSizeOk() (*int64, bool)`

GetRealSizeOk returns a tuple with the RealSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealSize

`func (o *VolumeSnapshot) SetRealSize(v int64)`

SetRealSize sets RealSize field to given value.

### HasRealSize

`func (o *VolumeSnapshot) HasRealSize() bool`

HasRealSize returns a boolean if a field has been set.

### GetRealSizeStatus

`func (o *VolumeSnapshot) GetRealSizeStatus() bool`

GetRealSizeStatus returns the RealSizeStatus field if non-nil, zero value otherwise.

### GetRealSizeStatusOk

`func (o *VolumeSnapshot) GetRealSizeStatusOk() (*bool, bool)`

GetRealSizeStatusOk returns a tuple with the RealSizeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealSizeStatus

`func (o *VolumeSnapshot) SetRealSizeStatus(v bool)`

SetRealSizeStatus sets RealSizeStatus field to given value.

### HasRealSizeStatus

`func (o *VolumeSnapshot) HasRealSizeStatus() bool`

HasRealSizeStatus returns a boolean if a field has been set.

### GetRevertedOn

`func (o *VolumeSnapshot) GetRevertedOn() string`

GetRevertedOn returns the RevertedOn field if non-nil, zero value otherwise.

### GetRevertedOnOk

`func (o *VolumeSnapshot) GetRevertedOnOk() (*string, bool)`

GetRevertedOnOk returns a tuple with the RevertedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertedOn

`func (o *VolumeSnapshot) SetRevertedOn(v string)`

SetRevertedOn sets RevertedOn field to given value.

### HasRevertedOn

`func (o *VolumeSnapshot) HasRevertedOn() bool`

HasRevertedOn returns a boolean if a field has been set.

### GetServerId

`func (o *VolumeSnapshot) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *VolumeSnapshot) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *VolumeSnapshot) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *VolumeSnapshot) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *VolumeSnapshot) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *VolumeSnapshot) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *VolumeSnapshot) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *VolumeSnapshot) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetSize

`func (o *VolumeSnapshot) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeSnapshot) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeSnapshot) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumeSnapshot) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *VolumeSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *VolumeSnapshot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VolumeSnapshot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VolumeSnapshot) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VolumeSnapshot) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVolumeId

`func (o *VolumeSnapshot) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeSnapshot) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeSnapshot) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *VolumeSnapshot) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.

### GetVolumeName

`func (o *VolumeSnapshot) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *VolumeSnapshot) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *VolumeSnapshot) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *VolumeSnapshot) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


