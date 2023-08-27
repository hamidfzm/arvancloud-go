# VolumeAttachDetachRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** |  | [optional] 
**VolumeId** | Pointer to **string** |  | [optional] 

## Methods

### NewVolumeAttachDetachRequest

`func NewVolumeAttachDetachRequest() *VolumeAttachDetachRequest`

NewVolumeAttachDetachRequest instantiates a new VolumeAttachDetachRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeAttachDetachRequestWithDefaults

`func NewVolumeAttachDetachRequestWithDefaults() *VolumeAttachDetachRequest`

NewVolumeAttachDetachRequestWithDefaults instantiates a new VolumeAttachDetachRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *VolumeAttachDetachRequest) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *VolumeAttachDetachRequest) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *VolumeAttachDetachRequest) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *VolumeAttachDetachRequest) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetVolumeId

`func (o *VolumeAttachDetachRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeAttachDetachRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeAttachDetachRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *VolumeAttachDetachRequest) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


