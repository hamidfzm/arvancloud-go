# ServerVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewServerVolume

`func NewServerVolume() *ServerVolume`

NewServerVolume instantiates a new ServerVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerVolumeWithDefaults

`func NewServerVolumeWithDefaults() *ServerVolume`

NewServerVolumeWithDefaults instantiates a new ServerVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServerVolume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerVolume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerVolume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerVolume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ServerVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ServerVolume) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ServerVolume) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ServerVolume) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ServerVolume) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


